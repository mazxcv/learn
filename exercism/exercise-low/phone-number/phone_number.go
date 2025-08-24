package phonenumber

import (
	"fmt"
	"strings"
)

func Number(phoneNumber string) (string, error) {

	var sb strings.Builder
	for _, v := range phoneNumber {
		if v <= '9' && v >= '0' {
			sb.WriteRune(v)
		}
	}
	phone := sb.String()

	if len(phone) < 10 || len(phone) > 11 {
		return "", fmt.Errorf("length phonenumber must be 10 or more digits")
	}
	if len(phone) == 11 && (phone[0] != '1' || phone[1] < '2' || phone[4] < '2') {
		return "", fmt.Errorf("phonenumber must be 1(NXX)-NXX-XXXX template")
	}
	if len(phone) == 10 && (phone[0] < '2' || phone[3] < '2') {
		return "", fmt.Errorf("phonenumber must be 1(NXX)-NXX-XXXX template")
	}

	return phone[len(phone)-10:], nil
}

func AreaCode(phoneNumber string) (string, error) {
	phone, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	return phone[:3], nil
}

func Format(phoneNumber string) (string, error) {
	phoneNumber, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("(%s) %s-%s", phoneNumber[:3], phoneNumber[3:6], phoneNumber[6:10]), nil
}
