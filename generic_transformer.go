package gothumb

import (
	"errors"
	"image"
	"image/draw"
)

type GenericTransformer struct {
	In  image.Image
	Out image.Image
}

func (t GenericTransformer) None() error {
	t.Out = t.In

	return nil
}

func (t GenericTransformer) FlipH() error {
	out, err := FlipH(t.In)

	if err != nil {
		return err
	}

	t.Out = out

	return nil
}

func (t GenericTransformer) FlipV() error {
	out, err := FlipV(t.In)

	if err != nil {
		return err
	}

	t.Out = out

	return nil
}

func (t GenericTransformer) Transpose() error {
	out, err := Transpose(t.In)

	if err != nil {
		return err
	}

	t.Out = out

	return nil
}

func (t GenericTransformer) Rotate90() error {
	out, err := Rotate(t.In, 90)

	if err != nil {
		return err
	}

	t.Out = out

	return nil
}

func (t GenericTransformer) Rotate180() error {
	out, err := Rotate(t.In, 180)

	if err != nil {
		return err
	}

	t.Out = out

	return nil
}

func (t GenericTransformer) Rotate270() error {
	out, err := Rotate(t.In, 270)

	if err != nil {
		return err
	}

	t.Out = out

	return nil
}

func (t GenericTransformer) Transverse() error {
	out, err := Transverse(t.In)

	if err != nil {
		return err
	}

	t.Out = out

	return nil
}

func toRGBA(src image.Image) *image.RGBA {
	m := image.NewRGBA(image.Rect(0, 0, src.Bounds().Dx(), src.Bounds().Dy()))
	draw.Draw(m, m.Bounds(), src, src.Bounds().Min, draw.Src)
	return m
}

func Rotate(s image.Image, deg int) (image.Image, error) {
	src := toRGBA(s)
	var d image.Rectangle
	switch deg {
	default:
		return nil, errors.New("Unsupported angle (90, 180, 270).")
	case 90, 270:
		d = image.Rect(0, 0, src.Bounds().Size().Y, src.Bounds().Size().X)
	case 180:
		d = image.Rect(0, 0, src.Bounds().Size().X, src.Bounds().Size().Y)
	}
	rv := image.NewRGBA(d)
	b := src.Bounds()
	/* switch outside of loops for performance reasons */
	switch deg {
	case 270:
		for y := 0; y < b.Size().Y; y++ {
			for x := 0; x < b.Size().X; x++ {
				s := x*rv.Stride + 4*(d.Size().X-y-1)
				p := y*src.Stride + x*4
				copy(rv.Pix[s:s+4], src.Pix[p:p+4])
			}
		}
	case 180:
		for y := 0; y < b.Size().Y; y++ {
			for x := 0; x < b.Size().X; x++ {
				s := (d.Size().Y-y-1)*rv.Stride + 4*d.Size().X - (x+1)*4
				p := y*src.Stride + x*4
				copy(rv.Pix[s:s+4], src.Pix[p:p+4])
			}
		}
	case 90:
		for y := 0; y < b.Size().Y; y++ {
			for x := 0; x < b.Size().X; x++ {
				s := (d.Size().Y-x-1)*rv.Stride + y*4
				p := y*src.Stride + x*4
				copy(rv.Pix[s:s+4], src.Pix[p:p+4])
			}
		}
	}
	return rv, nil
}

func FlipH(s image.Image) (image.Image, error) {
	src := toRGBA(s)
	d := image.Rect(0, 0, src.Bounds().Size().X, src.Bounds().Size().Y)
	rv := image.NewRGBA(d)
	b := src.Bounds()
	for y := 0; y < b.Size().Y; y++ {
		for x := 0; x < b.Size().X; x++ {
			s := y*rv.Stride + x*4
			p := y*src.Stride + (b.Size().X-x-1)*4
			copy(rv.Pix[s:s+4], src.Pix[p:p+4])
		}
	}
	return rv, nil
}

func FlipV(s image.Image) (image.Image, error) {
	src := toRGBA(s)
	d := image.Rect(0, 0, src.Bounds().Size().X, src.Bounds().Size().Y)
	rv := image.NewRGBA(d)
	b := src.Bounds()
	for y := 0; y < b.Size().Y; y++ {
		for x := 0; x < b.Size().X; x++ {
			s := y*rv.Stride + x*4
			p := (b.Size().Y-y-1)*src.Stride + x*4
			copy(rv.Pix[s:s+4], src.Pix[p:p+4])
		}
	}
	return rv, nil
}

func Transpose(s image.Image) (image.Image, error) {
	src := toRGBA(s)
	d := image.Rect(0, 0, src.Bounds().Size().Y, src.Bounds().Size().X)
	rv := image.NewRGBA(d)
	b := src.Bounds()
	for y := 0; y < b.Size().Y; y++ {
		for x := 0; x < b.Size().X; x++ {
			s := x*rv.Stride + y*4
			p := y*src.Stride + x*4
			copy(rv.Pix[s:s+4], src.Pix[p:p+4])
		}
	}
	return rv, nil
}

func Transverse(s image.Image) (image.Image, error) {
	src := toRGBA(s)
	d := image.Rect(0, 0, src.Bounds().Size().Y, src.Bounds().Size().X)
	rv := image.NewRGBA(d)
	b := src.Bounds()
	for y := 0; y < b.Size().Y; y++ {
		for x := 0; x < b.Size().X; x++ {
			s := (d.Size().Y-x-1)*rv.Stride + (d.Size().X-y-1)*4
			p := y*src.Stride + x*4
			copy(rv.Pix[s:s+4], src.Pix[p:p+4])
		}
	}
	return rv, nil
}
