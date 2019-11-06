package main

// Hashicorp Vault Kv Engine v2 API: https://www.vaultproject.io/api/secret/kv/kv-v2.html

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/go-resty/resty/v2"
)

func getKeyPaths(keypath string) []string {
	var allKeys []string
	var listPath string
	var fullPath string
	client := resty.New()

	// Setup our Keypath
	if len(keypath) > 0 {
		listPath = fmt.Sprintf("secret/metadata/%s", keypath)
	} else {
		listPath = "secret/metadata/"
	}

	resp, err := client.R().
		SetHeader("X-Vault-Token", string(vaultToken[:])).
		Execute("LIST", fmt.Sprintf("%s/v1/%s", vaultAddr, listPath))

	var response interface{}
	err = json.Unmarshal(resp.Body(), &response)
	if err != nil {
		fmt.Println("Error decoding json")
		fmt.Println(fmt.Sprintf("%#v\n", resp))
		os.Exit(1)
	}
	keys := response.(map[string]interface{})["data"].(map[string]interface{})["keys"].([]interface{})
	for key := range keys {
		if len(keypath) > 0 {
			fullPath = fmt.Sprintf("%s/%s", keypath, keys[key].(string))
		} else {
			fullPath = keys[key].(string)
		}
		if strings.HasSuffix(fullPath, "/") {
			subKeys := getKeyPaths(strings.TrimSuffix(fullPath, "/"))
			for skey := range subKeys {
				allKeys = append(allKeys, subKeys[skey])
			}
		} else {
			allKeys = append(allKeys, fullPath)
		}
	}
	return allKeys
}
func prettyPrint(keys []string) {

}

func jsonPrint(keys []string) {

}

func readFile(filePath string) ([]byte, error) {
	return ioutil.ReadFile(filePath)
}
