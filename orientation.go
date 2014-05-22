package gothumb

import (
	"github.com/rwcarlsen/goexif/exif"
	"io"
)

func Orientation(reader io.Reader) (orientation int, err error) {
	orientation = 1

	info, err := exif.Decode(reader)

	if err != nil {
		return
	}

	orientTag, err := info.Get(exif.Orientation)

	if err != nil {
		return 1, nil
	}

	orientation = int(orientTag.Int(0))

	return
}
