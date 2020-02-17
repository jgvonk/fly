package flying

import (
	"fmt"
)

func hi() {
	fmt.Println("hi")
}

type myDrone struct{}

func (d myDrone) start() {
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

func newMyDrone() *myDrone {
	return &myDrone{}
}
