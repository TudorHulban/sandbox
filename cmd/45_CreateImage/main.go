package main

import "fmt"

func main() {
	fmt.Println(
		CreateImage(
			&ParamsCreateImage{
				PathFont:  _fontPath,
				PathImage: "xxx.png",

				PixelWidth:  700,
				PixelHeight: 700,

				ColorText:       _textColor,
				ColorBackground: _backgroundColor,
			},
		),
	)
}
