package main

import (
	"bytes"
	"machine"
	"time"

	"tinygo.org/x/drivers/ssd1306"
)

// TODO: investigate and fix frame drops

var imageReceived []byte

const UART_CHUNK_SIZE = 128
const OLED_IMAGE_BUFFER_SIZE = 1024

var FRAME_SYNC_COMMAND = []byte("SYNC")

func main() {

	uart := machine.DefaultUART
	uart.Configure(machine.UARTConfig{})

	var uartReceived = make([]byte, UART_CHUNK_SIZE)

	machine.I2C0.Configure(machine.I2CConfig{Frequency: machine.TWI_FREQ_400KHZ})

	time.Sleep(time.Second * 1)
	display := ssd1306.NewI2C(machine.I2C0)
	display.Configure(ssd1306.Config{Width: 128, Height: 64, Address: ssd1306.Address_128_32, VccState: ssd1306.SWITCHCAPVCC})

	display.ClearBuffer()
	display.ClearDisplay()

	err := display.SetBuffer(image)
	if err != nil {
		println(err)
	}

	display.Display()

	// clear the initial buffer value
	if uart.Buffered() > 0 {
		uart.ReadByte()
	}

	for {
		time.Sleep(time.Millisecond * 20)

		if uart.Buffered() > 0 {
			received_bytes := uart.Buffered()
			// println("Received", received_bytes)
			uart.Read(uartReceived)
			imageReceived = append(imageReceived, uartReceived[0:received_bytes]...)

			syncIndex := bytes.Index(imageReceived, FRAME_SYNC_COMMAND)

			if syncIndex > -1 {
				imageReceived = append([]byte{}, imageReceived[syncIndex+len(FRAME_SYNC_COMMAND):]...)
			}

		}

		if len(imageReceived) >= OLED_IMAGE_BUFFER_SIZE {
			err := display.SetBuffer(imageReceived[:OLED_IMAGE_BUFFER_SIZE])
			if err != nil {
				println(err)
			}
			err = display.Display()
			if err != nil {
				println(err)
			}

			if len(imageReceived) > OLED_IMAGE_BUFFER_SIZE {
				imageReceived = append([]byte{}, imageReceived[OLED_IMAGE_BUFFER_SIZE:]...)
				continue
			}

			// reset image buffer after display
			imageReceived = nil
		}

	}
}
