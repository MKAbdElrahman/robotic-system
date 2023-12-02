package plugin

import "fmt"

type SensorIntegration struct{}

func (sip SensorIntegration) PerformAction() {
	fmt.Println("Executing Sensor Integration Plugin")
}
