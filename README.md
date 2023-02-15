# WASM QR code generator
This module generates QR codes and is based on https://github.com/skip2/go-qrcode.

# Usage
## Import initializer
First of all, include script `./wasm_exec.js` to your webpage.

## Load WASM
````js
function loadWasm(pathToWasm) {
    const go = new Go()
    return new Promise((resolve, reject) => {
        WebAssembly.instantiateStreaming(fetch(path), go.importObject)
            .then(result => {
                go.run(result.instance)
                resolve(result.instance)
            })
            .catch(error => {
                reject(error)
            })
    })
}
````

## Use exported function
When loaded, this package adds function `generateQrCode` to `window` object. Just call it with a single parameter containing URL or text to be encoded. The function will return base64 encoded string with QR code.
````js
await loadWasm('./build/go-qr-code-generator.wasm')
const base64QRCode = generateQrCode('https://ros-plata.ru')
````

## Set src attribute
````js
const image = document.querySelector('img#qr')
image.setAttribute('src', `data:image/png;base64,${base64QrCode}`)
````

# Development
- `git clone`
- run `go get`
- do changes
- run `GOOS=js GOARCH=wasm go build -o  ./build/go-qr-code-generator.wasm`
See [this article](https://golangbot.com/webassembly-using-go/) for details.
