$os = $env:GOOS
$env:GOOS = "js"
$arch = $env:GOARCH
$env:GOARCH = "wasm"

go build -o main.wasm .\main.go .\stopwatch.go

$env:GOOS = $os
$env:GOARCH = $arch