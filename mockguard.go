package mockguard

import (
	"flag"
	"path/filepath"
	"regexp"
	"strings"

	"golang.org/x/tools/go/analysis"
)

const doc = "mockguard checks if mockgen is used in conventional filename and options."

// Analyzer checks usage of the mockgen follows convention.
var Analyzer = &analysis.Analyzer{
	Name:     "mockguard",
	Doc:      doc,
	Run:      run,
	Requires: []*analysis.Analyzer{},
	Flags:    *flag.NewFlagSet("mockguard", flag.ExitOnError),
}

func init() {
	Analyzer.Flags.StringVar(&FileName, "FileName", FileName, "conventional file name for mockgen")
	Analyzer.Flags.StringVar(&Content, "Content", Content, "conventional comment content for mockgen")
}

func checkFileName(fileName string) (bool, string) {
	expected := FileName
	base := filepath.Base(fileName)
	return base == expected, expected
}

func checkOptions(comment string) (bool, string) {
	expected := Content
	after, found := strings.CutPrefix(comment, expected)
	if !found {
		// comment should starts with this prefix.
		return false, expected
	}
	// comment can have after (this is basically for unit test)
	mo, err := regexp.MatchString(`(\s*//.*)?`, after)
	if err != nil {
		return false, ""
	}
	return mo, expected
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		// Iterate through comments in the file
		for _, group := range file.Comments {
			for _, comment := range group.List {
				if strings.HasPrefix(comment.Text, "//go:generate mockgen") {
					fileName := pass.Fset.File(file.Pos()).Name()
					if ok, expected := checkFileName(fileName); !ok {
						pass.Reportf(comment.Pos(), "incorrect mockgen file, expected: %s", expected)
					}
					if ok, expected := checkOptions(comment.Text); !ok {
						pass.Reportf(comment.Pos(), "incorrect mockgen options, expected: %s", expected)
					}
					break
				}
			}
		}
	}
	return nil, nil
}
