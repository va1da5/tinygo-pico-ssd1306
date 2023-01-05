package main

import (
	"machine"
	"time"

	"tinygo.org/x/drivers/ssd1306"
)

func show(display ssd1306.Device, animation Animation, repeat int) {
	if repeat == 0 {
		repeat = 2
	}

	for x := 0; x < repeat; x++ {
		for i := 0; i < len(animation.Frames); i++ {
			time.Sleep(animation.Delay)
			err := display.SetBuffer(animation.Next())
			if err != nil {
				println(err)
			}
			display.Display()
		}
	}
}

func main() {

	// 400kHz gives a better refresh rate and also the display did not work with 100kHz mode
	machine.I2C0.Configure(machine.I2CConfig{Frequency: machine.TWI_FREQ_400KHZ})

	// the delay is needed for display start from a cold reboot, not sure why
	time.Sleep(time.Second * 1)

	display := ssd1306.NewI2C(machine.I2C0)
	display.Configure(ssd1306.Config{Width: 128, Height: 64, Address: ssd1306.Address_128_32, VccState: ssd1306.SWITCHCAPVCC})
	display.ClearDisplay()

	for {

		show(display, pepe, 10)

		show(display, dragon, 10)
	}

}
