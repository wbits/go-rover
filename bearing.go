package main

type bearing struct {
	direction int
}

type movable interface {
	up()
	right()
	down()
	left()
}

func newBearing() bearing {
	return bearing{0}
}

func (b bearing) String() string {
	compass := [4]string{"N", "E", "S", "W"}

	return compass[b.direction % 4]
}

func (b *bearing) turnRight() {
	b.direction++
}

func (b *bearing) turnLeft() {
	b.direction += 3
}

func (b bearing) move(o movable, forward bool) {
	move := [4]func(){o.up, o.right, o.down, o.left}
	index := b.direction
	if !forward {
		index += 2
	}

	move[index % 4]()
}
