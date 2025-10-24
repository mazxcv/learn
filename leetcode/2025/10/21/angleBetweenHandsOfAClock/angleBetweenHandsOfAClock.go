package anglebetweenhandsofaclock

// https://leetcode.com/problems/angle-between-hands-of-a-clock/description/

// Given two numbers, hour and minutes, return the smaller angle (in degrees) formed between the hour and the minute hand.

// Идея
// часовая стрелка вращается на 360 градусов в 12 часов -> 360 / 12 = 30 градусов в за 1 час
// часовая стрелка двигается на 30 градусов за 60 минут -> 30 / 60 = 0.5 градуса в за 1 минуту
// минутная стрелка двигается на 360 градусов в 60 минут -> 360 / 60 = 6 градусов в за 1 минуту
// разность между двумя стрелками надо из часовой стрелки отнять минутную по модулю 180 градусов
// часовую стрелку предварительно привести по модулю в 12

func angleClock(hour int, minutes int) float64 {
	hour = hour % 12
	minutes = minutes % 60

	hourAngle := float64(hour*30) + float64(minutes)*0.5
	minuteAngle := float64(minutes) * 6

	res := abs(hourAngle - minuteAngle)
	if res > 180 {
		res = 360 - res
	}
	return res
}

func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}
