$version = "0.0.1"
$os = "windows"
$arch = if ([Environment]::Is64BitOperatingSystem) { "amd64" } else { "386" }

Write-Host "Downloading and installing Gondalf $version for $os/$arch..."

$url = "https://github.com/developia-io/gondalf/releases/download/v$version/gondalf-windows-$arch.tar.gz"
Invoke-WebRequest -Uri $url -OutFile "$env:TEMP\gondalf.tar.gz"
Expand-Archive -Path "$env:TEMP\gondalf.tar.gz" -DestinationPath "$env:ProgramFiles\gondalf" -Force

Write-Host "Gondalf $version has been installed successfully!"
