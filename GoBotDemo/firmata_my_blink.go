// +build OMIT
package main

import (
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/firmata"
)

func main() { // OMIT START
	var i int
	//firmataAdaptor := firmata.NewAdaptor("/dev/ttyACM0")
	firmataAdaptor := firmata.NewAdaptor("/dev/tty.usbmodem14241")
	var led [6]*gpio.LedDriver
	var led7 [6]*gpio.LedDriver

	led[0] = gpio.NewLedDriver(firmataAdaptor, "2")
	led[1] = gpio.NewLedDriver(firmataAdaptor, "3")
	led[2] = gpio.NewLedDriver(firmataAdaptor, "4")
	led[3] = gpio.NewLedDriver(firmataAdaptor, "5")
	led[4] = gpio.NewLedDriver(firmataAdaptor, "6")
	led[5] = gpio.NewLedDriver(firmataAdaptor, "7")

	led7[0] = gpio.NewLedDriver(firmataAdaptor, "8")
	led7[1] = gpio.NewLedDriver(firmataAdaptor, "9")
	led7[2] = gpio.NewLedDriver(firmataAdaptor, "10")
	led7[3] = gpio.NewLedDriver(firmataAdaptor, "11")
	led7[4] = gpio.NewLedDriver(firmataAdaptor, "12")
	led7[5] = gpio.NewLedDriver(firmataAdaptor, "13")

	work := func() {
		gobot.Every(100000000*time.Nanosecond, func() {
			i = i % 6
			led[i].Toggle()
			led7[i].Toggle()
			i = i + 1
		})
	}

	robot := gobot.NewRobot("bot",
		[]gobot.Connection{firmataAdaptor},
		[]gobot.Device{led[0], led[1], led[2], led[3], led[4], led[5], led7[0], led7[1], led7[2], led7[3], led7[4], led7[5]},
		work,
	)

	robot.Start()
} // OMIT END
