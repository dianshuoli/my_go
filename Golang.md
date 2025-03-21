

#     Golang

code tips

```go
for i, l :=0, len(slice); i<l;i++{  #可以避免每层循环求长操作，减少开销

}
```

# 标准输出

在 Go 语言中，有几个主要的函数和方法用于输出文本信息到标准输出（通常是终端或控制台）。下面是这些函数和方法的详细介绍：

1. **`fmt.Print`**:

   - **功能**: `fmt.Print` 用于将一个或多个值打印到标准输出，不附加换行符。

   - **格式化**: 不支持格式化，直接输出给定的值。

   - 示例

     :

     ```
     goCopy codefmt.Print("Hello, ")
     fmt.Print("World!")
     // 输出: Hello, World!
     ```

2. **`fmt.Printf`**:

   - **功能**: `fmt.Printf` 用于将一个格式化字符串和参数打印到标准输出。

   - **格式化**: [支持格式化]()，可以在格式化字符串中使用占位符 `%` 来指定参数的格式。例如 `%s` 表示字符串，`%d` 表示整数。

   - 示例

     :

     ```
     goCopy codename := "Alice"
     age := 30
     fmt.Printf("Name: %s, Age: %d\n", name, age)
     // 输出: Name: Alice, Age: 30
     ```

3. **`fmt.Println`**:

   - **功能**: `fmt.Println` 用于将一个或多个值打印到标准输出，并在最后附加换行符。

   - **格式化**: 不支持格式化，直接输出给定的值。

   - 示例

     :

     ```
     goCopy codefmt.Println("Hello,")
     fmt.Println("World!")
     // 输出:
     // Hello,
     // World!
     ```

4. **`log.Print`**:

   - **功能**: `log.Print` 类似于 `fmt.Print`，用于将一个或多个值打印到标准输出，不附加换行符。它通常用于记录应用程序的运行日志。

   - **格式化**: 不支持格式化，直接输出给定的值。

   - 示例

     :

     ```
     goCopy codelog.Print("Logging some information")
     ```

5. **`log.Printf`**:

   - **功能**: `log.Printf` 类似于 `fmt.Printf`，用于将一个格式化字符串和参数打印到标准输出，通常用于记录格式化的运行日志。

   - **格式化**: 支持格式化，可以在格式化字符串中使用占位符 `%`。

   - 示例

     :

     ```
     goCopy codeerrorCode := 404
     log.Printf("Error: HTTP %d Not Found", errorCode)
     ```

6. **`log.Println`**:

   - **功能**: `log.Println` 类似于 `fmt.Println`，用于将一个或多个值打印到标准输出，并在最后附加换行符，通常用于记录日志信息。

   - **格式化**: 不支持格式化，直接输出给定的值。

   - 示例

     :

     ```go
     goCopy codelog.Println("This is an important log message.")
     ```

7. **`fmt.Fprint`** 和 **`fmt.Fprintf`**:

   - **功能**: `fmt.Fprint` 和 `fmt.Fprintf` 用于将文本输出到指定的 `io.Writer`，而不是标准输出。这使你可以将文本输出到文件、网络连接等。

   - **格式化**: 与 `fmt.Print` 和 `fmt.Printf` 相同，支持格式化。

   - 示例

     :

     ```go
     goCopy codefile, _ := os.Create("output.txt")
     fmt.Fprint(file, "This text is written to a file.")
     fmt.Fprintf(file, "Value: %d\n", 42)
     ```

8. **`fmt.Sprint` 和 `fmt.Sprintf`**:

   - **功能**: `fmt.Sprint` 和 `fmt.Sprintf` 与 `fmt.Print` 和 `fmt.Printf` 类似，但它们返回一个字符串，而不是将文本输出到标准输出。

   - **格式化**: 支持格式化。

   - 示例

     :

     ```go
     goCopy coderesult := fmt.Sprint("This is a string result.")
     formatted := fmt.Sprintf("Formatted: %d", 123)
     ```

这些函数和方法允许你在 Go 中执行各种文本输出操作，无论是简单的打印还是复杂的日志记录。你可以根据需要选择适合你的特定任务的函数或方法。

![1727516808655](C:\Users\admin\AppData\Roaming\Typora\typora-user-images\1727516808655.png)

## 不定长参数

加三个.(...)转为不定长参数

## 指针

## new和make

我们先来看一个例子：

```go
func main() {
    var a *int
    *a = 100
    fmt.Println(*a)

    var b map[string]int  
    b["测试"] = 100
    fmt.Println(b)
}
```

执行上面的代码会引发panic，为什么呢？ 在Go语言中对于引用类型的变量，我们在使用的时候不仅要声明它，还要为它分配内存空间，否则我们的值就没办法存储。而对于值类型的声明不需要分配内存空间，是因为它们在声明的时候已经默认分配好了内存空间。要分配内存，就引出来今天的new和make。 Go语言中new和make是内建的两个函数，主要用来分配内存

 ***var b map[string]int     //出现panic  引用类型的变量  声明不分配空间****

## new

new是一个内置的函数，它的函数签名如下：

```
    func new(Type) *Type
```

其中，

```
    1.Type表示类型，new函数只接受一个参数，这个参数是一个类型
    2.*Type表示类型指针，new函数返回一个指向该类型内存地址的指针。
```

new函数不太常用，使用new函数得到的是一个类型的指针，并且该指针对应的值为该类型的零值。举个例子：

```go
func main() {
    a := new(int)
    b := new(bool)
    fmt.Printf("%T\n", a) // *int
    fmt.Printf("%T\n", b) // *bool
    fmt.Println(*a)       // 0
    fmt.Println(*b)       // false
}    
```

本节开始的示例代码中`var a *int`只是声明了一个指针变量a但是没有初始化，指针作为引用类型需要初始化后才会拥有内存空间，才可以给它赋值。应该按照如下方式使用内置的new函数对a进行初始化之后就可以正常对其赋值了：

```go
func main() {
    var a *int
    a = new(int)
    *a = 10
    fmt.Println(*a)
}
```

## make

**make**也是用于内存分配的，区别于new，**make只用于slice、map以及chan的内存创建**，而且它返回的类型就是这三个类型本身，而不是他们的指针类型，因为这三种类型就是引用类型，所以就没有必要返回他们的指针了。make函数的函数签名如下：

```
func make(t Type, size ...IntegerType) Type
```

make函数是无可替代的，我们在使用slice、map以及channel的时候，都需要使用make进行初始化，然后才可以对它们进行操作。这个我们在上一章中都有说明，关于channel我们会在后续的章节详细说明。

本节开始的示例中`var b map[string]int`只是声明变量b是一个map类型的变量，需要像下面的示例代码一样使用make函数进行初始化操作之后，才能对其进行键值对赋值：

```go
func main() {
    var b map[string]int
    b = make(map[string]int, 10)
    b["测试"] = 100
    fmt.Println(b)
}
```

![1726651050021](C:\Users\admin\AppData\Roaming\Typora\typora-user-images\1726651050021.png)

### new与make的区别

```
    1.二者都是用来做内存分配的。
    2.make只用于slice、map以及channel的初始化，返回的还是这三个引用类型本身；
    3.而new用于类型的内存分配，并且内存对应的值为类型零值，返回的是指向类型的指针。
```

## 切片（slice）

是一个可变长度的序列。

切片的数据结构包含三个元素：

1. 指向底层数组的指针：用于访问实际存储元素的数据。 指向底层数组的起始位置。 
2. 切片的长度：表示当前切片中元素的数量。
3. 切片的容量：底层数组中可供使用的元素数量。

切片的扩容机制是这样的：
当向切片添加元素导致其长度超过当前容量时，切片会进行扩容。扩容的策略不是简单地增加固定的大小，而是根据当前容量的大小来决定新的容量。

通常，如果当前容量小于 1024 个元素，新容量会翻倍。例如，如果当前容量是 100 ，添加元素导致长度超过 100 时，新容量会变为 200 。

如果当前容量大于或等于 1024 个元素，新容量会增加约 25% 。例如，如果当前容量是 2000 ，新容量可能会增加到约 2500 。

## map

map中判断键值对是否存在

```go
if value, exist := m[5]; exist {
		fmt.Println(value)
	} else {
		fmt.Print("key is:", exist)
		fmt.Println(value)
	}
```

## Math包 常用函数

![1727513490777](C:\Users\admin\AppData\Roaming\Typora\typora-user-images\1727513490777.png)

![1727513778902](C:\Users\admin\AppData\Roaming\Typora\typora-user-images\1727513778902.png)

## 时间

![1727515119277](C:\Users\admin\AppData\Roaming\Typora\typora-user-images\1727515119277.png)

## interface 

在 Go 语言中，`interface{}` 类型的变量可以保存任何类型的值。当你需要从 `interface{}` 类型转换回具体类型时，需要使用类型断言（type assertion）来实现。

如果你有一个 `interface{}` 类型的变量 `v`，并且你知道 `v` 实际上保存的是一个 `string` 类型的值，你可以通过类型断言将其转换回来：

示例代码

```go
1package main
2
3import (
4	"fmt"
5)
6
7func main() {
8	// 创建一个 interface{} 类型的变量
9	var v interface{} = "Hello, World!"
10
11	// 类型断言：将 interface{} 转换为 string
12	if s, ok := v.(string); ok {
13		fmt.Println("Converted string:", s)
14	} else {
15		fmt.Println("Type assertion failed: expected string, got something else")
16	}
17}
```

详细解释

1. **类型断言**：
   - `if s, ok := v.(string); ok`：这是一个类型断言表达式。它尝试将 `v` 转换成 `string` 类型。
   - `ok` 是一个布尔值，表示类型断言是否成功。如果 `v` 实际上是 `string` 类型，则 `ok` 为 `true`，否则为 `false`。
2. **使用转换后的值**：
   - 如果类型断言成功（即 `ok` 为 `true`），则可以安全地使用转换后的值 `s`。
   - 如果类型断言失败（即 `ok` 为 `false`），则应该处理这种情况，避免程序因类型错误而崩溃。

注意事项

- **类型断言的必要性**：当从 `interface{}` 类型转换到具体类型时，必须使用类型断言来确保类型的安全性。
- **错误处理**：在类型断言失败的情况下，需要妥善处理，避免程序因类型错误而崩溃。
- **多重赋值**：类型断言可以使用多重赋值的形式，即 `if s, ok := v.(string); ok`，这样可以在一次赋值中同时获取转换后的值和断言的结果。

### 其他示例

如果你不确定 `v` 的具体类型，但想要处理多种情况，可以使用类型开关（type switch）：

```go
1package main
2
3import (
4	"fmt"
5)
6
7func main() {
8	// 创建一个 interface{} 类型的变量
9	var v interface{} = "Hello, World!"
10
11	switch s := v.(type) {
12	case string:
13		fmt.Println("Converted string:", s)
14	case int:
15		fmt.Println("Converted integer:", s)
16	default:
17		fmt.Println("Unknown type:", s)
18	}
19}
```

### 类型开关

- **类型开关**：使用 `switch` 语句可以针对不同的类型进行处理。每个 `case` 子句指定一个类型，如果类型断言成功，则执行该子句中的代码。
- **默认分支**：可以使用 `default` 子句来处理未匹配的情况。

通过类型断言和类型开关，你可以安全地从 `interface{}` 类型转换到具体类型，并处理可能出现的不同情况。

## String

![1726626705496](C:\Users\admin\AppData\Roaming\Typora\typora-user-images\1726626705496.png)

```go
//获取中文字符串字符长度的示例
package main

import (
	"fmt"
)

func main() {
	str := "你好，世界" // 中文字符串
	// 使用 len() 获取字节长度
	fmt.Println("Byte length:", len(str))

	// 将字符串转换为 rune 切片来获取字符长度
	fmt.Println("Character length:", len([]rune(str)))
    
    ch := "li李说i"
	fmt.Printf("char of ch[2] is :%c\n", []rune(ch)[2]) // 李
    for _, v := range ch {  //默认以rune类型遍历
		fmt.Printf("%c,", v)
	}  //l,i,李,说,i,  
    
    //下标形式以byte类型输出
    fmt.Println(ch[2], "len of ch is:", len(ch))  //230 len of ch is: 9
    
}

```

### 字符串读取格式区别

for range 默认以rune类型遍历

```go
ch := "li李说i"
for _, v := range ch {  //默认以rune类型遍历
		fmt.Printf("%c,", v)
	}  //l,i,李,说,i, 
```

下表形式

```go
ch := "li李说i"
fmt.Println(ch[2], "len of ch is:", len(ch))  //230 len of ch is: 9
```



### 字符串拼接

1、**使用 `+` 运算符**

这是 Go 中最直观的拼接方式，直接使用 `+` 运算符连接两个或多个字符串。

```go
go复制代码str := "Hello" + " " + "World"
```

- **性能开销**：
  - Go 中的字符串是不可变的，每次使用 `+` 进行拼接时，都会创建一个新的字符串。
  - 每次拼接都会在内存中分配新的空间，将旧字符串的数据复制到新的字符串中。
  - 如果在循环中频繁使用 `+`，将导致大量的内存分配和数据复制，性能较差，尤其是拼接大量或长字符串时。
- **适用场景**：
  - 适用于少量、简单的字符串拼接场景。
  - 不适合大量拼接操作，尤其是在循环中频繁拼接的情况。

2、**使用 `strings.Join`**

`strings.Join` 是 Go 标准库中提供的拼接方法，它将一个字符串切片通过指定的分隔符连接起来。

```go
go复制代码import "strings"

strSlice := []string{"Hello", "World"}
str := strings.Join(strSlice, " ")
```

- **性能开销**：
  - `strings.Join` 通过一次性分配目标字符串所需的内存来避免重复分配。
  - 性能较好，因为它在拼接之前会预先计算所有字符串的总长度，然后只做一次内存分配和拷贝。
  - 在需要拼接多个字符串时，它的性能明显优于 `+` 操作符。
- **适用场景**：
  - 适合拼接一个字符串数组，特别是当你已经有一个切片时。
  - 高效且适合在高性能要求的场景中使用。

3、**使用 `fmt.Sprintf`**

`fmt.Sprintf` 提供了一种格式化拼接字符串的方法，类似于 C 语言的 `printf`。

```
go复制代码import "fmt"

str := fmt.Sprintf("%s %s", "Hello", "World")
```

- **性能开销**：
  - `fmt.Sprintf` 提供了灵活的格式化功能，但这种灵活性带来了较大的性能开销。
  - 相比 `+` 运算符和 `strings.Join`，`fmt.Sprintf` 的性能较差，因为它需要解析格式化字符串并进行转换操作。
  - 性能劣于 `+` 和 `strings.Join`，通常不适合大量拼接操作。
- **适用场景**：
  - 适用于需要格式化输出的场景，尤其是需要混合不同数据类型的场合（如整数、浮点数、字符串等的拼接）。
  - 不建议在高性能要求的场景中频繁使用。

4、**使用 `strings.Builder`**

`strings.Builder` 是 Go 1.10 引入的一个高效的字符串构建器，专门用于优化多次拼接字符串的场景。

```
import "strings"

var builder strings.Builder
builder.WriteString("Hello")
builder.WriteString(" ")
builder.WriteString("World")
str := builder.String()
```

- **性能开销**：
  - `strings.Builder` 使用一个可动态扩展的缓冲区来存储拼接的字符串，避免了每次拼接时重新分配内存。
  - 内部自动处理内存的扩展和分配问题，性能非常高效，尤其是在大量拼接时。
  - 推荐用于需要频繁拼接的场景，因为它大幅减少了内存分配的次数。
- **适用场景**：
  - 适合频繁拼接或大量拼接字符串的场景，比如在循环中反复拼接。
  - 性能最佳，是应对大量拼接需求的首选方案。

5、**使用 `[]byte` 切片**

使用 `[]byte` 切片是另一种高效的拼接方式，它通过操作字节切片来拼接字符串，最后再将字节切片转换为字符串。

```
str1 := "Hello"
str2 := "World"
byteSlice := append([]byte(str1), ' ', []byte(str2)...)
str := string(byteSlice)
```

- **性能开销**：
  - `[]byte` 切片拼接相对高效，因为你直接操作字节数组而不是创建新的字符串对象。
  - 每次使用 `append` 时可能需要扩展切片，但相比于 `+` 运算符，它的内存分配次数更少。
  - 性能虽然不错，但手动处理字节切片增加了代码的复杂性。
- **适用场景**：
  - 适用于高级性能优化场景，比如对二进制数据或字符串进行混合处理的场景。
  - 不推荐作为日常使用的拼接方法，除非有明确的性能要求。

#### 性能对比与总结

| 拼接方法          | 简单性 | 性能 | 内存分配次数 | 适用场景                                 |
| ----------------- | ------ | ---- | ------------ | ---------------------------------------- |
| `+`               | 高     | 较低 | 高           | 简单拼接，少量字符串                     |
| `strings.Join`    | 中     | 高   | 低           | 拼接字符串切片，性能要求较高的场景       |
| `fmt.Sprintf`     | 中     | 较低 | 高           | 需要格式化的场景                         |
| `strings.Builder` | 中     | 极高 | 低           | 大量或频繁拼接，尤其是在循环中           |
| `[]byte`          | 低     | 高   | 较低         | 高级性能优化，处理二进制或混合数据的场景 |



## 数据转换

![1726629652715](C:\Users\admin\AppData\Roaming\Typora\typora-user-images\1726629652715.png)

## 泛型

![1727511607887](C:\Users\admin\AppData\Roaming\Typora\typora-user-images\1727511607887.png)

# 结构体

### 取结构体的地址实例化

使用&对结构体进行取地址操作相当于对该结构体类型进行了一次new实例化操作。

```go
p3 := &person{}
fmt.Printf("%T\n", p3)     //*main.person
fmt.Printf("p3=%#v\n", p3) //p3=&main.person{name:"", city:"", age:0}
p3.name = "博客"
p3.age = 30
p3.city = "成都"
fmt.Printf("p3=%#v\n", p3) //p3=&main.person{name:"博客", city:"成都", age:30}
```

p3.name = “博客”其实在底层是(*p3).name = “博客”，这是Go语言帮我们实现的语法糖。

### 结构体初始化

```go
type person struct {
    name string
    city string
    age  int8
}

func main() {
    var p4 person
    fmt.Printf("p4=%#v\n", p4) //p4=main.person{name:"", city:"", age:0}
} 
```

### 使用键值对初始化

使用键值对对结构体进行初始化时，键对应结构体的字段，值对应该字段的初始值。

```go
p5 := person{
    name: "pprof.cn",
    city: "北京",
    age:  18,
}
fmt.Printf("p5=%#v\n", p5) //p5=main.person{name:"pprof.cn", city:"北京", age:18}
```

也可以对结构体指针进行键值对初始化，例如：

```go
p6 := &person{
    name: "pprof.cn",
    city: "北京",
    age:  18,
}
fmt.Printf("p6=%#v\n", p6) //p6=&main.person{name:"pprof.cn", city:"北京", age:18} 
```

当某些字段没有初始值的时候，该字段可以不写。此时，没有指定初始值的字段的值就是该字段类型的零值。

```go
p7 := &person{
    city: "北京",
}
fmt.Printf("p7=%#v\n", p7) //p7=&main.person{name:"", city:"北京", age:0}
```

### 使用值的列表初始化

初始化结构体的时候可以简写，也就是初始化的时候不写键，直接写值：

```go
p8 := &person{
    "pprof.cn",
    "北京",
    18,
}
fmt.Printf("p8=%#v\n", p8) //p8=&main.person{name:"pprof.cn", city:"北京", age:18}
```

使用这种格式初始化时，需要注意：

```
    1.必须初始化结构体的所有字段。
    2.初始值的填充顺序必须与字段在结构体中的声明顺序一致。
    3.该方式不能和键值初始化方式混用。 
```

### 方法和接收者

### 指针类型的接收者

指针类型的接收者由一个结构体的指针组成，由于指针的特性，调用方法时修改接收者指针的任意成员变量，在方法结束后，修改都是有效的。这种方式就十分接近于其他语言中面向对象中的this或者self。 例如我们为Person添加一个SetAge方法，来修改实例变量的年龄。

```
    // SetAge 设置p的年龄
    // 使用指针接收者
    func (p *Person) SetAge(newAge int8) {
        p.age = newAge
    }
```

调用该方法：

```go
func main() {
    p1 := NewPers
    on("测试", 25)
    fmt.Println(p1.age) // 25
    p1.SetAge(30)
    fmt.Println(p1.age) // 30
} 
```

### 值类型的接收者

当方法作用于值类型接收者时，Go语言会在代码运行时将接收者的值复制一份。在值类型接收者的方法中可以获取接收者的成员值，但修改操作只是针对副本，无法修改接收者变量本身。

```go
// SetAge2 设置p的年龄
// 使用值接收者
func (p Person) SetAge2(newAge int8) {
    p.age = newAge
}

func main() {
    p1 := NewPerson("测试", 25)
    p1.Dream()
    fmt.Println(p1.age) // 25
    p1.SetAge2(30) // (*p1).SetAge2(30)
    fmt.Println(p1.age) // 25
} 
```

### 什么时候应该使用指针类型接收者

```
    1.需要修改接收者中的值
    2.接收者是拷贝代价比较大的大对象
    3.保证一致性，如果有某个方法使用了指针接收者，那么其他的方法也应该使用指针接收者。
```

### 任意类型添加方法

在Go语言中，接收者的类型可以是任何类型，不仅仅是结构体，任何类型都可以拥有方法。 举个例子，我们基于内置的int类型使用type关键字可以定义新的自定义类型，然后为我们的自定义类型添加方法。

```go
//MyInt 将int定义为自定义MyInt类型
type MyInt int

//SayHello 为MyInt添加一个SayHello的方法
func (m MyInt) SayHello() {
    fmt.Println("Hello, 我是一个int。")
}
func main() {
    var m1 MyInt
    m1.SayHello() //Hello, 我是一个int。
    m1 = 100
    fmt.Printf("%#v  %T\n", m1, m1) //100  main.MyInt
} 
```

注意事项： 非本地类型不能定义方法，也就是说我们不能给别的包的类型定义方法。





# select

## select语句

select 语句类似于 switch 语句，但是select会随机执行一个可运行的case。如果没有case可运行，它将阻塞，直到有case可运行。

select 是Go中的一个控制结构，类似于用于通信的switch语句。每个case必须是一个通信操作，要么是发送要么是接收。 select 随机执行一个可运行的case。如果没有case可运行，它将阻塞，直到有case可运行。一个默认的子句应该总是可运行的。

### 语法

Go 编程语言中 select 语句的语法如下：

```
select {
    case communication clause  :
       statement(s);      
    case communication clause  :
       statement(s);
    /* 你可以定义任意数量的 case */
    default : /* 可选 */
       statement(s);
}
```

以下描述了 select 语句的语法：

```
    每个case都必须是一个通信
    所有channel表达式都会被求值
    所有被发送的表达式都会被求值
    如果任意某个通信可以进行，它就执行；其他被忽略。
    如果有多个case都可以运行，Select会随机公平地选出一个执行。其他不会执行。
    否则：
    如果有default子句，则执行该语句。
    如果没有default字句，select将阻塞，直到某个通信可以运行；Go不会重新对channel或值进行求值。
```

### 实例：

```go
package main

import "fmt"

func main() {
   var c1, c2, c3 chan int
   var i1, i2 int
   select {
      case i1 = <-c1:
         fmt.Printf("received ", i1, " from c1\n")
      case c2 <- i2:
         fmt.Printf("sent ", i2, " to c2\n")
      case i3, ok := (<-c3):  // same as: i3, ok := <-c3
         if ok {
            fmt.Printf("received ", i3, " from c3\n")
         } else {
            fmt.Printf("c3 is closed\n")
         }
      default:
         fmt.Printf("no communication\n")
   }    
}
```

以上代码执行结果为：

```
    no communication
```

select可以监听channel的数据流动

select的用法与switch语法非常类似，由select开始的一个新的选择块，每个选择条件由case语句来描述

与switch语句可以选择任何使用相等比较的条件相比，select由比较多的限制，其中最大的一条限制就是每个case语句里必须是一个IO操作

```
    select { //不停的在这里检测
    case <-chanl : //检测有没有数据可以读
    //如果chanl成功读取到数据，则进行该case处理语句
    case chan2 <- 1 : //检测有没有可以写
    //如果成功向chan2写入数据，则进行该case处理语句


    //假如没有default，那么在以上两个条件都不成立的情况下，就会在此阻塞//一般default会不写在里面，select中的default子句总是可运行的，因为会很消耗CPU资源
    default:
    //如果以上都没有符合条件，那么则进行default处理流程
    }
```

在一个select语句中，Go会按顺序从头到尾评估每一个发送和接收的语句。

如果其中的任意一个语句可以继续执行（即没有被阻塞），那么就从那些可以执行的语句中任意选择一条来使用。 如果没有任意一条语句可以执行（即所有的通道都被阻塞），那么有两种可能的情况： ①如果给出了default语句，那么就会执行default的流程，同时程序的执行会从select语句后的语句中恢复。 ②如果没有default语句，那么select语句将被阻塞，直到至少有一个case可以进行下去。



# 函数

## 不同文件属于同一包

使用 `go run .` 或 `go build` 来编译或运行代码是 Go 语言的推荐做法，原因如下：

1. **编译整个包**：在一个目录下的多个 Go 文件被视为一个包。使用 `go run .` 或 `go build` 可以确保所有文件都被编译在一起，从而能够正确地解析函数和变量之间的引用。如果直接用 `go run main.go`，只有 `main.go` 文件会被编译，可能会导致未定义的错误。
2. **依赖管理**：Go 会自动处理包的依赖关系。通过运行整个目录，可以确保所有相关文件都被编译，避免遗漏某些文件。
3. **可执行文件生成**：使用 `go build` 会生成一个可执行文件，这样可以方便地在不同环境中运行，而不仅仅是在开发环境中。
4. **简化运行命令**：使用 `go run .` 或 `go build` 使得管理多个文件变得简单。你不需要逐个指定文件名，只需指定包（目录）。
5. **提高效率**：在大型项目中，逐个文件编译可能效率低下，而使用 `go run .` 或 `go build` 可以一次性编译所有内容，节省时间。

##  延迟调用（defer）

### Golang延迟调用：

#### defer特性：

```
    1. 关键字 defer 用于注册延迟调用。
    2. 这些调用直到 return 前才被执。因此，可以用来做资源清理。
    3. 多个defer语句，按先进后出的方式执行。
    4. defer语句中的变量，在defer声明时就决定了。
```

#### defer用途：

```
    1. 关闭文件句柄
    2. 锁资源释放
    3. 数据库连接释放
```

go语言 defer

go 语言的defer功能强大，对于资源管理非常方便，但是如果没用好，也会有陷阱。

defer 是先进后出, 这个很自然,后面的语句会依赖前面的资源，因此如果先前面的资源先释放了，后面的语句就没法执行了。

```go
package main

import "fmt"

func main() {
    var whatever [5]struct{}

    for i := range whatever {
        defer fmt.Println(i)
    }
} 
```

输出结果：

```
    4
    3
    2
    1
    0
```

#### defer 碰上闭包

```go
package main

import "fmt"

func main() {
    var whatever [5]struct{}
    for i := range whatever {
        defer func() { fmt.Println(i) }()
    }
} 
```

输出结果：

```
    4
    4
    4
    4
    4
```

其实go说的很清楚,我们一起来看看go spec如何说的

Each time a “defer” statement executes, the function value and parameters to the call are evaluated as usualand saved anew but the actual function is not invoked.

也就是说函数正常执行,由于闭包用到的变量 i 在执行的时候已经变成4,所以输出全都是4.

​		如果你希望每次 `defer` 的输出都对应 `i` 当时的值，可以在循环中将 `i` 的值作为参数传递给匿名函数，这样就能捕获 `i` 当时的值，而不是最终的引用。

```go
go复制代码func main() {
    var whatever [5]struct{}
    for i := range whatever {
        defer func(i int) { fmt.Println(i) }(i)
    }
}
```

在这个修正版代码中，`defer` 的匿名函数通过参数 `i int` 捕获了 `i` 当时的值，因此输出将会是：

```go
plaintext复制代码4
3
2
1
0
```

```go
//defer的一个例子
func Def() (i int) {
	i = 9
	defer func() { //打印函数返回前最后赋值完的值  “defer在函数返回前执行”
		fmt.Printf("1 rst i=%d\n", i)
	}()
	func() {
		fmt.Printf("second =:%d\n", i)
	}()
	defer func(i int) {
		fmt.Printf("third is:%d\n", i)
		i += 5
		fmt.Printf("third +5 is:%d\n", i)
	}(i) //defer将此时的i拷贝，与后边i的变化无关
	defer fmt.Printf("fourth i is:%d\n", i) //defer 后跟一个语句，则相关变量在注册defer时被拷贝或者计算
	i = 999
	return 666
}
```

### defer panic

panic 会执行什么：

1、逆序执行当前goroutine 的defer链（recover从这里介入）

2、打印错误信息和调用堆栈

3、调用exit（2）结束整个进程。

```go
func main() {
	defer fmt.Println("11111111")

	fmt.Println("AAAAAAAA")
	defer fmt.Println("2222222222")
	fmt.Println("BBBBBBBBB")
	defer fmt.Println("33333333")
	panic("OOOOOOOOOOps")
	defer fmt.Println("4444444")
	fmt.Println("CCCCCCCC")
}
```

输出：

```go
AAAAAAAA
BBBBBBBBB
33333333
2222222222
11111111
panic: OOOOOOOOOOps  

goroutine 1 [running]:
main.main()
        F:/mystudy/base/defer/main.go:12 +0x18c
exit status 2
```

即：遇到恐慌触发错误退出函数，在退出函数之前的最后一步，将该代码处之前注册好的defer执行完毕。（defer的执行顺序就是在函数退出前逆序执行）

### 闭包（Closure）

闭包是一个函数，**它引用了其外部作用域中的变量**。

闭包的特性：

- 闭包可以访问其外部函数的变量，即使这些变量在外部函数执行完毕后也依然存在。

- 闭包会“捕获”**外部变量的引用，而不是值**。这意味着如果外部变量的值在闭包创建后发生了变化，闭包访问的将是最新的值。

- ```go
  func main() {
      var x int = 10
  
      // 定义一个闭包，捕获了外部变量 x
      increment := func() int {
          x++
          return x
      }
  
      fmt.Println(increment()) // 输出：11
      fmt.Println(increment()) // 输出：12
  
      // 注意：x 在闭包之外的值也已经改变
      fmt.Println(x) // 输出：12
  }
  
  ```

  ![1726831914803](C:\Users\admin\AppData\Roaming\Typora\typora-user-images\1726831914803.png)

  ### 匿名函数（Anonymous Function）

  匿名函数，顾名思义，就是没有名字的函数。在 Go 中，你可以定义和使用匿名函数，就像使用普通函数一样。匿名函数可以立即执行，也可以赋值给变量或作为参数传递给其他函数。

  #### 示例：定义和使用匿名函数
  
  ```go
  go复制代码func main() {
      // 定义并立即执行匿名函数
      func() {
          fmt.Println("Hello, World!")
      }() // 这个括号表示立即执行函数
  
      // 将匿名函数赋值给变量
      add := func(a, b int) int {
          return a + b
      }
  
      fmt.Println(add(3, 4)) // 输出：7
  }
  ```

# 文件

## 打开

常见的两种打开文件的方式是使用`os`包提供的两个函数，`Open`函数返回值一个文件指针和一个错误，

```go
func Open(name string) (*File, error)
```

后者`OpenFile`能够提供更加细粒度的控制，实际上`Open`函数就是对`OpenFile`函数的一个简单封装。

```go
func OpenFile(name string, flag int, perm FileMode) (*File, error)
```

先来介绍第一种使用方法，直接提供对应的文件名即可，代码如下

```go
func main() {
   file, err := os.Open("README.txt")
   fmt.Println(file, err)
}
```

文件的查找路径默认为项目`go.mod`文件所在的路径，由于项目下并没有该文件，所以自然会返回一个错误。

```text
<nil> open README.txt: The system cannot find the file specified.
```

 因为IO错误的类型有很多，所以需要手动的去判断文件是否存在，同样的`os`包也为此提供了方便函数，修改后的代码如下 

```go
func main() {
	file, err := os.Open("README.txt")
	if os.IsNotExist(err) {
		fmt.Println("文件不存在")
	} else if err != nil {
		fmt.Println("文件访问异常")
	} else {
		fmt.Println("文件读取成功", file)
	}
}
```

事实上第一种函数读取的文件仅仅只是只读的，无法被修改，`Open`函数内部实现



```go
func Open(name string) (*File, error) {
	return OpenFile(name, O_RDONLY, 0)
}
```

通过`OpenFile`函数可以控制更多细节，例如修改文件描述符和文件权限，关于文件描述符，

### `os`包下提供了以下常量以供使用。

```go
const (
   // 只读，只写，读写 三种必须指定一个
   O_RDONLY int = syscall.O_RDONLY // 以只读的模式打开文件
   O_WRONLY int = syscall.O_WRONLY // 以只写的模式打开文件
   O_RDWR   int = syscall.O_RDWR   // 以读写的模式打开文件
   // 剩余的值用于控制行为
   O_APPEND int = syscall.O_APPEND // 当写入文件时，将数据添加到文件末尾
   O_CREATE int = syscall.O_CREAT  // 如果文件不存在则创建文件
   O_EXCL   int = syscall.O_EXCL   // 与O_CREATE一起使用, 文件必须不存在
   O_SYNC   int = syscall.O_SYNC   // 以同步IO的方式打开文件
   O_TRUNC  int = syscall.O_TRUNC  // 当打开的时候截断可写的文件
    //当指定了 os.O_TRUNC 标志时，如果文件已经存在，文件会在打开时被截断为零长度，即清空文件内容。如果文件不存在，则创建一个空文件
)
```

### 关于文件权限的则提供了以下常量。

```go
const (
   ModeDir        = fs.ModeDir        // d: 目录
   ModeAppend     = fs.ModeAppend     // a: 只能添加
   ModeExclusive  = fs.ModeExclusive  // l: 专用
   ModeTemporary  = fs.ModeTemporary  // T: 临时文件
   ModeSymlink    = fs.ModeSymlink    // L: 符号链接
   ModeDevice     = fs.ModeDevice     // D: 设备文件
   ModeNamedPipe  = fs.ModeNamedPipe  // p: 具名管道 (FIFO)
   ModeSocket     = fs.ModeSocket     // S: Unix 域套接字
   ModeSetuid     = fs.ModeSetuid     // u: setuid
   ModeSetgid     = fs.ModeSetgid     // g: setgid
   ModeCharDevice = fs.ModeCharDevice // c: Unix 字符设备, 前提是设置了 ModeDevice
   ModeSticky     = fs.ModeSticky     // t: 黏滞位
   ModeIrregular  = fs.ModeIrregular  // ?: 非常规文件

   // 类型位的掩码. 对于常规文件而言，什么都不会设置.
   ModeType = fs.ModeType

   ModePerm = fs.ModePerm // Unix 权限位, 0o777
)
```

提示

`truncates`意思即为将文件容量调整到合适的大小以容纳数据，不大也不小。

下面是一个以读写模式打开一个文件的代码例子，权限为`0666`，表示为所有人都可以对该文件进行读写，且不存在时会自动创建。

```go
func main() {
	file, err := os.OpenFile("README.txt", os.O_RDWR|os.O_CREATE, 0666)
	if os.IsNotExist(err) {
		fmt.Println("文件不存在")
	} else if err != nil {
		fmt.Println("文件访问异常")
	} else {
		fmt.Println("文件打开成功", file.Name())
		file.Close()
	}
}
```

输出如下

```text
文件打开成功 README.txt
```

倘若只是想获取该文件的一些信息，并不想读取该文件，可以使用`os.Lstat()`函数进行操作，代码示例如下

```go
func main() {
	fileInfo, err := os.Lstat("README.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(fmt.Sprintf("%+v", fileInfo))
	}
}
```

输出如下

```text
&{name:README.txt FileAttributes:32 CreationTime:{LowDateTime:3603459389 HighDateTime:31016791} LastAccessTime:{LowDateTime:3603459389 HighDateTime:31016791} LastWriteTime:{LowDateTime:3603459389 HighDateTime:31016791} FileSizeHigh
:0 FileSizeLow:0 Reserved0:0 filetype:0 Mutex:{state:0 sema:0} path:README.txt vol:0 idxhi:0 idxlo:0 appendNameToPath:false}
```

注意

打开一个文件后永远要记得关闭该文件，通常关闭操作会放在`defer`语句里

```go
defer file.Close()
```

## 关闭

以下是如何调整代码以解决此警告的示例：

```go
goCopy codepackage main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("打开文件时发生错误:", err)
		return
	}
	defer func() {
		if cerr := file.Close(); cerr != nil {
			fmt.Println("关闭文件时发生错误:", cerr)
		}
	}()

	// 其余读取或处理文件的代码
}
```

在这个修改后的例子中，`file.Close()` 被延迟到一个匿名函数中，并在该函数内部检查了由 `file.Close()` 返回的错误。这确保了在关闭操作期间发生的任何错误都会得到适当处理。

## [#](https://golang.xiniushu.com/语言入门/语法进阶/100.io.html#读取)读取

当成功的打开文件后，便可以进行读取操作了，关于读取文件的操作，`os.File`类型提供了以下几个公开的方法

```go
// 将文件读进传入的字节切片
func (f *File) Read(b []byte) (n int, err error) 

// 相较于第一种可以从指定偏移量读取
func (f *File) ReadAt(b []byte, off int64) (n int, err error) 
```

大多数情况第一种使用的较多。针对于第一种方法，需要自行编写逻辑来进行读取时切片的动态扩容，代码如下

```go
func ReadFile(file *os.File) ([]byte, error) {
	buffer := make([]byte, 0, 512)
	for {
		// 当容量不足时
		if len(buffer) == cap(buffer) {
			// 扩容
			buffer = append(buffer, 0)[:len(buffer)]
		}
		// 继续读取文件
		offset, err := file.Read(buffer[len(buffer):cap(buffer)])
		// 将已写入的数据归入切片
		buffer = buffer[:len(buffer)+offset]
		// 发生错误时
		if err != nil {
			if errors.Is(err, io.EOF) {
				err = nil
			}
			return buffer, err
		}
	}
}
```

剩余逻辑如下

```go
func main() {
   file, err := os.OpenFile("README.txt", os.O_RDWR|os.O_CREATE, 0666)
   if err != nil {
      fmt.Println("文件访问异常")
   } else {
      fmt.Println("文件打开成功", file.Name())
      bytes, err := ReadFile(file)
      if err != nil {
         fmt.Println("文件读取异常", err)
      } else {
         fmt.Println(string(bytes))
      }
      file.Close()
   }
}
```

输出为

```text
文件打开成功 README.txt
hello world!
```

   除此之外，还可以使用两个方便函数来进行文件读取，分别是`os`包下的`ReadFile`函数，以及`io`包下的`ReadAll`函数。对于`os.ReadFile`而言，只需要提供文件路径即可，而对于`io.ReadAll`而言，则需要提供一个`io.Reader`类型的实现，

### **os.ReadFile**

```go
func ReadFile(name string) ([]byte, error)
```

使用例子如下：

```go
func main() {
	bytes, err := os.ReadFile("README.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(bytes))
	}
}
```

输出如下

```text
hello world!
```

### **io.ReadAll**

```go
func ReadAll(r Reader) ([]byte, error)
```

使用例子如下

```go
func main() {

   file, err := os.OpenFile("README.txt", os.O_RDWR|os.O_CREATE, 0666)
   if err != nil {
      fmt.Println("文件访问异常")
   } else {
      fmt.Println("文件打开成功", file.Name())
      bytes, err := io.ReadAll(file)
      if err != nil {
         fmt.Println(err)
      } else {
         fmt.Println(string(bytes))
      }
      file.Close()
   }
}
```

```text
文件打开成功 README.txt
hello world!
```



## [#](https://golang.xiniushu.com/语言入门/语法进阶/100.io.html#写入)写入

`os.File`结构体提供了以下几种方法以供写入数据

### os.File

```go
// 写入字节切片
func (f *File) Write(b []byte) (n int, err error)

// 写入字符串
func (f *File) WriteString(s string) (n int, err error)

// 从指定位置开始写，当以os.O_APPEND模式打开时，会返回错误
func (f *File) WriteAt(b []byte, off int64) (n int, err error)
```

如果想要对一个文件写入数据，则必须以`O_WRONLY`或`O_RDWR`的模式打开，否则无法成功写入文件。下面是一个以`os.O_RDWR|os.O_CREATE|os.O_APPEND|os.O_TRUNC`模式打开文件，且权限为`0666`向指定写入数据的例子

```go
func main() {
	file, err := os.OpenFile("README.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("文件访问异常")
	} else {
		fmt.Println("文件打开成功", file.Name())
		for i := 0; i < 5; i++ {
			offset, err := file.WriteString("hello world!\n")
			if err != nil {
				fmt.Println(offset, err)
			}
		}
		fmt.Println(file.Close())
	}
}
```

由于是以`os.O_APPEND`模式打开的文件，所以在写入文件时会将数据添加到文件尾部，执行完毕后文件内容如下

```txt
hello world!
hello world!
hello world!
hello world!
hello world!
```

向文件写入字节切片也是类似的操作，就不再赘述。对于写入文件的操作标准库同样提供了方便函数，分别是`os.WriteFile`与`io.WriteString`

### **os.WriteFile**

```go
func WriteFile(name string, data []byte, perm FileMode) error
```

使用例子如下

```go
func main() {
	err := os.WriteFile("README.txt", []byte("hello world!\n"), 0666)
	if err != nil {
		fmt.Println(err)
	}
}
```

此时文件内容如下

```txt
hello world!
```

### **io.WriteString**

```go
func WriteString(w Writer, s string) (n int, err error) 
```

使用例子如下

```go
func main() {
   file, err := os.OpenFile("README.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND|os.O_TRUNC, 0666)
   if err != nil {
      fmt.Println("文件访问异常")
   } else {
      fmt.Println("文件打开成功", file.Name())
      for i := 0; i < 5; i++ {
         offset, err := io.WriteString(file, "hello world!\n")
         if err != nil {
            fmt.Println(offset, err)
         }
      }
      fmt.Println(file.Close())
   }
}
```

```text
hello world!
hello world!
hello world!
hello world!
hello world!
```

提示

`os.Create`函数用于创建文件，本质上也是对`OpenFile`的封装。

```go
func Create(name string) (*File, error) {
   return OpenFile(name, O_RDWR|O_CREATE|O_TRUNC, 0666)
}
```

注意:在创建一个文件时，如果其父目录不存在，将创建失败并会返回错误。

## [#](https://golang.xiniushu.com/语言入门/语法进阶/100.io.html#复制)复制

对于复制文件而言，需要同时打开两个文件，第一种方法是将原文件中的数据读取出来，然后写入目标文件中，代码示例如下

```go
func main() {
    // 从原文件中读取数据
	data, err := os.ReadFile("README.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
    // 写入目标文件
	err = os.WriteFile("README(1).txt", data, 0666)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("复制成功")
	}
}
```

### ***os.File.ReadFrom**

另一种方法是使用`os.File`提供的方法`ReadFrom`，打开文件时，一个只读，一个只写。

```go
func (f *File) ReadFrom(r io.Reader) (n int64, err error)
```

使用示例如下

```go
func main() {
    // 以只读的方式打开原文件
	origin, err := os.OpenFile("README.txt", os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
    defer origin.Close()
    // 以只写的方式打开副本文件
	target, err := os.OpenFile("README(1).txt", os.O_RDONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
    defer target.Close()
    // 从原文件中读取数据，然后写入副本文件
	offset, err := target.ReadFrom(origin)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("文件复制成功", offset)
}
```

### **io.Copy**

还有一种方法就是使用`io.Copy`方便函数

```go
func Copy(dst Writer, src Reader) (written int64, err error)
```

使用示例如下

```go
func main() {
	// 以只读的方式打开原文件
	origin, err := os.OpenFile("README.txt", os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
    defer origin.Close()
	// 以只写的方式打开副本文件
	target, err := os.OpenFile("README(1).txt", os.O_RDONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
    defer target.Close()
	// 复制
	written, err := io.Copy(target, origin)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(written)
	}
}
```

提示

`io.Copy`函数对于文件夹复制同样有效。