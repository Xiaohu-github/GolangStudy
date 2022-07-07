package main

import (
	_ "GolangStudy/5-init/lib1"      //匿名使用，即使不使用lib1内的任何方法，它同样会执行包内的init方法
	mylib2 "GolangStudy/5-init/lib2" //别名使用
	// . "GolangStudy/5-init/lib2" //直接使用：不需要通过 lib2.Lib2Test() 而是直接 使用 Lib2Test() ；尽量避免使用该方式，可能会导致两个包内方法重名报错
)

func main() {
	/*
		在执行lib包时，执行顺序为先执行完包内所有init方法之后再走其他方法
		->lib1Init
		->lib2Init
		->lib2Test
	*/

	// lib1.Lib1Test()
	// lib2.Lib2Test()
	mylib2.Lib2Test()
}
