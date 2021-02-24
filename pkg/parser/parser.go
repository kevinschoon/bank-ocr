package parser

import (
	// "bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"strings"
	"text/scanner"
)

var ErrBoundsInvalid = errors.New("document exceeds bounds")

const (
	zero = `
 _ 
| |
|_|
   
`

	one = `
   
  |
  |
   
`

	two = `
 _ 
 _|
|_ 
   
`

	three = `
 _ 
 _|
 _|
   
`
	four = `
   
|_|
  |
   
`
	five = `
 _ 
|_ 
 _|
   
`
	six = `
 _ 
|_ 
|_|
   
`
	seven = `
 _ 
  |
  |
   
`
	eight = `
 _ 
|_|
|_|
   
`
	nine = `
 _ 
|_|
 _|
   
`
)

var _ Parser = (*textParser)(nil)

// Parser parses a bank file
type Parser interface {
	// ReadAll reads all of the account numbers in a text file
	ReadAll(io.Reader) ([][9]int, error)
}

type textParser struct {
	numbers [][]byte
}

func (p textParser) readNumber(raw []byte) int {
	for i, other := range p.numbers {
		if bytes.Equal(raw, other) {
			return i
		}
	}
	return -1
}

func (p textParser) ReadAll(reader io.Reader) (accounts [][9]int, err error) {
	var s scanner.Scanner
	s.Init(reader)
	s.Whitespace = 1 << '\n' // ignore new lines
	digits := []*bytes.Buffer{
		bytes.NewBuffer(nil),
		bytes.NewBuffer(nil),
		bytes.NewBuffer(nil),
		bytes.NewBuffer(nil),
		bytes.NewBuffer(nil),
		bytes.NewBuffer(nil),
		bytes.NewBuffer(nil),
		bytes.NewBuffer(nil),
		bytes.NewBuffer(nil),
	}
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		// bounds checks
		if s.Column > 27 {
			return nil, ErrBoundsInvalid
		}
		switch s.Column {
		case 1, 2, 3:
			digits[0].WriteString(s.TokenText())
		case 4, 5, 6:
			digits[1].WriteString(s.TokenText())
		case 7, 8, 9:
			digits[2].WriteString(s.TokenText())
		case 10, 11, 12:
			digits[3].WriteString(s.TokenText())
		case 13, 14, 15:
			digits[4].WriteString(s.TokenText())
		case 16, 17, 18:
			digits[5].WriteString(s.TokenText())
		case 19, 20, 21:
			digits[6].WriteString(s.TokenText())
		case 22, 23, 24:
			digits[7].WriteString(s.TokenText())
		case 25, 26, 27:
			digits[8].WriteString(s.TokenText())
		}
		if s.Line%4 == 0 && s.Column == 27 {
			account := [9]int{
				p.readNumber(digits[0].Bytes()),
				p.readNumber(digits[1].Bytes()),
				p.readNumber(digits[2].Bytes()),
				p.readNumber(digits[3].Bytes()),
				p.readNumber(digits[4].Bytes()),
				p.readNumber(digits[5].Bytes()),
				p.readNumber(digits[6].Bytes()),
				p.readNumber(digits[7].Bytes()),
				p.readNumber(digits[8].Bytes()),
			}
			for i, entry := range account {
				if entry == -1 {
					return nil, fmt.Errorf("unable to parse number (%d)", i)
				}
			}
			accounts = append(accounts, account)
			for _, buf := range digits {
				buf.Reset()
			}
		}
	}
	return accounts, err
}

// New initializes a new Parser
func New() Parser {
	numbers := [][]byte{
		[]byte(strings.ReplaceAll(zero, "\n", "")),
		[]byte(strings.ReplaceAll(one, "\n", "")),
		[]byte(strings.ReplaceAll(two, "\n", "")),
		[]byte(strings.ReplaceAll(three, "\n", "")),
		[]byte(strings.ReplaceAll(four, "\n", "")),
		[]byte(strings.ReplaceAll(five, "\n", "")),
		[]byte(strings.ReplaceAll(six, "\n", "")),
		[]byte(strings.ReplaceAll(seven, "\n", "")),
		[]byte(strings.ReplaceAll(eight, "\n", "")),
		[]byte(strings.ReplaceAll(nine, "\n", "")),
	}
	return textParser{numbers}
}
