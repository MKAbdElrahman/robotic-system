package core

import (
	"context"
	"fmt"
	"sync"
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

// Execute runs the robot core and all attached plugins concurrently
func (rc *Robot) Execute() {
	fmt.Println("Executing Robot Core")

	var wg sync.WaitGroup
	wg.Add(len(rc.plugins))

	for _, plugin := range rc.plugins {
		go func(p RobotPlugin) {
			defer wg.Done()

			if err := p.Initialize(); err != nil {
				fmt.Printf("Error initializing plugin: %v\n", err)
				return
			}

			ctx := context.Background() // Use a context for passing shared data or services

			if err := p.PerformAction(ctx); err != nil {
				fmt.Printf("Error executing plugin: %v\n", err)
				return
			}

			results := p.ReportResults()
			fmt.Printf("Plugin Results: %s\n", results)
		}(plugin)
	}

	wg.Wait()
}
