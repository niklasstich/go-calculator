package util

import "strconv"

type TokenType = int
type Op = int32

const (
	TokenTypeOperand TokenType = iota
	TokenTypeOperator
)

const (
	OpLeftBracket Op = iota
	OpRightBracket
	OpFactorial
	OpExponentiation
	OpMultiplication
	OpDivision
	OpAddition
	OpSubtraction
)

// Operator contains Op which can be used to quickly get the type of operation, Char which is its textual representation
// as a character, a Precedence from 1 to 5, whether the operation is LeftAssociative and whether the Operator is
// a Bracket or not.
type Operator struct {
	Op
	Char            int32
	Precedence      int
	LeftAssociative bool
	Bracket         bool
}

func (o Operator) String() string {
	return string(o.Char)
}

// Token contains a TokenType, which denotes the type of the token, either TokenTypeOperand or TokenTypeOperator.
// Depending on this, either TokenOperand or TokenOperator can be expected to have valid values.
type Token struct {
	TokenType
	TokenOperator *Operator
	TokenOperand  float64
}

func (t Token) String() string {
	if t.TokenType == TokenTypeOperand {
		return strconv.FormatFloat(t.TokenOperand, 'g', -1, 64)
	} else {
		return t.TokenOperator.String()
	}
}
