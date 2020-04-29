package main

import "fmt"

func main() {
	status, ok := isName("dinda")
	if !ok {
		fmt.Println(status)
	} else {
		fmt.Println(status)
	}
}

func isName(name string) (string, bool) {
	arr := []string{"bagas", "dinda", "faisal", "alvin"}

	for i := 0; i < len(arr); i++ {
		if arr[i] == name {
			return "found", true
		}
	}
	return "not found", false
}
