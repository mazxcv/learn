// Package leap is a library to determine leap year (in the Gregorian calendar)
package leap

// IsLeapYear determine and returns if the passed in year is a leap year.
// In every year that is evenly divisible by 4.
// Unless the year is evenly divisible by 100,
// in which case it's only a leap year if the year is also evenly divisible by 400.
func IsLeapYear(year int) bool {
	divBy4 := year%4 == 0
	divBy100 := year%100 == 0
	divBy400 := year%400 == 0

	return divBy4 && (divBy400 || !divBy100 && !divBy400)
}
