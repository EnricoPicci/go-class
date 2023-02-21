package main

import (
	"os"
	"os/exec"
	"runtime"
	"testing"
)

// Test that the printContent function WORKS if it is passed a file that does exist
func TestPrintContentWorks(t *testing.T) {
	worksEnvVar := "I_WILL_WORK"
	var _, thisFile, _, _ = runtime.Caller(0)
	if os.Getenv(worksEnvVar) == "1" {
		printContent(thisFile)
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestPrintContentWorks")
	cmd.Env = append(os.Environ(), worksEnvVar+"=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		t.Fatalf("process ran with err %v, want exit status 1", err)
	}
}

// Test that the printContent function crashes if it is passed a file that does not exist
// https://stackoverflow.com/a/33404435
func TestPrintContentCrash(t *testing.T) {
	crashEnvVar := "I_WILL_CRASH"
	fileThatDoesNotExist := "fileThatDoesNotExists.gooooo"
	if os.Getenv(crashEnvVar) == "1" {
		printContent(fileThatDoesNotExist)
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestPrintContentCrash")
	cmd.Env = append(os.Environ(), crashEnvVar+"=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatalf("process ran with err %v, want exit status 1", err)
}
