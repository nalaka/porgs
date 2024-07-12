package main

import (
	"embed"
)

type Plugin struct{}

func (p *Plugin) GetName() string {
	return "core"
}

func (p *Plugin) GetDependencies() []string {
	return nil
}

//go:embed assets/*
//go:embed layouts/*.go.html
//go:embed views/*.go.html
var embeddedFS embed.FS

func (p *Plugin) GetFS() embed.FS {
	return embeddedFS
}
