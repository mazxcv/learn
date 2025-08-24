package wordy

import (
	"regexp"
	"strconv"
	"strings"
)

type Lexer struct {
}

// func Answer Parse and evaluate simple math word problems returning the answer as an integer
// What is 3 plus 2 multiplied by 3?
func Answer(question string) (res int, ok bool) {
	if !strings.HasPrefix(question, "What is ") {
		return
	}

	question = strings.TrimPrefix(question, "What is ")
	question = strings.Trim(question, "?")
	question = strings.ReplaceAll(question, "divided by", "dividedby")
	question = strings.ReplaceAll(question, "multiplied by", "multipliedby")

	if res, err := strconv.Atoi(question); err == nil {
		return res, true
	}

	for {
		re := regexp.MustCompile(`\-?\d+\s+\w+\s+\-?\d+`)
		sequence := re.FindString(question)
		if sequence != "" {
			replace, ok := Prepare(sequence)
			if !ok {
				return 0, false
			}
			question = strings.Replace(question, sequence, replace, 1)
		} else {
			res, err := strconv.Atoi(question)
			if err != nil {
				return 0, false
			}
			return res, true
		}
	}
}

func Prepare(sequence string) (string, bool) {
	re := regexp.MustCompile(`(\-?\d+)\s+(\w+)\s+(\-?\d+)`)
	values := re.FindStringSubmatch(sequence)

	if len(values) < 3 {
		return "", false
	}
	left, err := strconv.Atoi(values[1])
	if err != nil {
		return "", false
	}
	right, err := strconv.Atoi(values[3])
	if err != nil {
		return "", false
	}
	op := values[2]

	switch op {
	case "plus":
		return strconv.Itoa(left + right), true
	case "minus":
		return strconv.Itoa(left - right), true
	case "multipliedby":
		return strconv.Itoa(left * right), true
	case "dividedby":
		return strconv.Itoa(left / right), true
	default:
		return "", false
	}
}
