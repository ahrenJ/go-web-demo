package main

import (
	"lite-talk/models"
	"lite-talk/routers"
)

func main() {
	models.Setup()

	r := routers.InitRouter()
	r.Run()
}
