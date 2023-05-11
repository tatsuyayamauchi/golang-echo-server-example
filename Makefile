ENABLE_DEBUG := true

run: ## run application
	go run web/pkg/cmd/web/main.go \
	  --debug


show-flags: ## show application all flags
	go run web/pkg/cmd/web/main.go --help
