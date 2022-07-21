#### Go 初始化项目

1. 任意文件夹下创建一个项目 `mkdir -p $HOME/aceld/modoules_test`
1. 创建 go.mod 文件，同时起当前项目的模块名称 `go mod init github.com/aceld/module_test `
1. 创建完成后就会生成一个 go mod 文件
   ```
   module github.com/aceld/moudles_test
   go 1.14
   ```
1. 在项目编写源代码 如果源码依赖某个库（github.com/aceld/zinx/znet）有两种方式 down 下来
   1. 手动 down：`go get github.com/aceld/zinx/znet`
   1. 自动 down
1. down 下来后 go mod 文件会新增一行代码
   ```
   module github.com/aceld/moudles_test
   go 1.14
   require github.com/aceld/zinx v0.0.0-2021xxx-f0xxxxx //indirect
   ```
   含义： 当前模块依赖 `github.com/aceld/zinx` 依赖的版本为：`v0.0.0-2021xxx-f0xxxxx`
   `//indirect` ：表示间接依赖 `github.com/aceld/zinx/znet`因为项目直接依赖的是`znet`,`znet`是`zinx`的子包，所以这是间接依赖
1. 会生成一个 go.sum 文件 ，go.sum 文件的作用 罗列当前项目直接或间接的依赖的所有模块，保证今后项目依赖版本不会被篡改
   ```
   github.com/aceld/zinx v0.0.0-2021xxx-f0xxxxx h1:turbcEfboY81jsKVSQtGkqkturbcEfboY81jsKVSQtGkqk
   github.com/aceld/zinx v0.0.0-2021xxx-f0xxxxx/go.mod h1:turbcEfboY81jsKVSQtGkqkturbcEfboY81jsKVSQtGkqk
   ```
   1. `h1:hash`:表示整体文件的 zip 文件打开后的全部文件的 "校验和" 来生成的 hash,如果不存在，可能表示依赖的库用不上
   1. `xxx/go.mod h1:hash`: 表示用 go.mod 文件来生成的 hash
