package parser

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func TestTokenizer(t *testing.T) {
	//our testcases
	var tests = []struct {
		input string
		err   error
		want  []Token
	}{
		{"", nil, []Token{}},

		{"123", nil, []Token{{
			TokenType:    TokenTypeOperand,
			TokenOperand: 123,
		}}},

		{"777.7742", nil, []Token{{
			TokenType:    TokenTypeOperand,
			TokenOperand: 777.7742,
		}}},

		{"456 789 123.456", nil, []Token{
			{
				TokenType:    TokenTypeOperand,
				TokenOperand: 456,
			},
			{
				TokenType:    TokenTypeOperand,
				TokenOperand: 789,
			},
			{
				TokenType:    TokenTypeOperand,
				TokenOperand: 123.456,
			},
		}},

		{"+", nil, []Token{{
			TokenType: TokenTypeOperator,
			TokenOperator: &Operator{
				Op:              OpAddition,
				Char:            '+',
				Precedence:      1,
				LeftAssociative: true,
				Bracket:         false,
			},
		}}},

		{"+ - (456.789 * 123) / ^ !", nil, []Token{
			{
				TokenType: TokenTypeOperator,
				TokenOperator: &Operator{
					Op:              OpAddition,
					Char:            '+',
					Precedence:      1,
					LeftAssociative: true,
					Bracket:         false,
				},
			},
			{
				TokenType: TokenTypeOperator,
				TokenOperator: &Operator{
					Op:              OpSubtraction,
					Char:            '-',
					Precedence:      1,
					LeftAssociative: true,
					Bracket:         false,
				},
			},
			{
				TokenType: TokenTypeOperator,
				TokenOperator: &Operator{
					Op:              OpLeftBracket,
					Char:            '(',
					Precedence:      5,
					LeftAssociative: false,
					Bracket:         true,
				},
			},
			{
				TokenType:    TokenTypeOperand,
				TokenOperand: 456.789,
			},
			{
				TokenType: TokenTypeOperator,
				TokenOperator: &Operator{
					Op:              OpMultiplication,
					Char:            '*',
					Precedence:      2,
					LeftAssociative: true,
					Bracket:         false,
				},
			},
			{
				TokenType:    TokenTypeOperand,
				TokenOperand: 123,
			},
			{
				TokenType: TokenTypeOperator,
				TokenOperator: &Operator{
					Op:              OpRightBracket,
					Char:            ')',
					Precedence:      5,
					LeftAssociative: false,
					Bracket:         true,
				},
			},
			{
				TokenType: TokenTypeOperator,
				TokenOperator: &Operator{
					Op:              OpDivision,
					Char:            '/',
					Precedence:      2,
					LeftAssociative: true,
					Bracket:         false,
				},
			},
			{
				TokenType: TokenTypeOperator,
				TokenOperator: &Operator{
					Op:              OpExponentiation,
					Char:            '^',
					Precedence:      3,
					LeftAssociative: false,
					Bracket:         false,
				},
			},
			{
				TokenType: TokenTypeOperator,
				TokenOperator: &Operator{
					Op:              OpFactorial,
					Char:            '!',
					Precedence:      4,
					LeftAssociative: true,
					Bracket:         false,
				},
			},
		}},

		{"@", fmt.Errorf("%w: %c at pos %d", ErrInvalidToken, '@', 0), []Token{}},

		{"2..", errors.New("found malformed expression while cleaning up: strconv.ParseFloat: parsing " +
			"\"2..\": invalid syntax"), []Token{}},

		{"2.. 4 7 3", errors.New("malformed expression near ' ': strconv.ParseFloat: parsing " +
			"\"2..\": invalid syntax"), []Token{}},
	}

	for i, tt := range tests {
		testname := fmt.Sprintf("%d:\"%s\"", i+1, tt.input)
		t.Run(testname, func(t *testing.T) {
			got, err := TokenizeString(tt.input)

			//errors
			if tt.err == nil && err != nil {
				t.Errorf("Unexpected error, expected nil: %v", err)
			}
			if tt.err != nil && err == nil {
				t.Errorf("Expected error %v, got nil", tt.err)
			}
			if tt.err != nil && err != nil && tt.err.Error() != err.Error() {
				t.Errorf("Expected error %v, got %v", tt.err, err)
			}

			//ret values
			if len(tt.want) != len(got) {
				t.Errorf("Expected slice of len %d, got len %d", len(tt.want), len(got))
			}
			for i, w := range tt.want {
				got := got[i]
				if got.TokenType != w.TokenType {
					t.Errorf("i=%d:Expected TokenType %d, got %d", i, w.TokenType, got.TokenType)
				}
				if got.TokenOperand != w.TokenOperand {
					t.Errorf("i=%d:Expected TokenOperand %f, got %f", i, w.TokenOperand, got.TokenOperand)
				}
				if !reflect.DeepEqual(got.TokenOperator, w.TokenOperator) {
					t.Errorf("i=%d:Expected TokenOperator %v, got %v", i, w.TokenOperator, got.TokenOperator)
				}
			}
		})
	}
}
