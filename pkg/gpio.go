package pkg

import (
	"fmt"
	"os"

	"github.com/stianeikeland/go-rpio"
)

func ControllGIPO() {
	// Open and map memory to access gpio, check for errors
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Unmap gpio memory when done
	defer rpio.Close()
	pin := rpio.Pin(2)

	// Set pin to output mode
	pin.Output()

	pin.Toggle()
}
