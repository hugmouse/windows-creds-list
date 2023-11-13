package main

import (
	"fmt"
	"github.com/danieljoos/wincred"
)

const (
	redColor    = "\033[31m"
	greenColor  = "\033[32m"
	yellowColor = "\033[33m"
	resetColor  = "\033[0m"
)

func main() {
	creds, err := wincred.List()
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := range creds {
		fmt.Printf("%s%s%s: %s\n", yellowColor, creds[i].TargetName, resetColor, greenColor+string(creds[i].CredentialBlob)+resetColor)
	}
}
