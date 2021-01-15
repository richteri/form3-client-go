GO_RUN = env GO111MODULE=on go

test:
	$(GO_RUN) test ./... -v
