package cbuf

type CBuf[T any] struct {
	getEnd int
	putEnd int
	size   int
	data   []*T
}

func NewCBuf[T any](size int) *CBuf[T] {
	return &CBuf[T]{0, 0, size, make([]*T, size)}
}
func (ctx *CBuf[T]) Push(value *T) bool {
	if ctx.IsFull() {
		return false
	}
	ctx.data[ctx.putEnd] = value
	ctx.putEnd = (ctx.putEnd + 1) % ctx.size
	return true
}

func (ctx *CBuf[T]) Peek() *T {
	if ctx.IsEmpty() {
		return nil
	}
	return ctx.data[ctx.getEnd]
}

func (ctx *CBuf[T]) NumEntries() int {
	if ctx.getEnd < ctx.putEnd {
		return ctx.putEnd - ctx.getEnd
	}
	return ctx.size - ctx.getEnd + ctx.putEnd
}

func (ctx *CBuf[T]) Pop() *T {
	if ctx.IsEmpty() {
		return nil
	}
	var val = ctx.data[ctx.getEnd]
	ctx.getEnd = (ctx.getEnd + 1) % ctx.size
	return val
}

func (ctx *CBuf[T]) Clear() {
	ctx.getEnd = 0
	ctx.putEnd = 0
}

func (ctx *CBuf[T]) IsFull() bool {
	return (ctx.putEnd+1)%ctx.size == ctx.getEnd
}

func (ctx *CBuf[T]) IsEmpty() bool {
	return ctx.getEnd == ctx.putEnd
}

func (ctx *CBuf[T]) Size() int {
	return ctx.size
}
