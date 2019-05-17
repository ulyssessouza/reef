PLUGIN_NAME=reef

GO_FLAGS=CGO_ENABLED=0
GO_BUILD=$(GO_FLAGS) go build
OUTPUT=docker-$(PLUGIN_NAME)

build:
	$(GO_BUILD) -o $(OUTPUT)

mktargetdir:
	mkdir -p ~/.docker/cli-plugins

install: mktargetdir build
	cp -rf $(OUTPUT) ~/.docker/cli-plugins/$(OUTPUT)

link: mktargetdir build
	ln -sf ${CURDIR}/$(OUTPUT) ~/.docker/cli-plugins/$(OUTPUT)

.PHONY: build mktargetdir install link
.DEFAULT: build
