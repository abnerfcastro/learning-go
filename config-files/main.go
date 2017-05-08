package main

import (
	"flag"
	"fmt"

	"github.com/vharitonsky/iniflags"
)

// go run main.go -config="dev.ini"
func main() {

	// flags to be read from dev.ini
	greetingPtr := flag.String("greeting", "Hello", "This is how the user will be greeted.")
	treatmentPtr := flag.String("treatment", "Mr.", "This is the title the user will be addressed with.")
	namePtr := flag.String("name", "Abner Castro", "This is the name of the user")

	iniflags.Parse()

	fmt.Println(*greetingPtr, *treatmentPtr, *namePtr)

}
