package commandths

import (
	"fmt"

	"github.com/leep-frog/command"
)

type parserContext struct {
	head, cur *numericalTerm
	curOp     *operationTerm
	size      int
}

func (pc *parserContext) appendNumber(n int) {
	nt := &numericalTerm{
		value: n,
	}
	if pc.head == nil {
		pc.head = nt
		pc.cur = nt
		return
	}
	pc.curOp.next = nt
	nt.prev = pc.curOp
	pc.cur = nt
	pc.curOp = nil
}

func (pc *parserContext) appendOperation(op Operation[int]) {
	ot := &operationTerm{
		op:       op,
		position: pc.size,
	}
	pc.size++
	pc.cur.next = ot
	ot.prev = pc.cur
	pc.curOp = ot
	pc.cur = nil
}

func parse(seq *expressionSequence, inParen bool) (int, error) {

	var state parserState = &operationState{}
	ctx := &parserContext{}
	for !seq.done() {
		pt, err := toParserTerm(seq.next())
		if err != nil {
			return 0, err
		}

		if pt.symbolType == closeParenSymbol {
			if !inParen {
				return 0, fmt.Errorf("unexpected close parentheses")
			}
			break
		}

		if pt.symbolType == openParenSymbol {
			v, err := parse(seq, true)
			if err != nil {
				return 0, err
			}
			pt = &parserTerm{numberSymbol, nil, v}
		}

		// Symbol type guaranteed to either be number or operator at this point.
		state, err = state.processTerm(ctx, pt)
		if err != nil {
			return 0, err
		}
	}

	return ctx.head.evaluate(), nil
}

type symbolType int

const (
	openParenSymbol symbolType = iota
	closeParenSymbol
	operationSymbol
	numberSymbol
	// TODO evaluated number symbol [e.g. 4 + 5 (3 + 9) ===> 4 + 5 * 12 (not 4 + 5 12)]
)

type parserTerm struct {
	symbolType     symbolType
	operationValue Operation[int]
	numberValue    int
}

func toParserTerm(s string) (*parserTerm, error) {
	// Parse int first (incase we have a negative number)
	i, parseIntErr := command.ParseInt(s)
	if parseIntErr == nil {
		return &parserTerm{numberSymbol, nil, i}, nil
	}
	// Check if operation
	if op, ok := OperationMap[s]; ok {
		return &parserTerm{operationSymbol, op, 0}, nil
	}

	if s == "(" {
		return &parserTerm{openParenSymbol, nil, 0}, nil
	}
	if s == ")" {
		return &parserTerm{closeParenSymbol, nil, 0}, nil
	}

	// If here, then assume they have a misconfigured int
	return nil, fmt.Errorf("failed to parse int: %v", parseIntErr)
}

type expressionSequence struct {
	idx         int
	expressions []string
}

func newSequence(sl []string) *expressionSequence {
	return &expressionSequence{0, sl}
}

func (seq *expressionSequence) done() bool {
	return seq.idx >= len(seq.expressions)
}

func (seq *expressionSequence) next() string {
	seq.idx++
	return seq.expressions[seq.idx-1]
}

type parserState interface {
	processTerm(*parserContext, *parserTerm) (parserState, error)
}

type operationState struct{}

func (*operationState) processTerm(ctx *parserContext, term *parserTerm) (parserState, error) {
	if term.symbolType == operationSymbol {
		// We are
		// if slices.Contains(term.operationValue.Symbols(), "-") {
		// 	return nil,
		// }
		// if term.operationValue == m.
		// TODO: have error include term position and string value
		return nil, fmt.Errorf("consecutive operations")
	}
	if term.symbolType == closeParenSymbol {
		panic("RATS")
	}
	// Otherwise number symbol
	ctx.appendNumber(term.numberValue)
	return &numberState{}, nil
}

type numberState struct {
	negate bool
}

func (*numberState) processTerm(ctx *parserContext, term *parserTerm) (parserState, error) {
	if term.symbolType == numberSymbol {
		return nil, fmt.Errorf("unexpected number %d", term.numberValue)
	}
	// Otherwise operation symbol
	ctx.appendOperation(term.operationValue)
	return &operationState{}, nil
}
