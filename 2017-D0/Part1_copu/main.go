package main

import "strings"

var raw = "asd;gl2khasdpoi2asd;sdfaasdf;xzjcbglxmcngccccccccccccccopieadf22hgtaoeiwurthy2423908522jbsdtwetwe2dtsa2rt2weasd;gl2khasdpoi2yqwyerp2oiqewy22rl;xzjcbglxmcngopieadf22hgtaoeiwurthy2423908522jbsdtwetwe2dtsa2rt2weasd;gl2khasdpoi2yqwyerp2oiqewy22rl;xzjcbglxmcngopieadf22hgtaoeiwurthy2423908522jbsdtwetwe2dtsa2rt2weasd;gl2khasdpoi2yqwyerp2oiqewy22rl;xzjccccccccccccbglxmcngopieadf22hgtao222222222222eiwurthy2423908522jbsdtwetwe2dtsa2rt2weyqwyerp2oiqewy22rl;xzjcbglxmcngopieadf22hgtaoeiwurthy2423908522jbsdtwetwe2dtsa2rt2we"

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
