package main

import (
	_"bytes"
	"fmt"
	_"io"
	"os"
	_"strings"
)

//file table
var FileTable map[uint64]string

//file struct
type file struct {
	name string
	rootpath string
	filehash uint64
	filetype string
	content []byte
	size uint64
}

//filesystem struct
type filesystem struct {
	dir string
	rootpath string
	files map[string]*file
	dirs map[string]*filesystem
	parentdir *filesystem
}

//rootdir
var root *filesystem

//editing file
var editingFile *file

func makefs(dir string, rootpath string, parentdir *filesystem) *filesystem {
	return &filesystem{
		dir: dir,
		rootpath: rootpath,
		files: make(map[string]*file),
		dirs: make(map[string]*filesystem),
		parentdir: parentdir,
	}
}

func fileCont(name string) []byte {
	data, _ := os.ReadFile(name)
	return data
}

//init filesystem
func initfs() *filesystem {
	//make an empty fs
	root = makefs(".", ".", nil)
	fmt.Println("Welcome to sakuni FS")
	return root
}

func (fs *filesystem) pwd() {
	fmt.Println(fs.rootpath)
}

//resets the fs
func (root *filesystem) resetfs(){
	fmt.Println("resetting...")
}

//file offset index
var index int

func (fs *filesystem) touch(name string) bool {
	if _, exists := fs.files[name]; exists {
		fmt.Println("touch: filename already exists")
		return false
	}
	newfile := &file{
		name: name,
		rootpath: fs.rootpath + "/" + name,
	}
	fs.files[name] = newfile

	return true
}

func (fs *filesystem) mkdir(dir string) bool {
	if _, exists := fs.dirs[dir]; exists {
		fmt.Println("mkdir: dir already exists")
		return false
	}
	newdir := makefs(dir, fs.rootpath + "/" + dir, fs)
	fs.dirs[dir] = newdir

	return true
}

func (fs *filesystem) ls() {
	if fs.files != nil {
		fmt.Println("Files:")
		for _, file := range fs.files {
			fmt.Println("\t%s\n", file.name)
		}
	}

	if fs.dirs != nil {
		fmt.Println("Dirs:")
		for _, dir := range fs.dirs {
			fmt.Println("\t%s\n", dir.dir)
		}
	}
}

//exit session
func (root *filesystem) quit() {
	fmt.Println("exiting sakuni FS")
}

//save current state of FS
func (root *filesystem) saveState() {
	fmt.Println("Saving the current state of the FS")
}

func (fs * filesystem) usage(comms []string) bool {
	switch comms[0] {
	case "mkdir":
		if len(comms) < 2 {
			fmt.Println("Usage : mkdir [list of directories to make]")
			return false
		}
	}
	return true
}

// execute runs the commands passed into it.
func (fs * filesystem) execute(command []string) (*filesystem, bool){

	if fs.usage(command) == false {
		return fs, false
	}

	switch command[0] {
		case "pwd":
		fs.pwd()
		case "touch": 
		fs.touch(command[1])
		case "mkdir":
		fs.mkdir(command[1])
		case "ls":
		fs.ls()
		case "quit":
		os.Exit(1)
		default:
		fmt.Println(command[0], ": Command not found")
	}

	return fs, true
}