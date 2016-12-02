package filepathh

import (
	"os"
	"strings"
	"testing"
)

func TestExtractExt(t *testing.T) {
	type testElement struct {
		p   string
		ext string
		rem string
	}

	tests := []testElement{
		{"", "", ""},
		{".ext", ".ext", ""},
		{".", ".", ""},
		{"name", "", "name"},
		{"name.", ".", "name"},
		{"name.ext", ".ext", "name"},
		{"/", "", "/"},
		{"/.", ".", "/"},
		{"./", "", "./"},
		{"./.", ".", "./"},
		{"pa.th/", "", "pa.th/"},
		{"pa.th/name", "", "pa.th/name"},
		{"pa.th/name.", ".", "pa.th/name"},
		{"pa.th/name.ext", ".ext", "pa.th/name"},
		{"/name", "", "/name"},
		{"/name.", ".", "/name"},
		{"/name.ext", ".ext", "/name"},
	}

	// Adopt to runtime OS separator
	for i := range tests {
		tests[i].p = strings.Replace(tests[i].p, "/", string(os.PathSeparator), -1)
		tests[i].ext = strings.Replace(tests[i].ext, "/", string(os.PathSeparator), -1)
		tests[i].rem = strings.Replace(tests[i].rem, "/", string(os.PathSeparator), -1)
	}

	for _, v := range tests {
		if rem, ext := ExtractExt(v.p); rem != v.rem || ext != v.ext {
			t.Errorf("expect '%v' '%v', got '%v' '%v'", v.rem, v.ext, rem, ext)
		}
	}
}
