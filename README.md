# vault-tree

Recurse Hashicorp Vault and return a list of Key paths

## Usage

Given a `.vault-token` file living in the $HOME directory containing a valid
token to access the vault, and an environment variable VAULT_ADDR containing
the URL to the Vault server, this small utility will recurse through the
key paths of the Vault server, optionally starting at a specified key path.

## References

- [Hashicorp Vault Kv Engine v2 API](https://www.vaultproject.io/api/secret/kv/kv-v2.html)
