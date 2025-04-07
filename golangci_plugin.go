package mockguard

import (
	"github.com/golangci/plugin-module-register/register"
	"golang.org/x/tools/go/analysis"
)

func init() {
	RegisterPlugin()
}

func RegisterPlugin() {
	// https://golangci-lint.run/plugins/module-plugins/
	register.Plugin("mockguard", newPlugin)
}

func newPlugin(conf any) (register.LinterPlugin, error) {
	s, err := register.DecodeSettings[settings](conf)
	if err != nil {
		return nil, err
	}

	return &plugin{settings: &s}, nil
}

type settings struct {
	FileName string
	Content  string
}

type plugin struct {
	settings *settings
}

func (p *plugin) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	if p.settings.FileName != "" {
		FileName = p.settings.FileName
	}
	if p.settings.Content != "" {
		Content = p.settings.Content
	}
	return []*analysis.Analyzer{
		Analyzer,
	}, nil
}

func (p *plugin) GetLoadMode() string {
	return register.LoadModeSyntax
}

var _ register.LinterPlugin = &plugin{}
