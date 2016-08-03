package gothumb

import (
	"github.com/koofr/resize"
	"image"
	_ "image/gif"
	"image/jpeg"
	_ "image/png"
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

	img, _, err := image.Decode(reader)

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

	opts := &jpeg.Options{
		Quality: quality,
	}

	jpeg.Encode(writer, thumb, opts)

	return
}
