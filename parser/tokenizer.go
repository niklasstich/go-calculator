package parser

import (
	"errors"
	"fmt"
	"github.com/niklasstich/calculator/util"
	"strconv"
)

var ErrInvalidToken = errors.New("expression contains invalid token")

var opLookUp = map[int32]*util.Operator{
	'!': {
		Char:            '!',
		Precedence:      4,
		LeftAssociative: true,
		Bracket:         false,
		Op:              util.OpFactorial,
	},
	'(': {
		Char:            '(',
		Precedence:      5,
		LeftAssociative: false,
		Bracket:         true,
		Op:              util.OpLeftBracket,
	},
	')': {
		Char:            ')',
		Precedence:      5,
		LeftAssociative: false,
		Bracket:         true,
		Op:              util.OpRightBracket,
	},
	'^': {
		Char:            '^',
		Precedence:      3,
		LeftAssociative: false,
		Bracket:         false,
		Op:              util.OpExponentiation,
	},
	'*': {
		Char:            '*',
		Precedence:      2,
		LeftAssociative: true,
		Bracket:         false,
		Op:              util.OpMultiplication,
	},
	'/': {
		Char:            '/',
		Precedence:      2,
		LeftAssociative: true,
		Bracket:         false,
		Op:              util.OpDivision,
	},
	'+': {
		Char:            '+',
		Precedence:      1,
		LeftAssociative: true,
		Bracket:         false,
		Op:              util.OpAddition,
	},
	'-': {
		Char:            '-',
		Precedence:      1,
		LeftAssociative: true,
		Bracket:         false,
		Op:              util.OpSubtraction,
	},
}

// TokenizeString takes a string in arbitrary notation (infix, polish, reverse polish) and returns a slice of Token
func TokenizeString(input string) (tokens []util.Token, err error) {
	//prepare return value
	tokens = make([]util.Token, 0, 20)
	var numbuf string
	numQueued := false
	//iterate over all characters in the string, see if they are numerical or an operator
	for i, c := range input {
		if isNumerical(c) || isDot(c) {
			//append new digit and remember that we have a number queued
			numbuf += string(c)
			numQueued = true
		} else {
			//if numQueued is true, we need to push a new symbol with the number first
			if numQueued {
				numQueued = false
				num, err := strconv.ParseFloat(numbuf, 64)
				if err != nil {
					return nil, fmt.Errorf("malformed expression near '%c': %w", c, err)
				}
				tokens = append(tokens, util.Token{
					TokenType:    util.TokenTypeOperand,
					TokenOperand: num,
				})
				numbuf = ""
			}
			if c == ' ' || c == '\n' {
				continue
			}
			operator := opLookUp[c]
			if operator == nil {
				return nil, fmt.Errorf("%w: %c at pos %d", ErrInvalidToken, c, i)
			}
			tokens = append(tokens, util.Token{
				TokenType:     util.TokenTypeOperator,
				TokenOperator: operator,
			})
		}
	}
	//check for remaining number
	if numQueued {
		num, err := strconv.ParseFloat(numbuf, 64)
		if err != nil {
			return nil, fmt.Errorf("found malformed expression while cleaning up: %w", err)
		}
		tokens = append(tokens, util.Token{
			TokenType:    util.TokenTypeOperand,
			TokenOperand: num,
		})
	}
	return
}

func isDot(c int32) bool {
	return c == '.'

}

func isNumerical(c int32) bool {
	return c >= '0' && c <= '9'
}
