package gothumb

import (
	"github.com/jkmcnk/goepeg"
)

func EpegThumbnail(input string, output string, size int, quality int, scaleType goepeg.ScaleType) (err error) {
	return goepeg.Thumbnail(input, output, size, quality, scaleType)
}
