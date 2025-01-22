package main;

import (
	"fmt"
	"os"
)

func run() string {
	return "Hello World";
}

func main(){
	fmt.Println( run() )
	os.Exit(0)
}
