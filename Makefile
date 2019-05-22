PLUGIN_NAME=reef

GO_FLAGS=CGO_ENABLED=0
OUTPUT=docker-$(PLUGIN_NAME)

build:
	$(GO_FLAGS) go build -o $(OUTPUT) cmd/cli/*.go

mktargetdir:
	mkdir -p ~/.docker/cli-plugins

install: mktargetdir build
	cp -rf $(OUTPUT) ~/.docker/cli-plugins/$(OUTPUT)

link: mktargetdir build
	ln -sf ${CURDIR}/$(OUTPUT) ~/.docker/cli-plugins/$(OUTPUT)

.PHONY: build mktargetdir install link
.DEFAULT: build
