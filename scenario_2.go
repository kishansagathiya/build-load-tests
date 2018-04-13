package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

/*
Make sure that Jenkins is running. (i.e., If it is running, we are good. If not unidle it)
Wait for 45 minutes and check if it is idled.
*/
func main() {

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

// Status :
type Status struct {
	IsIdle bool `json:"is_idle"`
}

func errorHandler(err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
}

func IsIdle() bool {
	url := "http://localhost:8080/api/idler/isidle/ksagathi-preview-jenkins?openshift_api_url=https%3A%2F%2Fapi.starter-us-east-2a.openshift.com%2F"

	req, err := http.NewRequest("GET", url, nil)
	errorHandler(err)

	req.Header.Add("cache-control", "no-cache")

	res, err := http.DefaultClient.Do(req)
	errorHandler(err)

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	errorHandler(err)

	var status Status
	err = json.Unmarshal(body, &status)
	errorHandler(err)

	return status.IsIdle
}

func UnIdle() {
	url := "http://localhost:8080/api/idler/unidle/ksagathi-preview-jenkins?openshift_api_url=https%3A%2F%2Fapi.starter-us-east-2a.openshift.com%2F"

	req, err := http.NewRequest("GET", url, nil)
	errorHandler(err)

	req.Header.Add("cache-control", "no-cache")

	res, err := http.DefaultClient.Do(req)
	errorHandler(err)

	defer res.Body.Close()

	fmt.Println(res.Status)

	return
}

func Idle() {
	url := "http://localhost:8080/api/idler/idle/ksagathi-preview-jenkins?openshift_api_url=https%3A%2F%2Fapi.starter-us-east-2a.openshift.com%2F"

	req, err := http.NewRequest("GET", url, nil)
	errorHandler(err)

	req.Header.Add("cache-control", "no-cache")

	res, err := http.DefaultClient.Do(req)
	errorHandler(err)

	defer res.Body.Close()

	fmt.Print(res.Status)

	return
}
