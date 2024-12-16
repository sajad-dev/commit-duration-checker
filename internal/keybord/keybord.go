package keybord

import (

	"github.com/robotn/gohook"
)

func CheckEndKey(keypress chan int) bool {

	select {
	case key := <-keypress:
		if key == 11111111 {
			hook.End()
			return true
		}
	default:
		return false
	}
	return false
}

var ev = hook.Start()

func Keyboard(keypress chan int) {

	for eventChan := range ev {
		switch {
		case eventChan.Kind == hook.KeyDown:

			if eventChan.Keychar == 27 {

				keypress <- 11111111
				return

			}

			keypress <- int(eventChan.Keychar)

		}
	}
}
