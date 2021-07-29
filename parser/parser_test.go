package parser

import (
	"fmt"
	"github.com/niklasstich/calculator/util"
	"reflect"
	"testing"
)

func TestParser(t *testing.T) {
	var tests = []struct {
		testName string
		input    []util.Token
		err      error
		want     []util.Token
	}{
		{
			"2+4",
			[]util.Token{
				{
					TokenType:    util.TokenTypeOperand,
					TokenOperand: 2,
				},
				{
					TokenType: util.TokenTypeOperator,
					TokenOperator: &util.Operator{
						Op:              util.OpAddition,
						Char:            '+',
						Precedence:      1,
						LeftAssociative: true,
						Bracket:         false,
					},
				},
				{
					TokenType:    util.TokenTypeOperand,
					TokenOperand: 4,
				},
			},
			nil,
			[]util.Token{
				{
					TokenType:    util.TokenTypeOperand,
					TokenOperand: 2,
				},
				{
					TokenType:    util.TokenTypeOperand,
					TokenOperand: 4,
				},
				{
					TokenType: util.TokenTypeOperator,
					TokenOperator: &util.Operator{
						Op:              util.OpAddition,
						Char:            '+',
						Precedence:      1,
						LeftAssociative: true,
						Bracket:         false,
					},
				},
			},
		},
		{
			"2+4*3",
			[]util.Token{
				{
					TokenType:    util.TokenTypeOperand,
					TokenOperand: 2,
				},
				{
					TokenType: util.TokenTypeOperator,
					TokenOperator: &util.Operator{
						Op:              util.OpAddition,
						Char:            '+',
						Precedence:      1,
						LeftAssociative: true,
						Bracket:         false,
					},
				},
				{
					TokenType:    util.TokenTypeOperand,
					TokenOperand: 4,
				},
				{
					TokenType: util.TokenTypeOperator,
					TokenOperator: &util.Operator{
						Op:              util.OpMultiplication,
						Char:            '*',
						Precedence:      2,
						LeftAssociative: true,
						Bracket:         false,
					},
				},
				{
					TokenType:    util.TokenTypeOperand,
					TokenOperand: 3,
				},
			},
			nil,
			[]util.Token{
				{
					TokenType:    util.TokenTypeOperand,
					TokenOperand: 2,
				},
				{
					TokenType:    util.TokenTypeOperand,
					TokenOperand: 4,
				},
				{
					TokenType:    util.TokenTypeOperand,
					TokenOperand: 3,
				},
				{
					TokenType: util.TokenTypeOperator,
					TokenOperator: &util.Operator{
						Op:              util.OpMultiplication,
						Char:            '*',
						Precedence:      2,
						LeftAssociative: true,
						Bracket:         false,
					},
				},
				{
					TokenType: util.TokenTypeOperator,
					TokenOperator: &util.Operator{
						Op:              util.OpAddition,
						Char:            '+',
						Precedence:      1,
						LeftAssociative: true,
						Bracket:         false,
					},
				},
			},
		},
		{
			"2*4*4+3",
			[]util.Token{
				{
					TokenType:    util.TokenTypeOperand,
					TokenOperand: 2,
				},
				{
					TokenType: util.TokenTypeOperator,
					TokenOperator: &util.Operator{
						Op:              util.OpMultiplication,
						Char:            '*',
						Precedence:      2,
						LeftAssociative: true,
						Bracket:         false,
					},
				},
				{
					TokenType:    util.TokenTypeOperand,
					TokenOperand: 4,
				},
				{
					TokenType: util.TokenTypeOperator,
					TokenOperator: &util.Operator{
						Op:              util.OpMultiplication,
						Char:            '*',
						Precedence:      2,
						LeftAssociative: true,
						Bracket:         false,
					},
				},
				{
					TokenType:    util.TokenTypeOperand,
					TokenOperand: 4,
				},
				{
					TokenType: util.TokenTypeOperator,
					TokenOperator: &util.Operator{
						Op:              util.OpAddition,
						Char:            '+',
						Precedence:      1,
						LeftAssociative: true,
						Bracket:         false,
					},
				},
				{
					TokenType:    util.TokenTypeOperand,
					TokenOperand: 3,
				},
			},
			nil,
			[]util.Token{
				{
					TokenType:    util.TokenTypeOperand,
					TokenOperand: 2,
				},
				{
					TokenType:    util.TokenTypeOperand,
					TokenOperand: 4,
				},
				{
					TokenType: util.TokenTypeOperator,
					TokenOperator: &util.Operator{
						Op:              util.OpMultiplication,
						Char:            '*',
						Precedence:      2,
						LeftAssociative: true,
						Bracket:         false,
					},
				},
				{
					TokenType:    util.TokenTypeOperand,
					TokenOperand: 4,
				},
				{
					TokenType: util.TokenTypeOperator,
					TokenOperator: &util.Operator{
						Op:              util.OpMultiplication,
						Char:            '*',
						Precedence:      2,
						LeftAssociative: true,
						Bracket:         false,
					},
				},
				{
					TokenType:    util.TokenTypeOperand,
					TokenOperand: 3,
				},
				{
					TokenType: util.TokenTypeOperator,
					TokenOperator: &util.Operator{
						Op:              util.OpAddition,
						Char:            '+',
						Precedence:      1,
						LeftAssociative: true,
						Bracket:         false,
					},
				},
			},
		},

		{
			"2*(4+3.7)",
			[]util.Token{
				{
					TokenType:    util.TokenTypeOperand,
					TokenOperand: 2,
				},
				{
					TokenType: util.TokenTypeOperator,
					TokenOperator: &util.Operator{
						Op:              util.OpMultiplication,
						Char:            '*',
						Precedence:      2,
						LeftAssociative: true,
						Bracket:         false,
					},
				},
				{
					TokenType: util.TokenTypeOperator,
					TokenOperator: &util.Operator{
						Op:              util.OpLeftBracket,
						Char:            '(',
						Precedence:      5,
						LeftAssociative: false,
						Bracket:         true,
					},
				},
				{
					TokenType:    util.TokenTypeOperand,
					TokenOperand: 4,
				},
				{
					TokenType: util.TokenTypeOperator,
					TokenOperator: &util.Operator{
						Op:              util.OpAddition,
						Char:            '+',
						Precedence:      1,
						LeftAssociative: true,
						Bracket:         false,
					},
				},
				{
					TokenType:    util.TokenTypeOperand,
					TokenOperand: 3.7,
				},
				{
					TokenType: util.TokenTypeOperator,
					TokenOperator: &util.Operator{
						Op:              util.OpRightBracket,
						Char:            ')',
						Precedence:      5,
						LeftAssociative: false,
						Bracket:         true,
					},
				},
			},
			nil,
			[]util.Token{
				{
					TokenType:    util.TokenTypeOperand,
					TokenOperand: 2,
				},
				{
					TokenType:    util.TokenTypeOperand,
					TokenOperand: 4,
				},
				{
					TokenType:    util.TokenTypeOperand,
					TokenOperand: 3.7,
				},
				{
					TokenType: util.TokenTypeOperator,
					TokenOperator: &util.Operator{
						Op:              util.OpAddition,
						Char:            '+',
						Precedence:      1,
						LeftAssociative: true,
						Bracket:         false,
					},
				},
				{
					TokenType: util.TokenTypeOperator,
					TokenOperator: &util.Operator{
						Op:              util.OpMultiplication,
						Char:            '*',
						Precedence:      2,
						LeftAssociative: true,
						Bracket:         false,
					},
				},
			},
		},

		{
			"(2+4",
			[]util.Token{
				{
					TokenType: util.TokenTypeOperator,
					TokenOperator: &util.Operator{
						Op:              util.OpLeftBracket,
						Char:            '(',
						Precedence:      5,
						LeftAssociative: false,
						Bracket:         true,
					},
				},
				{
					TokenType:    util.TokenTypeOperand,
					TokenOperand: 2,
				},
				{
					TokenType: util.TokenTypeOperator,
					TokenOperator: &util.Operator{
						Op:              util.OpAddition,
						Char:            '+',
						Precedence:      1,
						LeftAssociative: true,
						Bracket:         false,
					},
				},
				{
					TokenType:    util.TokenTypeOperand,
					TokenOperand: 4,
				},
			},
			fmt.Errorf("%v: Missing right bracket", ErrUnmatchedParenthesis),
			nil,
		},
		{
			"2+4)",
			[]util.Token{
				{
					TokenType:    util.TokenTypeOperand,
					TokenOperand: 2,
				},
				{
					TokenType: util.TokenTypeOperator,
					TokenOperator: &util.Operator{
						Op:              util.OpAddition,
						Char:            '+',
						Precedence:      1,
						LeftAssociative: true,
						Bracket:         false,
					},
				},
				{
					TokenType:    util.TokenTypeOperand,
					TokenOperand: 4,
				},
				{
					TokenType: util.TokenTypeOperator,
					TokenOperator: &util.Operator{
						Op:              util.OpRightBracket,
						Char:            ')',
						Precedence:      5,
						LeftAssociative: false,
						Bracket:         true,
					},
				},
			},
			fmt.Errorf("%v: Missing left bracket", ErrUnmatchedParenthesis),
			nil,
		},
	}

	for i, tt := range tests {
		testname := fmt.Sprintf("%d: %s", i+1, tt.testName)
		t.Run(testname, func(t *testing.T) {
			got, err := ReformToRPN(tt.input)

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

func TestTokenizerIntoParser(t *testing.T) {
	var tests = []struct {
		input, want string
		err1, err2  error
	}{
		{
			"3+4*2/(1-5)^2^3",
			"[3 4 2 * 1 5 - 2 3 ^ ^ / +]",
			nil, nil,
		},
		//TODO: some more cases here, longer and more complex inputs
	}

	for i, tt := range tests {
		testname := fmt.Sprintf("%d: %s", i+1, tt.input)
		t.Run(testname, func(t *testing.T) {
			tokens, err := TokenizeString(tt.input)
			if err == nil && tt.err1 != nil {
				t.Errorf("Unexpected error, expected nil %v", err)
			}
			if err != nil && tt.err1 == nil {
				t.Errorf("Expected error %v, got nil", err)
			}
			if tt.err1 != nil && err != nil && tt.err1.Error() != err.Error() {
				t.Errorf("Expected error %v, got %v", tt.err1, err)
			}

			tokens, err = ReformToRPN(tokens)
			if err == nil && tt.err2 != nil {
				t.Errorf("Unexpected error, expected nil %v", err)
			}
			if err != nil && tt.err2 == nil {
				t.Errorf("Expected error %v, got nil", err)
			}
			if tt.err2 != nil && err != nil && tt.err2.Error() != err.Error() {
				t.Errorf("Expected error %v, got %v", tt.err2, err)
			}

			got := fmt.Sprint(tokens)
			if got != tt.want {
				t.Errorf("Wanted %s, got %s", tt.want, got)
			}
		})
	}
}
