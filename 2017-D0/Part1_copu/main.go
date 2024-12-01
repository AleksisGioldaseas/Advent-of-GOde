package main

import "strings"

var raw = ""

func main() {

}

func slow() {

	var builder strings.Builder
	for _, r := range raw {
		switch r {
		case 'a':
			builder.WriteRune('A')
		case 'A':
			builder.WriteRune('a')
		case 'c':
			builder.WriteRune('C')
		case 'C':
			builder.WriteRune('c')
		case '2':
			builder.WriteRune('%')
		case '%':
			builder.WriteRune('2')
		default:
			builder.WriteRune(r)
		}
	}
	raw = builder.String()

}

func fast() {
	replacer := strings.NewReplacer("a", "A", "A", "a", "c", "C", "C", "c", "2", "%", "%", "2")
	for range 1 {
		raw = replacer.Replace(raw)
	}
}
