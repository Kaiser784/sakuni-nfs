package main

import (
	"fmt"
	"os"
	"runtime"
	"os/exec"
	_"io"
	"strings"
)

//shell object
type shell struct {
	env map[string]string //environment variables
}

//env variable stores curr dir vars
var env map[string]string

//initialize shell object
func initShell() *shell {
	env = make(map[string]string)
	return &shell{
		env: env,
	}
}

//clear terminal
func (shell *shell) clearScreen() {
	clear := make(map[string]func())
	clear["Linux"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["Windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	cls, ok := clear[runtime.GOOS]
	if ok {
		cls()
	}
}

//dir check
func (shell *shell) dirCheck(dirName string, fs *filesystem) bool {
	if _, found := fs.dirs[dirName]; found {
		return true
	}
	return false
}

//verifies path in dir
func (shell *shell) verifyPath(dirName string, fs *filesystem) *filesystem {
	checker := shell.handleRoot(dirName, fs)
	segments := strings.Split(dirName, "/")

	for _, segment := range segments {
		if len(segment) == 0 {
			continue
		}
		if segment == ".." {
			if checker.parentdir == nil {
				continue
			} else if shell.dirCheck(segment, checker) == true {
				checker = checker.dirs[segment]
			} else {
				fmt.Println("Error: %s doesn't exist", dirName)
				return fs
			}
		}
	}
	return checker
}

func (shell *shell) handleRoot(dirName string, fs *filesystem) *filesystem {
	if dirName[0] == '/' {
		return root
	}

	return fs
}

func (shell * shell) cd(dirName string, fs *filesystem) *filesystem {
	if dirName == "/" {
		return root
	}

	return shell.verifyPath(dirName, fs)
}

func (shell * shell) open(filename string, fs *filesystem) error {
	segments := strings.Split(filename, "/")
	if len(segments) == 1 {
		if _, exists := fs.files[filename]; exists {
			editingFile = fs.files[filename]
			// editor()
		} else {
			fmt.Println(filename, ": doesn't exist")
		}
	} else {	
		dirPath := strings.Join(segments[:len(segments)-2],"/")
		tmp := shell.verifyPath(dirPath, fs)

		if _, exists := tmp.files[segments[len(segments)-1]]; exists {
			editingFile = tmp.files[segments[len(segments)-1]]
			// editor()
		} else {
			fmt.Println(filename, ": doesn't exist")
		}

	}
	return nil
}

func (shell *shell) cat(filename string, fs *filesystem) {
	
	segments := strings.Split(filename, "/")
	if len(segments) == 1 {
		if _, exists := fs.files[filename]; exists {
			content := string(fs.files[filename].content)
			fmt.Println(content)
		} else {
			fmt.Println("cat: file doesn't exist")
		}
	} else {
		dirPath := strings.Join(segments[:len(segments)-2],"/")
		tmp := shell.verifyPath(dirPath, fs)
		
		if _, exists := tmp.files[segments[len(segments)-1]]; exists {
			content := string(tmp.files[segments[len(segments)-1]].content)
			fmt.Println(content)
		} else {
			fmt.Println(filename, ": doesn't exist")
		}

	}
}

func (s * shell) usage(comms[] string) bool {
	switch comms[0] {
	case "cd":
		if len(comms) != 2 {
			fmt.Println("Usage : cd [target directory")
			return false
		}
	case "cat":
		if len(comms) != 2 {
			fmt.Println("Usage : cat [target file]")
			return false
		}

	case "open":
		if len(comms) != 2 {
			fmt.Println("Usage : open [file name]")
			return false
		}
	}
	return true
}

func (s * shell) execute(comms []string, fs *filesystem) (*filesystem, bool) {

	if s.usage(comms) == false {
		return fs, false
	}
	switch comms[0] {
	case "open":
		s.open(comms[1], fs)
	case "cd":
		fs = s.cd(comms[1], fs)
	case "cat":
		s.cat(comms[1], fs)
	case "clear":
		s.clearScreen()
	default:
		return fs, false
	}
	return fs, true
}