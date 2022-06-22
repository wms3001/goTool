package goTool

import (
	"fmt"
	"testing"
)

func TestGoTool_JsonToMap(t *testing.T) {
	var goTool = &GoTool{}
	jsonStr := `{"Code":1,"Message":"connected","Data":""}`
	mapRe := goTool.JsonToMap(jsonStr)
	fmt.Println(mapRe)
}
