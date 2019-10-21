package main

import (
	"fmt"
	"log"
	"os/exec"
)

func cmdRun(args []string) {
	cmd := exec.Command(args[0], args[1:]...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Print("cmd.Run() returned an error: ", err)
	}
	fmt.Println("combined output: ", string(out))
}

func main() {
	fmt.Println("Hello, eBPF World!")
	cmdRun([]string{"ls", "-lrtah"})
}
