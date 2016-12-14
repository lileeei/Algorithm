/*
 *此文档由[参考](http://blog.cyeam.com/golang/2014/08/15/go_largenumberx)修改得到，原程序当输入为`91*21`时得到的结果为`911`而非`1911`
 *待看[文档](http://blog.csdn.net/chhuach2005/article/details/21168179)
**/

package main

import (
	"fmt"
)

func LargeNumberMultiplication(a string, b string) (reuslt string) {
	a = Reverse(a)
	b = Reverse(b)
	c := make([]byte, len(a)+len(b))

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			c[i+j] += (a[i] - '0') * (b[j] - '0')
			fmt.Println("i+j: ", i+j, "c[i+j: ]", c[i + j])
		}
	}

	fmt.Printf("c: %#v\n", c)

	var plus byte = 0
	for i := 0; i < len(c); i++ {
		temp := c[i] + plus
		if temp == 0 {

			break
		}
	
		plus = 0
		if temp > 9 {
			plus = temp / 10
			reuslt += string(temp - plus*10 + '0')
		} else {
			reuslt += string(temp + '0')
		}

	}
	fmt.Printf("reuslt: %#v\n", reuslt)
	return Reverse(reuslt)
}

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func main() {
	a := "91"
	fmt.Println(LargeNumberMultiplication(a, "21"))
}
