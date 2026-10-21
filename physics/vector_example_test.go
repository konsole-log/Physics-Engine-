package physics_test

import (
	"fmt"

	"github.com/konsole-log/Physics-Engine-.git/physics"
)

func ExampleNewVector3() {
	v := physics.NewVector3(1, 2, 3)
	fmt.Println(v.X, v.Y, v.Z)
	// Output: 1 2 3
}
