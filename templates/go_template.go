/*
Template for comparing the speed of two functions - AI Time Checker
Repo Link: https://github.com/eric-hamilton/ai-time-checker
Created by Eric Hamilton
GitHub: https://github.com/eric-hamilton
*/

package main

import (
	"fmt"
	"time"
)

// User-defined function
func myFunction() {
	// Add your own code here
}

// AI-generated function
func aiFunction() {
	// Add AI code here
}

const numRuns = 3 // Number of times to run each function

func main() {
	var myTimes []time.Duration
	var aiTimes []time.Duration

	var myTotalTime time.Duration
	var aiTotalTime time.Duration

	for i := 0; i < numRuns; i++ {
		start := time.Now()
		myFunction() // Add any required arguments here
		myTime := time.Since(start)

		start = time.Now()
		aiFunction() // Add any required arguments here
		aiTime := time.Since(start)

		myTotalTime += myTime
		aiTotalTime += aiTime
		myTimes = append(myTimes, myTime)
		aiTimes = append(aiTimes, aiTime)
	}

	myAvgTime := float64(myTotalTime) / float64(numRuns*time.Millisecond)
	aiAvgTime := float64(aiTotalTime) / float64(numRuns*time.Millisecond)

	fmt.Printf("My average time: %.10f milliseconds\n", myAvgTime)
	fmt.Printf("AI average time: %.10f milliseconds\n", aiAvgTime)
	fmt.Println()

	fmt.Println("My execution times:")
	for i, myTime := range myTimes {
		fmt.Printf("Run %d: %.10f milliseconds\n", i+1, float64(myTime)/float64(time.Millisecond))
	}

	fmt.Println("AI execution times:")
	for i, aiTime := range aiTimes {
		fmt.Printf("Run %d: %.10f milliseconds\n", i+1, float64(aiTime)/float64(time.Millisecond))
	}

	if myTotalTime > aiTotalTime {
		timeFactor := float64(myTotalTime) / float64(aiTotalTime)
		fmt.Printf("AI is, on average, %.10f times faster.\n", timeFactor)
	} else {
		timeFactor := float64(aiTotalTime) / float64(myTotalTime)
		fmt.Printf("Mine is, on average, %.10f times faster.\n", timeFactor)
	}

	fmt.Println("Press Enter to exit")
	fmt.Scanln()
}
