package main

import (
	"bytes"
	"fmt"
	"github.com/davyxu/golexer"
)

const (
	Token_EOF = iota
	Token_Unknown
	Token_Numeral
	Token_String
	Token_WhiteSpace
	Token_LineEnd
	Token_UnixStyleComment
	Token_Identifier
	Token_Dot
	Token_Go
	Token_XX
	Token_Every
	Token_Week
	Token_Semicolon
	Token_BackTicker
	Token_New
	Token_NewClass
)

type CustomParser struct {
	*golexer.Parser
}

func NewCustomParser() *CustomParser {
	l := golexer.NewLexer()
	// 匹配顺序从高到低
	l.AddMatcher(golexer.NewNumeralMatcher(Token_Numeral))
	l.AddMatcher(golexer.NewStringMatcher(Token_String))
	l.AddIgnoreMatcher(golexer.NewWhiteSpaceMatcher(Token_WhiteSpace))
	l.AddIgnoreMatcher(golexer.NewLineEndMatcher(Token_LineEnd))
	l.AddIgnoreMatcher(golexer.NewUnixStyleCommentMatcher(Token_UnixStyleComment))
	l.AddMatcher(golexer.NewSignMatcher(Token_Semicolon, ";"))
	l.AddMatcher(golexer.NewSignMatcher(Token_Dot, "."))
	l.AddMatcher(golexer.NewKeywordMatcher(Token_Go, "go"))
	l.AddMatcher(golexer.NewKeywordMatcher(Token_XX, "xx"))
	l.AddMatcher(golexer.NewKeywordMatcher(Token_Every, "等"))
	l.AddMatcher(golexer.NewKeywordMatcher(Token_Week, "秒"))
	l.AddMatcher(golexer.NewBackTicksMatcher(Token_BackTicker))
	l.AddMatcher(golexer.NewKeywordMatcher(Token_NewClass, "newclass"))
	l.AddMatcher(golexer.NewKeywordMatcher(Token_New, "new"))
	l.AddMatcher(golexer.NewIdentifierMatcher(Token_Identifier))
	l.AddMatcher(golexer.NewUnknownMatcher(Token_Unknown))
	return &CustomParser{
		Parser: golexer.NewParser(l, "custom"),
	}
}

func ErrorCatcher(errFunc func(error)) {
	err := recover()
	switch err.(type) {
	// 运行时错误
	case interface {
		RuntimeError()
	}:
		// 继续外抛， 方便调试
		panic(err)
	case error:
		errFunc(err.(error))
	case nil:
	default:
		panic(err)
	}
}

func main() {
	//反引号内容
	//MatcherName: 'KeywordMatcher' Value: 'go'
	//MatcherName: 'SignMatcher' Value: '.'
	//MatcherName: 'BackTicksMatcher' Value: 'xxx'
	//TestBackTicks()
	//测试词分离，将长的串放在前部先匹配
	//TestSplite()
	//综合测试
	TestParser()
}

func TestBackTicks() {
	p := NewCustomParser()
	defer ErrorCatcher(func(err error) {
		fmt.Println(err)
	})
	p.Lexer().Start("go . `xxx` ")
	p.NextToken()
	for p.TokenID() != 0 {
		fmt.Println(fmt.Sprintf("MatcherName: '%s' Value: '%s'\n", p.MatcherName(), p.TokenValue()))
		p.NextToken()
	}
}

func TestSplite() {
	p := NewCustomParser()
	defer ErrorCatcher(func(err error) {
		fmt.Println(err)
	})
	p.Lexer().Start("newclass new")
	p.NextToken()
	if p.TokenID() != Token_NewClass {
		fmt.Println("expect newclass, got ", p.TokenValue())
		return
	}
	p.NextToken()
	if p.TokenID() != Token_New {
		fmt.Println("expect new, got ", p.TokenValue())
		return
	}
}

func TestParser() {
	p := NewCustomParser()
	defer ErrorCatcher(func(err error) {
		fmt.Println(err)
	})
	p.Lexer().Start(`"a"
		123.3;
		-1
		Base64Text
		gonew.xx
		_id # comment
		等2秒
		"\'\""
		""
		;
		'b'
		`)
	p.NextToken()
	rightAnswer := `===
MatcherName: 'StringMatcher' Value: 'a'
MatcherName: 'NumeralMatcher' Value: '123.3'
MatcherName: 'SignMatcher' Value: ';'
MatcherName: 'NumeralMatcher' Value: '-1'
MatcherName: 'IdentifierMatcher' Value: 'Base64Text'
MatcherName: 'KeywordMatcher' Value: 'go'
MatcherName: 'IdentifierMatcher' Value: 'new'
MatcherName: 'SignMatcher' Value: '.'
MatcherName: 'KeywordMatcher' Value: 'xx'
MatcherName: 'IdentifierMatcher' Value: '_id'
MatcherName: 'KeywordMatcher' Value: '等'
MatcherName: 'NumeralMatcher' Value: '2'
MatcherName: 'KeywordMatcher' Value: '秒'
MatcherName: 'StringMatcher' Value: ''"'
MatcherName: 'StringMatcher' Value: ''
MatcherName: 'SignMatcher' Value: ';'
MatcherName: 'StringMatcher' Value: 'b'
===
`
	var b bytes.Buffer
	b.WriteString("===\n")
	for p.TokenID() != 0 {
		b.WriteString(fmt.Sprintf("MatcherName: '%s' Value: '%s'\n", p.MatcherName(), p.TokenValue()))
		p.NextToken()
	}
	b.WriteString("===\n")
	fmt.Println(b.String())
	if b.String() != rightAnswer {
		fmt.Println("is not same")
	}
}
