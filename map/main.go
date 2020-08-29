package main

import "fmt"

func main() {

	colors := map[string]string{
		"0": "Red",
		"1": "Blue",
	}

	printmap(colors)
}

func printmap(c map[string]string) {
	for key, value := range c {
		fmt.Println("key : ", key+"value:", value)
	}

}
