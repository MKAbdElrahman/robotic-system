package main

import (
	"robot/core"
	"robot/plugin"
)

func main() {
	// Create an instance of the robot core
	robot := core.Robot{}

	// Register and add plugins to the robot core
	motionControl := &plugin.MotionControl{}
	sensorIntegration := &plugin.SensorIntegration{}

	robot.AddPlugin(motionControl)
	robot.AddPlugin(sensorIntegration)

	// Execute the robot core and all attached plugins
	robot.Execute()
}
