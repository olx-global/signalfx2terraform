// (c) Copyright 2017-2019 OLX
package main

//https://godoc.org/github.com/hashicorp/hcl2/hclwrite#Block.BuildTokens

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/hashicorp/hcl2/hclwrite"
	b "github.com/olx-global/signalfx2terraform/builder"
	"github.com/signalfx/signalfx-go"
)

var (
	signalfxAPIToken string
	dashboardID      string
	detectorID       string
	version          string
)

const (
	// APIURL : default entrypoint for customer's requests
	APIURL = "https://api.eu0.signalfx.com"
)

func main() {

	signalfxAPITokenPtr := flag.String("token", "", "SignalFX Token")
	dashboardIDPtr := flag.String("dashboard", "", "Dashboard ID, without URL")
	detectorIDPtr := flag.String("detector", "", "Detector ID, without URL")
	versionPtr := flag.Bool("v", false, "Prints current exporter version")
	versionLongPtr := flag.Bool("version", false, "Prints current exporter version")
	flag.Parse()
	if *versionPtr || *versionLongPtr {
		fmt.Printf("Version: %s\r\n", version)
		os.Exit(0)
	}

	signalfxAPIToken = *signalfxAPITokenPtr
	if signalfxAPIToken == "" {
		usagePrint("Can't find token")
	}
	// TODO: Simplify that
	dashboardID = *dashboardIDPtr
	detectorID = *detectorIDPtr

	if dashboardID != "" && detectorID != "" {
		usagePrint("Please choose one - dashboard or detector.")
	}

	if dashboardID == "" && detectorID == "" {
		usagePrint("Can't find any objects for exporting, use dashboard or detector at least. ")
	}

	if dashboardID != "" {
		dashboardProcessor(dashboardID)
		os.Exit(0) // TODO: Maybe better use else here.
	}

	if detectorID != "" {
		detectorProcessor(detectorID, signalfxAPIToken)
		os.Exit(0)
	}
}

func usagePrint(message string) {
	log.Println(message)
	fmt.Println(`Usage:
  -dashboard string
        Dashboard ID, without URL
  -detector string
        Detector ID, without URL
  -httptest.serve string
        if non-empty, httptest.NewServer serves on this address and blocks
  -token string
        SignalFX Token
  -v    Prints current exporter version
  -version
        Prints current exporter version`)
	log.Fatal("Exit")
}

func dashboardProcessor(dashboardID string) {

	client, err := signalfx.NewClient(signalfxAPIToken, signalfx.APIUrl(APIURL))

	if err != nil {
		log.Fatal("Something wrong with API client")
	}
	dashboard, err := client.GetDashboard(dashboardID)

	if err != nil {
		log.Printf("Dashboard error: %v", err)
		log.Fatal("Can't fetch dashboard")
	}
	charts := dashboard.Charts

	f := hclwrite.NewEmptyFile()

	b.CreateDashboard(f, dashboard, client)

	for _, v := range charts {
		chart, err := client.GetChart(v.ChartId)
		if err != nil {
			log.Printf("Chart error: %v", err)
			log.Fatal("Can't get chart")
		}

		switch types := chart.Options.Type; types {
		case "SingleValue":
			b.SingleValueChart(f, chart)

		case "Heatmap":
			b.HeatMapChart(f, chart)

		case "TimeSeriesChart":
			b.TimeSeriesChart(f, chart)

		case "List":
			b.ListChart(f, chart)

		case "Text":
			b.TextChart(f, chart)

		default:
			continue
		}

	}
	fmt.Printf("%s", f.Bytes())

}

func detectorProcessor(detectorID string, token string) {
	client, err := signalfx.NewClient(signalfxAPIToken, signalfx.APIUrl(APIURL))
	if err != nil {
		log.Fatal("Something wrong with API client")
	}
	detector, err := client.GetDetector(detectorID)
	if err != nil {
		// Failover for V1 detector
		// TODO: Implement reliable check
		//log.Printf("Detector error: %v", err)
		//log.Println("Can't fetch detector with V2 API, trying failover method...")
		f := hclwrite.NewEmptyFile()
		b.CreateDetectorV1(f, APIURL, detectorID, token)
		fmt.Printf("%s", f.Bytes())

	} else {
		f := hclwrite.NewEmptyFile()
		b.CreateDetector(f, detector)
		fmt.Printf("%s", f.Bytes())

	}

}
