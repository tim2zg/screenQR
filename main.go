package main

import (
	"fmt"
	"github.com/kbinani/screenshot"
	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
)

func main() {
	captureDisplay, err := screenshot.CaptureDisplay(0) // capture primary display
	if err != nil {
		panic(err)
	}

	// prepare BinaryBitmap
	bmp, _ := gozxing.NewBinaryBitmapFromImage(captureDisplay)

	// decode image
	qrReader := qrcode.NewQRCodeReader()
	result, _ := qrReader.Decode(bmp, nil)

	fmt.Println(result)
}
