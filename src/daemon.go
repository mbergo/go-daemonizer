package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("Please provide a command to run as first argument")
	}

	// get the command to run from the command line first argument
	cmd := os.Args[1]

	// get the arguments to pass to the command
	args := os.Args[2:]

	// get the absolute path to the command
	path, err := exec.LookPath(cmd)
	if err != nil {
		log.Fatalln(err)
	}

	// create a daemon process
	err = daemon(path, args...)
	if err != nil {
		log.Fatalln(err)
	}
}

// daemon creates a daemon process with the given command and arguments
func daemon(cmd string, args ...string) error {
	// create a new process
	process, err := os.StartProcess(cmd, args, &os.ProcAttr{
		Dir:   filepath.Dir(cmd),
		Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
		Sys: &syscall.SysProcAttr{
			Setsid: true,
		},
	})
	if err != nil {
		return err
	}

	// wait for the process to exit
	_, err = process.Wait()
	if err != nil {
		return err
	}

	fmt.Println("Process exited")

	return nil
}
