package main

import "fmt"

type musical interface {
	play()
	listen()
}

type instrument struct {
	instru string
}

func (i instrument) play() {
	fmt.Println("Play " + i.instru)
}

func (i instrument) listen() {
	fmt.Println("Listen " + i.instru)
}

func main() {
	musical.play(instrument{instru: "guitar"})
	musical.listen(instrument{instru: "piano"})
}
