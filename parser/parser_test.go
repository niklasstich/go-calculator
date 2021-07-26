package parser

import (
	"reflect"
	"testing"
)

func TestParser(t *testing.T) {
	var tests = []struct {
		testName string
		input    []Token
		err      error
		want     []Token
	}{
		{
			"2+4",
			[]Token{
				{
					TokenType:    TokenTypeOperand,
					TokenOperand: 2,
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
					TokenOperand: 4,
				},
			},
			nil,
			[]Token{
				{
					TokenType:    TokenTypeOperand,
					TokenOperand: 2,
				},
				{
					TokenType:    TokenTypeOperand,
					TokenOperand: 4,
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
			}},

		{
			"2*(4+3.7)",
			[]Token{
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
					TokenOperand: 4,
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
					TokenOperand: 3.7,
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
			},
			nil,
			[]Token{
				{
					TokenType:    TokenTypeOperand,
					TokenOperand: 2,
				},
				{
					TokenType:    TokenTypeOperand,
					TokenOperand: 4,
				},
				{
					TokenType:    TokenTypeOperand,
					TokenOperand: 3.7,
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
					TokenType: TokenTypeOperator,
					TokenOperator: &Operator{
						Op:              OpMultiplication,
						Char:            '*',
						Precedence:      2,
						LeftAssociative: true,
						Bracket:         false,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			got, err := ReformToBNF(tt.input)

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
