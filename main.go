package main

import (
	"fmt"

	_ "github.com/sdboyer/dep-test"
	_ "github.com/sdboyer/deptest"
	_ "github.com/sdboyer/deptestdos"
)

func main() {
	fmt.Println("go!")
}
