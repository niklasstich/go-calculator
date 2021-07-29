package util

type TokenStack struct {
	top *tokenWrapper
}

type tokenWrapper struct {
	prev *tokenWrapper
	Token
}

func (stack *TokenStack) Push(t Token) {
	wrap := tokenWrapper{
		prev:  stack.top,
		Token: t,
	}
	stack.top = &wrap
}

func (stack *TokenStack) Peek() *Token {
	if stack.top == nil {
		return nil
	}
	return &stack.top.Token
}

func (stack *TokenStack) Pop() (t *Token) {
	if stack.top == nil {
		return nil
	}
	t = &stack.top.Token
	stack.top = stack.top.prev
	return
}

func (stack *TokenStack) HasElements() bool {
	return stack.top != nil
}
