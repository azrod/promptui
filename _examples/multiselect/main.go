package main

import (
	"fmt"

	"github.com/azrod/promptui"
)

func main() {
	prompt := promptui.MultiSelect{
		Label: "Select Day",
		Items: []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday",
			"Saturday", "Sunday"},
	}

	i, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	fmt.Printf("You choose %v\n", i)
}
