package context

import "html/template"

type App struct {
	Config *Config
	IndexTpl *template.Template
	ErrorTpl *template.Template
}

type Config struct {
	DisplayName string `yaml:"displayName"`
	DisplayDescription string `yaml:"displayDescription"`
	UploadPath string `yaml:"uploadPath"`
	Port int `yaml:"port"`
}
