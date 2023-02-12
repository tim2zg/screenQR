package main

import (
	"fmt"
	"github.com/gen2brain/beeep"
	"github.com/kbinani/screenshot"
	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
	"golang.design/x/clipboard"
)

func main() {
	err := clipboard.Init()
	if err != nil {
		panic(err)
	}

	captureDisplay, err := screenshot.CaptureDisplay(0) // capture primary display
	if err != nil {
		panic(err)
	}

	// prepare BinaryBitmap
	bmp, _ := gozxing.NewBinaryBitmapFromImage(captureDisplay)

	// decode image
	qrReader := qrcode.NewQRCodeReader()
	result, _ := qrReader.Decode(bmp, nil)

	if result != nil {
		fmt.Println(result.String())
		err := beeep.Notify("QR Code", result.String(), "assets/qr.png")
		if err != nil {
			return
		}
		clipboard.Write(clipboard.FmtText, []byte(result.String()))
	}
}
