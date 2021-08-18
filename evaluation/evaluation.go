package evaluation

import (
	"errors"
	"fmt"
	"github.com/niklasstich/calculator/parser"
	"github.com/niklasstich/calculator/util"
	"math"
)

var ErrDivByZero = errors.New("division by 0")
var ErrNotImplemented = errors.New("factorial is not yet implemented")
var ErrInvalidExpression = errors.New("provided expression is not valid")

var funcLookup = map[util.Op]func(stack *util.TokenStack) error{
	util.OpAddition: func(stack *util.TokenStack) error {
		op1 := stack.Pop()
		op2 := stack.Pop()
		stack.Push(util.Token{
			TokenType:    util.TokenTypeOperand,
			TokenOperand: op1.TokenOperand + op2.TokenOperand,
		})
		return nil
	},
	util.OpSubtraction: func(stack *util.TokenStack) error {
		op1 := stack.Pop()
		op2 := stack.Pop()
		stack.Push(util.Token{
			TokenType:    util.TokenTypeOperand,
			TokenOperand: op2.TokenOperand - op1.TokenOperand,
		})
		return nil
	},
	util.OpMultiplication: func(stack *util.TokenStack) error {
		op1 := stack.Pop()
		op2 := stack.Pop()
		stack.Push(util.Token{
			TokenType:    util.TokenTypeOperand,
			TokenOperand: op1.TokenOperand * op2.TokenOperand,
		})
		return nil
	},
	util.OpDivision: func(stack *util.TokenStack) error {
		op1 := stack.Pop()
		op2 := stack.Pop()
		if op2.TokenOperand == 0 {
			return ErrDivByZero
		}
		stack.Push(util.Token{
			TokenType:    util.TokenTypeOperand,
			TokenOperand: op2.TokenOperand / op1.TokenOperand,
		})
		return nil
	},
	util.OpExponentiation: func(stack *util.TokenStack) error {
		op1 := stack.Pop()
		op2 := stack.Pop()
		stack.Push(util.Token{
			TokenType:    util.TokenTypeOperand,
			TokenOperand: math.Pow(op2.TokenOperand, op1.TokenOperand),
		})
		return nil
	},
	util.OpFactorial: func(stack *util.TokenStack) error {
		//factorial is more complicated than i thought, because we first need to assure that the token we pop is an int
		return ErrNotImplemented
	},
}

func EvaluateRPNExpression(expression parser.RPNExpression) (result *util.Token, err error) {
	stack := util.TokenStack{}
	for _, token := range expression {
		if token.TokenType == util.TokenTypeOperand {
			stack.Push(token)
		} else {
			err = funcLookup[token.TokenOperator.Op](&stack)
			if err != nil {
				return nil, err
			}
		}
	}

	//pop element, assure that it is an operand, otherwise our expression was malformed
	result = stack.Pop()
	if result.TokenType != util.TokenTypeOperand {
		return nil, fmt.Errorf("%v: Top token after expression evaluation was not an operand", ErrInvalidExpression)
	}

	//if there are still tokens on the stack, again the expression is malformed
	if stack.HasElements() {
		return nil, fmt.Errorf("%v: There were extra tokens on the stack after evaluation of expression", ErrInvalidExpression)
	}

	return
}
