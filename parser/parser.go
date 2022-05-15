package goresp

import (
	"io"
)

type RespDataType byte

const (
	RespSimpleString RespDataType = '+'
	RespError        RespDataType = '-'
	RespInteger      RespDataType = ':'
	RespBulkString   RespDataType = '$'
	RespArray        RespDataType = '*'

	RespPreTerminalByte RespDataType = '\r'
	RespTerminalByte    RespDataType = '\n'
)

type parserState byte

const (
	consumeChar parserState = iota
	consumeInt
	consumeTerminalByte
	terminate
)

type ParserError string

func (e ParserError) Error() string {
	return string(e)
}

const (
	ErrDifferentTypes    ParserError = "error trying to parse: type trying to be parsed is different from incoming data"
	ErrWrongTerminalByte ParserError = "error reading the last byte from the stream: it's different than the protocol specification"
)

type RespBasicTypesConstrainsts interface {
	string | int
}

type Parser[T RespBasicTypesConstrainsts] interface {
	Parse(r io.Reader) (T, error)
}

type RespParser[T RespBasicTypesConstrainsts] struct {
	Parser[T]
}

func NewParser[T RespBasicTypesConstrainsts](p Parser[T]) *RespParser[T] {
	return &RespParser[T]{
		Parser: p,
	}
}
