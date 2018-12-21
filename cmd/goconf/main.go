package main

import (
	"github.com/Dabuti/goconf/camera"
	"gocv.io/x/gocv"
)

func main() {
	webcam, _ := camera.NewWebCam()
	window := gocv.NewWindow("Hello")
	img := gocv.NewMat()

	for {
		webcam.Read(&img)
		window.IMShow(img)
		window.WaitKey(1)
	}
}
