test:
	go test -timeout 1s -json -skip /pkg/test -v ./... 2>&1 | gotestfmt

vulns-check:
	govulncheck ./...