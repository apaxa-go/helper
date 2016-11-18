package ioutilh

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func TestReadDirNames(t *testing.T) {

	//check dir is empty
	nameDir, errTempDir := ioutil.TempDir("", "temp")
	if errTempDir != nil {
		t.Errorf("TestReadDirNames. Got error while creating temp dir: %v", errTempDir)
	}
	s, err := ReadDirNames(nameDir, true)
	if err != nil {
		t.Errorf("TestReadDirNames. Got error: %v", err)
	}
	if len(s) != 0 {
		t.Errorf("TestReadDirNames. Slice should be empty but len(s): %v", s)
	}

	//check dir contains dirs
	errChdir := os.Chdir(nameDir)
	if errChdir != nil {
		t.Errorf("TestReadDirNames. Got error while changing dir: %v", errChdir)
	}
	dirs := []string{"bla", "bla1", "bla2", "bla3"}
	for i, v := range dirs {
		errMkdir := os.Mkdir(v, 0777)
		if errMkdir != nil {
			t.Errorf("TestReadDirNames. Got error while making dir №%v: %v", i, errMkdir)
		}
	}
	s, err = ReadDirNames(nameDir, true)
	if err != nil {
		t.Errorf("TestReadDirNames. Got error: %v", err)
	}
	if !reflect.DeepEqual(s, dirs) {
		t.Errorf("TestReadDirNames. Wrong dir names. Expected: %v, got: %v", dirs, s)
	}

	//check dir with files
	nameFile := "bla_file"
	_, errFile := os.Create(nameFile)
	if errFile != nil {
		t.Errorf("TestReadDirNames. Got error while creating file: %v", errFile)
	}
	entries := append(dirs, nameFile)
	s, err = ReadDirNames(nameDir, true)
	if err != nil {
		t.Errorf("TestReadDirNames. Got error: %v", err)
	}
	if !reflect.DeepEqual(s, entries) {
		t.Errorf("TestReadDirNames. Wrong dir/file's names. Expected: %v, got: %v", entries, s)
	}

	//func (f *File) Chmod(mode FileMode) error
	nameFile = "bla_file1"
	F, errFile1 := os.Create(nameFile)
	if errFile1 != nil {
		t.Errorf("TestReadDirNames. Got error while creating file: %v", errFile1)
	}
	errChmod := F.Chmod(0000)
	if errChmod != nil {
		t.Errorf("TestReadDirNames. Got error: %v", errChmod)
	}
	entries = append(entries, nameFile)
	s, err = ReadDirNames(nameDir, true)
	if err != nil {
		t.Errorf("TestReadDirNames. Got error: %v", err)
	}
	if !reflect.DeepEqual(s, entries) {
		t.Errorf("TestReadDirNames. Wrong dir/file's names. Expected: %v, got: %v", entries, s)
	}
	errChmod = F.Chmod(0777)
	if errChmod != nil {
		t.Errorf("TestReadDirNames. Got error: %v", errChmod)
	}
	s, err = ReadDirNames(nameDir, true)
	if err != nil {
		t.Errorf("TestReadDirNames. Got error: %v", err)
	}
	if !reflect.DeepEqual(s, entries) {
		t.Errorf("TestReadDirNames. Wrong dir/file's names. Expected: %v, got: %v", entries, s)
	}
	errChmod = F.Chmod(0333)
	if errChmod != nil {
		t.Errorf("TestReadDirNames. Got error: %v", errChmod)
	}
	s, err = ReadDirNames(nameDir, true)
	if err != nil {
		t.Errorf("TestReadDirNames. Got error: %v", err)
	}
	if !reflect.DeepEqual(s, entries) {
		t.Errorf("TestReadDirNames. Wrong dir/file's names. Expected: %v, got: %v", entries, s)
	}
	s, err = ReadDirNames(".", true)
	if err != nil {
		t.Errorf("TestReadDirNames. Got error: %v", err)
	}
	if !reflect.DeepEqual(s, entries) {
		t.Errorf("TestReadDirNames. Wrong dir/file's names. Expected: %v, got: %v", entries, s)
	}

	errChdir1 := os.Chdir("bla1")
	if errChdir1 != nil {
		t.Errorf("TestReadDirNames. Got error while change dir: %v", errChdir1)
	}
	s, err = ReadDirNames("..", true)
	if err != nil {
		t.Errorf("TestReadDirNames. Got error: %v", err)
	}
	if !reflect.DeepEqual(s, entries) {
		t.Errorf("TestReadDirNames. Wrong dir/file's names. Expected: %v, got: %v", entries, s)
	}

	// remove dir
	errRemove := os.RemoveAll(nameDir)
	if errRemove != nil {
		t.Errorf("TestReadDirNames. Got error while removing dir: %v", errRemove)
	}

	//check nonexistent dir
	_, errN := ReadDirNames(nameDir, true)
	if errN == nil {
		t.Error("\nTestReadDirNames.\nError expected but got nil")
	}
}

func TestIsEmpty(t *testing.T) {
	//check empty dir
	nameDir, errTempDir := ioutil.TempDir("", "temp")
	if errTempDir != nil {
		t.Errorf("TestIsEmpty. Got error while creating temp dir: %v", errTempDir)
	}
	b, err := IsDirEmpty(nameDir)
	if err != nil {
		t.Errorf("\nTestIsEmpty.\nError expected: %v, got: %v", nil, err)
	} else if !b {
		t.Error("\nTestIsEmpty.\nError: expected empty dir")
	}
	//check not empty dir
	_, errTempFile := ioutil.TempFile(nameDir, "temp")
	if errTempFile != nil {
		t.Errorf("TestIsEmpty. Got error while creating temp file: %v", errTempFile)
	}

	b1, err1 := IsDirEmpty(nameDir)
	if err1 != nil {
		t.Errorf("\nTestIsEmpty.\nError expected: %v, got: %v", nil, err1)
	} else if b1 {
		t.Error("\nTestIsEmpty.\nError: expected empty dir")
	}
	//check nonexistent dir
	_, err2 := IsDirEmpty("")
	if err2 == nil {
		t.Error("\nTestIsEmpty.\nError expected, got: nil")
	}
}
