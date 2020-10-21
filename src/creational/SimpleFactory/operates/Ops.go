package operates

type Ops interface {
	SetData(data ...interface{})
	GetResult() float64
}
