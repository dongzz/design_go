package operates

type SubOps struct {
	NumOne,NumTwo float64
}

func (ops *SubOps) SetData(data ...interface{}) {
	if len(data) != 2 {
		panic("参数需为两个")
	}
	if _, ok := data[0].(float64); !ok {
		panic("参数 1 需为 float64 类型")
	}

	if _, ok := data[1].(float64); !ok {
		panic("参数 2 需为 float64 类型")
	}
	ops.NumOne = data[0].(float64)
	ops.NumTwo = data[1].(float64)
}

func (ops *SubOps) GetResult() float64 {
	return ops.NumOne - ops.NumTwo
}
