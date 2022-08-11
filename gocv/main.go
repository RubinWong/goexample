package main

import (
	"fmt"
	"time"

	"gocv.io/x/gocv"
)

func main() {
	// Capture(24)

	// c := make(chan gocv.Mat, 100)
	// go Capture2(&c)
	// Play(&c)

	Capture3()
}

func Capture(framerate int64) {
	webcam, err := gocv.VideoCaptureDevice(0)
	if err != nil {
		fmt.Println("Error opening video capture device: ", err)
		return
	}
	window := gocv.NewWindow("Hello")
	img := gocv.NewMat()

	to := time.NewTicker(time.Duration(int64(time.Second) / framerate))
	to2 := time.NewTicker(time.Second)
	var frames int64
	for {
		select {
		case <-to.C:
			if ok := webcam.Read(&img); !ok {
				fmt.Println("Error reading from video capture device")
				return
			}
			if img.Empty() {
				fmt.Println("empty image")
				return
			}
			// window.IMShow(img)
			frames++
		case <-to2.C:
			fmt.Println("FPS: ", frames)
			frames = 0
		}
		key := window.WaitKey(1)
		if key == 27 {
			return
		}
	}
}

func Capture2(c *chan gocv.Mat) {
	webcam, err := gocv.VideoCaptureDevice(0)
	if err != nil {
		fmt.Println("Error opening video capture device: ", err)
		return
	}
	webcam.Set(gocv.VideoCaptureFPS, 30)
	webcam.Set(gocv.VideoCaptureFrameWidth, 1280)
	webcam.Set(gocv.VideoCaptureFrameHeight, 720)
	img := gocv.NewMat()

	to := time.NewTicker(time.Second)
	var frames int64
	for {
		if ok := webcam.Read(&img); !ok {
			fmt.Println("Error reading from video capture device")
			return
		}
		if img.Empty() {
			fmt.Println("empty image")
			return
		}
		*c <- img
		select {
		case <-to.C:
			fmt.Println("FPS: ", frames)
			frames = 0
		default:
			frames++
		}
	}
}

func Play(c *chan gocv.Mat) {
	window := gocv.NewWindow("Hello")
	for {
		select {
		case img := <-*c:
			window.IMShow(img)
		}
		key := window.WaitKey(1)
		if key == 27 {
			return
		}
	}
}

func Capture3() {
	webcam, err := gocv.VideoCaptureDevice(0)
	if err != nil {
		fmt.Println("Error opening video capture device: ", err)
		return
	}
	webcam.Set(gocv.VideoCaptureFPS, 30)
	webcam.Set(gocv.VideoCaptureFrameWidth, 640)
	webcam.Set(gocv.VideoCaptureFrameHeight, 480)

	window := gocv.NewWindow("Hello")
	img := gocv.NewMat()

	to := time.NewTicker(time.Second)
	var frames int64
	for {
		if ok := webcam.Read(&img); !ok {
			fmt.Println("Error reading from video capture device")
			return
		}
		if img.Empty() {
			fmt.Println("empty image")
			return
		}
		// window.IMShow(img)
		// frames++
		select {
		case <-to.C:
			fmt.Println("FPS: ", frames)
			frames = 0
		default:
			frames++
		}
		key := window.WaitKey(1)
		if key == 27 {
			return
		}
	}
}
