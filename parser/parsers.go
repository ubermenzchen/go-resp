package goresp

import "io"

type IntegerParser struct{}

func (p *IntegerParser) Parse(r io.Reader) (int, error) {
	buffer := make([]byte, 1)
	_, err := r.Read(buffer)
	if err != nil {
		return 0, err
	}

	if RespDataType(buffer[0]) != RespInteger {
		return 0, ErrDifferentTypes
	}

	state := consumeInt
	buffer[0] = 0
	result := 0
	decimal := 0

	for state != terminate {
		_, err := r.Read(buffer)
		if err != nil {
			return 0, err
		}

		switch state {
		case consumeInt:
			if buffer[0] < '0' || buffer[0] > '9' {
				if buffer[0] == '\r' {
					state = consumeTerminalByte
					continue
				}
				return 0, ErrDifferentTypes
			}
			result = 10*result + int(buffer[0]-byte(48))
			decimal++

		case consumeTerminalByte:
			if buffer[0] != '\n' {
				return 0, ErrWrongTerminalByte
			}
			state = terminate
		}
	}

	return result, nil
}

type SimpleStringParser struct{}

func (p *SimpleStringParser) Parse(r io.Reader) (string, error) {
	buffer := make([]byte, 1)
	_, err := r.Read(buffer)
	if err != nil {
		return "", err
	}

	if RespDataType(buffer[0]) != RespSimpleString {
		return "", ErrDifferentTypes
	}

	state := consumeChar
	buffer[0] = 0
	result := ""

	for state != terminate {
		_, err := r.Read(buffer)
		if err != nil {
			return "", err
		}

		switch state {
		case consumeChar:
			if buffer[0] == '\r' {
				state = consumeTerminalByte
				continue
			}
			result += string(buffer)

		case consumeTerminalByte:
			if buffer[0] != '\n' {
				return "", ErrWrongTerminalByte
			}
			state = terminate
		}
	}

	return result, nil
}

type ErrorParser struct{}

func (p *ErrorParser) Parse(r io.Reader) (string, error) {
	buffer := make([]byte, 1)
	_, err := r.Read(buffer)
	if err != nil {
		return "", err
	}

	if RespDataType(buffer[0]) != RespError {
		return "", ErrDifferentTypes
	}

	state := consumeChar
	buffer[0] = 0
	result := ""

	for state != terminate {
		_, err := r.Read(buffer)
		if err != nil {
			return "", err
		}

		switch state {
		case consumeChar:
			if buffer[0] == '\r' {
				state = consumeTerminalByte
				continue
			}
			result += string(buffer)

		case consumeTerminalByte:
			if buffer[0] != '\n' {
				return "", ErrWrongTerminalByte
			}
			state = terminate
		}
	}

	return result, nil
}
