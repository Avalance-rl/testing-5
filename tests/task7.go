package tests

import "time"

func GetGreetingMessage(time time.Time) string {
	currentHour := time.Hour()
	if currentHour < 12 {
		return "Good morning!"
	} else if currentHour < 18 {
		return "Good afternoon!"
	}
	return "Good evening!"
}
