package reader

import "regexp"

var (
	reg = regexp.MustCompile("[\\s,]*(~@|[\\[\\]{}()'`~^@]|\"(?:\\\\.|[^\\\\\"])*\"?|;.*|[^\\s\\[\\]{}('\"`,;)]*)")
)

func tokenize(s string) []Token {
	matches := reg.FindAllStringSubmatch(s, -1)
	if len(matches) == 0 {
		return nil
	}

	tokens := make([]Token, len(matches))
	for i, match := range matches {
		tokens[i] = Token(match[1])
	}
	return tokens
}
