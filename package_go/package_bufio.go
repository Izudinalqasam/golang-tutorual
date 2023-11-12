package packagego

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func MainPackageBufio() {
	input := strings.NewReader("this is long string\nwith new line\n")
	writer := bufio.NewWriter(os.Stdout)
	reader := bufio.NewReader(input)

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		fmt.Println(string(line))
	}

	_, _ = writer.WriteString("Hello world\n")
	_, _ = writer.WriteString("Let's Study\n")
	writer.Flush()

}
