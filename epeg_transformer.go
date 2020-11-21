package gothumb

import (
	"github.com/williamhaley/goepeg"
)

type EpegTransformer struct {
	In  string
	Out string
}

func NewEpegTransformer(input string, output string) *EpegTransformer {
	return &EpegTransformer{input, output}
}

func (t EpegTransformer) None() error {
	return goepeg.Transform(t.In, t.Out, goepeg.TransformNone)
}

func (t EpegTransformer) FlipH() error {
	return goepeg.Transform(t.In, t.Out, goepeg.TransformFlipH)
}

func (t EpegTransformer) FlipV() error {
	return goepeg.Transform(t.In, t.Out, goepeg.TransformFlipV)
}

func (t EpegTransformer) Transpose() error {
	return goepeg.Transform(t.In, t.Out, goepeg.TransformTranspose)
}

func (t EpegTransformer) Rotate90() error {
	return goepeg.Transform(t.In, t.Out, goepeg.TransformRot90)
}

func (t EpegTransformer) Rotate180() error {
	return goepeg.Transform(t.In, t.Out, goepeg.TransformRot180)
}

func (t EpegTransformer) Rotate270() error {
	return goepeg.Transform(t.In, t.Out, goepeg.TransformRot270)
}

func (t EpegTransformer) Transverse() error {
	return goepeg.Transform(t.In, t.Out, goepeg.TransformTransverse)
}
