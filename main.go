package main

import (
	"math/rand"
	"runtime"
	"time"

	"./app"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	rand.Seed(time.Now().UnixNano())
	app := app.App{}
	app.Init()
}
