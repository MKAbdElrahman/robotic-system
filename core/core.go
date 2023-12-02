package core

import "fmt"

type RobotPlugin interface {
	PerformAction()
}

type Robot struct {
	plugins []RobotPlugin
}

func (rc *Robot) AddPlugin(p RobotPlugin) {
	rc.plugins = append(rc.plugins, p)
}

func (rc *Robot) Execute() {
	fmt.Println("Executing Robot Core")

	for _, plugin := range rc.plugins {
		plugin.PerformAction()
	}
}
