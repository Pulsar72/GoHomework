/*
设计一个游戏人物的结构体，包含人物的名字、等级、经验值、血量和攻击力等信息。然后设计一个接口，该接口有一个 `Attack` 方法，表示该角色攻击另一个角色。该方法接受一个 `Person` 类型的参数，表示被攻击的角色。实现该接口的结构体需要具备攻击能力，即能够将被攻击角色的血量减少相应的攻击力值。

- 接口的 `Attack` 方法需要在被攻击角色的血量上减去攻击者的攻击力值；
- 结构体的名字、等级、经验值、血量和攻击力等信息需要在创建时初始化。
- 在 `main` 函数中创建两个角色，然后让其中一个角色攻击另一个角色，打印被攻击角色的血量变化情况。
*/
package main

import "fmt"

type Person struct {
	name            string
	level           int
	experienceValue int
	bloodVolume     int
	aggressivity    int
}

func (p Person) Attack(a Person) Person {
	a.bloodVolume -= p.aggressivity
	return a
}

type game interface {
	Attack(a Person) Person
}

func main() {
	jack := Person{
		name:            "jack",
		level:           10,
		experienceValue: 54,
		bloodVolume:     100,
		aggressivity:    5,
	}
	kite := Person{
		name:            "kite",
		level:           20,
		experienceValue: 23,
		bloodVolume:     200,
		aggressivity:    10,
	}
	var a game
	a = jack
	fmt.Printf("jack攻击kite前kite的血量为%d\n", kite.bloodVolume)
	kite = a.Attack(kite)
	fmt.Printf("攻击后kite的血量为%d\n", kite.bloodVolume)
}
