package exech

import (
	"errors"
	"io/ioutil"
	"os/exec"
)

// Exec execute binary named by name with arguments args.
// stdin used as standard input.
// Exec returns standard output as slice of bytes and error.
// TODO describe behaviour better (what about stderr, what if exit code).
func Exec(name string, stdin []byte, args ...string) (stdout []byte, err error) {
	//Init command
	//TODO check for error?
	cmd := exec.Command(name, args...)

	//Init std{in,out,err} pipes
	stdinPipe, err := cmd.StdinPipe()
	if err != nil {
		return
	}
	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		return
	}
	stderrPipe, err := cmd.StderrPipe()
	if err != nil {
		return
	}

	//Start command
	if err = cmd.Start(); err != nil {
		return
	}

	//Write original data
	_, err = stdinPipe.Write(stdin) //Write should return error if written not all bytes
	if err != nil {
		return
	}
	stdinPipe.Close()

	//Read compiled data
	stdout, err = ioutil.ReadAll(stdoutPipe)
	if err != nil {
		return
	}

	//Read stderr
	errMessage, err := ioutil.ReadAll(stderrPipe)
	if err != nil {
		return
	}
	if len(errMessage) > 0 {
		return stdout, errors.New(string(errMessage))
	}

	//Wait for command done
	if err = cmd.Wait(); err != nil {
		return
	}

	return
}
