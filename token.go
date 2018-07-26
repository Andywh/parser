package main

import "strconv"

type token struct {
	token_type int
	value interface{}
}

func Token(token_type int, value interface{}) token {
	d := map[string]int{
		":": colon,
		",": comma,
		"{": braceLeft,
		"}": braceRight,
		"[": bracketLeft,
		"]": bracketRight,
	}
	t := token{}
	if token_type == auto {
		v := value.(string)
		t.token_type = d[v]
	} else {
		t.token_type = token_type
	}
	t.value = value
	return t
}

func (t token) String() string {
	s := ""
	if v, ok := t.value.(string); ok {
		s = v
	} else if v, ok := t.value.(int); ok {
		s = strconv.Itoa(v)
	}
	return s
}


