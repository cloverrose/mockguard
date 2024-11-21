package main

import (
	"golang.org/x/tools/go/analysis/unitchecker"

	"github.com/cloverrose/mockguard"
)

func main() { unitchecker.Main(mockguard.Analyzer) }
