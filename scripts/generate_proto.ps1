# PowerShell script to generate Go files from proto using namely/protoc-all

$Current_Path = Get-Location
$Proto_Files = "${Current_Path}\proto"
$Proto_Files_Docker = $Proto_Files -replace "\\", "/"
$Generated_Files = "${Proto_Files}\generated"

# Ensure output directory exists
if (-Not (Test-Path -Path $Generated_Files)) {
    New-Item -ItemType Directory -Path $Generated_Files
}

# Accepted languages
$Languages = @("golang")
$Language_Abr = @("go")

# Prompt for language
$Language = "golang"  # Default to Go for automation
if ($Languages -contains $Language) {
    Write-Host "Generating code for $Language"
    if (-Not (Test-Path -Path $Proto_Files)) {
        Write-Host "Proto files not found"
        exit
    }

    # Run Docker command
    docker run --rm -v "${Proto_Files_Docker}:/defs" -v "${Generated_Files}:/out" namely/protoc-all -d /defs -o /out -l $Language_Abr --with-grpc

    Write-Host "Code generated successfully in $Generated_Files"

    # Install dependencies if needed
    $Ok = Read-Host "Do you want to install necessary dependencies? (y/n)"
    if ($Ok -eq "y") {
        Write-Host "Installing missed dependencies"
        go get -u google.golang.org/grpc
        go get -u google.golang.org/protobuf/...
        Write-Host "Dependencies installed successfully"
    } else {
        Write-Host "Dependencies not installed"
    }
} else {
    Write-Host "Language not supported"
    exit
}
