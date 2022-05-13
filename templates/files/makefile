all: helm-template

.PHONY: helm-template
helm-template:
	@rm -rf ./build
	@mkdir -p build
	helm template --debug -g --output-dir build .
	@# . > generated.yaml
