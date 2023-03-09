package main

import (
	"gocv.io/x/gocv"
)

func main() {
	webcam, err := gocv.VideoCaptureDevice(0)
	if err != nil {
		panic(err)
		println("Error opening webcam")
		return
	}
	defer webcam.Close()

	window := gocv.NewWindow("Detector")
	defer window.Close()

	classfier := gocv.NewCascadeClassifier()
	defer classfier.Close()

	classfier.Load("haarcascade_frontalface_default.xml")

	for {
		img := gocv.NewMat()
		if ok := webcam.Read(&img); !ok {
			println("Cannot read from webcam")
			return
		}

	}
}
