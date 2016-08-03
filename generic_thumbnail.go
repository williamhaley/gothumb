package gothumb

import (
	"github.com/koofr/resize"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"github.com/jkmcnk/goepeg"
)

func GenericThumbnail(input string, output string, size int, quality int, scaleType goepeg.ScaleType) (err error) {
	reader, err := os.Open(input)

	if err != nil {
		return err
	}

	defer reader.Close()

	writer, err := os.OpenFile(output, os.O_CREATE|os.O_WRONLY, 0600)

	if err != nil {
		return err
	}

	defer writer.Close()

	img, fmt, err := image.Decode(reader)

	if err != nil {
		return err
	}

	var thumb image.Image

	if scaleType == goepeg.ScaleTypeFitMin {
		if img.Bounds().Size().X >= img.Bounds().Size().Y {
			if img.Bounds().Size().X > size {
				thumb = resize.Resize(0, uint(size), img, resize.NearestNeighbor)
			} else {
				thumb = img
			}
		} else {
			if img.Bounds().Size().Y > size {
				thumb = resize.Resize(uint(size), 0, img, resize.NearestNeighbor)
			} else {
				thumb = img
			}
		}
	} else {
		if img.Bounds().Size().X >= img.Bounds().Size().Y {
			if img.Bounds().Size().X > size {
				thumb = resize.Resize(uint(size), 0, img, resize.NearestNeighbor)
			} else {
				thumb = img
			}
		} else {
			if img.Bounds().Size().Y > size {
				thumb = resize.Resize(0, uint(size), img, resize.NearestNeighbor)
			} else {
				thumb = img
			}
		}
	}

	if fmt == "png" {
		png.Encode(writer, thumb)
	} else if fmt == "gif" {
		opts := &gif.Options{
			NumColors: 256,
		}
		gif.Encode(writer, thumb, opts)
	} else {
		opts := &jpeg.Options{
			Quality: quality,
		}
		jpeg.Encode(writer, thumb, opts)
	}

	return
}
