package main

import (
	"fmt"
	"log"
	"time"

	"github.com/tarm/serial"
)

var FRAME_SYNC_COMMAND = []byte("SYNC")

const UART_DEVICE = "/dev/ttyUSB0"
const UART_BAUDRATE = 115200

func chunkSlice(slice []byte, chunkSize int) [][]byte {
	var chunks [][]byte
	for {
		if len(slice) == 0 {
			break
		}

		// necessary check to avoid slicing beyond
		// slice capacity
		if len(slice) < chunkSize {
			chunkSize = len(slice)
		}

		chunks = append(chunks, slice[0:chunkSize])
		slice = slice[chunkSize:]
	}

	return chunks
}

func main() {

	c := &serial.Config{Name: UART_DEVICE, Baud: UART_BAUDRATE}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}

	for {

		for _, frame := range dragon.Frames {
			s.Write(FRAME_SYNC_COMMAND)
			fmt.Println("Frame BEGIN")

			for _, chunk := range chunkSlice(frame, 128) {
				time.Sleep(time.Millisecond * 50)
				_, err = s.Write(chunk)
				if err != nil {
					log.Fatal(err)
				}

				fmt.Print(".")

			}
		}
	}

}
