package cbuf

type CBuf[T any] struct {
	get_end int
	put_end int
	size    int
	data    []*T
}

func NewCBuf[T any](size int) *CBuf[T] {
	return &CBuf[T]{0, 0, size, make([]*T, size)}
}
func (ctx *CBuf[T]) Push(value *T) bool {
	if ctx.IsFull() {
		return false
	}
	ctx.data[ctx.put_end] = value
	ctx.put_end = (ctx.put_end + 1) % ctx.size
	return true
}

func (ctx *CBuf[T]) Peek() *T {
	if ctx.IsEmpty() {
		return nil
	}
	return ctx.data[ctx.get_end]
}

func (ctx *CBuf[T]) NumEntries() int {
	if ctx.get_end < ctx.put_end {
		return ctx.put_end - ctx.get_end
	}
	return ctx.size - ctx.get_end + ctx.put_end
}

func (ctx *CBuf[T]) Pop() *T {
	if ctx.IsEmpty() {
		return nil
	}
	var val = ctx.data[ctx.get_end]
	ctx.get_end = (ctx.get_end + 1) % ctx.size
	return val
}

func (ctx *CBuf[T]) Clear() {
	ctx.get_end = 0
	ctx.put_end = 0
}

func (ctx *CBuf[T]) IsFull() bool {
	return (ctx.put_end+1)%ctx.size == ctx.get_end
}

func (ctx *CBuf[T]) IsEmpty() bool {
	return ctx.get_end == ctx.put_end
}

func (ctx *CBuf[T]) Size() int {
	return ctx.size
}



