package osh

import (
	"io/ioutil"
	"math"
	"testing"
	"time"
)

func TestModTime(t *testing.T) {
	const inaccuracySeconds = 5

	// Test non-existing file
	if _, err := ModTime("/asdfasdfasdfadsfasfkjfdkjajkhjasdf"); err == nil {
		t.Error("Expected error, but got nil")
	}

	// Test self-created file
	now := time.Now()
	f, errTemp := ioutil.TempFile("", "temp")
	if errTemp != nil {
		t.Errorf("Got error while creating temp file: %v", errTemp)
	}

	if modtime, err := ModTime(f.Name()); err != nil {
		t.Errorf("Error expected: %v, got: %v", nil, err)
	} else if math.Abs(float64(modtime.Sub(now))) < inaccuracySeconds {
		t.Errorf("Expected modtime\n%v\ngot\n%v\n", now, modtime)
	}
}
