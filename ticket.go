package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {

	cias := [3]string{"Space Adventures", "SpaceX", "Virgin Galactic"}

	fmt.Println("Spaceline          Days Trip type  Price")
	fmt.Println("========================================")

	randNumCycles := randNewInt(1, 10)
	for i := randNumCycles; i >= 0; i-- {

		spaceCia := getSpaceCompany(cias)
		totalDays := getTotalDaysOfTrip(getSpeedOfShip())
		isRoundTrip := isRoundTrip()
		price := getPriceOfTrip(totalDays)
		if isRoundTrip {
			price = price * 2
		}

		fmt.Printf("%s    %d  %v  %d$\n", spaceCia, totalDays, isRoundTrip, price)
	}

}

func getSpeedOfShip() int {
	//Randomly choose the speed the ship will travel, from 16 to 30 km/s
	return randNewInt(16, 31)
}

func getDepartureDate() string {
	return "October 13, 2020"
}

func getSpaceCompany(companies [3]string) string {
	return companies[randNewInt(0, 2)]
}

func randNewInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func isRoundTrip() bool {
	return randNewInt(0, 2) > 0
}

func getTotalDaysOfTrip(speed int) int {
	return 621000 / speed
}

func getPriceOfTrip(totalDays int) int {
	return totalDays * 5
}
