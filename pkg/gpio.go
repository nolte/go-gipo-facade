package pkg

import "github.com/stianeikeland/go-rpio"

func ControllGIPO() error {
	err := rpio.Open()
	if err != nil {
		return err
	}

	pin := rpio.Pin(2)

	pin.Toggle()
	return nil
}
