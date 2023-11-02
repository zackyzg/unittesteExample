# unittesteExample

Go语言单元测试示例

### 不写测试的开发不是好程序员

# GO TEST 工具

go 语言中的测试依赖`go test`命令。编写测试代码和编写普通的Go代码过程类似的，并不需要学习新的语法、规则和工具。

go test 命令是一个按照一定的约定和组织的测试代码的驱动程序。
在包目录内，所有以`_test.go`为后缀名的源代码都是`go test`测试的一部分，不会被`go build`编译到最终的可执行文件中。

在`*_test.go`文件中又三种类型的函数，单元测试函数、基准测试函数和示例函数。

| 类型   |格式|作用|
|:-----|:----|:----|
| 测试函数 | 函数前缀名为Test | 测试程序的一些逻辑行为是否正确 |
| 基准函数 | 函数名前缀为Benchmark | 测试函数的性能 |
| 示例函数 | 函数名前缀为Example | 为文档提供示例文档 |

go test 命令会遍历所有的`*_test.go`文件中符合上述命名规制的函数，然后生成一个临时的main包用于调用
相应的测试函数，然后构建并运行、报告测试结果，最后清理测试中生成的临时文件

## 测试函数的格式
每个测试函数必须导入testing包，测试函数的基本格式（签名）如下：
> func TestName(t *testing.T){
> 
>      // ...
> 
> }

测试函数的名字必须以`Test`开头，可选的后缀名必须以大写字母开头，举几个例子:
> func TestAdd(t *testing.T){ ... }
> 
> func TestSum(t *testing.T){ ... }
> 
> func TestLog(t *testing.T){ ... }

其中参数t用于报告测试失败和附加的日志信息。 testing.T的拥有的方法如下
> func (c *T) Error(args ...interface{})
> 
> func (c *T) Errorf(format string, args ...interface{}
> 
> func (c *T) Fail()
> 
> func (c *T) FailNow()
> 
> func (c *T) Failed() bool
> 
> func (c *T) Fatal(args ...interface{})
> 
> func (c *T) Fatalf(format string, args ...interface{})
> 
> func (c *T) Log(args ...interface{})
> 
> func (c *T) Logf(format string, args ...interface{})
> 
> func (c *T) Name() string
> 
> func (t *T) Parallel()
> 
> func (c *T) Name() string
> 
> func (t *T) Run(name string, f func(t *T)) bool
> 
> func (c *T) Skip(args ...interface{})
> 
> func (c *T) SkipNow()
> 
> func (c *T) Skipf(format string, args ...interface{})
> 
> func (c *T) Skipped() bool
> 