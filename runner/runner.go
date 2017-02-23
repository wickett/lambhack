package runner

import (
	"bytes"
	"log"
	"net/url"
	"os/exec"
	"strings"
)

func Run(cmd string) string {
	var out bytes.Buffer
	unescaped, _ := url.QueryUnescape(cmd)
	commands := strings.Split(unescaped, " ")
	command := string(commands[0])
	args := append(commands[:0], commands[1:]...)

	c := exec.Command(command, args...)
	c.Stdout = &out
	err := c.Run()
	if err != nil {
		log.Fatal(err)
	}
	//	fmt.Printf("out: %q\n", out.String())
	return out.String()
}
