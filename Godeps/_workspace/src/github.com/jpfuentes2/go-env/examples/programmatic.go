package main

import (
	"fmt"
	"github.com/FoxComm/libs/Godeps/_workspace/src/github.com/jpfuentes2/go-env"
	"os"
	"path"
)

// Note this is exactly what autoload.go does
func main() {
	pwd, _ := os.Getwd()
	env.ReadEnv(path.Join(pwd, ".env"))
	for _, v := range os.Environ() {
		fmt.Println(v)
	}
}
