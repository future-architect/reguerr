package gen

import "strings"

type Verb string

var verbs = []string{
	"%v",
	"%+v",
	"%#v",
	"%T",
	"%t",
	"%b",
	"%c",
	"%d",
	"%o",
	"%O",
	"%q",
	"%x",
	"%X",
	"%U",
	"%b",
	"%e",
	"%E",
	"%f",
	"%F",
	"%g",
	"%G",
	"%x",
	"%X",
	"%s",
	"%q",
	"%x",
	"%X",
	"%p",
}

func Analyze(format string) []Verb {
	var containVerbs []Verb

	start := 0
	for start <= len(format) {
		indexOf := strings.Index(format[start:], "%")
		if indexOf == -1 {
			break
		}

		originalStart := start
		for _, verb := range verbs {
			if strings.HasPrefix(format[start + indexOf:], verb) {
				containVerbs = append(containVerbs, Verb(verb))
				start += indexOf + len(verb)
				break
			}
		}

		if start == originalStart {
			start += indexOf + 1
		}
	}

	return containVerbs
}
