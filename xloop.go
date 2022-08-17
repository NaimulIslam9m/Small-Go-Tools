package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

func main() {
	raw_cmd := os.Args
	bin := raw_cmd[1]
	buf, err := exec.Command(bin, raw_cmd[2:]...).CombinedOutput()

	if err == nil {
		for i := 1; i < 100; i++ {
			fmt.Println(string(buf))
			time.Sleep(time.Second)
			cls := exec.Command("clear")
			cls.Stdout = os.Stdout
			cls.Run()
		}
	} else {
		log.Fatal(err)
	}
}
