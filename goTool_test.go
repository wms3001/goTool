package goTool

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGoTool_JsonToMap(t *testing.T) {
	var goTool = &GoTool{}
	jsonStr := `{"Code":1,"Message":"connected","Data":"","t":{"a":1,"b":"32423"},"de":[{"c":"23ewrwe"},{"c":"2342rerre"}]}`
	mapRe := goTool.JsonToMap(jsonStr)
	fmt.Println(mapRe)
	fmt.Println(mapRe["Code"])
	ty := reflect.TypeOf(mapRe["de"])
	fmt.Println(ty)
	vl := reflect.ValueOf(mapRe["de"])
	fmt.Println(vl)
	tty := reflect.TypeOf(mapRe["t"])
	fmt.Println(tty)
	vvl := reflect.ValueOf(mapRe["t"])
	fmt.Println(vvl)
	a := mapRe["t"]
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.ValueOf(a))

	aa := a.(map[string]interface{})
	fmt.Println(aa)
	fmt.Println(aa["a"])

	//bb := mapRe["de"]
	//for _, value := range bb {
	//	fmt.Println(value)
	//}
	//fmt.Println(mapRe["t"])
	//tty1 := reflect.TypeOf(vvl)
	//fmt.Println(tty1)
	//vvl1 := reflect.ValueOf(vvl)
	//fmt.Println(vvl1)
	//m := vvl.(map[string]interface{})
	//fmt.Println(m)
	//mqpRee := goTool.JsonToMap(vvl)
}
