package main

import (
	"GroupieTracker/controller"
)

func main() {
	controller.Init()
	controller.GetDataByID(0)
}
