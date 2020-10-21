package operates

type OpsFactory struct {}

func (factory *OpsFactory) CreateOperate(strOperate string) Ops {
	switch strOperate {
	case "+":
		return new(AddOps)
	case "-":
		return new(SubOps)
	}
	return nil
}