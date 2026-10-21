package main

import (
	"github.com/konsole-log/Physics-Engine-.git/physics"
)

func main() {
	v := physics.NewVector3(10, 20, 30)
	a := physics.NewVector3(10, 20, 30)
	result := v.AddScaledVector(a, 0.5)
	w := physics.Vector3{X:0,Y:0,Z:0}
	w.Print()
	result.Print()
}
