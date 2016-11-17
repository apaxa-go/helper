package strconvh

import "testing"

func TestParseInt(t *testing.T) {

	type testElement struct {
		s   string
		i   int
		err bool
	}

	test := []testElement{
		// 0
		testElement{
			"10",
			10,
			false,
		},

		// 1
		testElement{
			"",
			0,
			true,
		},

		// 2
		testElement{
			"-10",
			-10,
			false,
		},

		// 3
		testElement{
			"0",
			0,
			false,
		},

		// 4
		testElement{
			"10.05",
			0,
			true,
		},

		// 5
		testElement{
			"-10.05",
			0,
			true,
		},

		// 6
		testElement{
			"9223372036854775807",
			9223372036854775807,
			false,
		},

		// 7
		testElement{
			"-9223372036854775808",
			-9223372036854775808,
			false,
		},

		// 8
		testElement{
			"-9223372036854775809",
			0,
			true,
		},

		// 9
		testElement{
			"9223372036854775808",
			0,
			true,
		},

		// 10
		testElement{
			"0x01",
			1,
			true,
		},
	}

	for j, v := range test {
		i, err := ParseInt(v.s)
		if (err != nil) != v.err {
			t.Errorf("Test-%v\nError expected: %v, got: %v", j, v.err, err)
		}
		if !v.err && (err == nil) {
			if i != v.i {
				t.Errorf("Test-%v.Wrong parse. Expected int: %v, got: %v", j, v.i, i)
			}
		}
	}
}

func TestParseInt8(t *testing.T) {

	type testElement struct {
		s   string
		i   int8
		err bool
	}

	test := []testElement{
		// 0
		testElement{
			"10",
			10,
			false,
		},

		// 1
		testElement{
			"",
			0,
			true,
		},

		// 2
		testElement{
			"-10",
			-10,
			false,
		},

		// 3
		testElement{
			"0",
			0,
			false,
		},

		// 4
		testElement{
			"10.05",
			0,
			true,
		},

		// 5
		testElement{
			"10.05",
			0,
			true,
		},

		// 6
		testElement{
			"-128",
			-128,
			false,
		},

		// 7
		testElement{
			"-129",
			0,
			true,
		},

		// 8
		testElement{
			"127",
			127,
			false,
		},

		// 9
		testElement{
			"128",
			0,
			true,
		},
	}

	for i, v := range test {
		i8, err := ParseInt8(v.s)
		if (err != nil) != v.err {
			t.Errorf("Test-%v\nError expected: %v, got: %v", i, v.err, err)
		}
		if !v.err && (err == nil) {
			if i8 != v.i {
				t.Errorf("Test-%v. Wrong parse. Expected int: %v, got: %v", i, v.i, i8)
			}
		}
	}
}

func TestParseInt16(t *testing.T) {

	type testElement struct {
		s   string
		i   int16
		err bool
	}

	test := []testElement{
		// 0
		testElement{
			"10",
			10,
			false,
		},

		// 1
		testElement{
			"",
			0,
			true,
		},

		// 2
		testElement{
			"-10",
			-10,
			false,
		},

		// 3
		testElement{
			"0",
			0,
			false,
		},

		// 4
		testElement{
			"10.05",
			0,
			true,
		},

		// 5
		testElement{
			"10.05",
			0,
			true,
		},

		// 6
		testElement{
			"32767",
			32767,
			false,
		},

		// 7
		testElement{
			"-32768",
			-32768,
			false,
		},

		// 8
		testElement{
			"32768",
			0,
			true,
		},

		// 9
		testElement{
			"-32769",
			0,
			true,
		},
	}

	for i, v := range test {
		i16, err := ParseInt16(v.s)
		if (err != nil) != v.err {
			t.Errorf("Test-%v\nError expected: %v, got: %v", i, v.err, err)
		}
		if !v.err && (err == nil) {
			if i16 != v.i {
				t.Errorf("Test-%v. Wrong parse. Expected int: %v, got: %v", i, v.i, i16)
			}
		}
	}
}

func TestParseInt32(t *testing.T) {

	type testElement struct {
		s   string
		i   int32
		err bool
	}

	test := []testElement{
		// 0
		testElement{
			"10",
			10,
			false,
		},

		// 1
		testElement{
			"",
			0,
			true,
		},

		// 2
		testElement{
			"-10",
			-10,
			false,
		},

		// 3
		testElement{
			"0",
			0,
			false,
		},

		// 4
		testElement{
			"10.05",
			0,
			true,
		},

		// 5
		testElement{
			"10.05",
			0,
			true,
		},

		// 6
		testElement{
			"2147483647",
			2147483647,
			false,
		},

		// 7
		testElement{
			"-2147483648",
			-2147483648,
			false,
		},

		// 8
		testElement{
			"2147483648",
			0,
			true,
		},

		// 9
		testElement{
			"-2147483649",
			0,
			true,
		},
	}

	for i, v := range test {
		i32, err := ParseInt32(v.s)
		if (err != nil) != v.err {
			t.Errorf("Test-%v\nError expected: %v, got: %v", i, v.err, err)
		}
		if !v.err && (err == nil) {
			if i32 != v.i {
				t.Errorf("Test-%v. Wrong parse. Expected int: %v, got: %v", i, v.i, i32)
			}
		}
	}
}

func TestParseInt64(t *testing.T) {

	type testElement struct {
		s   string
		i   int64
		err bool
	}

	test := []testElement{
		// 0
		testElement{
			"10",
			10,
			false,
		},

		// 1
		testElement{
			"",
			0,
			true,
		},

		// 2
		testElement{
			"-10",
			-10,
			false,
		},

		// 3
		testElement{
			"0",
			0,
			false,
		},

		// 4
		testElement{
			"10.05",
			0,
			true,
		},

		// 5
		testElement{
			"10.05",
			0,
			true,
		},

		// 6
		testElement{
			"9223372036854775807",
			9223372036854775807,
			false,
		},

		// 7
		testElement{
			"-9223372036854775808",
			-9223372036854775808,
			false,
		},

		// 8
		testElement{
			"9223372036854775808",
			0,
			true,
		},

		// 9
		testElement{
			"-9223372036854775809",
			0,
			true,
		},
	}

	for i, v := range test {
		i64, err := ParseInt64(v.s)
		if (err != nil) != v.err {
			t.Errorf("Test-%v\nError expected: %v, got: %v", i, v.err, err)
		}
		if !v.err && (err == nil) {
			if i64 != v.i {
				t.Errorf("Test-%v. Wrong parse. Expected int: %v, got: %v", i, v.i, i64)
			}
		}
	}
}

func TestParseUint(t *testing.T) {

	type testElement struct {
		s   string
		u   uint
		err bool
	}

	test := []testElement{
		// 0
		testElement{
			"10",
			10,
			false,
		},

		// 1
		testElement{
			"",
			0,
			true,
		},

		// 2
		testElement{
			"-10",
			0,
			true,
		},

		// 3
		testElement{
			"0",
			0,
			false,
		},

		// 4
		testElement{
			"10.05",
			0,
			true,
		},

		// 5
		testElement{
			"-10.05",
			0,
			true,
		},

		// 6
		testElement{
			"18446744073709551615",
			18446744073709551615,
			false,
		},

		// 7
		testElement{
			"18446744073709551616",
			0,
			true,
		},

		// 8
		testElement{
			"4294967295",
			4294967295,
			false,
		},

		// 9
		testElement{
			"0x01",
			1,
			true,
		},
	}

	for j, v := range test {
		u, err := ParseUint(v.s)
		if (err != nil) != v.err {
			t.Errorf("Test-%v\nError expected: %v, got: %v", j, v.err, err)
		}
		if !v.err && (err == nil) {
			if u != v.u {
				t.Errorf("Test-%v. Wrong parse. Expected int: %v, got: %v", j, v.u, u)
			}
		}
	}
}

func TestParseUint8(t *testing.T) {
	type testElement struct {
		s   string
		u   uint8
		err bool
	}
	test := []testElement{
		// 0
		testElement{
			"10",
			10,
			false,
		},

		// 1
		testElement{
			"",
			0,
			true,
		},

		// 2
		testElement{
			"-10",
			0,
			true,
		},

		// 3
		testElement{
			"0",
			0,
			false,
		},

		// 4
		testElement{
			"10.05",
			0,
			true,
		},

		// 5
		testElement{
			"10.05",
			0,
			true,
		},

		// 6
		testElement{
			"-129",
			0,
			true,
		},

		// 7
		testElement{
			"128",
			128,
			false,
		},

		// 8
		testElement{
			"255",
			255,
			false,
		},

		// 9
		testElement{
			"256",
			0,
			true,
		},
	}

	for i, v := range test {
		u8, err := ParseUint8(v.s)
		if (err != nil) != v.err {
			t.Errorf("Test-%v\nError expected: %v, got: %v", i, v.err, err)
		}
		if !v.err && (err == nil) {
			if u8 != v.u {
				t.Errorf("Test-%v. Wrong parse. Expected int: %v, got: %v", i, v.u, u8)
			}
		}
	}
}

func TestParseUint16(t *testing.T) {

	type testElement struct {
		s   string
		u   uint16
		err bool
	}

	test := []testElement{
		// 0
		testElement{
			"10",
			10,
			false,
		},

		// 1
		testElement{
			"",
			0,
			true,
		},

		// 2
		testElement{
			"-10",
			0,
			true,
		},

		// 3
		testElement{
			"0",
			0,
			false,
		},

		// 4
		testElement{
			"10.05",
			0,
			true,
		},

		// 5
		testElement{
			"10.05",
			0,
			true,
		},

		// 6
		testElement{
			"-32768",
			0,
			true,
		},

		// 7
		testElement{
			"32768",
			32768,
			false,
		},

		// 8
		testElement{
			"65535",
			65535,
			false,
		},

		// 9
		testElement{
			"65536",
			0,
			true,
		},
	}

	for i, v := range test {
		u16, err := ParseUint16(v.s)
		if (err != nil) != v.err {
			t.Errorf("Test-%v\nError expected: %v, got: %v", i, v.err, err)
		}
		if !v.err && (err == nil) {
			if u16 != v.u {
				t.Errorf("Test-%v. Wrong parse. Expected int: %v, got: %v", i, v.u, u16)
			}
		}
	}
}

func TestParseUint32(t *testing.T) {

	type testElement struct {
		s   string
		u   uint32
		err bool
	}

	test := []testElement{
		// 0
		testElement{
			"10",
			10,
			false,
		},

		// 1
		testElement{
			"",
			0,
			true,
		},

		// 2
		testElement{
			"-10",
			0,
			true,
		},

		// 3
		testElement{
			"0",
			0,
			false,
		},

		// 4
		testElement{
			"10.05",
			0,
			true,
		},

		// 5
		testElement{
			"10.05",
			0,
			true,
		},

		// 6
		testElement{
			"-2147483648",
			0,
			true,
		},

		// 7
		testElement{
			"2147483648",
			2147483648,
			false,
		},

		// 8
		testElement{
			"4294967295",
			4294967295,
			false,
		},

		// 9
		testElement{
			"4294967296",
			0,
			true,
		},
	}

	for i, v := range test {
		u32, err := ParseUint32(v.s)
		if (err != nil) != v.err {
			t.Errorf("Test-%v\nError expected: %v, got: %v", i, v.err, err)
		}
		if !v.err && (err == nil) {
			if u32 != v.u {
				t.Errorf("Test-%v. Wrong parse. Expected int: %v, got: %v", i, v.u, u32)
			}
		}
	}
}

func TestParseUint64(t *testing.T) {

	type testElement struct {
		s   string
		u   uint64
		err bool
	}

	test := []testElement{
		// 0
		testElement{
			"10",
			10,
			false,
		},

		// 1
		testElement{
			"",
			0,
			true,
		},

		// 2
		testElement{
			"-10",
			0,
			true,
		},

		// 3
		testElement{
			"0",
			0,
			false,
		},

		// 4
		testElement{
			"10.05",
			0,
			true,
		},

		// 5
		testElement{
			"10.05",
			0,
			true,
		},

		// 6
		testElement{
			"9223372036854775807",
			9223372036854775807,
			false,
		},

		// 7
		testElement{
			"-9223372036854775808",
			0,
			true,
		},

		// 8
		testElement{
			"18446744073709551615",
			18446744073709551615,
			false,
		},

		// 9
		testElement{
			"18446744073709551616",
			0,
			true,
		},
	}

	for i, v := range test {
		u64, err := ParseUint64(v.s)
		if (err != nil) != v.err {
			t.Errorf("Test-%v\nError expected: %v, got: %v", i, v.err, err)
		}
		if !v.err && (err == nil) {
			if u64 != v.u {
				t.Errorf("Test-%v. Wrong parse. Expected uint64: %v, got: %v", i, v.u, u64)
			}
		}
	}
}
