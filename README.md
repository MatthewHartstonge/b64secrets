# b64secrets
> An extremely simple globbing util to transform k8s Secret definition file values to base64 
> encoded secrets ready for upload via kubectl.

I love Infrastructure as Code (IaC). It drove me nuts that k8s secret values in 
secret files had to have values in base64 encoding.

This little utility globs recursively with the pattern `./**/*.yml` writing the encoded secrets 
to `./**/*.base64.yml` ready for uploading to a k8s cluster via applicable `kubectl` commands.

## Usage
in a given directory containing `*.yml` files, simply run:

```sh
$ b64secrets
INFO[0000] b64secrets is converting secrets..            method=main
INFO[0000] Created conformed secrets file                conformedPath="super-secrets-development.base64.yml" method=createSecretsFile originalPath="config\\super-secrets-development.yml"
INFO[0000] Created conformed secrets file                conformedPath="/secrets/within/a/folder/finds-love.base64.yml" method=createSecretsFile originalPath="secrets/within/a/folder/finds-love.yml"
INFO[0000] b64secrets file conversions completed!        method=main
```
