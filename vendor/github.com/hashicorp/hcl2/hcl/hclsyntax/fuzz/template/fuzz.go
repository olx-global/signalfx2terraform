package fuzztemplate

import (
	"github.com/hashicorp/hcl2/hcl"
	"github.com/hashicorp/hcl2/hcl/hclsyntax"
)

func Fuzz(data []byte) int {
	_, diags := hclsyntax.ParseTemplate(data, "<fuzz-tmpl>", hcl.Pos{Line: 1, Column: 1})

	if diags.HasErrors() {
		return 0
	}

	return 1
}
