package flying

import (
	"fmt"
)

type myDrone struct{}

func (d myDrone) Start() {
	fmt.Println("starting the drone")
	// drone := tello.NewDriver("8888")

	// work := func() {
	// 	drone.TakeOff()

	// 	gobot.After(5*time.Second, func() {
	// 		drone.Land()
	// 	})
	// }

	// robot := gobot.NewRobot("tello",
	// 	[]gobot.Connection{},
	// 	[]gobot.Device{drone},
	// 	work,
	// )

	// robot.Start()
}

func NewMyDrone() *myDrone {
	return &myDrone{}
}
