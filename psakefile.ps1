properties {
  [Diagnostics.CodeAnalysis.SuppressMessageAttribute("PSUseDeclaredVarsMoreThanAssignments", "current_directory", Justification = "Variable is used in multiple tasks.")]
  $current_directory = Get-Location
}

task default -Depends list

task list {
  Get-PSakeScriptTasks
} -Description "List all tasks in the script"

task go-deps-ls {
  go list -u -m -f '{{if not .Indirect}}{{.}}{{end}}' all
} -Description "List all direct dependencies in go.mod with their versions"

task test {
  Write-Host "Modules..."
  go mod download
  go mod tidy
  if ($LastExitCode -ne 0) { throw "Failed to download modules" }
  Write-Host "Formatting..."
  go fmt ./...
  if ($LastExitCode -ne 0) { throw "Failed to format code" }
  Write-Host "Vetting..."
  go vet ./...
  if ($LastExitCode -ne 0) { throw "Failed to vet code" }
  Write-Host "Linting..."
  golangci-lint run ./...
  if ($LastExitCode -ne 0) { throw "Failed to lint code" }
  Write-Host "Testing..."
  gotestsum --format testdox -- -v ./...
  if ($LastExitCode -ne 0) { throw "Failed to test code" }
} -Description "Run all tests"

task build {
  go build -ldflags "-s -w" -o ./bin/spawnforge.exe .
  if ($LastExitCode -ne 0) { throw "Failed to build binaries" }
} -Description "Build all binaries"