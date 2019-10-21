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
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Println("cmd.Run() returned an error: ", err)
	}
	fmt.Println("combined output: ", string(out))
}

func writeFile(dir string) *os.File {
	f, err := ioutil.TempFile(dir, "helloworld-*.tmp")
	if err != nil {
		log.Println("unable to create temp file: ", err)
		return nil
	}
	_, err = f.Write([]byte(`A quick brown fox jumps over the lazy little dog.`))
	if err != nil {
		log.Println("unable to write to temp file: ", err)
		return nil
	}

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
	fmt.Println("Hello, eBPF World!")
	cmdRun([]string{"ls", "-lrtah"})
	f := writeFile("/")
	defer f.Close()

	if f != nil {
		readFile(f.Name())
	}
}
