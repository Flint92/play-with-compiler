package main

import (
	"bufio"
	"fmt"
	"github.com/flint92/play-with-compiler/play"
	"io"
	"os"
	"strings"
)

// >>2;
// 2
// >>2+3;
// 5
// >>1+2*3;
// 7
// >>1+3%2;
// 2
// >>exit();
// Bye!
func main() {
	fmt.Println("Enjoy PlayScript!!!")
	Run(bufio.NewReader(os.Stdin), os.Stdout)
}

func Run(in io.Reader, out io.Writer) {
	defer func() {
		if err := recover(); err != nil {
			_, _ = fmt.Fprintf(out, "Error: %v\n", err)
		}
	}()

	scanner := bufio.NewScanner(in)
	_, _ = fmt.Fprint(out, "\n>>")
	scriptText := ""
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		if line == "exit();" {
			_, _ = fmt.Fprintf(out, "%s\n", "Bye!")
			break
		}

		scriptText += line + "\n"
		if strings.HasSuffix(line, ";") {
			result := play.Eval(scriptText)
			_, _ = fmt.Fprintf(out, "%d\n>>", result)
			scriptText = ""
		}

	}
}
