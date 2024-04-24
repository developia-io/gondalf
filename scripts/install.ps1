$version = "0.0.1"
$os = "windows"
$arch = if ([Environment]::Is64BitOperatingSystem) { "amd64" } else { "386" }

Write-Host "Downloading and installing Gondalf $version for $os/$arch..."

$url = "https://github.com/developia-io/gondalf/releases/download/v$version/gondalf-windows-$arch.zip"
Invoke-WebRequest -Uri $url -OutFile "$env:TEMP\gondalf.zip"
Expand-Archive -Path "$env:TEMP\gondalf.zip" -DestinationPath "$env:ProgramFiles\gondalf" -Force

$gondalfPath = Join-Path -Path "$env:ProgramFiles\gondalf" -ChildPath "gondalf.exe"
$existingPath = [Environment]::GetEnvironmentVariable("Path", "Machine")
if ($existingPath -notlike "*$gondalfPath*") {
    $newPath = "$existingPath;$gondalfPath"
    [Environment]::SetEnvironmentVariable("Path", $newPath, "Machine")
}

Write-Host "Gondalf $version has been installed successfully!"
