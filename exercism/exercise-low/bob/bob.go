package bob

import "strings"

type Remark struct {
	Remark string
}

func NewRemark(remark string) Remark {
	return Remark{strings.TrimSpace(remark)}
}

func (r Remark) IsScream() bool {
	maybeScream := false
	for a := 'A'; a <= 'Z'; a++ {
		if strings.ContainsRune(r.Remark, a) {
			maybeScream = true
			break
		}
	}

	return strings.ToUpper(r.Remark) == r.Remark && maybeScream
}

func (r Remark) IsQuestion() bool {
	return strings.HasSuffix(r.Remark, "?")
}

func (r Remark) isSilence() bool {
	return r.Remark == ""
}

// func  Hey imitated teenager answer
func Hey(remark string) string {
	r := NewRemark(remark)

	switch {
	case r.isSilence():
		return "Fine. Be that way!"
	case r.IsScream() && r.IsQuestion():
		return "Calm down, I know what I'm doing!"
	case r.IsScream() && !r.IsQuestion():
		return "Whoa, chill out!"
	case r.IsQuestion():
		return "Sure."
	default:
		return "Whatever."
	}
}
