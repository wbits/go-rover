package main

import (
	"strconv"
)

type rover struct {
	position
	bearing bearing
}

func newRover() rover {
	return rover{position{0, 0}, newBearing()}
}

func (r rover) exec(commands string) map[string]string {
	for i := range commands {
		r.executeCommand(commands[i])
	}

	return map[string]string{
		"x": strconv.Itoa(r.x),
		"y": strconv.Itoa(r.y),
		"h": r.bearing.String(),
	}
}

func (r *rover) executeCommand(command byte) {
	switch command {
	case 'f':
		r.bearing.move(r,true)
	case 'b':
		r.bearing.move(r, false)
	case 'r':
		r.bearing.turnRight()
	case 'l':
		r.bearing.turnLeft()
	}
}

func main() {

}
