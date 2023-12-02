package plugin

import "fmt"

type MotionControl struct{}

func (mcp MotionControl) PerformAction() {
	fmt.Println("Executing Motion Control Plugin")
}
