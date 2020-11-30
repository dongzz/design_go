package main

import "fmt"

type Company interface {
	Add(company Company)
	Remove(company Company)
	Display(dept int)
	LineOfDuty()
}

type ConcreteCompany struct {
	name            string
	childrenCompany []Company
}

func (company *ConcreteCompany) Add(c Company) {
	company.childrenCompany = append(company.childrenCompany, c)
}

func (company *ConcreteCompany) Remove(c Company) {
	for i, c2 := range company.childrenCompany {
		if c2 == c {
			company.childrenCompany = append(company.childrenCompany[:i], company.childrenCompany[i+1:]...)
		}
	}
}

func (company *ConcreteCompany) Display(dept int) {
	for i := dept; i > 0; i-- {
		fmt.Printf("-")
	}
	fmt.Println(company.name)

	for _, c := range company.childrenCompany {
		c.Display(dept + 2)
	}
}

func (company *ConcreteCompany) LineOfDuty() {
	for _, c := range company.childrenCompany {
		c.LineOfDuty()
	}
}

type HRDepartment struct {
	name            string
	childrenCompany []Company
}

func (company *HRDepartment) Add(c Company) {
	company.childrenCompany = append(company.childrenCompany, c)
}

func (company *HRDepartment) Remove(c Company) {
	for i, c2 := range company.childrenCompany {
		if c2 == c {
			company.childrenCompany = append(company.childrenCompany[:i], company.childrenCompany[i+1:]...)
		}
	}
}

func (company *HRDepartment) Display(dept int) {
	for i := dept; i > 0; i-- {
		fmt.Printf("-")
	}
	fmt.Println(company.name)

	for _, c := range company.childrenCompany {
		c.Display(dept + 2)
	}
}

func (company *HRDepartment) LineOfDuty() {
	fmt.Printf("%s 员工招聘培训管理\n", company.name)
}

type FinanceDepartment struct {
	name string
}

func (company *FinanceDepartment) Add(c Company) {

}

func (company *FinanceDepartment) Remove(c Company) {

}

func (company *FinanceDepartment) Display(dept int) {
	for i := dept; i > 0; i-- {
		fmt.Printf("-")
	}
	fmt.Println(company.name)
}

func (company *FinanceDepartment) LineOfDuty() {
	fmt.Printf("%s 公司财务收支管理\n", company.name)
}

func main() {
	company := ConcreteCompany{name: "新锐和达"}
	company.Add(&HRDepartment{name: "人事部"})
	company.Add(&FinanceDepartment{name: "财务部"})

	company2 := ConcreteCompany{name: "大通关"}
	company2.Add(&HRDepartment{name: "大通关人事部"})
	company2.Add(&FinanceDepartment{name: "大通关财务部"})
	company.Add(&company2)

	company3 := ConcreteCompany{name: "云游网"}
	company3.Add(&HRDepartment{name: "云游网人事部"})
	company3.Add(&FinanceDepartment{name: "云游网财务部"})
	company.Add(&company3)

	fmt.Println("结构图")
	company.Display(1)
	fmt.Println("职责")
	company.LineOfDuty()
}
