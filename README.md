# snapword

Choose a random snap word.

Snap words are common sight words that early readers learn to recognize
"in a snap", without sounding them out. snapword prints one at random,
which is handy for flash-card practice.

## Usage

```sh
go run .
```

```
curious
```

## Development

```sh
make check    # gofmt, goimports, go vet, golangci-lint, gosec
make build    # cross-compile into ./bin
go test ./...
```
