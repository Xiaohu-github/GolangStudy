package main //程序的包名，和文件名可以没关系。

/*
import "fmt"
import "time"
导包建议使用下方这种形式
*/
import (
	"fmt"
	"time"
)

//main 函数
//函数{ 必须和函数名同行，这是语法规则否则编译报错。
func main() {
	//go语言对于语句末尾分号无要求，建议不加
	time.Sleep(1 * time.Second)

	fmt.Println("hello Go！！")
}
