package evaluation

import (
	"fmt"
	"github.com/niklasstich/calculator/parser"
	"github.com/niklasstich/calculator/util"
	"testing"
)

//integration
func TestEvaluateRPNExpression(t *testing.T) {
	var tests = []struct {
		expr parser.RPNExpression
		want string
		err  error
	}{
		{
			expr: []util.Token{
				{
					TokenType:    util.TokenTypeOperand,
					TokenOperand: 4,
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
			},
			want: "6",
			err:  nil,
		},
	}

	for i, tt := range tests {
		testName := fmt.Sprintf("%d:%v", i, tt.expr)
		t.Run(testName, func(t *testing.T) {
			got, err := EvaluateRPNExpression(tt.expr)

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

			if got.String() != tt.want {
				t.Errorf("Expected %s, got %s", tt.want, got)
			}
		})
	}
}

func TestTokenizerIntoParserIntoEvaluation(t *testing.T) {
	var tests = []struct {
		input, result string
	}{
		{
			"1+2+3",
			"6",
		},
		{
			"9+10",
			"19",
		},
		{
			"3+4*2/(1-5)^2",
			"3.5",
		},
		{
			"3    +   3",
			"6",
		},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("%s=%s", tt.input, tt.result)
		t.Run(testName, func(t *testing.T) {
			tokens, err := parser.TokenizeString(tt.input)
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			tokens, err = parser.ReformToRPN(tokens)
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			result, err := EvaluateRPNExpression(tokens)
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			resultString := fmt.Sprint(result)
			if resultString != tt.result {
				t.Errorf("Expected %s, got %s", tt.result, resultString)
			}
		})
	}
}
