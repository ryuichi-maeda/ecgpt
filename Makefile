help: ## makeコマンド一覧
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

test: ## go test
	go test -v ./...

release-local: ## ローカルでGoReleaseの実行
	goreleaser release --snapshot -clean

release-v%: ## 本番リリース
	git tag ${@:release-%=%}
	git push origin ${@:release-%=%}
