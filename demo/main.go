package main

import (
	"fmt"

	"github.com/konsole-log/Physics-Engine-.git/physics"
)

func main() {
	v:=physics.NewVector3(10,20,30)
	v.Normalize()
	l:=v.Magnitude()
	fmt.Print(l)
	v.Print()
}
