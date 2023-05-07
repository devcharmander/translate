lint:
	golangci-lint run --build-tags integration --disable-all -E revive -E staticcheck -E unused -E gocritic -E gocyclo -E gofmt -E misspell -E stylecheck -E unconvert -E unparam

