package main

import (
	"fmt"
	"os"
	"strings"
)

var home string
var vaultToken []byte
var vaultAddr string

func main() {

	var err error
	var startPath string
	var outputOption string

	home = os.Getenv("HOME")
	vaultToken, err = readFile(fmt.Sprintf("%s/.vault-token", home))
	vaultAddr = os.Getenv("VAULT_ADDR")
	if err != nil {
		fmt.Println("Could not read the vault-token")
	}

	// TODO: Add "command" processing
	// TODO: Add --options

	startPath = ""
	outputOption = "json"
	for arg := 1; arg < len(os.Args); arg++ {
		if strings.HasPrefix(os.Args[arg], "-o=") {
			outputOption = strings.ToLower(strings.TrimPrefix(os.Args[arg], "-o="))
		}
		if arg > 0 && !strings.HasPrefix(os.Args[arg], "-") {
			startPath = os.Args[arg]
		}
	}

	allkeys := getKeyPaths(startPath)

	treePrint(allkeys, outputOption)

}
