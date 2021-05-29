// Package parser implements parsing of KeyValue values in text format.
package parser

import (
	"fmt"
	"io"
	"strings"
	"text/scanner"
	"unicode"
)

const (
	tokObjectStart rune = '{'
	tokObjectEnd   rune = '}'
	tokQuote       rune = '"'
	tokComment     rune = '/'
)

var (
	identRanges = []*unicode.RangeTable{
		unicode.Number,
		unicode.Letter,
		unicode.Punct,
	}

	escapeSeqs = [][2]string{
		{`\\`, `\`},
		{`\n`, "\n"},
		{`\t`, "\t"},
		{`\"`, `"`},
	}
)

func unquoteToken(t string) string {
	q := string(tokQuote)
	// strings.Trim eats too much in the case "hello \"world\""
	// strconv.Unquote errors with non-quoted strings
	t = strings.TrimPrefix(t, q)
	t = strings.TrimSuffix(t, q)

	return t
}

func unescapeToken(t string) string {
	for _, esc := range escapeSeqs {
		t = strings.ReplaceAll(t, esc[0], esc[1])
	}

	return t
}

func parseToken(t string) string {
	t = unquoteToken(t)
	t = unescapeToken(t)

	return t
}

type namer interface {
	Name() string
}

// TextParser is a parser for KeyValue in text format.
type TextParser struct {
	s *scanner.Scanner
}

// NewTextParser creates a TextParser.
func NewTextParser(fname string, r io.Reader) *TextParser {
	s := (&scanner.Scanner{}).Init(r)

	s.Whitespace = 1<<' ' | 1<<'\t' | 1<<'\r' | 1<<'\n'
	s.Mode = scanner.ScanIdents |
		scanner.ScanStrings |
		scanner.ScanComments |
		scanner.SkipComments

	s.IsIdentRune = func(ch rune, _ int) bool {
		return unicode.In(ch, identRanges...) &&
			ch != tokObjectStart &&
			ch != tokObjectEnd &&
			ch != tokQuote &&
			ch != tokComment
	}

	if fname == "" {
		if n, ok := r.(namer); ok {
			fname = n.Name()
		}
	}

	s.Filename = fname

	return &TextParser{s: s}
}

// Parse reads parses the text-encoded KeyValue values from the input stream, generating an AST
// tree.
func (p *TextParser) Parse() (*Node, error) {
	var scanErr error

	root := &Node{}
	scope := root
	node := root

	p.s.Error = func(s *scanner.Scanner, msg string) {
		scanErr = fmt.Errorf("kv: error at %s: %s", s.Pos(), msg)
	}

	for {
		tok := p.s.Scan()

		if tok == scanner.EOF {
			if node.Type != Object && node.Key != "" && node.Value == "" {
				return root, fmt.Errorf("kv: %s: unexpected EOF", p.s.Pos())
			}

			if scope != nil {
				return root, fmt.Errorf("kv: %s: unexpected EOF", p.s.Pos())
			}

			break
		}

		if scanErr != nil {
			return root, scanErr
		}

		text := p.s.TokenText()

		switch {
		case node.Key == "":
			switch tok {
			case tokObjectEnd:
				scope = scope.Parent
			case scanner.String, scanner.Ident:
				node.Key = parseToken(text)
			default:
				return root, fmt.Errorf(
					"kv: %s: unexpected token %s",
					p.s.Pos(),
					scanner.TokenString(tok),
				)
			}
		case node.Value == "":
			switch tok {
			case tokObjectStart:
				node.Type = Object
			case scanner.String:
				node.Type = Field
				node.Value = parseToken(text)
			case scanner.Ident:
				node.Type = Field
				node.Value = text
			default:
				return root, fmt.Errorf(
					"kv: %s: unexpected token %s",
					p.s.Pos(),
					scanner.TokenString(tok),
				)
			}

			if node != scope {
				scope.addChild(node)
			}

			if node.Type == Object {
				scope = node
			}

			node = &Node{}
		}
	}

	return root, nil
}
