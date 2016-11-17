package osh

import (
	"io/ioutil"
	"math"
	"testing"
	"time"
)

func TestModTime(t *testing.T) {
	const inaccuracySeconds = 5
	type testElement struct {
		file    string
		modtime string
		err     bool
	}
	test := []testElement{
		// 0
		testElement{
			"",
			"2006-01-02T15:04:05Z",
			true,
		},
		//TODO add check exist file
		/*
			// 1
			testElement{
				"modtime.go",
				"2015-11-11T17:20:52Z",
				false,
			},
		*/
	}
	for i, v := range test {
		//RFC3339     = "2006-01-02T15:04:05Z07:00"
		expModTime, errParse := time.Parse(time.RFC3339, v.modtime)
		if errParse != nil {
			t.Errorf("Got error while parsing string %s: %v", v.modtime, errParse)
		}
		modtime, err := ModTime(v.file)
		if (err != nil) != v.err {
			t.Errorf("Test-%d\nError expected: %v, got: %v", i, v.err, err)
		}

		if !v.err && (err == nil) {
			if !modtime.Equal(expModTime) {
				t.Errorf("Test-%d\nExpected modtime\n%v\ngot\n%v\n", i, expModTime, modtime)
			}
		}
	}

	now := time.Now()
	f, errTemp := ioutil.TempFile("", "temp")
	if errTemp != nil {
		t.Errorf("Got error while creating temp file: %v", errTemp)
	}
	modtime1, err1 := ModTime(f.Name())
	if err1 != nil {
		t.Errorf("Error expected: %v, got: %v", nil, err1)
	} else if math.Abs(float64(modtime1.Sub(now))) < inaccuracySeconds {
		t.Errorf("Expected modtime\n%v\ngot\n%v\n", now, modtime1)
	}
}
