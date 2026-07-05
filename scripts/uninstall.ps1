$path = Join-Path $env:USERPROFILE ".local\bin\leaserage.exe"
if (Test-Path -LiteralPath $path) {
  Remove-Item -LiteralPath $path -Force
}
Write-Host "removed $path"
