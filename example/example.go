package main

import (
	"gothumb"
	"io"
	"os"
)

func main() {
	input := "image.jpg"
	output := "thumb.jpg"
	size := 1024
	quality := 85

	in, err := os.Open(input)

	if err != nil {
		panic(err)
	}

	out, err := gothumb.Thumbnail(in, size, quality)

	if err != nil {
		panic(err)
	}

	defer out.Close()

	outputFile, err := os.OpenFile(output, os.O_CREATE|os.O_WRONLY, 0664)

	if err != nil {
		panic(err)
	}

	_, err = io.Copy(outputFile, out)

	if err != nil {
		panic(err)
	}
}
