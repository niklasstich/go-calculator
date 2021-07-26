package parser

import "errors"

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

func (stack *tokenStack) Count() (i int) {
	p := stack.top
	for p != nil {
		p = p.prev
		i++
	}
	return
}

var ErrUnmatchedParenthesis = errors.New("there were unmatched parenthesis in the expression")

// ReformToBNF uses the Shunting-yard algorithm by Dijkstra to convert a tokenized infix expression to BNF
func ReformToBNF(expression []Token) (bnf []Token, err error) {
	bnf = make([]Token, 0, len(expression))
	opStack := tokenStack{}
	for _, t := range expression {
		switch t.TokenType {
		case TokenTypeOperand:
			{
				bnf = append(bnf, t)
				break
			}
		case TokenTypeOperator: //TODO: refactor this mess so it is easier to read and understand
			{
				o2 := opStack.Peek()
				if t.TokenOperator.Bracket {
					if t.TokenOperator.Op == OpLeftBracket {
						opStack.Push(t)
					} else if t.TokenOperator.Op == OpRightBracket {
						for {
							if o2.TokenOperator.Op == OpLeftBracket {
								opStack.Pop()
								break
							} else {
								opStack.Pop()
								if !opStack.HasElements() {
									return nil, ErrUnmatchedParenthesis
								}
								bnf = append(bnf, *o2)
								o2 = opStack.Peek()
							}
						}
					}
				} else {
					for o2 != nil && !o2.TokenOperator.Bracket && (o2.TokenOperator.Precedence > t.TokenOperator.Precedence ||
						(o2.TokenOperator.Precedence == t.TokenOperator.Precedence && o2.TokenOperator.LeftAssociative)) {
						//pop element and push to output
						opStack.Pop()
						bnf = append(bnf, *o2)
						o2 = opStack.Peek()
					}
					opStack.Push(t)
				}
			}
		}
	}
	for opStack.HasElements() {
		op := opStack.Pop()
		if op.TokenOperator.Bracket {
			return nil, ErrUnmatchedParenthesis
		}
		bnf = append(bnf, *op)
	}
	return
}
