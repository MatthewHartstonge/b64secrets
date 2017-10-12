# b64secrets
> An extremely simple globbing util to transform k8s Secret definition file values to base64 
> encoded secrets ready for upload via kubectl.

I love Infrastructure as Code (IaC). It drove me nuts that k8s secret values in 
secret files had to have values in base64 encoding.

For now, this little utility serves my purposes by globbing recursively with the pattern 
`./**/*.yml` writing the encoded secrets to `./**/*.base64.yml` ready for uploading to a k8s
cluster via `kubectl` commands.

## Usage
in a given directory containing `secrets-*.yml` files, simply run:

```sh
$ b64secrets
```

## Disclaimer
Use at your own risk!