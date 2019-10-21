package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

func cmdRun(args []string) {
	cmd := exec.Command(args[0], args[1:]...)
	_, err := cmd.CombinedOutput()
	if err != nil {
		log.Println("cmd.Run() returned an error: ", err)
	}
	//fmt.Println("combined output: ", string(out))
}

func writeFile(dir string, contents string) *os.File {
	log.Println("writing file with contents: ", contents)

	f, err := ioutil.TempFile(dir, "charlie-*.tmp")
	if err != nil {
		log.Println("unable to create temp file: ", err)
		return nil
	}
	_, err = f.Write([]byte(contents))
	if err != nil {
		log.Println("unable to write to temp file: ", err)
		return nil
	}

	log.Println("file written to: ", f.Name())
	return f
}

func readFile(fileName string) {
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Println("unable to read file: ", err)
	}
	log.Println("read file contents: ", string(b))
}

func main() {
	fmt.Println("Hello, eBPF World from within a container!")

	// Write a file to a protected dir "/"
	f := writeFile("/", `A quick brown fox jumps over the lazy little dog.`)
	defer f.Close()
	if f != nil {
		log.Println("reading file from: ", f.Name())
		readFile(f.Name())
	}

	// Exec a command inside
	cmdRun([]string{"/bin/sh"})
}
