package parser

import (
	"strconv"
	"bytes"
)

//keyword,
//identifier,
//literal, 字面量
//symbol,
//paren,
//semicolon,// ;
//other,
//eof
const (
	KEYWORD    = iota
	IDENTIFIER
	LITERAL
	SYMBOL
	PAREN
	SEMICOLON
	EOF
)

var symbol = []byte{'+', '-', '*', '/', '='}

var keywords = map[string]int{
	"insert": 1,
	"select":1,
	"into":1,
	"update": 1,
	"values": 3,
	"set":3,
	"from":3,
	"where":4,
}

type token struct {
	kind int
	val  interface{}
}
type tokenizer struct {
	store    []byte
	pos      int
	tmp      int
	curToken []token
	err      error
}

func newTokenizer(store string) *tokenizer {
	return &tokenizer{[]byte(store), -1, -1, []token{}, nil}
}

func (tkn *tokenizer) getByte() byte {
	return tkn.store[tkn.pos]
}

func (tkn *tokenizer) peekByte() (b byte, isEof bool) {
	tkn.pos++
	isEof = tkn.pos >= len(tkn.store)
	if (isEof) {
		tkn.pos = len(tkn.store)-1
	}
	return tkn.store[tkn.pos], isEof
}

func (tkn *tokenizer) mark() {
	tkn.tmp = tkn.pos
}
func (tkn *tokenizer) track() {
	tkn.pos = tkn.tmp
	tkn.tmp = -1
}

func (tkn *tokenizer) getDigit() token {
	var buf []byte
	var isFloat bool
	tk := token{kind: LITERAL}
	p := tkn.getByte()
	buf = append(buf, p)
	for {
		tkn.mark()
		b, eof := tkn.peekByte()
		if (b == '.') {
			isFloat = true
			buf = append(buf, b)
			continue
		}
		if eof || isSpace(b) || !isDigit(b) {
			tkn.track()
			break
		}
		buf = append(buf, b)
	}
	if isFloat {
		val, _ := strconv.ParseFloat(string(buf), 64)
		tk.val = val
	} else {
		val, _ := strconv.ParseInt(string(buf), 10, 32)
		tk.val = val
	}
	return tk
}

func (tkn *tokenizer) getString() token {
	var buf []byte
	p := tkn.getByte()
	buf = append(buf, p)
	for {
		tkn.mark()
		b, eof := tkn.peekByte()
		buf = append(buf, b)
		if b=='\''{
			// TODO:
			break
		}
		if eof {
			// error

		}
	}
	return token{LITERAL,string(buf)}
}

func (tkn *tokenizer) getIdentifier() token {
	var buf []byte
	var tk token
	p := tkn.getByte()
	buf = append(buf, p)
	for {
		tkn.mark()
		b, eof := tkn.peekByte()
		if eof || isSpace(b) || isSymbol(b) {
			tkn.track()
			break
		}
		buf = append(buf, b)
	}
	if _, ok := keywords[string(buf)]; ok {
		tk.kind = KEYWORD
		tk.val = string(buf)
	} else {
		tk.kind = IDENTIFIER
		tk.val = string(buf)
	}
	return tk
}

func (tkn *tokenizer) getSymbol() token {
	var buf []byte
	p := tkn.getByte()
	buf = append(buf, p)
	for {
		tkn.mark()
		b, eof := tkn.peekByte()
		if eof || isSpace(b) || isAlpha(b)||statSymbol(b) {
			tkn.track()
			break
		}
		buf = append(buf, b)
	}
	return token{SYMBOL, string(buf)}
}

func (tkn *tokenizer) getToken() {
for{
	a:=next(tkn)
	tkn.curToken=append(tkn.curToken,a)
	if(a.kind==EOF){
		break
	}
}
	tkn.pos,tkn.tmp=0,0
}
// 推进
func (tkn *tokenizer) popToken() {
	tkn.pos++
}
// 获得前向
func (tkn *tokenizer) PeekToken()token {
	tkn.tmp++
	return tkn.curToken[tkn.tmp]
}
// pos->tem 重置前向
func (tkn *tokenizer) token() token{
	tkn.mark()
	return tkn.curToken[tkn.pos]
}

func next(tkn *tokenizer) token{
	for {
		b, eof := tkn.peekByte()
		if (eof) {
			return token{EOF,nil}
		}
		if isSpace(b) {
			continue
		}
		if isDigit(b) {
			return tkn.getDigit()
		}
		if isAlpha(b) {
			return tkn.getIdentifier()
		}
		if isSymbol(b) {
			return tkn.getSymbol()
		}
		if b == '\'' {
			return tkn.getString()
		}
		if b =='('||b==')'{
			return token{PAREN,string(b)}
		}
	}
}

func isSpace(b byte) bool {
	return b == ' '
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func isAlpha(b byte) bool {
	return b >= 'a' && b <= 'z' || b >= 'A' && b <= 'Z'
}

func isSymbol(b byte) bool {
	return bytes.Contains(symbol, []byte{b})
}

func statSymbol(b byte) bool  {
	return b=='\''||b=='`'
}


func isEof(tk token) bool {
	return tk.kind==EOF
}

func isKeyword(tk token) bool {
	return tk.kind==KEYWORD
}