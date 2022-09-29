package goTool

import (
	"encoding/base64"
	"encoding/json"
	"strconv"
)

type GoTool struct {
}

/**
byte数组转base64字符串
*/
func (goTool *GoTool) ByteToBase64(bt []byte) string {
	return base64.StdEncoding.EncodeToString(bt)
}

/**
base64转byte数组
*/
func (goTool *GoTool) Base64ToByte(baseStr string) []byte {
	res, _ := base64.StdEncoding.DecodeString(baseStr)
	return res
}

/**
字符串转base64字符串
*/
func (goTool *GoTool) StrToBase64(str string) string {
	return goTool.ByteToBase64([]byte(str))
}

/**
base64转字符串
*/
func (goTool *GoTool) Base64ToStr(baseStr string) string {
	return string(goTool.Base64ToByte(baseStr))
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

func (goTool *GoTool) JsonToMap1(jsonStr string) map[string]string {
	var mapRe map[string]string
	err := json.Unmarshal([]byte(jsonStr), &mapRe)
	if err != nil {
		mapRe["code"] = "-1"
		mapRe["message"] = err.Error()
	}
	return mapRe
}

func (goTool *GoTool) StrToJson() {

}

func (goTool *GoTool) Strval(value interface{}) string {
	var key string
	if value == nil {
		return key
	}

	switch value.(type) {
	case float64:
		ft := value.(float64)
		key = strconv.FormatFloat(ft, 'f', -1, 64)
	case float32:
		ft := value.(float32)
		key = strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case int:
		it := value.(int)
		key = strconv.Itoa(it)
	case uint:
		it := value.(uint)
		key = strconv.Itoa(int(it))
	case int8:
		it := value.(int8)
		key = strconv.Itoa(int(it))
	case uint8:
		it := value.(uint8)
		key = strconv.Itoa(int(it))
	case int16:
		it := value.(int16)
		key = strconv.Itoa(int(it))
	case uint16:
		it := value.(uint16)
		key = strconv.Itoa(int(it))
	case int32:
		it := value.(int32)
		key = strconv.Itoa(int(it))
	case uint32:
		it := value.(uint32)
		key = strconv.Itoa(int(it))
	case int64:
		it := value.(int64)
		key = strconv.FormatInt(it, 10)
	case uint64:
		it := value.(uint64)
		key = strconv.FormatUint(it, 10)
	case string:
		key = value.(string)
	case []byte:
		key = string(value.([]byte))
	default:
		newValue, _ := json.Marshal(value)
		key = string(newValue)
	}

	return key
}
