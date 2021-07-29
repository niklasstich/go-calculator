package parser

import (
	"errors"
	"fmt"
)

type tokenStack struct {
	top *tokenWrapper
}

type tokenWrapper struct {
	prev *tokenWrapper
	Token
}

func (stack *tokenStack) Push(t Token) {
	wrap := tokenWrapper{
		prev:  stack.top,
		Token: t,
	}
	stack.top = &wrap
}

func (stack *tokenStack) Peek() *Token {
	if stack.top == nil {
		return nil
	}
	return &stack.top.Token
}

func (stack *tokenStack) Pop() (t *Token) {
	if stack.top == nil {
		return nil
	}
	t = &stack.top.Token
	stack.top = stack.top.prev
	return
}

func (stack *tokenStack) HasElements() bool {
	return stack.top != nil
}

var ErrUnmatchedParenthesis = errors.New("there were unmatched parenthesis in the expression")

// ReformToRPN uses the Shunting-yard algorithm by Dijkstra to convert a tokenized infix expression to RPN
func ReformToRPN(expression []Token) (bnf []Token, err error) {
	bnf = make([]Token, 0, len(expression))
	opStack := tokenStack{}
	for _, t := range expression {
		if t.TokenType == TokenTypeOperand {
			//we can just push all operands straight to the output
			bnf = append(bnf, t)
		} else {
			o2 := opStack.Peek()
			switch t.TokenOperator.Op {
			case OpLeftBracket:
				{
					opStack.Push(t)
				}
			case OpRightBracket:
				{
					for {
						//if o2 is a left bracket, discard both brackets
						if o2.TokenOperator.Op == OpLeftBracket {
							opStack.Pop()
							break
						} else {
							//otherwise keep popping operators from the stack to the output until we find a left bracket
							opStack.Pop()
							if !opStack.HasElements() {
								return nil, fmt.Errorf("%v: Missing left bracket", ErrUnmatchedParenthesis)
							}
							bnf = append(bnf, *o2)
							o2 = opStack.Peek()
						}
					}
				}
			default:
				{
					//keep popping ops into output while:
					for o2 != nil && //there are ops on the stack
						!o2.TokenOperator.Bracket && //and they aren't brackets
						(o2.TokenOperator.Precedence > t.TokenOperator.Precedence || //and they have higher precedence or
							//they have the same precedence and are left associative
							(o2.TokenOperator.Precedence == t.TokenOperator.Precedence && o2.TokenOperator.LeftAssociative)) {
						opStack.Pop()
						bnf = append(bnf, *o2)
						o2 = opStack.Peek()
					}
					opStack.Push(t)
				}
			}
		}
	}
	//pop remaining operators
	for opStack.HasElements() {
		op := opStack.Pop()
		//if we still have a bracket on the op stack, we had mismatched parenthesis, missing a right bracket
		if op.TokenOperator.Bracket {
			return nil, fmt.Errorf("%v: Missing right bracket", ErrUnmatchedParenthesis)
		}
		bnf = append(bnf, *op)
	}
	return
}
