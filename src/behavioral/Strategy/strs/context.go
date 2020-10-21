package strs

type Context struct {
	strategy Str
}

func (context *Context) Context(strategy int, total int, price float64) {
	switch strategy {
	case 1:
		context.strategy = new(EightStr)
	case 2:
		context.strategy = new(SevenStr)
	case 3:
		context.strategy = new(FiveStr)
	case 4:
		context.strategy = new(ThreeHumdrumSubThirtyStr)
	case 0:
		fallthrough
	default:
		context.strategy = new(NormalStr)
	}
	context.strategy.SetData(float64(total) * price)
}

func (context *Context) ContextInterface() float64 {
	return context.strategy.Algorithm()
}
