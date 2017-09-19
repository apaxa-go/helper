package boundaryh

//go:generate go run ./internal/test-generator/main.go ../internal/ucd-data

type ucdTest struct {
	runes      []rune
	boundaries []Boundary
}
