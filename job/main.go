package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("job completed........")

	// important to return exit code 0, if not kubernetes treats it as failure
	os.Exit(0)
}
