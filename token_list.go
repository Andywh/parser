package main

import (
	"io/ioutil"
	"fmt"
	"strings"
	"strconv"
)

func load_file(path string) string {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Print(err)
	}
	str := string(b)
	return str
}

func number_end(s1 []rune, offset int) int {
	// 现在只能解析普通数字
	// 小数 负数 都不能解析
	digits := "1234567890"
	for i, c := range s1[offset:] {
		if !strings.Contains(digits, string(c)) {
			return i
		}
	}
	fmt.Println("** 错误, 数字解析错误")
	return -1
}

func string_end(s1 []rune, offset int) (string, int){
	bs := map[string]string{
		"b": "\b",
		"f": "\f",
		"n": "\n",
		"r": "\r",
		"t": "\t",
		"/": "/",
		"\"": "\"",
		"\\": "\\",
	}
	res := ""
	i := offset
	for i := offset; i < len(s1); {
		a := s1[i]
		if string(a) == "\"" {
			return res, i
		} else if string(a) == "\\" {
			b := s1[i + 1]
			if c, ok := bs[string(b)]; ok {
				res += c
				i += 2
			} else {
				fmt.Println("** 错误, 不合法的转义字符: %s", string(a) + c)
				example := "\\b \\f \\n \\r \\t \\/ \\\" \\\\"
				fmt.Println("合法的转义字符是: %s", example)
				return "", -1
			}
		} else {
			res += string(a)
			i += 1
		}
	}
	return res, i
}

func loads(s string) []token {
	code :=[]rune(s)
	i := 0
	length := len(code)
	tokens := make([]token, 0)
	spaces := " \b\f\n\r\t"
	digits := "0123456789"
	is_open := true
	for ; i < length; {
		c := code[i]
		i++
		if strings.ContainsRune(spaces, c){
			continue
		} else if strings.Contains(":,[]{}", string(c)) {
			t := Token(auto, string(c))
			tokens = append(tokens, t)
		} else if string(c) == "\"" && is_open {
			// 字符串处理
			result, index := string_end(code, i)
			if index != -1 {
				t := Token(str, result)
				i = index
				tokens = append(tokens, t)
				is_open = !is_open
			} else {
				return nil
			}
		} else if string(c) == "\"" && !is_open {
			is_open = !is_open
			continue
		} else if strings.Contains(digits, string(c)) {
			offset := number_end(code, i)
			value, err := strconv.Atoi(string(code[i-1:i+offset]))
			if err != nil {
				fmt.Println("字符串转换成整数失败")
			}
			t := Token(number, value)
			i += offset
			tokens = append(tokens, t)
		}else if strings.Contains("tfn", string(c)) {
			kvs := map[string]string {
				"t": "true",
				"f": "false",
				"n": "null",
			}
			// 要判断一下是否真的是 true false null
			t := Token(keyword, kvs[string(c)])
			tokens = append(tokens, t)
			i += len(kvs[string(c)])
		} else {
			fmt.Println("*** 错误 %s %s", string(c), string(code[i:i+10]))
			return nil
		}
	}
	return tokens
}

func token_list(path string) []token {
	s := load_file(path)
	ts := loads(s)
	return ts
}
