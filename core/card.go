package core

import "errors"

type Card struct {
	Color Color
	Value Value
}

var (
	ErrInvalidColor    = errors.New("invalid color")
	ErrInvalidValue    = errors.New("invalid value")
	ErrInvalidCardCode = errors.New("invalid card code")
)

func (c *Card) IsWild() bool {
	return c.Color == CWild || c.Value == VWild || c.Value == VWildDrawFour
}

func (c *Card) HasEffect() bool {
	return c.Value == VDrawTwo || c.Value == VSkip || c.Value == VReverse || c.Value == VWild || c.Value == VWildDrawFour
}

func (c *Card) CanPlayOn(other Card) bool {
	if c.IsWild() {
		return true
	}

	return c.Color == other.Color || c.Value == other.Value
}

func (c *Card) String() string {
	return c.Color.String() + c.Value.String()
}

func (c *Card) FromString(code string) error {
	if len(code) < 2 {
		return ErrInvalidCardCode
	}

	colorLetter := code[0:1]
	valueLetter := code[1:2]

	var color Color
	var value Value
	// determines the color
	switch colorLetter {
	case "B":
		color = CBlue
	case "G":
		color = CGreen
	case "R":
		color = CRed
	case "Y":
		color = CYellow
	case "W":
		color = CWild
	default:
		return ErrInvalidColor
	}

	switch valueLetter {
	case "0":
		value = VZero
	case "1":
		value = VOne
	case "2":
		value = VTwo
	case "3":
		value = VThree
	case "4":
		value = VFour
	case "5":
		value = VFive
	case "6":
		value = VSix
	case "7":
		value = VSeven
	case "8":
		value = VEight
	case "9":
		value = VNine
	case "S":
		value = VSkip
	case "R":
		value = VReverse
	case "T":
		value = VDrawTwo
	case "W":
		value = VWild
	case "F":
		value = VWildDrawFour
	default:
		return ErrInvalidValue
	}

	c.Color = color
	c.Value = value

	return nil
}

type Color int
type Value int

const (
	CRed Color = iota + 1
	CYellow
	CGreen
	CBlue
	CWild // Used for Wild and Wild Draw Four
)

const (
	VZero Value = iota
	VOne
	VTwo
	VThree
	VFour
	VFive
	VSix
	VSeven
	VEight
	VNine
	VSkip
	VReverse
	VDrawTwo
	VWild         // Wild
	VWildDrawFour // Wild Draw Four
)

func (c Color) String() string {
	switch c {
	case CBlue:
		return "B"
	case CGreen:
		return "G"
	case CRed:
		return "R"
	case CYellow:
		return "Y"
	case CWild:
		return "W"
	default:
		return "U"
	}
}

func (v Value) String() string {
	switch v {
	case VZero:
		return "0"
	case VOne:
		return "1"
	case VTwo:
		return "2"
	case VThree:
		return "3"
	case VFour:
		return "4"
	case VFive:
		return "5"
	case VSix:
		return "6"
	case VSeven:
		return "7"
	case VEight:
		return "8"
	case VNine:
		return "9"
	case VSkip:
		return "S"
	case VReverse:
		return "R"
	case VDrawTwo:
		return "T"
	case VWild:
		return "W"
	case VWildDrawFour:
		return "F"
	default:
		return "U"
	}
}
