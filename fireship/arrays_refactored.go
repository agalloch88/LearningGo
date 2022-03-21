package main

import "fmt"

func main5() {
	DeploymentOptions := [...]string{"R-pi", "AWS", "GCP", "Azure"}

	for, index, option := range DeploymentOptions {
		fmt.Println(index, option)
	}
}