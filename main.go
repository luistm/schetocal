package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// from https://brandur.org/fragments/go-days-in-month
func daysIn(m time.Month, year int) int {
	return time.Date(year, m+1, 0, 0, 0, 0, 0, time.UTC).Day()
}

func scheduleToCall(rawSchedule string) {
	fmt.Println("Building calendar files for the following schedule")
	fmt.Println(">> ", rawSchedule)

	dailyShifts := strings.Split(rawSchedule, ",")

	// Check if the number of shifts matches the number of days in the month
	_, month, year := time.Now().Date()
	daysInMonth := daysIn(month, year)
	if len(dailyShifts) != daysInMonth {
		fmt.Println("Number of shifts is not equal to the number of days in the current month")
		fmt.Printf("This month has %d days!\n", daysInMonth)
		os.Exit(0)
	}

	// Create a calendar file for each shift
	for _, v := range dailyShifts {
		fmt.Printf("%s\n", v)
	}

}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Invalid number of arguments: ", os.Args[1:])
		fmt.Println("Exiting!")
		os.Exit(0)
	}
	scheduleToCall(os.Args[1])
}
