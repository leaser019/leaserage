param(
  [string]$Version = "latest",
  [string]$Repo = $(if ($env:LEASERAGE_REPO) { $env:LEASERAGE_REPO } else { "vomkhang/leaserage" }),
  [Parameter(ValueFromRemainingArguments = $true)]
  [string[]]$CliArgs = @()
)

$ErrorActionPreference = "Stop"
$arch = if ([System.Runtime.InteropServices.RuntimeInformation]::OSArchitecture -eq "Arm64") { "arm64" } else { "amd64" }
$name = "leaserage-windows-$arch.zip"
$url = "https://github.com/$Repo/releases/$Version/download/$name"
$tmp = Join-Path ([System.IO.Path]::GetTempPath()) ("leaserage-" + [guid]::NewGuid())
$bin = Join-Path $env:USERPROFILE ".local\bin"

New-Item -ItemType Directory -Path $tmp -Force | Out-Null
New-Item -ItemType Directory -Path $bin -Force | Out-Null
Invoke-WebRequest -Uri $url -OutFile (Join-Path $tmp $name)
Expand-Archive -LiteralPath (Join-Path $tmp $name) -DestinationPath $tmp -Force
Copy-Item -LiteralPath (Join-Path $tmp "leaserage.exe") -Destination (Join-Path $bin "leaserage.exe") -Force
& (Join-Path $bin "leaserage.exe") @CliArgs
