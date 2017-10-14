# b64secrets changelog
## v0.2.2
- Stops generating secret files if it contains `.base64.yml` in the filepath.
- Now only logs the files that have actually been generated

## v0.2.1
- Fixes logging false negative write errors
- Fixes creating b64 encoded secret files for secret types other than 'Opaque'
- Updates readme to be more useful

## v0.2.0
- Added logging
- Changed to use a recursive glob from `.`
- Created build script to ease shipping new releases

## v0.1.0
- It's alive!