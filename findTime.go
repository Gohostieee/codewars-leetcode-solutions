package main

import "strconv"

func HumanReadableTime(seconds int) string {
	// your code here
	time := ""
	minutes := 0
	hours := 0
	parsed := 0
	for seconds > 0 {
		seconds--
		parsed++
		if parsed == 60 {
			minutes++
			if minutes == 60 {
				hours++
				minutes = 0
			}
			parsed = 0

		}

	}
	charMinutes := ""
	charSeconds := ""
	charHour := ""
	if minutes < 10 {
		charMinutes = "0" + strconv.Itoa(minutes)
	} else {
		charMinutes = strconv.Itoa(minutes)
	}
	if parsed < 10 {
		charSeconds = "0" + strconv.Itoa(parsed)
	} else {
		charSeconds = strconv.Itoa(parsed)
	}
	if hours < 10 {
		charHour = "0" + strconv.Itoa(hours)
	} else {
		charHour = strconv.Itoa(hours)
	}
	time = charHour + ":" + charMinutes + ":" + charSeconds
	return time
}

func main() {

}
