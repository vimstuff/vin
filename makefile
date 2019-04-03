FILES = vin

help:
	@echo "You can perform the following:"
	@echo ""
	@echo "  check  Format, lint, vet, and test Go code"
	@echo "  deploy Copy to ~/bin"
	@echo "  local  Build for local OS"
	@echo "  win    Build for Windows"

check:
	@echo 'Formatting, linting, vetting, and testing Go code.'
	go fmt ./...
	golint ./...
	go vet ./...
	go test ./...

win:
	env GOOS=windows GOARCH=amd64 go build -o dist/vin.exe

local:
	env go build -o dist/vin

deploy: local
	cp dist/vin ~/bin
