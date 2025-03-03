EXECUTABLE=quic-config
GO111MODULE=on

WINDOWS	= $(EXECUTABLE)-windows-amd64
LINUX		= $(EXECUTABLE)-linux-amd64
DARWIN	= $(EXECUTABLE)-darwin-amd64

VERSION = $(if $(DRONE_TAG),$(DRONE_TAG),dirty)

ifeq ($(OS), Windows_NT)
	PLATFORM=windows
else
	OS := $(shell uname)
	ifeq ($(OS), Darwin)
		PLATFORM=darwin
	else
		PLATFORM=linux
	endif
endif

MODE ?= release
ifeq ($(MODE), debug)
	MODE = debug
else
	MODE = release
endif

BUILD_FLAGS = -ldflags="-X main.version=$(VERSION) -X main.mode=$(MODE) -X main.githubToken=$(GITHUB_TOKEN)"

.PHONY: all clean web $(EXECUTABLE)-$(PLATFORM)-amd64
all: windows linux darwin
	@echo version: $(VERSION)

install: pkg/statik
	go install $(BUILD_FLAGS) ./cmd/quic-config

windows: $(WINDOWS).zip
$(WINDOWS).zip: $(WINDOWS).exe
	@zip -r $(WINDOWS).zip $(WINDOWS).exe

$(WINDOWS).exe: cmd/quic-config/rsrc.syso pkg/statik
	env CGO_ENABLED=1 GOOS=windows GOARCH=amd64 go build $(BUILD_FLAGS) -o $(WINDOWS).exe ./cmd/quic-config

cmd/quic-config/rsrc.syso: web/public/favicon.ico
	go get github.com/akavel/rsrc
	rsrc -arch amd64 -ico web/public/favicon.ico -manifest script/quic-config-windows-amd64.exe.manifest -o cmd/quic-config/rsrc.syso

linux: $(LINUX).zip
$(LINUX).zip: $(LINUX)
	@zip -r $(LINUX).zip $(LINUX)

$(LINUX): pkg/statik
	env CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build $(BUILD_FLAGS) -o $(LINUX) ./cmd/quic-config

darwin: $(DARWIN).zip
$(DARWIN).zip: $(DARWIN).app
	@zip -r $(DARWIN).zip $(DARWIN).app

$(DARWIN).app: $(DARWIN)
	go get github.com/machinebox/appify
	appify -name "$(DARWIN)" -icon ./web/src/assets/logo.png $(DARWIN)

$(DARWIN): pkg/statik
	env CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 go build $(BUILD_FLAGS) -o $(DARWIN) ./cmd/quic-config

serve-web:
	npm --prefix web run serve

serve-reload:
	CompileDaemon -build="make MODE=debug $(PLATFORM)" -command="./$(EXECUTABLE)-$(PLATFORM)-amd64"

web/node_modules: web/package.json
	cd web && npm install && cd ..

web: web/dist
web/dist: web/node_modules
	cd web && npm run build && cd ..

pkg/statik: web/dist
	@go get github.com/rakyll/statik
	@go generate ./...

clean:
	@rm -rf web/node_modules web/dist || true
	@rm -rf pkg/statik cmd/quic-config/rsrc.syso || true
	@rm -r $(WINDOWS)* $(LINUX)* $(DARWIN)* || true