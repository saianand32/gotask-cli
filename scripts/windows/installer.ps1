<#
.SYNOPSIS
gotask-cli Windows Installer

.DESCRIPTION
Installs, uninstalls, or updates gotask-cli for Windows.
#>

param (
    [switch]$install,
    [switch]$uninstall,
    [switch]$update,
    [switch]$help
)

$ScriptDir = Split-Path -Parent $MyInvocation.MyCommand.Path
$ProjectRoot = (Resolve-Path "$ScriptDir\..\..").Path
$InstallDir = "$env:LOCALAPPDATA\gotask"
$DestBin = Join-Path $InstallDir "gotask.exe"

function Show-Help {
    Write-Host "==============================="
    Write-Host "   gotask-cli Windows Installer  "
    Write-Host "==============================="
    Write-Host "Usage: .\installer.ps1 [-install] [-uninstall] [-update]"
    Write-Host ""
    Write-Host "Commands:"
    Write-Host "  -install    Install gotask-cli"
    Write-Host "  -uninstall  Uninstall gotask-cli"
    Write-Host "  -update     Update gotask-cli to the latest built version"
    Write-Host "  -help       Show this help message"
}

function Build-And-Install([string]$action) {
    Write-Host "Changing to project root: $ProjectRoot"
    Set-Location -Path $ProjectRoot

    Write-Host "Building the Windows binary..."
    
    # Try using make first
    try {
        make build-win
        if ($LASTEXITCODE -ne 0) { throw "Make failed." }
    } catch {
        Write-Host "Make not available or failed. Falling back to native go build..."
        $env:GOOS="windows"
        $env:GOARCH="amd64"
        if (-not (Test-Path "build\bin")) { New-Item -ItemType Directory -Force -Path "build\bin" | Out-Null }
        go build -o build\bin\gotask.exe .\cmd\app\main.go
        if ($LASTEXITCODE -ne 0) {
            Write-Host "Error: Build failed."
            exit 1
        }
    }

    $BinaryPath = "build\bin\gotask.exe"

    if (-not (Test-Path $BinaryPath)) {
        Write-Host "Error: Binary not found at $BinaryPath"
        exit 1
    }

    if ($action -eq "install") {
        Write-Host "Installing gotask-cli to $InstallDir..."
    } else {
        Write-Host "Updating gotask-cli in $InstallDir..."
    }

    if (-not (Test-Path $InstallDir)) {
        New-Item -ItemType Directory -Force -Path $InstallDir | Out-Null
    }

    Copy-Item -Force -Path $BinaryPath -Destination $DestBin

    # Add to User PATH if not already there
    $UserPath = [Environment]::GetEnvironmentVariable("PATH", "User")
    if ($UserPath -notmatch [regex]::Escape($InstallDir)) {
        [Environment]::SetEnvironmentVariable("PATH", "$UserPath;$InstallDir", "User")
        Write-Host "Added $InstallDir to your PATH. You may need to restart your terminal."
    }

    Write-Host "==============================="
    if ($action -eq "install") {
        Write-Host "Installation completed successfully!"
    } else {
        Write-Host "Update completed successfully!"
    }
    Write-Host "You can now run 'gotask' from your terminal."
    Write-Host "Try running: gotask help"
    Write-Host "==============================="
}

function Uninstall-Gotask {
    if (-not (Test-Path $DestBin)) {
        Write-Host "gotask-cli is not installed at $DestBin."
        exit 0
    }

    Write-Host "Uninstalling gotask-cli from $InstallDir..."
    Remove-Item -Force $DestBin

    $UserPath = [Environment]::GetEnvironmentVariable("PATH", "User")
    if ($UserPath -match [regex]::Escape(";$InstallDir")) {
        $NewPath = $UserPath -replace [regex]::Escape(";$InstallDir"), ""
        [Environment]::SetEnvironmentVariable("PATH", $NewPath, "User")
    } elseif ($UserPath -match [regex]::Escape("$InstallDir;")) {
        $NewPath = $UserPath -replace [regex]::Escape("$InstallDir;"), ""
        [Environment]::SetEnvironmentVariable("PATH", $NewPath, "User")
    }

    Write-Host "==============================="
    Write-Host "Uninstallation completed successfully!"
    Write-Host "==============================="
}

if ($help -or (-not $install -and -not $uninstall -and -not $update)) {
    Show-Help
} elseif ($install) {
    Build-And-Install "install"
} elseif ($uninstall) {
    Uninstall-Gotask
} elseif ($update) {
    Build-And-Install "update"
}
