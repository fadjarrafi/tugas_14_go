package main

import (
	"fmt"
	"os/exec"
)

func main() {
	var output1, _ = exec.Command("help").Output()
	fmt.Println("help", string(output1))

	var output2, _ = exec.Command("ipconfig").Output()
	fmt.Println("ipconfig", string(output2))

	var output3, _ = exec.Command("git", "config", "user.name").Output()
	fmt.Printf(" -> git config user.name\n%s\n", string(output3))
}
