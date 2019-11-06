package main

import (
	"fmt"
	"os"
)

var home string
var vaultToken []byte
var vaultAddr string

func main() {

	var err error
	var startPath string

	home = os.Getenv("HOME")
	vaultToken, err = readFile(fmt.Sprintf("%s/.vault-token", home))
	vaultAddr = os.Getenv("VAULT_ADDR")
	if err != nil {
		fmt.Println("Could not read the vault-token")
	}

	// TODO: Add "command" processing
	// [ ] LIST - List all the KEYS at a KeyPath
	// [ ] LISTALL - List all the KEYS starting at a KeyPath
	// [ ] GET - Get the Value of a unique Key
	// [ ] GETALL - Get the Values of all the unique Keys starting at a KeyPath
	// Examples:
	//  - vault-tree list         # lists keys contained in root
	//  - vault-tree list azure   # lists keys contained at azure/
	//  - vault-tree listall      # lists all keys starting at root
	//  - vault-tree get          # return the value for a specific key
	//  - vault-tree getall       # return the value for all keys starting at root
	// TODO: Add --options
	// [ ] --output=json|yaml - Return the data in a specific format.

	if len(os.Args) > 1 {
		startPath = os.Args[1]
	} else {
		startPath = ""
	}
	allkeys := getKeyPaths(startPath)

	for key := range allkeys {
		fmt.Println(allkeys[key])
	}
}
