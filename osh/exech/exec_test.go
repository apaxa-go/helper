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
		// 0
		testElement{
			"grep",
			[]byte{0x01, 0x02},
			[]string{""},
			[]byte{0x01, 0x02, 0x0a},
			false,
		},

		// 1
		testElement{
			"grep",
			[]byte{0x31, 0x0a, 0x32},
			[]string{"2"},
			[]byte{0x32, 0x0a},
			false,
		},

		// 2
		// negative test
		// invalid argument
		testElement{
			"grep",
			[]byte{0x31, 0x0a, 0x32},
			[]string{"-Ccccccccc 2"},
			[]byte{0x32, 0x0a},
			true,
		},

		// 3
		testElement{
			"grep",
			[]byte{0x31, 0x0a, 0x32},
			[]string{"2", "-v"},
			[]byte{0x31, 0x0a},
			false,
		},

		// 4
		// negative test
		testElement{
			"kdsfkjfhskjhf",
			[]byte{},
			[]string{},
			[]byte{},
			true,
		},

		// 5
		testElement{
			"grep",
			[]byte{},
			[]string{"2"},
			[]byte{},
			true,
		},

		// 6
		testElement{
			"grep",
			[]byte{0x31, 0x0a, 0x32},
			[]string{"25"},
			[]byte{},
			true,
		},

		// 7
		testElement{
			"",
			[]byte{0x31, 0x0a, 0x32},
			[]string{"25"},
			[]byte{},
			true,
		},
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
