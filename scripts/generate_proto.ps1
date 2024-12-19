# PowerShell script to generate Go files from proto using namely/protoc-all

param (
    [string]$protoFile = "parser.proto",
    [string]$protoPath = "./proto",
    [string]$outputPath = "./proto/generated"
)

# Ensure Docker is installed
if (-Not (Get-Command docker -ErrorAction SilentlyContinue)) {
    Write-Error "Docker is not installed. Please install Docker and try again."
    exit 1
}

# Ensure output directory exists
if (-Not (Test-Path $outputPath)) {
    New-Item -ItemType Directory -Path $outputPath
}

# Run Docker command
try {
    docker run --rm -v "$(Get-Location)/${protoPath}:/protos" -v "$(Get-Location)/${outputPath}:/out" namely/protoc-all:1.44_1 `
        -f "${protoFile}" `
        -l go `
        -o /out `
        --with-grpc
         
    Write-Host "Protobuf files generated successfully in $outputPath"
} catch {
    Write-Error "An error occurred while generating protobuf files: $_"
    exit 1
}