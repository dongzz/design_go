package strs

type Context struct {
	strategy Str
}

func (context *Context) Context(strategy int, total int, price float64) {
	switch strategy {
	case 0:
		context.strategy = new(NormalStr)
		context.strategy.SetData(total, price)
	}
}

func (context *Context) ContextInterface() float64 {
	return context.strategy.Algorithm()
}
