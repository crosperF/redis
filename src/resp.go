package main

import (
	"bufio"
	"io"
	"strconv"
)

const (
	STRING  = "+"
	ERROR   = '-'
	INTEGER = ':'
	BULK    = '$'
	ARRAY   = '*'
)

type Value struct {
	typ   string
	str   string
	num   int
	bulk  string
	array []Value
}

type Resp struct {
	reader *bufio.Reader
}

func (r *Resp) readLine() (line []byte, n int, err error) {
	for {
		by, e := r.reader.ReadByte()
		if e != nil {
			return nil, n_out, err
		}
		n += 1
		line = append(line, by)

		if by == '\r' {
			n -= 1
			break
		}
	}
	return line[:len(line)-1], n, nil
}

func (r *Resp) readInteger() (x int, n int, err error) {
	byt, n, err := r.readLine()

	if err != nil {
		return 0, 0, err
	}

	i64, err := strconv.ParseInt(string(byt), 10, 64)

	if err != nil {
		return 0, n, err
	}

	return int(i64), n, nil
}

func newResp(read io.Reader) *Resp {
	return &Resp{reader: bufio.NewReader(read)}
}
