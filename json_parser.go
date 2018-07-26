package main

import "fmt"

func main() {
	t0 := "data.json"
	ts := token_list(t0)
	o := parse(ts)

	if value, ok := o.(map[token]interface{}); ok {
		for k, v := range value {
			fmt.Println(k,":\t", v)
		}
	}
}
