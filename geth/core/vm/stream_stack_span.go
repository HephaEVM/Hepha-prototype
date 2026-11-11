package vm

import "github.com/holiman/uint256"

// Only use these prices in the canonical Prague operation table. A failed
// precondition leaves its instruction to the stock interpreter at the original PC.
var stackSpanStart [256]bool
var stackSpanRules *JumpTable

func init() {
	stackSpanRules = &pragueInstructionSet
	for i := PUSH0; i <= SWAP16; i++ {
		stackSpanStart[i] = true
	}
	for _, o := range []OpCode{POP, JUMPDEST, JUMP, JUMPI, ADD, SUB, MUL, AND, OR, XOR, NOT, LT, GT, EQ, ISZERO, SHL, SHR} {
		stackSpanStart[o] = true
	}
}
func stackSpanDestination(c *Contract, d uint64) bool {
	if d >= uint64(len(c.Code)) || OpCode(c.Code[d]) != JUMPDEST {
		return false
	}
	if c.analysis != nil {
		return c.analysis.codeSegment(d)
	}
	return c.isCode(d)
}

// executeStackSpan commits the valid prefix when it reaches an unsupported opcode,
// an entry marker, or an exceptional precondition. On the latter the ordinary
// interpreter resumes at the failing PC with every earlier charge intact.
func executeStackSpan(e *EVM, c *Contract, stack *Stack, view []byte, pc uint64) (uint64, bool) {
	data := stack.data
	gas := c.Gas
	var executed uint64
loop:
	for pc < uint64(len(view)) {
		op := OpCode(view[pc])
		n := len(data)
		switch {
		case op >= PUSH1 && op <= PUSH32:
			if n >= 1024 || gas < 3 {
				break loop
			}
			width := uint64(op-PUSH1) + 1
			end := pc + 1 + width
			if end > uint64(len(c.Code)) {
				break loop
			}
			arg := c.Code[pc+1 : end]
			if width <= 2 && end < uint64(len(view)) {
				branch := OpCode(view[end])
				if (branch == JUMP || branch == JUMPI && n > 0) && !e.abort.Load() {
					cost := uint64(11)
					taken := true
					if branch == JUMPI {
						cost = 13
						taken = !data[n-1].IsZero()
					}
					d := uint64(arg[0])
					if width == 2 {
						d = d<<8 | uint64(arg[1])
					}
					if gas >= cost && (!taken || stackSpanDestination(c, d)) {
						if branch == JUMPI {
							data = data[:n-1]
						}
						pc = end + 1
						if taken {
							pc = d
						}
						gas -= cost
						if paperCountInstructions {
							executed += 2
						}
						continue
					}
				}
			}
			var x uint256.Int
			switch width {
			case 1:
				x.SetUint64(uint64(arg[0]))
			case 2:
				x.SetBytes2(arg)
			default:
				x.SetBytes(arg)
			}
			data = append(data, x)
			gas -= 3
			pc = end
			if paperCountInstructions {
				executed++
			}
			continue
		case op >= DUP1 && op <= DUP16:
			src := n - 1 - int(op-DUP1)
			if src < 0 || n >= 1024 || gas < 3 {
				break loop
			}
			data = append(data, data[src])
			gas -= 3
		case op >= SWAP1 && op <= SWAP16:
			src := n - 2 - int(op-SWAP1)
			if src < 0 || gas < 3 {
				break loop
			}
			data[n-1], data[src] = data[src], data[n-1]
			gas -= 3
		default:
			switch op {
			case JUMPDEST:
				if gas < 1 {
					break loop
				}
				gas--
			case PUSH0:
				if n >= 1024 || gas < 2 {
					break loop
				}
				data = append(data, uint256.Int{})
				gas -= 2
			case POP:
				if n < 1 || gas < 2 {
					break loop
				}
				data = data[:n-1]
				gas -= 2
			case JUMP, JUMPI:
				pop := 1
				cost := uint64(8)
				taken := true
				if op == JUMPI {
					pop = 2
					cost = 10
				}
				if n < pop || gas < cost || e.abort.Load() {
					break loop
				}
				if op == JUMPI {
					taken = !data[n-2].IsZero()
				}
				d, overflow := data[n-1].Uint64WithOverflow()
				if taken && (overflow || !stackSpanDestination(c, d)) {
					break loop
				}
				data = data[:n-pop]
				gas -= cost
				if taken {
					pc = d
					if paperCountInstructions {
						executed++
					}
					continue
				}
			case NOT, ISZERO:
				if n == 0 || gas < 3 {
					break loop
				}
				x := &data[n-1]
				if op == NOT {
					x.Not(x)
				} else if x.IsZero() {
					x.SetOne()
				} else {
					x.Clear()
				}
				gas -= 3
			case ADD, SUB, MUL, AND, OR, XOR, LT, GT, EQ, SHL, SHR:
				cost := uint64(3)
				if op == MUL {
					cost = 5
				}
				if n < 2 || gas < cost {
					break loop
				}
				x, y := &data[n-1], &data[n-2]
				switch op {
				case ADD:
					y.Add(x, y)
				case SUB:
					y.Sub(x, y)
				case MUL:
					y.Mul(x, y)
				case AND:
					y.And(x, y)
				case OR:
					y.Or(x, y)
				case XOR:
					y.Xor(x, y)
				case LT:
					if x.Lt(y) {
						y.SetOne()
					} else {
						y.Clear()
					}
				case GT:
					if x.Gt(y) {
						y.SetOne()
					} else {
						y.Clear()
					}
				case EQ:
					if x.Eq(y) {
						y.SetOne()
					} else {
						y.Clear()
					}
				case SHL:
					if x.LtUint64(256) {
						y.Lsh(y, uint(x.Uint64()))
					} else {
						y.Clear()
					}
				case SHR:
					if x.LtUint64(256) {
						y.Rsh(y, uint(x.Uint64()))
					} else {
						y.Clear()
					}
				}
				data = data[:n-1]
				gas -= cost
			default:
				break loop
			}
		}
		pc++
		if paperCountInstructions {
			executed++
		}
	}
	// Every admitted opcode charges positive gas. This progress check also works on 32-bit hosts.
	moved := gas != c.Gas
	stack.data = data
	c.Gas = gas
	if paperCountInstructions {
		paperInterpreted += executed
	}
	return pc, moved
}
