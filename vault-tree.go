package main

// Hashicorp Vault Kv Engine v2 API: https://www.vaultproject.io/api/secret/kv/kv-v2.html

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/go-resty/resty/v2"
	"sigs.k8s.io/yaml"
	// "sigs.k8s.io/yaml"
	// "github.com/go-resty/resty/v2"
)

// DataKeys - structure of the response data for a Hashicorp Vault LIST object
// {"request_id":"8dcc5db1-2090-835c-a8b1-a0d3b449f2fd","lease_id":"","renewable":false,"lease_duration":0,"data":{"keys":["config/"]},"wrap_info":null,"warnings":null,"auth":null}
type DataKeys struct {
	RequestID     string `json:"request_id"`
	LeaseID       string `json:"lease_id"`
	Renewable     bool   `json:"renewable"`
	LeaseDuration int64  `json:"lease_duration"`
	Data          struct {
		Keys []string `json:"keys"`
	} `json:"data"`
	WrapInfo string `json:"wrap_info"`
	Warning  string `json:"warnings"`
	Auth     string `json:"auth"`
}

func getKeyPaths(keypath string) map[string]interface{} {
	var pathStructure map[string]interface{}
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

	var response DataKeys
	err = json.Unmarshal(resp.Body(), &response)
	if err != nil {
		fmt.Println("Error decoding json")
		fmt.Println(fmt.Sprintf("%#v\n", resp.Body()))
		os.Exit(1)
	}

	keys := response.Data.Keys
	for key := range keys {
		if len(keypath) > 0 {
			fullPath = fmt.Sprintf("%s/%s", keypath, keys[key])
		} else {
			fullPath = keys[key]
		}
		if strings.HasSuffix(fullPath, "/") {
			subKeys := getKeyPaths(strings.TrimSuffix(fullPath, "/"))
			pathStructure = CoalesceTables(pathStructure, subKeys)
		} else {
			allContainers := strings.Split(fullPath, "/")
			lastMap := make(map[string]interface{})

			for container := (len(allContainers) - 2); container >= 0; container-- {
				if container < (len(allContainers) - 2) {
					nv := make(map[string]interface{})
					nv = map[string]interface{}{
						fmt.Sprintf("%s/", allContainers[container]): lastMap,
					}
					lastMap = nv
				} else {
					lastMap = map[string]interface{}{
						fmt.Sprintf("%s/", allContainers[container]): map[string]interface{}{
							allContainers[(container + 1)]: fullPath,
						},
					}
				}
			}
			if len(allContainers) == 1 {
				lastMap = map[string]interface{}{
					allContainers[0]: fullPath,
				}
			}
			pathStructure = CoalesceTables(pathStructure, lastMap)
		}
	}
	return pathStructure
}

func treePrint(keys map[string]interface{}, option string) {

	if option == "yaml" {
		raw, err := yaml.Marshal(keys)
		if err == nil {
			str := string(raw[:])
			fmt.Println(str)
		}
	}
	if option == "json" {
		raw, err := json.Marshal(keys)
		if err == nil {
			str := string(raw[:])
			fmt.Println(str)
		}
	}
}

// this function is not used anymore since we are converting to a map[string]interface{}
// leaving it here as a learning function for people new to Go.
func indentString(spaces int, line string) string {
	return fmt.Sprintf("%s%s", strings.Repeat(" ", spaces), line)
}

func nextLevelIsPath(key string, level string) (bool, string) {
	remaining := strings.TrimPrefix(key, level)
	// fmt.Println(fmt.Sprintf("key: %s, level: %s, remaining: %s", key, level, remaining))
	if strings.Contains(remaining, "/") {
		return true, fmt.Sprintf("%s/", strings.Split(remaining, "/")[0])
	}
	return false, remaining
}

func readFile(filePath string) ([]byte, error) {
	return ioutil.ReadFile(filePath)
}
