package main

import (
	"encoding/json"
	"errors"
	"os"
	"os/exec"
)

type CurlRequest struct {
	Name   string `json:"name"`
	URL    string `json:"url"`
	Method string `json:"method"`
	Data   string `json:"data"`
}

func unwrap(e error) {
	if e != nil {
		panic(e)
	}
}

func getCurlRequest(crs []CurlRequest, name string) (CurlRequest, error) {
	for _, cr := range crs {
		if cr.Name == name {
			return cr, nil
		}
	}
	return CurlRequest{}, errors.New("curl request not found")
}

func main() {
	args := os.Args[1:]
	dat, err := os.ReadFile("./simpl-curl.json")
	unwrap(err)

	var curlRequests []CurlRequest
	err = json.Unmarshal(dat, &curlRequests)
	unwrap(err)

	curlRequest, err := getCurlRequest(curlRequests, args[0])

	curlCmd := exec.Command("curl")
	if curlRequest.Method != "" {
		curlCmd.Args = append(curlCmd.Args, "-X")
		curlCmd.Args = append(curlCmd.Args, curlRequest.Method)
	}
	if curlRequest.Data != "" {
		curlCmd.Args = append(curlCmd.Args, "-d")
		curlCmd.Args = append(curlCmd.Args, curlRequest.Data)
	}
	curlCmd.Args = append(curlCmd.Args, curlRequest.URL)
	curlCmd.Stdin = os.Stdin
	curlCmd.Stdout = os.Stdout
	curlCmd.Stderr = os.Stderr

	_ = curlCmd.Run()
}
