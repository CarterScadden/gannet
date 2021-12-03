package main

import "fmt"

func main() {
	fmt.Printf("Hello world! %s\n", GetName()+" is great!")
}

// https://gist.github.com/barberta/2b9253a8f92ae475fa322ad725f67fe1

func GetName() string {
	return "Gannet"
}
