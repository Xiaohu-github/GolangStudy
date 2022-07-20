#### Go 中如何进行项目模块及依赖管理

- **GOROOT**：Go 的源码包路径
- **GOPATH (bin pkg src) 工作模式弊端：**
  - 无版本控制概念
  - 无法同步一致第三方版本号
  - 无法指定当前项目引用哪个第三方版本号

##### **GOMODULES 模式**：

go 在`V1.11`以上版本开始引用 gomodule 模式，所以要使用`go mod` 版本必须在`v1.11`以上。
通过 `go mod help` 来进行查看

- `download` 下载 go.mod 文件中指明的所有依赖
- `edit` 编辑 go.mod 文件
- `graph` 查看现有的依赖结构
- `init` 生成 go.mod 文件
- `tidy` 整理现有依赖
- `vendor` 导出项目所有的依赖到 vendor 目录
- `verify` 校验一个模块是否被篡改
- `why` 查看为什么需要依赖某模块

##### **go mod 环境变量**：

通过`go env`命令进行查看

```
$ go env

GO111MODULE="auto"
GOPROXY="http://proxy.golang,org,direct"
GONOPROXY=""
GOSUMDB="sum.golang.org"
GONOSUMDB=""
GOPRIVATE=""
......

```

- **`GO111MODULE`** 该变量作为 Go Modules 的开关，包含以下参数：
  - `auto`：只要项目包含 go.mod 文件就代表启用 Go Modules。
  - `on`：启用 Go Modules
  - `off`：禁用 Go Modules，不推荐
- **`GOPROXY`** 该变量用于设置 GO 模块代理，用于使 Go 在后续拉去镜像版本使直接通过镜像站点来拉取，国内镜像:
  - 阿里云：`https://mirrors.aliyun.com/goproxy`
  - 七牛云：`https://goproxy.cn,direct`
  ```
    设置方法:
    $ go env -w GOPROXY=https://mirrors.aliyun.com/goproxy
  ```
- **`GOSUMDB`** 该变量用于校验包版本的完整性，校验地址就是这个值设置的地址，默认会走`GOPROXY`设置的值。
- **`GOPRIVATE`** 设置的地址不会走`GOSUMDB`下载和校验
  ```
    设置方法:
    $ go env -w GOPRIVATE="*.example.com"
  ```
  表示所有模块路径为 example.com 的子域名，比如 git.example.com、 hello.example.com 等都不会进行`GOSUMDB`下载和校验
