package filepathh

import (
	"errors"
	"github.com/apaxa-go/helper/ioh/ioutilh"
	"github.com/apaxa-go/helper/strconvh"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

var renameNum = 1

func rename(path string, info os.FileInfo, err error) error {
	err2 := os.Rename(path, path+strconvh.FormatInt(renameNum))
	renameNum++
	return err2
}

func TestWalk(t *testing.T) {
	//create tempDir
	nameDir, errTempDir := ioutil.TempDir("", "temp")
	if errTempDir != nil {
		t.Errorf("got error while creating temp dir: %v", errTempDir)
	}
	errChdir := os.Chdir(nameDir)
	if errChdir != nil {
		t.Errorf("got error while changing dir: %v", errChdir)
	}
	dirs := []string{"dir4", "dir1", "dir3", "dir2"}
	for i, v := range dirs {
		errMkdir := os.Mkdir(v, 0777)
		if errMkdir != nil {
			t.Errorf("got error while making dir #%v: %v", i, errMkdir)
		}
	}
	//call Walk
	err := Walk(nameDir, rename, true, true)
	if err != nil {
		t.Errorf("walk error: %v", err)
	}
	//check dirs
	sNew := []string{"dir11", "dir22", "dir33", "dir44"}
	s, err1 := ioutilh.ReadDirNames(nameDir+"5", true)
	if !reflect.DeepEqual(s, sNew) {
		t.Errorf("check dir names: %v, error: %v", s, err1)
	}
	// remove dir
	errRemove := os.RemoveAll(nameDir + "1")
	if errRemove != nil {
		t.Errorf("got error while removing dir: %v", errRemove)
	}
	err = Walk("", rename, true, true)
	if err == nil {
		t.Error("error expected but got nil")
	}
}

func rename2(path string, info os.FileInfo, err error) error {
	err2 := os.Rename(path, path+"1")
	return err2
}

func TestWalk2(t *testing.T) {
	//create tempDir
	nameDir, errTempDir := ioutil.TempDir("", "temp")
	if errTempDir != nil {
		t.Errorf("got error while creating temp dir: %v", errTempDir)
	}
	errChdir := os.Chdir(nameDir)
	if errChdir != nil {
		t.Errorf("got error while changing dir: %v", errChdir)
	}
	dirs := []string{"dir4", "dir3", "dir2", "dir1"}
	for i, v := range dirs {
		errMkdir := os.Mkdir(v, 0777)
		if errMkdir != nil {
			t.Errorf("got error while making dir №%v: %v", i, errMkdir)
		}
	}
	//call Walk
	err := Walk(nameDir, rename2, false, true)
	if err != nil {
		t.Errorf("walk error: %v", err)
	}
	//check dirs
	sNew := []string{"dir11", "dir21", "dir31", "dir41"}
	s, err1 := ioutilh.ReadDirNames(nameDir, true)
	if !reflect.DeepEqual(s, sNew) {
		t.Errorf("check dir names:%v\nerror: %v", s, err1)
	}
	// remove dir
	errRemove := os.RemoveAll(nameDir)
	if errRemove != nil {
		t.Errorf("got error while removing dir: %v", errRemove)
	}
}

func rename3(path string, info os.FileInfo, err error) error {
	return errors.New("func for return error")
}

func TestWalk3(t *testing.T) {
	//create tempDir
	nameDir, errTempDir := ioutil.TempDir("", "temp")
	if errTempDir != nil {
		t.Errorf("got error while creating temp dir: %v", errTempDir)
	}
	//call Walk
	err := Walk(nameDir, rename3, true, true)
	if err == nil {
		t.Error("error expected but got nil")
	}
	// remove dir
	errRemove := os.RemoveAll(nameDir)
	if errRemove != nil {
		t.Errorf("got error while removing dir: %v", errRemove)
	}
}

func TestWalk4(t *testing.T) {
	if err := Walk("/sadsadfjhsaduirqewnvxz", nil, false, false); err == nil {
		t.Error("expect error, got nil")
	} else if _, ok := err.(*os.PathError); !ok {
		t.Errorf("expect error of type os.PathError, got %v", err)
	}
}

func TestWalk5(t *testing.T) {
	var tmpDir string
	var err error
	if tmpDir, err = ioutil.TempDir("", "temp"); err != nil {
		t.Error(err)
	}
	defer func() {
		err := os.RemoveAll(tmpDir)
		if err != nil {
			t.Error(err)
		}
	}()

	if err = os.Mkdir(tmpDir+"/denied", 0x0); err != nil {
		t.Error(err)
	}

	//call Walk
	wf := func(path string, info os.FileInfo, err error) error { return err }
	if err = Walk(tmpDir, wf, true, true); err == nil {
		t.Error("error expected but got nil")
	}

}

func TestWalk6(t *testing.T) {
	var tmpDir string
	var err error
	if tmpDir, err = ioutil.TempDir("", "temp"); err != nil {
		t.Error(err)
	}
	defer func() {
		err := os.RemoveAll(tmpDir)
		if err != nil {
			t.Error(err)
		}
	}()

	if f, err := os.Create(tmpDir + "/file1"); err != nil {
		t.Error(err)
	} else {
		if err = f.Close(); err != nil {
			t.Error(err)
		}
	}
	if f, err := os.Create(tmpDir + "/file2"); err != nil {
		t.Error(err)
	} else {
		if err = f.Close(); err != nil {
			t.Error(err)
		}
	}

	//call Walk
	wf := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.Name() == "file1" {
			return os.Remove(tmpDir + "/file2")
		}
		return nil
	}
	if err = Walk(tmpDir, wf, true, true); err == nil {
		t.Error("error expected but got nil")
	}

}
