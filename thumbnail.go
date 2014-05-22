package gothumb

import (
	"github.com/koofr/go-ioutils"
	"io"
	"io/ioutil"
	"os"
)

func Thumbnail(in io.Reader, size int, quality int) (out io.ReadCloser, err error) {
	input, err := ioutil.TempFile("", "gothumb-input-")

	if err != nil {
		return
	}

	defer os.Remove(input.Name())

	_, err = io.Copy(input, in)

	if err != nil {
		return
	}

	err = input.Close()

	if err != nil {
		return
	}

	thumbnail, err := ioutil.TempFile("", "gothumb-thumbnail-")

	if err != nil {
		return
	}

	defer os.Remove(thumbnail.Name())

	err = thumbnail.Close()

	if err != nil {
		return
	}

	output, err := ioutil.TempFile("", "gothumb-output-")

	if err != nil {
		return
	}

	err = output.Close()

	if err != nil {
		return
	}

	exifReader, err := os.Open(input.Name())

	if err != nil {
		return
	}

	defer exifReader.Close()

	orientation, err := Orientation(exifReader)

	isJpeg := err == nil

	if isJpeg {
		err = EpegThumbnail(input.Name(), thumbnail.Name(), size, quality)
	} else {
		err = GenericThumbnail(input.Name(), thumbnail.Name(), size, quality)
	}

	if err != nil {
		return
	}

	if isJpeg && orientation != 1 {
		transformer := NewEpegTransformer(thumbnail.Name(), output.Name())

		err = Transform(orientation, transformer)

		if err != nil {
			return
		}
	} else {
		os.Rename(thumbnail.Name(), output.Name())
	}

	outputFile, err := os.Open(output.Name())

	if err != nil {
		return
	}

	out = ioutils.NewFileRemoveReader(outputFile)

	return
}
