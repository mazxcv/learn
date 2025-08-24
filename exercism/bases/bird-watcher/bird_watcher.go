package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	total := 0
	for _, v := range birdsPerDay {
		total += v
	}
	return total
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	const weekCountDay = 7

	index := (week - 1) * weekCountDay
	lastIndex := index + weekCountDay
	if lastIndex > len(birdsPerDay) {
		lastIndex = len(birdsPerDay)
	}
	total := 0
	if index > len(birdsPerDay) {
		return 0
	} else {
		for i := index; i < lastIndex; i++ {
			total += birdsPerDay[i]
		}
		return total
	}
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i, v := range birdsPerDay {
		if i%2 != 0 {
			birdsPerDay[i] = v
		} else {
			birdsPerDay[i] = v + 1
		}
	}
	return birdsPerDay
}
