package main

import "fmt"

func main() {
	var expression string
	_, err := fmt.Scan(&expression)
	if err != nil {
		return
	}
	calculate(expression)

}

func calculate(expression string) {
	var a int32 //a是未知数名
	f := 1
	now := 1
	var k, b, x int //f初始化为正，now初始为左
	var r bool      //用来判是否有数字读入
	for _, c := range expression {
		if c == '-' {
			b += now * f * x
			x = 0
			f = -1
			r = false
		}
		if c == '+' {
			b += now * f * x
			x = 0
			f = 1
			r = false
		}
		if c == '=' {
			b += now * f * x
			x = 0
			f = 1
			now = -1
			r = false
		}
		if c >= 'a' && c <= 'z' {
			if r == true {
				k += now * f * x
				x = 0
			} else {
				k += now * f
			}
			a = c
			r = false
		}
		if c >= '0' && c <= '9' {
			x = x*10 + int(c) - '0'
			r = true
		}
	}
	b += now * f * x
	ans := float64(-b) / float64(k)
	if ans == -0.0 {
		ans = 0
	} //特判
	fmt.Printf("%c=%.3f", a, ans)

}
