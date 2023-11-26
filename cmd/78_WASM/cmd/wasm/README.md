## Compile 

Go
```sh
	GOOS=js GOARCH=wasm go build -o  ../../assets/test.wasm
```

TinyGo
```sh
	tinygo build -o ../../assets/test.wasm -target wasm ./main.go
```