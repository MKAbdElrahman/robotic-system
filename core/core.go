package core

import (
	"context"
	"fmt"
)

type RobotPlugin interface {
	Initialize() error
	PerformAction(ctx context.Context) error
	ReportResults() string
}

type Robot struct {
	plugins []RobotPlugin
}

func (rc *Robot) AddPlugin(p RobotPlugin) {
	rc.plugins = append(rc.plugins, p)
}

// Execute runs the robot core and all attached plugins
func (rc *Robot) Execute() {
	fmt.Println("Executing Robot Core")

	for _, plugin := range rc.plugins {
		if err := plugin.Initialize(); err != nil {
			fmt.Printf("Error initializing plugin: %v\n", err)
			continue
		}

		ctx := context.Background() // Use a context for passing shared data or services

		if err := plugin.PerformAction(ctx); err != nil {
			fmt.Printf("Error executing plugin: %v\n", err)
			continue
		}

		results := plugin.ReportResults()
		fmt.Printf("Plugin Results: %s\n", results)
	}
}
