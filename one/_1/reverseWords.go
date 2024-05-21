package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func reverseWords(s string) string {
	t := strings.Fields(s)
	str := ""
	for i := len(t) - 1; i >= 0; i-- {
		str += t[i] + " "
	}
	str = strings.TrimSpace(str)
	return str
}

func main() {
	var name string
	reader := bufio.NewReader(os.Stdin)
	name, _ = reader.ReadString('\n')
	fmt.Println(reverseWords(name))
}
