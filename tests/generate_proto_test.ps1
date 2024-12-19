# Test script for generate_proto.ps1

$scriptPath = "../scripts/generate_proto.ps1"

Write-Host "Running tests for generate_proto.ps1..."

# Test 1: Verify script execution with default parameters
try {
    . $scriptPath
    Write-Host "Test 1 passed: Script executed successfully with default parameters."
} catch {
    Write-Error "Test 1 failed: Script failed with default parameters."
    exit 1
}

# Test 2: Verify custom parameters
$customOutput = "./tests/generated"
try {
    . $scriptPath -protoFile "parser.proto" -protoPath "./proto" -outputPath $customOutput
    Write-Host "Test 2 passed: Script executed successfully with custom parameters."
    if (Test-Path $customOutput) {
        Write-Host "Custom output directory verified: $customOutput"
        Remove-Item -Recurse -Force $customOutput
    } else {
        Write-Error "Test 2 failed: Custom output directory not found."
        exit 1
    }
} catch {
    Write-Error "Test 2 failed: Script failed with custom parameters."
    exit 1
}

Write-Host "All tests for generate_proto.ps1 passed."
