// package triangle is small library for determine type triangle
package triangle

// type Kind - type of triangle
type Kind int

const (
	NaT Kind = iota // not a triangle
	Equ             // equilateral
	Iso             // isosceles
	Sca             // scalene
)

// func KindFromSides takes the value of the three sides of a triangle and returns the type of the triangle
func KindFromSides(a, b, c float64) Kind {
	switch {
	case a+b <= c || a+c <= b || b+c <= a:
		return NaT
	case a == b && a == c:
		return Equ
	case a == b || a == c || b == c:
		return Iso
	default:
		return Sca
	}
}
