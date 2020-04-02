package util

import (
	"bytes"
	"errors"
	jsoniter "github.com/json-iterator/go"
	"log"
	"strconv"
)

// object to JSON str
func ObjectToJson(src interface{}) (string, error) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	if result, err := json.Marshal(src); err != nil {
		return "", errors.New("object to Json str throw exception: " + err.Error())
	} else {
		return string(result), nil
	}
}

// JSON str to object
func JsonToObject(src string, target interface{}) error {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	if err := json.Unmarshal([]byte(src), target); err != nil {
		return errors.New("Json str to object throw exception: " + err.Error())
	}
	return nil
}

// High-performance splice String
func AddStr(input ...interface{}) string {
	if input == nil || len(input) == 0 {
		return ""
	}
	var rstr bytes.Buffer
	for e := range input {
		s := input[e]
		if v, b := s.(string); b {
			rstr.WriteString(v)
		} else if v, b := s.(error); b {
			rstr.WriteString(v.Error())
		} else if v, b := s.(bool); b {
			if v {
				rstr.WriteString("true")
			} else {
				rstr.WriteString("false")
			}
		} else {
			rstr.WriteString(AnyToStr(s))
		}
	}
	return rstr.String()
}

//  int uint float string bool
//  json
func AnyToStr(any interface{}) string {
	if any == nil {
		return ""
	}
	if str, ok := any.(string); ok {
		return str
	} else if str, ok := any.(int); ok {
		return strconv.FormatInt(int64(str), 10)
	} else if str, ok := any.(int8); ok {
		return strconv.FormatInt(int64(str), 10)
	} else if str, ok := any.(int16); ok {
		return strconv.FormatInt(int64(str), 10)
	} else if str, ok := any.(int32); ok {
		return strconv.FormatInt(int64(str), 10)
	} else if str, ok := any.(int64); ok {
		return strconv.FormatInt(int64(str), 10)
	} else if str, ok := any.(float32); ok {
		return strconv.FormatFloat(float64(str), 'f', 0, 64)
	} else if str, ok := any.(float64); ok {
		return strconv.FormatFloat(float64(str), 'f', 0, 64)
	} else if str, ok := any.(uint); ok {
		return strconv.FormatUint(uint64(str), 10)
	} else if str, ok := any.(uint8); ok {
		return strconv.FormatUint(uint64(str), 10)
	} else if str, ok := any.(uint16); ok {
		return strconv.FormatUint(uint64(str), 10)
	} else if str, ok := any.(uint32); ok {
		return strconv.FormatUint(uint64(str), 10)
	} else if str, ok := any.(uint64); ok {
		return strconv.FormatUint(uint64(str), 10)
	} else if str, ok := any.(bool); ok {
		if str {
			return "True"
		}
		return "False"
	} else {
		if ret, err := ObjectToJson(any); err != nil {
			log.Println("any to json fail: ", err.Error())
			return ""
		} else {
			return ret
		}
	}
	return ""
}