package repl

import (
	"bufio"
	"fmt"
	"github.com/Hyla96/whirlpool/lexer"
	"github.com/Hyla96/whirlpool/token"
	"io"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {

	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprint(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()

		if line == "quit" {
			break
		}

		l := lexer.New(line)
		tok := l.NextToken()

		for tok.Type != token.EOF {
			fmt.Fprintf(out, "%+v\n", tok)
			tok = l.NextToken()
		}

	}

}
