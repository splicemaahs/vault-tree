# TODO List

## main

### TODO: Add "command" processing

[ ] LIST - List all the KEYS at a KeyPath
[ ] LISTALL - List all the KEYS starting at a KeyPath
[ ] GET - Get the Value of a unique Key
[ ] GETALL - Get the Values of all the unique Keys starting at a KeyPath

#### Examples

- vault-tree list         # lists keys contained in root
- vault-tree list azure   # lists keys contained at azure/
- vault-tree listall      # lists all keys starting at root
- vault-tree get          # return the value for a specific key
- vault-tree getall       # return the value for all keys starting at root

### TODO: Add --options

[x] --output=json|yaml - Return the data in a specific format.

- implemented as '-o={yaml|json}', defaults to json

[ ] --withvalues - Replace the "fullPath" with the actual value from Vault.
