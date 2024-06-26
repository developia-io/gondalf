$version = "0.0.2"
$os = "windows"
$arch = if ([Environment]::Is64BitOperatingSystem) { "amd64" } else { "386" }

Write-Host "Downloading and installing gondalf $version for $os/$arch..."

$url = "https://github.com/developia-io/gondalf/releases/download/v$version/gondalf-windows-$arch.zip"
Invoke-WebRequest -Uri $url -OutFile "$env:TEMP\gondalf.zip"
Expand-Archive -Path "$env:TEMP\gondalf.zip" -DestinationPath "$env:ProgramFiles\gondalf" -Force

$gondalfDirectory = "$env:ProgramFiles\gondalf"
$existingPath = [Environment]::GetEnvironmentVariable("Path", "Machine")
if ($existingPath -notlike "*$gondalfDirectory*") {
    $newPath = "$existingPath;$gondalfDirectory"
    [Environment]::SetEnvironmentVariable("Path", $newPath, "Machine")
}

Write-Host "gondalf $version has been installed successfully!"
