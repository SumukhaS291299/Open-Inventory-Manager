package qrutils

import (
	"image/color"
	"openinventorymanager/logger"

	qrcode "github.com/skip2/go-qrcode"
)

func Generate(url string) {

	// png, err := qrcode.Encode(url, qrcode.Medium, 256)
	err := qrcode.WriteColorFile(url, qrcode.Medium, 256, color.Black, color.White, "qr.png")

	if err != nil {
		logger.Logger.Error(err.Error())
	}
	// fmt.Println(png)
}
