package main

import (
	"encoding/json"
	"fmt"
)

type EquipType struct {
	Id   int64  `json:"id"`
	Code string `json:"code"'`
	Name string `json:"name"`
}

func main() {

	var equip EquipType
	equip.Id = 1
	equip.Code = "equip01"
	equip.Name = "武器"

	json, _ := json.Marshal(equip)
	fmt.Println(json)
	fmt.Println(string(json))
	equip.cannotChange() //不会修改外部的值，结构体是值传递
	fmt.Println(equip)
	(&equip).change() //等价于equip.change(),引用传递修改
	fmt.Println(equip)
}

func (e EquipType) cannotChange() {
	e.Name = "22"
	fmt.Println(e)
}

func (e *EquipType) change() {
	(*e).Code = "修改后的code"
	e.Name = "修改后的名称"
}
