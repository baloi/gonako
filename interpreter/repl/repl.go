package repl

import (
	"io"
	"fmt"
	"bufio"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		if line == ".quit" {
			return
		}

		fmt.Println(line)
	}
}