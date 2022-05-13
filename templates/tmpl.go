package templates

import "embed"

//go:embed *
var Tmpl embed.FS

type Template struct {
	Name string `yaml:"name"`
	File File   `yaml:"file"`
	Type string `yaml:"type"`
}

type File struct {
	SourceName      string `yaml:"source_name"`
	DestinationName string `yaml:"destination_name"`
}
