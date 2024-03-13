package main

import (
	"fmt"
	"os/exec"
)

func main() {
	fmt.Println("Hello world!")

	// option 1
	//c := make(chan string)
	//go runApp(c)
	//fmt.Println(<-c)

	// option 2
	runApp2()

	fmt.Println("Good bye!")
}

func runApp(c chan string) {
	cmd := exec.Command("C:\\dev\\projects\\go\\fynehw\\fynehw.exe")
	fmt.Println("Running app and waiting for it to complete")
	err := cmd.Run()
	if err != nil {
		fmt.Println(err.Error())
	}
	c <- "Done running app"
}

func runApp2() {
	cmd := exec.Command("C:\\dev\\projects\\go\\fynehw\\fynehw.exe")
	err := cmd.Run()
	if err != nil {
		fmt.Println(err.Error())
	}
}
