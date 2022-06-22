package goTool

import "encoding/json"

type GoTool struct {
}

func (goTool *GoTool) JsonToMap(jsonStr string) map[string]interface{} {
	var mapRe map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &mapRe)
	if err != nil {
		mapRe["code"] = -1
		mapRe["message"] = err.Error()
	}
	return mapRe
}

func (goTool *GoTool) StrToJson() {

}
