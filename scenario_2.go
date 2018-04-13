package main

import (
	"fmt"
	"time"
)

/*
Make sure that Jenkins is running. (i.e., If it is running, we are good. If not unidle it)
Wait for 45 minutes and check if it is idled.
*/
func scenario2() {

	fmt.Println(IsIdle())
	if IsIdle() {
		UnIdle()
		// Wait for 5 min.
		// Within this time Jenkins should be unidled
		fmt.Println("Waiting for 5 minutes")
		time.Sleep(5 * time.Minute)
	} else {
		fmt.Println("Jenkins is already unidle")
	}
	if IsIdle() {
		fmt.Println("Error: Did not unidle Jenkins in 5 minutes")
	}
	// Wait for 45 Min
	// Within this time Jenkins should be idled
	time.Sleep(45 * time.Minute)
	fmt.Println("Waiting for 45 minutes.....")
	if IsIdle() {
		fmt.Println("Waited for 45 minutes. And now Jenkins is idle")
	} else {
		fmt.Println("Waited for 45 minutes, but Jenkins is still unidle")
	}
}
