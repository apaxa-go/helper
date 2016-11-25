package exech

import (
	"reflect"
	"testing"
)

func TestExec(t *testing.T) {

	type testElement struct {
		name   string
		stdin  []byte
		arg    []string
		stdout []byte
		err    bool
	}

	test := []testElement{
		{name: "grep", stdin: []byte{0x01, 0x02}, arg: []string{""}, stdout: []byte{0x01, 0x02, 0x0a}, err: false},
		{name: "grep", stdin: []byte{0x31, 0x0a, 0x32}, arg: []string{"2"}, stdout: []byte{0x32, 0x0a}, err: false},
		// negative test
		// invalid argument
		{name: "grep", stdin: []byte{0x31, 0x0a, 0x32}, arg: []string{"-Ccccccccc 2"}, stdout: []byte{0x32, 0x0a}, err: true},
		{name: "grep", stdin: []byte{0x31, 0x0a, 0x32}, arg: []string{"2", "-v"}, stdout: []byte{0x31, 0x0a}, err: false},
		// negative test
		{name: "kdsfkjfhskjhf", stdin: []byte{}, arg: []string{}, stdout: []byte{}, err: true},
		{name: "grep", stdin: []byte{}, arg: []string{"2"}, stdout: []byte{}, err: true},
		{name: "grep", stdin: []byte{0x31, 0x0a, 0x32}, arg: []string{"25"}, stdout: []byte{}, err: true},
		{name: "", stdin: []byte{0x31, 0x0a, 0x32}, arg: []string{"25"}, stdout: []byte{}, err: true},
	}

	for i, v := range test {
		stdout, err := Exec(v.name, v.stdin, v.arg...)
		if (err != nil) != v.err {
			t.Errorf("Test-%v.\nError expected: %v, got: %v", i, v.err, err)
		}
		if !v.err && (err == nil) {
			if !reflect.DeepEqual(stdout, v.stdout) {
				t.Errorf("Test-%v.\nExpected stdout: %v\ngot: %v", i, v.stdout, stdout)
			}
		}
	}
}
