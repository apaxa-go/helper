package osh

import (
	"io/ioutil"
	"math"
	"os"
	"testing"
	"time"
)

const nonExistingFile = "/asdfasdfasdfadsfasfkjfdkjajkhjasdf"

func TestModTime(t *testing.T) {
	const inaccuracySeconds = 5

	// Test non-existing file
	if _, err := ModTime(nonExistingFile); err == nil {
		t.Error("expect error, but got nil")
	}

	// Test self-created file
	now := time.Now()
	f, errTemp := ioutil.TempFile("", "temp")
	if errTemp != nil {
		t.Errorf("error while creating temp file: %v", errTemp)
	}
	defer os.Remove(f.Name())

	if modtime, err := ModTime(f.Name()); err != nil {
		t.Errorf("error expected: %v, got: %v", nil, err)
	} else if math.Abs(float64(modtime.Sub(now))) < inaccuracySeconds {
		t.Errorf("expected modtime: %v, got %v", now, modtime)
	}
}

func TestExists(t *testing.T) {
	if Exists(nonExistingFile) {
		t.Error("expect false")
	}

	f, errTemp := ioutil.TempFile("", "temp")
	if errTemp != nil {
		t.Errorf("error while creating temp file: %v", errTemp)
	}
	defer os.Remove(f.Name())

	if !Exists(f.Name()) {
		t.Error("expect true")
	}
}
