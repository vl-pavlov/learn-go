package main

import (
	"fmt"
	"time"
)

func main() {

	//--	Key Functions

	// time.Now(): Returns the current time.
	// time.Date(): Creates a new time.Time object for a specific date and time.
	// time.Parse(layout, value): Parses a formatted string into a time.Time object.
	// time.Format(layout): Formats a time.Time object into a string according to the layout.
	// time.Add(duration): Adds a duration to a time.Time object.
	// time.Sub(time): Subtracts two time.Time objects to get the duration between them.
	// time.Sleep(duration): Pauses the current goroutine for the specified duration.

	//--	Getting the Current Time

	currentTime := time.Now()
	fmt.Println("Current time:", currentTime)

	//--	Creating a Time Object for a Specific Date

	someonesBirthday := time.Date(1990, time.May, 10, 23, 12, 5, 3, time.UTC)
	fmt.Println("Birthday:", someonesBirthday)

	//--	Formatting and Parsing Times

	layout := "2006-01-02"
	dateString := "2022-09-15"
	parsedTime, _ := time.Parse(layout, dateString)
	formattedTime := parsedTime.Format("Jan 2, 2006")
	fmt.Println("Formatted Time:", formattedTime)

	//--	Adding and Subtracting Time

	currentTime = time.Now()
	futureTime := currentTime.Add(24 * time.Hour)
	fmt.Println("Future Time:", futureTime)

	duration := futureTime.Sub(currentTime)
	fmt.Println("Duration:", duration)

	//--	Time Zones

	// In the time package in Go, time zones are handled using the Location type and the LoadLocation function.

	// Loading a location:
	loc, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		fmt.Println("Error loading location:", err)
		return
	}

	// Creating a time object with a specific location:
	t := time.Date(2022, time.January, 1, 12, 0, 0, 0, loc)
	fmt.Println("Time in Moscow:", t)

	// Converting time to a different location:
	locUTC, _ := time.LoadLocation("UTC")
	tUTC := t.In(locUTC)
	fmt.Println("Time in UTC:", tUTC)

	//--	Different Formats

	fmt.Println("Current Time in String: ", currentTime.String())

	fmt.Println("MM-DD-YYYY : ", currentTime.Format("01-02-2006"))

	fmt.Println("YYYY-MM-DD : ", currentTime.Format("2006-01-02"))

	fmt.Println("YYYY.MM.DD : ", currentTime.Format("2006.01.02 15:04:05"))

	fmt.Println("YYYY#MM#DD {Special Character} : ", currentTime.Format("2006#01#02"))

	fmt.Println("YYYY-MM-DD hh:mm:ss : ", currentTime.Format("2006-01-02 15:04:05"))

	fmt.Println("Time with MicroSeconds: ", currentTime.Format("2006-01-02 15:04:05.000000"))

	fmt.Println("Time with NanoSeconds: ", currentTime.Format("2006-01-02 15:04:05.000000000"))

	fmt.Println("ShortNum Month : ", currentTime.Format("2006-1-02"))

	fmt.Println("LongMonth : ", currentTime.Format("2006-January-02"))

	fmt.Println("ShortMonth : ", currentTime.Format("2006-Jan-02"))

	fmt.Println("ShortYear : ", currentTime.Format("06-Jan-02"))

	fmt.Println("LongWeekDay : ", currentTime.Format("2006-01-02 15:04:05 Monday"))

	fmt.Println("ShortWeek Day : ", currentTime.Format("2006-01-02 Mon"))

	fmt.Println("ShortDay : ", currentTime.Format("Mon 2006-01-2"))

	fmt.Println("Short Hour Minute Second: ", currentTime.Format("2006-01-02 3:4:5"))

	fmt.Println("Short Hour Minute Second: ", currentTime.Format("2006-01-02 3:4:5 PM"))

	fmt.Println("Short Hour Minute Second: ", currentTime.Format("2006-01-02 3:4:5 pm"))

}
