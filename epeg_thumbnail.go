package gothumb

import (
	"github.com/koofr/goepeg"
)

func EpegThumbnail(input string, output string, size int, quality int) (err error) {
	return goepeg.Thumbnail(input, output, size, quality)
}
