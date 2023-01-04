package main

import (
	"fmt"
	"os"

	"github.com/cvetkovski98/zvax-qrcode/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
