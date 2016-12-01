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
		t.Error("expect error, but got nil")
	}

	// Test self-created file
	now := time.Now()
	f, errTemp := ioutil.TempFile("", "temp")
	if errTemp != nil {
		t.Errorf("error while creating temp file: %v", errTemp)
	}

	if modtime, err := ModTime(f.Name()); err != nil {
		t.Errorf("error expected: %v, got: %v", nil, err)
	} else if math.Abs(float64(modtime.Sub(now))) < inaccuracySeconds {
		t.Errorf("expected modtime: %v, got %v", now, modtime)
	}
}
