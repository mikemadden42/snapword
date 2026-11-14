# TODO

Findings from a project review (2026-11-14). `go vet` and `go test -race` pass on Go 1.27.4.

## Bugs

- [x] **The last word can never be picked** (`main.go:18`). `rand.Intn(length - 1)` returns 0 to len−2, so "enormous" never prints. A one-word list calls `Intn(0)` and an empty list calls `Intn(-1)`; both panic. Fix: `rand.Intn(len(words))` plus a check for an empty list.
- [x] **The test can't catch this** (`snapword_test.go`). It only checks that the result isn't `""`. Instead, call `getWord` many times on a small list and check every word comes back, and test the one-word case.

## Outdated code

- [x] **Remove `rand.Seed`** (`main.go:16`). Deprecated since Go 1.20 and a no-op since Go 1.24. Delete it and the `time` import, or switch to `math/rand/v2` and `rand.IntN`.
- [ ] **Match the Go version in `go.mod` to CI.** `go.mod` says `go 1.17`, but CI tests 1.20–1.22, all out of support. Raise `go.mod` to a supported version (1.22+ is needed for `math/rand/v2`) and test a current matrix such as `stable` and `oldstable`.

## Makefile

- [x] **Rename the binaries** from `gofile_*` to `snapword_*`.
- [x] **Make failures stop `make`.** Commands are joined with `; \`, so a failed step is ignored unless it's the last one. The last line also ends with a stray `; \`. Put each command on its own line, or use `&&`.
- [x] **Add `bin/` to `.gitignore`.** Only `*.exe` is ignored today, so darwin and linux builds could get committed.
- [x] **Clean up the `check` target.** Dropped `golint` (deprecated since 2021). Still open: `golangci-lint` runs with defaults now that `.golangci.yml` is gone, and `release` runs `goreleaser` with no `.goreleaser.yml`; confirm both are intended.

## Minor

- [ ] **CI:** add `go vet ./...` and `go test -race`.
- [ ] **README:** add a usage example and a note on what snap words are (common sight words for early readers).
