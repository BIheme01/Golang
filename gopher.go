// Playing with fmt and time using Go

package main

import (
	"fmt"
	t "time"
)

func main() {

	dateTime := t.Now()

	fmt.Println("\nThis is my gopher:")
	fmt.Println("\n       `.-::::::::-.`")
	fmt.Println("   .:-::::::::::::::::-:.")
	fmt.Println("   `_:::     ::     :::_`")
	fmt.Println("    .:(  ^   ::  ^   ):.")
	fmt.Println("    `:::    (..)    :::.")
	fmt.Println("    `::::::::UU::::::::`")
	fmt.Println("    .::::::::::::::::::.")
	fmt.Println("    0::::::::::::::::::0")
	fmt.Println("    -::::::::::::::::::-")
	fmt.Println("    `::::::::::::::::::`")
	fmt.Println("     .::::::::::::::::.")
	fmt.Println("       oO:::::::::Oo")
	fmt.Println("\n This ASCII art was created on:", dateTime.Format("Mon 01-02-2006 15:04:05"))

}
