package main

const (
	auto = iota		// auto 的时候, c 是 : , {} [] 其中之一, 要自己判断
	colon
	comma
	braceLeft
	braceRight
	bracketLeft
	bracketRight
	keyword
	number
	str
)