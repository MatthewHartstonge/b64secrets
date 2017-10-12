# Enables simple cross-compilation and packaging of Go applications via Windows.
#
# Author: Matthew Hartstonge
# Written: 2017-10-12


# Creates binaries packaged and ready to go in a distibution folder
function main() {
  $operatingSystems = ("windows", "darwin", "linux", "openbsd", "freebsd", "solaris")
  $architechtures = ("amd64")
  $packagePath = "github.com\MatthewHartstonge\b64secrets"
  $appName = "b64secrets"
  $distPath = "./dist"
  
  # Obtain the version from the user. A.K.A: "don't piss around trying to understand what you think the user wants via Git, just ask..."
  $version = Read-Host -Prompt "Enter release version: "

  # Create dist directory, assuming it doesn't already exist
  $distPathExists = Test-Path $distPath
  if (!$distPathExists) {
    New-Item -Path "./dist" -ItemType "directory"
  }

  foreach ($os in $operatingSystems) {
    foreach ($arch in $architechtures) {
      # Build that fabulous idea
      $binaryFp = buildGoApp -OS $os -Arch $arch -AppName $appName -PackagePath $packagePath -DestinationPath $distPath
      
      # zip that sucker
      $zipName = "$appname-$version-$os-$arch"
      zipFile -FilePath $binaryFp -Destination $distPath -FileName $zipName
      
      # Make sure to clean up after yourself
      Remove-Item -Path $binaryFp
    }
  } 
}

# buildGoApp creates a go binary in DestinationPath
function buildGoApp($OS, $Arch, $AppName, $PackagePath, $DestinationPath) {
  $env:GOOS = $OS
  $env:GOARCH = $Arch
  
  $out = $DestinationPath + "/" + $AppName
  if ($os -eq "windows") {
    $out = $out + ".exe"
  } 
  go build -o $out $PackagePath

  return $out
}

# zipFile will create a zip of a given file, and give it the provided name.
function zipFile($FilePath, $Destination, $FileName) {
  Compress-Archive -Path $FilePath -CompressionLevel "Optimal" -DestinationPath "$Destination/$FileName.zip"
}

main
