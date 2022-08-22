package chapter09

import "fmt"

type Linter struct {
	stack   *Stack[rune]
	matches map[rune]rune
}

func NewLinter() *Linter {
	return &Linter{
		stack: NewStack([]rune{}),
		matches: map[rune]rune{
			'{': '}',
			'(': ')',
			'[': ']',
		},
	}
}

func (l *Linter) Lint(text string) error {
	for _, c := range text {
		if l.Opening(c) {
			l.stack.Push(c)
		} else if l.Closing(c) {
			last, err := l.stack.Pop()
			if err != nil {
				return fmt.Errorf("%s didn't have an opening", string(c))
			}
			if !l.IsMatch(last, c) {
				return fmt.Errorf("last opening %s did not match encounter closing %s", string(last), string(c))
			}
		}
	}
	if v, err := l.stack.Read(); err == nil {
		return fmt.Errorf("stack was not empty at the end of the string, contained: %s", string(v))
	}
	return nil
}

func (l *Linter) IsMatch(opening, closing rune) bool {
	if v, ok := l.matches[opening]; ok {
		if v == closing {
			return true
		}
	}
	return false
}

func (l *Linter) Opening(c rune) bool {
	return l.check([]byte{'{', '(', '['}, c)
}

func (l *Linter) Closing(c rune) bool {
	return l.check([]byte{'}', ')', ']'}, c)
}

func (l *Linter) check(slice []byte, c rune) bool {
	for _, b := range slice {
		if byte(c) == b {
			return true
		}
	}
	return false
}
