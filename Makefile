.PHONY: tidy
tidy:
	@go mod tidy

.PHONY: tools
tools: tidy
	@./tools/tools.sh

.PHONY: proto
proto: tools
	@./proto/generate.sh
