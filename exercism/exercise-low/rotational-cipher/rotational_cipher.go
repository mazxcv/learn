package rotationalcipher

import (
	"unicode"
)

// type RingAlphabet realized scope runes [Alpha, Omega]
// Rotate - shift
// Middle - [Alpha, Middle]: rune + shift; (Middle, Omega]: rune - count + shift
type RingAlphabet struct {
	Alpha    rune
	Omega    rune
	Middle   rune
	Rotate   int
	Distance int
}

func NewAlphabet(begin rune, end rune, shiftKey int) RingAlphabet {
	distance := end - begin + 1
	rotate := shiftKey % int(distance)
	return RingAlphabet{
		Alpha:    begin,
		Omega:    end,
		Middle:   end - rune(rotate),
		Rotate:   rotate,
		Distance: int(distance),
	}
}

func (r RingAlphabet) rotate(a rune) rune {
	if a <= r.Omega && a >= r.Alpha {
		if a <= r.Middle {
			return a + rune(r.Rotate)
		}
		return a - (rune(r.Distance) - rune(r.Rotate))
	} else {
		return a
	}

}

func RotationalCipher(plain string, shiftKey int) (text string) {

	r := NewAlphabet('a', 'z', shiftKey)
	R := NewAlphabet('A', 'Z', shiftKey)
	for _, v := range plain {
		if unicode.IsUpper(v) {
			text += string(R.rotate(v))
		} else {
			text += string(r.rotate(v))
		}

	}

	return
}
