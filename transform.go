package gothumb

import (
	"fmt"
)

type Transformer interface {
	None() error
	FlipH() error
	FlipV() error
	Transpose() error
	Rotate90() error
	Rotate180() error
	Rotate270() error
	Transverse() error
}

func Transform(orientation int, transformer Transformer) error {
	switch orientation {
	case 1:
		return transformer.None()
	case 2:
		return transformer.FlipH()
	case 3:
		return transformer.Rotate180()
	case 4:
		return transformer.FlipV()
	case 5:
		return transformer.Transpose()
	case 6:
		return transformer.Rotate270()
	case 7:
		return transformer.Transverse()
	case 8:
		return transformer.Rotate90()
	}

	return fmt.Errorf("Invalid orientation: %d", orientation)
}
