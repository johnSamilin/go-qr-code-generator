package main

import (
	"encoding/base64"
	"syscall/js"

	qrcode "github.com/skip2/go-qrcode"
)

func generateQrCode(this js.Value, args []js.Value) interface{} {
	var png []byte
	png, err := qrcode.Encode(args[0].String(), qrcode.Medium, 256)

	if err != nil {
		return err
	}

	return base64.StdEncoding.EncodeToString(png)
}

func registerCallbacks() {
	js.Global().Set("generateQrCode", js.FuncOf(generateQrCode))
}

func main() {
	c := make(chan bool)
	registerCallbacks()

	<-c
}
