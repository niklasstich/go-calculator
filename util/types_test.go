package util

import (
	"fmt"
	"testing"
)

func TestStringer(t *testing.T) {
	var tests = []struct {
		d    interface{}
		want string
	}{
		{
			Token{
				TokenType:     TokenTypeOperand,
				TokenOperator: nil,
				TokenOperand:  5,
			}, "5",
		},

		{
			Token{
				TokenType:     TokenTypeOperand,
				TokenOperator: nil,
				TokenOperand:  1234.5678901234567890,
			}, "1234.567890123457",
		},

		{
			Token{
				TokenType: TokenTypeOperator,
				TokenOperator: &Operator{
					Op:              OpAddition,
					Char:            '+',
					Precedence:      1,
					LeftAssociative: true,
					Bracket:         false,
				},
			}, "+",
		},

		{
			[]Token{
				{
					TokenType:    TokenTypeOperand,
					TokenOperand: 1,
				},
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
					TokenType:    TokenTypeOperand,
					TokenOperand: 2,
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
					TokenOperand: 3,
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
					TokenType:    TokenTypeOperand,
					TokenOperand: 4,
				},
			}, "[1 + 2 * 3 - 4]",
		},
	}

	for i, tt := range tests {
		testname := fmt.Sprintf("%d: %v", i+1, tt.want)
		t.Run(testname, func(t *testing.T) {
			got := fmt.Sprint(tt.d)
			if got != tt.want {
				t.Errorf("Wanted %s, got %s", tt.want, got)
			}
		})
	}
}
