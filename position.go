package main

type position struct {
	x, y int
}

func (p *position) up() {
	p.y++
}

func (p *position) right() {
	p.x++
}

func (p *position) down() {
	p.y--
}

func (p *position) left() {
	p.x--
}