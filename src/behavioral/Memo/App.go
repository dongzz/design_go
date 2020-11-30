package main

import "fmt"

type GameRole struct {
	vit int // 生命力
	atk int // 攻击力
	def int // 防御力
}

func (role *GameRole) StateDisplay() {
	fmt.Printf("生命力：%d\n攻击力：%d\n防御力：%d\n", role.vit, role.def, role.def)
}

func (role *GameRole) InitState() {
	fmt.Println("初始化状态：")
	role.vit = 100
	role.atk = 100
	role.def = 100
}

func (role *GameRole) FightWithBoss() {
	fmt.Println("单挑boss")
	role.vit = 1
	role.atk = 0
	role.def = 0
}

func (role *GameRole) CreateMemento() RoleStateMemento {
	fmt.Println("游戏存档")
	return RoleStateMemento{vit: role.vit, atk: role.atk, def: role.def}
}

func (role *GameRole) RecoveryMemento(memento RoleStateMemento) {
	fmt.Println("游戏读档")
	role.vit = memento.vit
	role.atk = memento.atk
	role.def = memento.def
}

type RoleStateMemento struct {
	vit int // 生命力
	atk int // 攻击力
	def int // 防御力
}

type RoleStateMementoCaretaker struct {
	memento RoleStateMemento
}

func main() {
	// 大战前
	role := GameRole{}
	role.InitState()
	role.StateDisplay()

	// 保存进度
	caretaker := RoleStateMementoCaretaker{memento: role.CreateMemento()}

	// 大战boss
	role.FightWithBoss()
	role.StateDisplay()

	// 重新读档
	role.RecoveryMemento(caretaker.memento)
	role.StateDisplay()
}
