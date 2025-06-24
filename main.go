package main

import (
	"bufio"
	"io"
	"os"
	"strings"
)

func readfiles(filename string, desfile *os.File, ch *chan bool) {
	sorfile, err := os.Open(filename)
	checknilerror(err)
	defer sorfile.Close()

	scanner := bufio.NewScanner(sorfile)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "ERROR") {
			io.WriteString(desfile, line+"\n")

		}
	}
	*ch <- true

}

func main() {
	filename := []string{"log1.log", "log2.log", "log3.log", "log4.log", "log5.log"}

	desfile, err := os.Create("output.log")
	checknilerror(err)

	defer desfile.Close()

	ch := make(chan bool)
	for _, val := range filename {
		go readfiles(val, desfile, &ch)
	}
	for i := 0; i < len(filename); i++ {
		<-ch
	}
}
func checknilerror(err error) {
	if err != nil {
		panic(err)
	}
}
