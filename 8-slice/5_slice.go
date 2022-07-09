/*
 切片截取元素:变量名[开始位置：结束位置]
 切片拷贝：copy(目标切片,源切片)
*/
package main

import "fmt"

func main() {
	s := make([]int, 3)
	s[0] = 100
	s[1] = 200
	fmt.Println("s=", s)

	fmt.Println("=== s1=s[1:3] ===")
	//切片截取 1~3 元素不包含3 [0=>100, |1=>200 ,2=>0|] ==> [200,0]
	s1 := s[1:3] //此时s1[0]==s[1]
	//s1是指向s的地址，所以s和s1是同一片空间，修改s1的值同时s也会改变
	// s1[1] = 22
	// fmt.Println("s=", s)
	fmt.Println("s1 = ", s1)

	//将s 拷贝覆盖到 s2，根据s2 len，多余的值会舍弃
	fmt.Println("===copy(s2, s)====")
	s2 := make([]int, 1, 3)
	s2[0] = 111
	s2 = append(s2, 222)
	fmt.Println("s2=", s2)
	copy(s2, s1)
	fmt.Println("copy(s2,s) s2=", s2)
}
