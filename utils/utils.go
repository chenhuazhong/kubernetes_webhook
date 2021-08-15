package utils

import (
	"fmt"
	"os"
	"strconv"
)

func GetDefaultEnv(key , default_value string) (value string){
	value = os.Getenv(key)
	if value != ""{
		return
	} else {
		value = default_value
		return
	}
}


func Typefun(string_value string ,value interface{}) interface{} {
	switch value.(type) {
	case string:
		return string_value
	case int:
        int_value, err := strconv.Atoi(string_value)
        if err != nil{
        	return int_value
		} else{
			fmt.Println(err)
			panic(err)
		}
	case float32:
		return string_value
	case float64:
		return string_value
	default:
		return string_value
	}
}