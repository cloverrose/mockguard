package mockguard_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/cloverrose/mockguard"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, mockguard.Analyzer, "a", "b", "c")
}
