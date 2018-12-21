package camera

import (
	"gocv.io/x/gocv"
)

func NewWebCam() (vc *gocv.VideoCapture, err error) {
	webcam, err := gocv.VideoCaptureDevice(0)
	return webcam, err
}
