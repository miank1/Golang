# Copilot instructions — MoveZerosToEnd

## Purpose
Provide concise, actionable guidance so an AI coding agent can be productive immediately in this repository.

## Big picture 🔍
- Single-file Go program: `main.go` is the only source file and the package is `main`.
- Behavior: the `moveZeros(nums []int)` function mutates the provided slice *in-place* and prints the resulting slice using `fmt.Println`. Example input from `main.go`: `[]int{0,1,0,3,12}` -> printed output: `[1 3 12 0 0]`.
- No module (`go.mod`) or external dependencies are present.

## How to run / build / debug ⚙️
- Run quickly: in this folder run (PowerShell):
  - `go run main.go`
- Build an executable:
  - `go build -o MoveZerosToEnd.exe` (Windows)
- Debugging:
  - Lightweight: use `fmt.Println` or temporarily add logging to `main.go`.
  - With Delve (if installed): `dlv debug` from the repository root.

## Tests and how to add them ✅
- Currently there are no tests. Tests should follow Go idioms:
  - Create `move_zeros_test.go` with `package main` and table-driven tests.
  - Tests should verify the in-place mutation behavior by copying inputs before calling `moveZeros` and comparing with `reflect.DeepEqual`.

Example test skeleton:

```go
func TestMoveZeros(t *testing.T) {
    cases := []struct{ in, want []int }{
        {[]int{0,1,0,3,12}, []int{1,3,12,0,0}},
    }
    for _, c := range cases {
        nums := make([]int, len(c.in)); copy(nums, c.in)
        moveZeros(nums)
        if !reflect.DeepEqual(nums, c.want) {
            t.Fatalf("moveZeros(%v) = %v, want %v", c.in, nums, c.want)
        }
    }
}
```

- Run tests: `go test ./...`

## Project-specific conventions & expectations 🔧
- Current functions mutate input slices in-place (see `moveZeros`). When adding new helpers, follow the same mutation pattern unless there is a clear reason to introduce a new functional/returning API — in that case, document the reason in the PR.
- Keep changes minimal and focused: this is a tiny educational/algorithmic repo; prefer clarity over over-engineering.
- Formatting: run `gofmt`/`go fmt` before submitting a PR.

## When adding dependencies or expanding the project 📦
- Add a `go.mod` at the repo root using `go mod init github.com/<org>/MoveZerosToEnd` and commit it.
- Add unit tests for any new behavior and ensure `go test` passes.

## PR checklist ✅
- Code compiles (`go build`).
- Tests added/updated and `go test` passes.
- Code formatted with `gofmt`.
- Update this instruction file if you introduce new project structure (modules, packages, CI, etc.).

---

If anything in this doc is unclear or you'd like additional examples (e.g., a sample test file added automatically), tell me which sections to expand or modify and I'll iterate. 🙌