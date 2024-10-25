# Go compendium
Go language compendium


---
### Index
- [Why Go?](#why-go)
- [Why Go Modules Replace GOPATH?](#why-go-modules-replace-gopath)
- [Go Commands summary](#go-commands-summary)
- [Module paths for downloadable packages](#module-paths-for-downloadable-packages)
- [Packages, Variables, and Constants](#packages-variables-and-constants)
- [Data Types](#data-types)
- [Pointers](#pointers)
- [Arrays and Slices](#arrays-and-slices)
- [Loops](#loops)
- [Conditionals](#conditionals)
- [Operators](#operators)
- [Switch Case](#switch-case)
- [Memory Stack and Heap](#memory-stack-and-heap)
- [Functions](#functions)
- [Packages and Scopes](#packages-and-scopes)
- [Maps](#maps)
- [Structs](#structs)
- [`goroutines` - Concurrency](#goroutines---concurrency)
- [Interfaces and Structs](#interfaces-and-structs)
- [Data types and recommendations on when to pass them by value or reference](#data-types-and-recommendations-on-when-to-pass-them-by-value-or-reference)
- [Error handling](#error-handling)
- [Functional programming](#functional-programming)
- [Generics](#generics)
- [Variadic functions](#variadic-functions)
- [Runes](#runes)
- [Testing](#testing)
- [Templates](#templates)
- [Further Reading](#further-reading)

---
### Why Go?

- **Built-in Concurrency Mechanism** that was **designed to run on multiple cores** and to support concurrency.
- **Concurrency is cheap and easy** compared to other languages like C++ or Java.
- For **Performant** applications and running on **scaled, distributed systems**.
- E.g. *Docker and K8s* are written in GO.
- Faster than interpreted languages like Python. 


---
### Why Go Modules Replace GOPATH:

Go modules use the `go.mod` file to manage dependencies for each project, and you can create Go projects anywhere in your file system, not just within the `GOPATH` directory.
Each project is self-contained, meaning that each project has its own module file (`go.mod`) to manage its dependencies independently of other projects.

*By default, `GOPATH` is set to `~/go`, but this is only relevant if you're using legacy projects that rely on the `GOPATH` structure*.
For modern Go development using modules, there's no need to manually manage `GOPATH` unless you're working on older projects.

Only need to set `GOPATH` in the following cases:
- You're working on older Go projects that still rely on the `GOPATH` workspace model.
- You prefer to have a central location where Go binaries (go install) and packages are installed, *but even this is optional since modules handle dependencies locally*.


---
### Go Commands summary:

| **Command**           | **Usage**                          | **Description**                                                                 | **Example**                                                                 |
|-----------------------|------------------------------------|---------------------------------------------------------------------------------|-----------------------------------------------------------------------------|
| `go mod init`         | `go mod init module-name`           | Initializes a new Go module. Creates a `go.mod` file for dependency management. | `go mod init github.com/yourusername/gocomp`                                |
| `go mod tidy`         | `go mod tidy`                       | Cleans up `go.mod` and `go.sum` by removing unused dependencies and adding missing ones. | `go mod tidy`                                                               |
| `go build`            | `go build [package]`                | Compiles Go source code into an executable binary.                              | `go build` (builds the project in the current directory)                    |
| `go run`              | `go run file.go`                    | Compiles and runs the specified Go source file(s).                              | `go run main.go`                                                            |
| `go test`             | `go test [package]`                 | Runs tests in the specified package or directory.                               | `go test ./...` (runs all tests in the current directory and subdirectories)|
| `go fmt`              | `go fmt [files or packages]`        | Formats Go source code according to the standard Go style.                       | `go fmt ./...` (formats all Go files in the current directory and subdirectories) |
| `go get`              | `go get package`                    | Downloads and installs the specified package along with its dependencies.       | `go get github.com/some/package`                                            |
| `go install`          | `go install package`                | Compiles and installs the specified Go package.                                 | `go install github.com/yourusername/gocomp` (installs the `gocomp` package) |
| `go list`             | `go list [package]`                 | Lists Go packages, including their module path and metadata.                    | `go list github.com/yourusername/gocomp`                                    |
| `go mod edit`         | `go mod edit -require=module@version` | Edits the `go.mod` file to add or remove dependencies.                          | `go mod edit -require=github.com/some/package@v1.2.3`                      |
| `go env`              | `go env`                            | Displays Go environment variables and configuration.                            | `go env`                                                                     |
| `go doc`              | `go doc [package]`                  | Displays documentation for the specified package or symbol.                     | `go doc github.com/yourusername/gocomp`                                     |
| `go version`          | `go version`                        | Shows the installed Go version.                                                  | `go version`                                                                 |
| `go mod why`          | `go mod why package`                | Explains why a specific module is needed by your module.                        | `go mod why github.com/some/package`                                        |
| `go mod vendor`       | `go mod vendor`                     | Creates a `vendor` directory with copies of all dependencies for offline builds. | `go mod vendor`                                                              |


---
### Module paths for downloadable packages:

If you’re creating a project which can be downloaded and used by other people and programs, then it’s good practice for your module path to equal the location that the code can be downloaded from.

For instance, if your package is hosted at https://github.com/foo/bar then the module path for the project should be github.com/foo/bar.


---
### Packages, Variables, and Constants:

- All code must belong to **packages**, i.e. all code is organized in **packages**. The first statement in a file should be a package declaration. **A package is a set of related source files** with their functions, e.g. [`fmt`](https://pkg.go.dev/fmt) and its functions to print in different formats like `Println` or `Printf` using printing verbs like `%v`, `%s`, `%d`, `%t`, etc.
- **Variables** are used to store values and reuse/update values like *containers*. Variables are declared with the `var` keyword.
- **Go Compile Errors to enforce better code quality**, e.g. leaving a `var` without a call/usage highlights an error (variables must be used) or trying to update a `const` value from 50 to 30. 
- **Go is a Statically Typed** language, i.e. Go Compiler will throw an error if a variable is not declared with a type, this is called **Type Checking**, unless the type of the variable could be inferred from its assigned value during the declaration, e.g. `a := 10`, `var a = 10`, `var a int = 10` are valid declarations in Go but `var a // declaration without assignment and after few lines of code being assigned with a = 10` is an error. 
**Type Inference** is when Go compiler infers the type of variable based on the assigned value. *Note `:=` can only be used in variables and NO to `const`*


---
### Data Types:

| Data Type   | Description                                  | Example                                  |
|-------------|----------------------------------------------|------------------------------------------|
| **bool**    | Represents a Boolean value (true or false).  | `var isActive bool = true`               |
| **string**  | Sequence of characters.                      | `var name string = "GoLang"`             |
| **int**     | Signed integer (size depends on platform).   | `var age int = 30`                       |
| **int8**    | 8-bit signed integer (-128 to 127).          | `var smallNum int8 = -10`                |
| **int16**   | 16-bit signed integer (-32,768 to 32,767).   | `var mediumNum int16 = 300`              |
| **int32**   | 32-bit signed integer (-2^31 to 2^31-1).     | `var largeNum int32 = 100000`            |
| **int64**   | 64-bit signed integer (-2^63 to 2^63-1).     | `var bigNum int64 = 100000000000`        |
| **uint**    | Unsigned integer (size depends on platform). | `var index uint = 10`                    |
| **uint8**   | 8-bit unsigned integer (0 to 255).           | `var byteVal uint8 = 255`                |
| **uint16**  | 16-bit unsigned integer (0 to 65,535).       | `var smallIndex uint16 = 500`            |
| **uint32**  | 32-bit unsigned integer (0 to 4,294,967,295).| `var mediumIndex uint32 = 100000`        |
| **uint64**  | 64-bit unsigned integer (0 to 2^64-1).       | `var largeIndex uint64 = 1000000000`     |
| **float32** | 32-bit floating-point number.                | `var price float32 = 9.99`               |
| **float64** | 64-bit floating-point number.                | `var pi float64 = 3.14159`               |
| **complex64** | Complex number with float32 real and imaginary parts. | `var c complex64 = 2 + 3i`      |
| **complex128**| Complex number with float64 real and imaginary parts. | `var c complex128 = 2 + 3i`   |
| **byte**    | Alias for `uint8`, represents a byte.        | `var b byte = 255`                       |
| **rune**    | Alias for `int32`, represents a Unicode character. | `var char rune = 'A'`               |
| **array**   | Fixed-size sequence of elements of the same type. | `var nums [3]int = [3]int{1, 2, 3}`   |
| **slice**   | Dynamic-size sequence of elements of the same type. | `var nums []int = []int{1, 2, 3}`   |
| **map**     | Key-value pairs, where keys and values can be of any type. | `var dict map[string]int = map[string]int{"a": 1}` |
| **struct**  | Collection of fields.                        | `type Person struct { Name string; Age int }` |
| **pointer** | Holds the memory address of a value.         | `var ptr *int = &age`                    |
| **interface** | Abstract type to represent any type.       | `var i interface{} = "hello"`            |
| **function** | A function signature with parameters and return values. | `func add(a int, b int) int { return a + b }` |


---
### Pointers:

Variables are stored in memory but when we reference a variable we are actually pointing to its memory address, i.e.
**a Pointer is a Variable that points to the Memory Address of another Variable** -*special variable* that holds the memory address of the other variable, e.g. printing a pointer returns the memory address of the variable pointed while printing a variable returns the actual value.

![Pointer](./img/0-pointer.png?raw=true)

```go
fmt.Println("Pointer of soldTickets is", &soldTickets, "and its Variable value", soldTickets)	
```	
```text
// output console
Pointer of soldTickets is 0xc0000ac00b and its Variable value 20
```


---
### Arrays and Slices:

**Arrays and Slices** are data structures to **store collections of elements in a Single Variable**.

- **An Array has a Fixed size**, i.e. how many elements the array can hold, 
e.g. `var nums [3]int`, `var nums = [3]int{1, 2, 3}`, `var nums = [3]int{1}` and `var nums = [3]int{}` are valid declarations.
**Arrays are indexed starting from 0**, e.g. `nums[0] = 1` and `nums[1] = 10`. Examples of updating an array inside a function: 
```go
func main() {
    var conferenceName = "Go Conference"
    const totalTickets uint8 = 50
    var remainingTickets uint8 = 30
    soldTickets := calculateSoldTickets(remainingTickets, totalTickets)
    var bookings [totalTickets]string
    initBookings(totalTickets, soldTickets, &bookings)
    // more code here
}

// Must pass booking Pointer otherwise array will be passed by copy and original array values won't be updated.
// An alternative approach is to return the copy of the bookings array and then replacing the original after the function is called.
func initBookings(totalTickets uint8, soldTickets uint8, bookings *[50]string) {
    for i := 0; i < int(totalTickets); i++ {
        if i < int(soldTickets) {
            // must update the actual array value not the pointer, i.e. doing dereference:
            (*bookings)[i] = "SOLD"
        } else {
            // Go has syntactic sugar that simplifies working with pointers to arrays.
            // When you use bookings[i], Go automatically dereferences the pointer for you to access the array element,
            // so you don’t need to explicitly write (*bookings)[i].
            bookings[i] = "AVAILABLE"
        }
    }
}
```

- **Slice is an abstraction of Array that has Dynamic size**, more flexible and powerful, i.e. slices are also **index-based** and **have a size, but is Resized when Needed**. More flexible and powerful than arrays, i.e. **Variable-length** or get a sub-array of its own.
Valid declarations of Slices are `var bookings []string`, `var bookings = []string{}` or `bookings := []string{}` never specifying the size. Also, **using `make()` to declare slices in Go is considered good practice because it allows you to initialize slices with a specific length and capacity, which improves memory management and performance**, e.g.: `var bookings = make([]string, 20, 50) // len=20 and cap=50`.
```go
func main() {
	var conferenceName = "Go Conference"
	const totalTickets uint8 = 50
	var remainingTickets uint8 = 30
	soldTickets := calculateSoldTickets(remainingTickets, totalTickets)
	bookings := make([]string, 0, totalTickets)                  // len=0, cap=50
	bookings = initBookings(totalTickets, soldTickets, bookings) // len=20, cap=50
    // more code here
}

// Bookiing Slice by reference, i.e. the Slice Descriptor pointing to the underlying array.
func initBookings(totalTickets uint8, soldTickets uint8, bookings []string) []string {
	for i := 0; i < int(totalTickets); i++ {
		if i < int(soldTickets) {
			// The append built-in function appends elements to the end of a slice.
			// If it has sufficient capacity, the destination is resliced to accommodate the new elements.
			// If it does not, a new underlying array will be allocated
			bookings = append(bookings, "SOLD")
		} else {
			// leave empty avaialble indexes so the len can be smaller that the slice capacity.
			break
		}
	}
	return bookings
}
```	

⚠️ **Note passing a Slice as a parameter is different from passing an Array as a parameter as GO passes Slices by Reference and Arrays by Value**, i.e.:

**1. Slices Are Passed by Reference**

**A slice in Go is a descriptor** that contains:
- **A pointer to the underlying array**.
- A length (the number of elements in the slice).
- A capacity (the size of the underlying array).

When you pass a slice to a function, Go passes this descriptor (which includes the pointer to the underlying array). This means:
- **Modifications to the slice elements inside the function** will affect the original slice since it references the same underlying array.
- The slice header (pointer, length, capacity) itself is passed by value, but the pointer still points to the same underlying array, allowing shared access to the elements.

Example:
```go
func modifySlice(s []int) {
    s[0] = 100 // This modifies the original array
}

func main() {
    numbers := []int{1, 2, 3}
    modifySlice(numbers) // Modifies the original slice
    fmt.Println(numbers) // Output: [100 2 3]
}
```

**2. Arrays Are Passed by Value**

In contrast, when you pass an array to a function, Go passes **a copy** of the entire array. This means:
- Modifications to the array inside the function will not affect the original array outside the function because the function is working with a **copy** of the original array.
- **Arrays are fixed in size and passing them can be inefficient for large arrays since the entire array is copied**.

Example:
```go
func modifyArray(a [3]int) {
    a[0] = 100 // Modifies only the copy of the array
}

func main() {
    numbers := [3]int{1, 2, 3}
    modifyArray(numbers) // Passes a copy of the array
    fmt.Println(numbers) // Output: [1 2 3] - Original array is unchanged
}
```

**Key Differences Between Array and Slice in General:**

| Aspect               | Slice                                      | Array                                       |
|----------------------|--------------------------------------------|---------------------------------------------|
| **Passing Behavior**  | Passed by reference (pointer to array)     | Passed by value (entire array is copied)    |
| **Size**              | Dynamic, can change size (flexible length) | Fixed size (defined at declaration)         |
| **Modification**      | Changes inside the function affect original | Changes inside the function affect only the copy |
| **Efficiency**        | More efficient for large data since only the descriptor is passed | Less efficient for large arrays as the entire array is copied |
| **Underlying Data**   | Slices reference an underlying array       | Arrays do not have a separate underlying structure |
| **Use Case**          | Preferred when working with collections of unknown or variable size | Used for fixed-size collections              |

**Key Differences Between Array and Slice 'Length':**

| Aspect              | Array                                        | Slice                                      |
|---------------------|----------------------------------------------|--------------------------------------------|
| **Size**            | Fixed size, set at declaration.              | Dynamic size, can grow or shrink.          |
| **`len()` Result**  | Always returns the fixed size of the array.  | Returns the current number of elements.    |
| **Mutability**      | The length cannot be changed after creation. | The length can change as elements are added or removed. |
| **Capacity**        | The length is equals to the Capacity, as it's a fixed size.| Capacity is separate from length and can be larger than the length. |


---
### Loops

A **loop** statement allows you to **execute a block of code multiple times**, in a loop.

**Loops are simplified in GO providing only one of them to use it for every possible case**, i.e. `for` loop -No `while`, `do-while` or `for-each` loops exist.
- **Infinite case**:
```go
for {
    // infinite loop
}
```
- **For by Index case**:
```go
for index := 0; index < 10; index++ {
    fmt.Println(index)
}
```
- **For Range case (for each)**:
Range **iterates over elements for different data structures** (not only arrays and slices).
In array and slices, **range provides the Index and the Value 'for each' element**, i.e.: 
```go
for index, value := range array {
    fmt.Println(i, v)
}  
```
⚠️ Note in go `_` is also known as **Blank identifier** so in case of requiring the Value without the Index setting `_` instead of `index` is a valid sentence, i.e.:
```go
for _, value := range array {
    fmt.Println(v)
}  
```


---
### Conditionals

The **expression that evaluate** to either `true` or `false` **is called a condition**.

**`if-else`:**
```go
if condition {
    // code to execute if condition is true
} else if anotherCondition {
    // code to execute if anotherCondition is true
} else {
    // code to execute if none of the above conditions are true
}
```

**Short Statement with Condition:**

You can also include a short statement before the condition, i.e. ⚠️ **declare a variable in the same line as `if`**:

```go
if y := 5; y > 0 {
    fmt.Println("y is positive")
}
```

Here, `y` is declared and initialized only within the scope of the `if` block.

⚠️ A `true` or `false` boolean conditional can also be applied to **loops**, e.g.:
```go
// infinite loop for ==> for "true", i.e.: 
for { }
	
// loop with true false condition:
for remainingTickets > 0 { }
```


---
### Operators

1. **Arithmetic Operators**

| **Operator** | **Description**     | **Example**        | **Explanation**                                 |
|--------------|---------------------|--------------------|-------------------------------------------------|
| `+`          | Addition             | `x + y`            | Adds `x` and `y`.                               |
| `-`          | Subtraction          | `x - y`            | Subtracts `y` from `x`.                         |
| `*`          | Multiplication       | `x * y`            | Multiplies `x` by `y`.                          |
| `/`          | Division             | `x / y`            | Divides `x` by `y`.                             |
| `%`          | Modulo               | `x % y`            | Remainder of `x` divided by `y`.                |

2. **Comparison Operators**

| **Operator** | **Description**       | **Example**        | **Explanation**                                 |
|--------------|-----------------------|--------------------|-------------------------------------------------|
| `==`         | Equal to               | `x == y`           | True if `x` is equal to `y`.                    |
| `!=`         | Not equal to           | `x != y`           | True if `x` is not equal to `y`.                |
| `>`          | Greater than           | `x > y`            | True if `x` is greater than `y`.                |
| `<`          | Less than              | `x < y`            | True if `x` is less than `y`.                   |
| `>=`         | Greater than or equal  | `x >= y`           | True if `x` is greater than or equal to `y`.    |
| `<=`         | Less than or equal     | `x <= y`           | True if `x` is less than or equal to `y`.       |

3. **Logical Operators**

| **Operator** | **Description**         | **Example**        | **Explanation**                                         |
|--------------|-------------------------|--------------------|---------------------------------------------------------|
| `&&`         | Logical AND              | `x > 0 && y < 10`  | True if both conditions are true.                       |
| `\|\|`         | Logical OR               | `x > 0 \|\| y < 10`  | True if at least one condition is true.                 |
| `!`          | Logical NOT              | `!x`               | Negates the condition (true becomes false, and vice versa). |

4. **Bitwise Operators**

| **Operator** | **Description**         | **Example**        | **Explanation**                                      |
|--------------|-------------------------|--------------------|------------------------------------------------------|
| `&`          | Bitwise AND              | `x & y`            | Performs AND on each pair of bits.                   |
| `\|`          | Bitwise OR               | `x \| y`            | Performs OR on each pair of bits.                    |
| `^`          | Bitwise XOR              | `x ^ y`            | Performs XOR on each pair of bits.                   |
| `&^`         | AND NOT (bit clear)      | `x &^ y`           | Clears bits of `x` where `y` is 1.                   |
| `<<`         | Left shift               | `x << 2`           | Shifts `x` to the left by 2 bits (multiply by 4).     |
| `>>`         | Right shift              | `x >> 2`           | Shifts `x` to the right by 2 bits (divide by 4).      |

5. **Assignment Operators**

| **Operator** | **Description**          | **Example**         | **Explanation**                                        |
|--------------|--------------------------|---------------------|--------------------------------------------------------|
| `=`          | Assignment                | `x = y`             | Assigns the value of `y` to `x`.                       |
| `+=`         | Add and assign            | `x += y`            | Adds `y` to `x` and assigns the result to `x`.         |
| `-=`         | Subtract and assign       | `x -= y`            | Subtracts `y` from `x` and assigns the result to `x`.  |
| `*=`         | Multiply and assign       | `x *= y`            | Multiplies `x` by `y` and assigns the result to `x`.   |
| `/=`         | Divide and assign         | `x /= y`            | Divides `x` by `y` and assigns the result to `x`.      |
| `%=`         | Modulo and assign         | `x %= y`            | Finds the remainder of `x` divided by `y`, assigns it to `x`. |
| `&=`         | Bitwise AND and assign    | `x &= y`            | Performs bitwise AND and assigns the result to `x`.    |
| `\|=`         | Bitwise OR and assign     | `x \|= y`            | Performs bitwise OR and assigns the result to `x`.     |
| `^=`         | Bitwise XOR and assign    | `x ^= y`            | Performs bitwise XOR and assigns the result to `x`.    |
| `<<=`        | Left shift and assign     | `x <<= 2`           | Shifts `x` left by 2 bits and assigns the result to `x`. |
| `>>=`        | Right shift and assign    | `x >>= 2`           | Shifts `x` right by 2 bits and assigns the result to `x`. |

6. **Miscellaneous Operators**

| **Operator** | **Description**         | **Example**         | **Explanation**                                        |
|--------------|-------------------------|---------------------|--------------------------------------------------------|
| `++`         | Increment               | `x++`               | Increases `x` by 1 (post-increment).                   |
| `--`         | Decrement               | `x--`               | Decreases `x` by 1 (post-decrement).                   |


---
### Switch Case:

The `switch` statement in Go is used to evaluate an expression against multiple cases, making it cleaner than using multiple `if-else` statements.
- ⚠️ No need for `break` (cases automatically break).
- Can evaluate expressions, not just constants.
- Optional `default` case handles unmatched cases.

In this example, the value of `day` is matched with the cases, and **Wednesday** is printed.
```go
package main

import "fmt"

func main() {
    day := 3
    switch day {
    case 1:
        fmt.Println("Monday")
    case 2:
        fmt.Println("Tuesday")
    case 3:
        fmt.Println("Wednesday")
    default:
        fmt.Println("Invalid day")
    }
}
```


---
### Memory Stack and Heap
In Go, memory management is straightforward thanks to its garbage collector, but understanding how stack and heap memory work can still be useful:

**Memory Stack**

- **Purpose**: The stack is used for managing function calls, local variables, and control flow.
- **Structure**: Operates in a **Last-In-First-Out (LIFO)** manner. Each function call pushes a new **frame** onto the stack, and **the frame is removed when the function returns**.
- **Usage**:
  - **Function Frames**: Stores local variables and function parameters.
  - **Automatic Allocation**: Local variables and parameters are allocated on the stack.
- **Characteristics**:
  - **Size**: Typically limited in size (stack size is managed by the Go runtime).
  - **Allocation/Deallocation**: Fast and automatic. The Go runtime manages this.
  - **Lifetime**: Memory is reclaimed automatically when the function returns.

**Heap**

- **Purpose**: The heap is used for **dynamic memory allocation**. Objects and data structures that need to persist beyond the scope of function calls are allocated here, e.g. **pointer**.
- **Structure**: The heap allows for more flexible allocation and deallocation. Memory is managed by the Go runtime's garbage collector.
- **Usage**:
  - **Dynamic Allocation**: Used for data that needs to live beyond the lifetime of a single function call or needs to be shared among different parts of a program.
  - **Garbage Collection**: Go automatically manages memory allocation and deallocation, so you don't need to manually free memory. The garbage collector reclaims memory that is no longer in use.
- **Characteristics**:
  - **Size**: Typically larger and more flexible compared to the stack.
  - **Allocation/Deallocation**: Handled by the garbage collector, which may introduce some overhead.
  - **Lifetime**: Managed automatically; memory is freed when there are no references to it.

Here are 2 examples illustrating stack and heap usage:

```go
package main

import "fmt"

func main() {
    a := 10      // Stack allocation
    b := &a      // Heap allocation (a reference to stack variable `a`)

    fmt.Println(a) // Stack
    fmt.Println(*b) // Dereferencing heap reference
}
```

![stack vs heap](./img/1-stack-heap.png)

***Summary***

- **Stack**: Used for local variables and function call management. Automatic and fast allocation/deallocation.
- **Heap**: Used for dynamic data and objects that persist beyond function scopes. Managed by Go’s garbage collector, which automates memory management but may add overhead.

Understanding how Go handles memory can help you write more efficient and effective code, especially when working with large data structures or managing performance.


---
### Functions

A function encapsulates code into its own container, which logically belong together.
Functions are declared using the `func` keyword and should be declared with a descriptive name.
A call of the function by its name is performed to execute the block of code that belong to it.
Every program has at least one function, i.e. the `main` function.
**Function calls are a way of reusing the same behavior**.

**Multiple returns are allowed in functions**, e.g.:
```go
func initConference() (uint8, uint8, uint8, []string) {
	const totalTickets uint8 = 50
	var remainingTickets uint8 = 30
	soldTickets := calculateSoldTickets(remainingTickets, totalTickets)
	bookings := make([]string, 0, totalTickets)                  // len=0, cap=50
	bookings = initBookings(totalTickets, soldTickets, bookings) // len=20, cap=50
	return totalTickets, remainingTickets, soldTickets, bookings
}
```

**Lower case for Private members or functions while Upper case for Public members or functions.** 

⚠️ **Package Level Variables** are defined at the top **outside all functions**., e.g.:
```go
package main

import (
	"fmt"
	"strings"
)

// package level variables:
var (
	remainingTickets uint8 = 30
	bookings               = make([]string, 0, totalTickets) // len=0 (no elements), cap=50
)

const (
	conferenceName       = "Go Conference"
	totalTickets   uint8 = 50
)

func main() {
    // more code here
}	
```

⚠️ **Best Practice is to define Variables as "local as possible"**, i.e. create the variable when you need it (the previous example is considered bad practice). **Local variables** are defined **inside a function or block**. They can be accessed only inside the function or block of code.


---
### Packages and Scopes

Modularization can be achieved using Packages, i.e. GO programs are organized into packages.
A **Package is a collection of GO files**, a.k.a. a single module. 

- ⚠️ **Scope: Variables and Functions**, defined outside any function, **can be accessed in all other files within the same package**. 
- **Sharing across packages**: Using Upper Camel case for Public members or functions while Lower Camel Case for Private members or functions, i.e. **Upper Camel Case** in case of **exposing a member or function outside the package**, i.e. **exporting a function or variables** simply **Capitalizing the first letter**.
- ⚠️ Importing a package is done using `import` keyword, but the import path should be **a valid path** in case of reusing a package of the own program, i.e. `import "github.com/yourusername/yourgoproject"` or `import "github.com/yourusername/yourgoproject/yourmodule"`, e.g. `"import github.com/paguerre3/gocomp/common"`. ***Note doing only `yourmodule` (`commom`), will not work*** as GO will search for the package in GO internals, so instead you must place your own project/main module domain at the beginning as it is defined under `go.mod`, e.g. `module github.com/paguerre3/gocomp`.
```go
package main

import (
	// GO "internal" module:
	"fmt"
	// Own module/GO extrenal package:
	"github.com/paguerre3/gocomp/common"
)

// package level variables:
var (
	remainingTickets uint8 = 30
	bookings               = make([]string, 0, totalTickets) // len=0 (no elements), cap=50
)

const (
	conferenceName       = "Go Conference"
	totalTickets   uint8 = 50
)

func main() {
    soldTickets := initConference()
    common.DisplayBookings(bookings)
        
    // more code here
}	
```

**3 Levels of Scope:**

1. **Local member:** 
- Declaration **within function**, i.e. **Can be used** only within that function.
- Declaration **within block**, i.e. **Cannot be used** only within block, e.g. `for`, `if-else`, etc.
    ```go
    func GetBookingsByPeopleNames(bookings []string) []string {
        ns := []string{} // ns is a local variable declared within "the function"
        // this is valid as no resize is done in the slice (simply a slice iteration to display values instead or appending)
        // but if there is a resize/update then a new array is built and changes won't be reflected until because a
        // variable is in the Local Function Frame of the Memory Stack unless a "new return" is providid
        // and then updating the package level variable (otherwise remove passing the slice as a reference
        // so the local/shadowed variable problem is avoided):
        for _, booking := range bookings {
            // strings.Fields(booking):
            // Fields splits the string s around each instance of one or more consecutive white space characters,
            // as defined by unicode.IsSpace, returning a slice of substrings of s or an empty slice if s contains only white space.
            fields := strings.Fields(booking) // fields is a local variable declared within the "for" block
            if len(fields) > 0 && fields[0] != "SOLD" {
                ns = append(ns, fmt.Sprintf("%s %s", fields[1], fields[2]))
            }
        }
        return ns
    }
    ```

2. **Package member:**
- Declaration **outside all functions** in same module, i.e. ⚠️ **Can be used in all other files within the same package**.
    ```go
    // package level variables:
    var (
        remainingTickets uint8 = 30
        bookings               = make([]string, 0, totalTickets) // len=0 (no elements), cap=50
    )

    // package level function starts with lower case letter:
    func calculateSoldTickets(remainingTickets uint8, totalTickets uint8) uint8 {
        return totalTickets - remainingTickets
    }
    ```	 

3. **Global member:**
- Declaration **outside all functions & upper case 1st letter**, i.e. **Can be used everywhere across all packages** (and its files).
    ```go
    func GetBookingsByPeopleNames(bookings []string) []string { 
        // more code here 
    }
    ```

***Variable Scope is the region of a program where a defined variable can be accessed***

- Encapsulated code into functions.
- Divided code into multiple files.


---
### Maps

Maps are built-in data structures that **store key-value pairs**. A map is created using the `make` function or by using a map literal. The keys must be of a type that supports equality comparison, and the values can be of any type, e.g.:
```go
myMap := make(map[string]int)
myMap["age"] = 30
```
Or using a map literal:
```go
myMap := map[string]int{"age": 30}
```
In this example, `"age"` is the key, and `30` is the value. Maps allow quick lookups, inserts, and deletions by key.

Other example:
```go
var userData = make(map[string]string)
userData["firstName"] = userFirstName
userData["lastName"] = userLastName
userData["email"] = userEmail
userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10) // 10 is used for decimal formatting
for i := soldTickets; i < lastBookIndex; i++ {
    bookings = append(bookings, userData)
}
```


---
### Structs
A struct is a composite data type that groups together fields (variables) under one name. 
Each field has a name and a type, and **structs are commonly used to define custom types**, e.g.:
```go
type Person struct {
    Name string
    Age  int
}

p := Person{Name: "Alice", Age: 30}
```
Here, `Person` is a struct with `Name` and `Age` fields, and `p` is an instance of `Person`. Structs help organize related data.

- **Mixed data types** can be defined within a struct (as the opposite to maps that support only one type of values defined during declaration).

- *Go typically encourages passing structs by reference (pointer) when mutation or large data is involved, while small structs that do not need modification can be passed by value to avoid unintended side effects.*


---
### `goroutines` - Concurrency

`goroutine` are used for achieving concurrency and/or parallelism, allowing you to run functions independently of each other. **They are lightweight compared to OS threads, as Go's runtime handles their scheduling and management**. `goroutines` are created by prefixing a function call with the `go` keyword, **allowing it to run concurrently**, e.g.:
```go
func sayHello() {
    fmt.Println("Hello")
}

func main() {
    go sayHello()  // Starts a new goroutine
    fmt.Println("World")
}
```
In this example, both `sayHello()` and `fmt.Println("World")` run concurrently. *The order of execution between them is non-deterministic*.

**Channels: Communication between `Goroutines`**

**Channels** are used to **safely communicate data between `goroutine`**, ensuring data consistency **without the need for explicit locking** mechanisms like `mutex`.

A **channel** is a typed **conduit through which you can send and receive values between `goroutines`**.

1. **Creating a Channel:**

Channels are created using the `make` function:
```go
ch := make(chan int)  // Creates a channel for int values
```

2. **Sending and Receiving Values:**

To **send data** into a channel, use the `<-` operator:
```go
ch <- 42  // Sends the value 42 into the channel
```

To **receive data** from a channel:
```go
value := <-ch  // Receives/obtains a value from the channel
```

*Example: `Goroutines` with Channels*
```go
package main

import (
    "fmt"
)

func sum(a int, b int, ch chan int) {
    result := a + b
    // Send result to the channel
    // (this technique avoids adding `Delta` with a `WaitGroup` and also doing `defer Done` for signaling that the `goroutine` is complete)
    ch <- result
}

func main() {
    ch := make(chan int)  // Create a channel for int values
    // Closing a channel is optional as it doesn't free resources, instead they are garbage collected 
    // (closing channel should only done when it’s important to signal to receivers that no more data will be sent)
    defer close(ch)

    // Start two concurrent goroutines
    go sum(2, 3, ch)
    go sum(5, 7, ch)

    // Receive/obtain result from chan and print the results from both goroutines
    // (no need of using a `WaitGroup Wait` as <-c during consumption is a blocking wait)
    result1 := <-ch
    result2 := <-ch

    fmt.Println("Results:", result1, result2)
}
```
In this example:
- Two `goroutines` are launched to perform the `sum` function concurrently.
- A shared channel `ch` is used to send the results from the `goroutines` back to the main function.
- *The `main` function ***waits*** for the values from both `goroutines` using `<-ch`*.

**Buffered vs. Unbuffered Channels:**

- **Unbuffered channels**: The sender and receiver must be ready at the same time. The `goroutine` sending data **will block until the receiver is ready to receive the value**.
- **Buffered channels**: You can **define a buffer size**, allowing the channel **to store multiple values before blocking**.

*Example of a Buffered Channel*:
```go
ch := make(chan int, 2)  // Create a buffered channel with size 2
ch <- 1
ch <- 2  // Both sends are non-blocking as the buffer is not full
```

⚠️ **Buffered channels allow sending and receiving of values without immediate synchronization between `goroutines` if the buffer is not full**.

**Closing Channels:**

Channels can be **closed** when no more values will be sent:
```go
close(ch)
```
**Once a channel is closed**, no more values can be sent, but **remaining values can still be received**. *Receiving from a closed channel returns the zero value for the channel’s type*, e.g.:
```go
ch := make(chan int, 2)
ch <- 10
close(ch)

val, ok := <-ch  // ok will be true if value was received
fmt.Println(val, ok)

val, ok = <-ch  // ok will be false as the channel is closed
fmt.Println(val, ok)
```

**Select Statement:**

The `select` statement is **used to wait on multiple channel operations**. It **blocks until one of its cases can proceed**, making it **useful for handling multiple concurrent channels**, e.g.:
```go
func main() {
    ch1 := make(chan string)
    ch2 := make(chan string)

    go func() {
        ch1 <- "Message from ch1"
    }()

    go func() {
        ch2 <- "Message from ch2"
    }()

    select {
    case msg1 := <-ch1:
        fmt.Println(msg1)
    case msg2 := <-ch2:
        fmt.Println(msg2)
    }
}
```
Here, `select` listens to both channels, and ⚠️ **whichever one receives a message First will trigger its corresponding case**. This is **useful for non-blocking concurrent operations**.

***Summary***

- **`Goroutines`** allow functions to run concurrently, managed efficiently by Go's runtime.
- **Channels** provide a way for `goroutines` to communicate and synchronize.
- **Unbuffered channels** require **simultaneous sending and receiving**, while **buffered channels** can **store values**.
- The `select` statement is used **for multiplexing channels**, allowing you to wait for multiple channel operations concurrently.
- **Closing a channel is optional** and should be **done when it’s important to signal to receivers that no more data will be sent**. Closing should be avoided when multiple senders are involved, and it's unnecessary if the receiver doesn't care about the end of transmission. **Closing a channel does not free resources**, i.e. channels are garbage collected when they are no longer referenced, so closing them is not necessary for resource cleanup. **Only the sender should close the channel**, i.e. it’s the responsibility of the goroutine that sends values to close the channel and receivers should not attempt to close the channel, as this would cause a runtime panic.

**Concurrency and Parallelism** in terms of CPU cores:

| **Aspect**              | **Concurrency**                                             | **Parallelism**                                          |
|-------------------------|-------------------------------------------------------------|----------------------------------------------------------|
| **Definition**           | Managing multiple tasks by interleaving their execution.    | Running multiple tasks at the same exact time.            |
| **CPU Cores**           | Can be done on a single core by rapidly switching between tasks (context switching). | Requires multiple cores to run tasks simultaneously.      |
| **Task Execution**       | Tasks progress independently but not necessarily simultaneously. | Tasks execute simultaneously on different cores.          |
| **Example**             | `Goroutines` sharing a single core, switching between them.    | `Goroutines` running on multiple cores, executing in parallel. |
| **System Requirement**   | Doesn't require multiple cores; can be achieved with one core. | Requires a multi-core CPU to achieve true parallelism.     |
| **Analogy**             | Juggling multiple tasks but only working on one at a time.   | Multiple workers performing tasks simultaneously.         |
| **Go Example**          | Multiple `goroutines` on a single core, interleaving execution. | `Goroutines` running on different cores, using `GOMAXPROCS` to utilize multiple cores. |

***The main advantage in Go programming is that Concurrency is cheap and Easy*** 

**Synchronization:**

The **main thread doesn't wait for the other threads to finish** their execution and for that **synchronization is needed**.

`WaitGroup` waits for the launched `goroutines` to finish their execution. Package `sync` provides basic synchronization functions.
- **Add(n):** Sets the number of `goroutines` to wait for (increases the counter by the provided value).
- **Wait:** Blocks until the `WaitGroup` counter is "0".
- **Done:** Decrements the `WaitGroup` counter by "1", this is called by the `goroutine` to indicate that it's finished.
```go
var wg sync.WaitGroup

wg.Add(2)  // Expecting two goroutines

go func() {
    // Mark goroutine as done, it awlays executes a the end of the function.
    // Same as writting it at the end but if there is a en error it can be bypassed so a good practice is to defer the call:
    defer wg.Done()  
    // Do some work
}()

go func() {
    defer wg.Done()
    // Do some work
}()

wg.Wait()  // Wait for both goroutines to finish
```
In this example, the program waits for both `goroutines` to complete before proceeding.

***Note Defer is used to ensure that a function call is performed later in a program’s execution***, usually for purposes of cleanup. defer is often used where e.g. ensure and finally would be used in other languages.

**Go is using, what's called "Green thread"**, i.e. an abstraction of an actual thread.

| **Aspect**                | **`Goroutines`**                                           | **OS Threads**                                           |
|---------------------------|----------------------------------------------------------|----------------------------------------------------------|
| **Creation Overhead**     | Lightweight, minimal overhead.                          | Heavier, more resource-intensive.                        |
| **Scheduling**            | Managed by Go runtime, efficient with many `goroutines`.  | Managed by the OS, less efficient for large numbers.     |
| **Context Switching**     | Fast and low-cost due to smaller stack size.            | Slower and more expensive due to larger stack size.      |
| **Memory Usage**          | Low memory footprint, small stack size (grows dynamically). | Higher memory usage per thread due to fixed stack size. |
| **Ease of Use**           | Simple to create and manage with the `go` keyword.       | More complex to manage and requires explicit thread management. |
| **Concurrency Model**     | Designed for high concurrency with low cost.            | Suited for parallelism but can be more expensive for concurrency. |
| **Synchronization**       | Channels provide built-in synchronization and communication. | Requires explicit use of synchronization primitives (mutexes, etc.). |
| **Scaling**               | Scales well with a large number of `goroutines`.          | Limited by the number of available system threads and resources. |


---
### Interfaces and Structs

In Go, the relationship between **structs** and **interfaces** is foundational to the language’s approach to type systems and polymorphism.

1. **Structs**:

    A **struct** is a collection of fields that can hold data. Structs are **used to define custom data types** that group related data together, e.g.:
    ```go
    type Car struct {
        brand string
        year  int
    }
    ```
    In the above example, `Car` is a struct with two fields, `brand` and `year`.

2. **Interfaces**:

    An **interface** is a type that **specifies a set of method signatures**. **"Any type" that implements these methods is said to "satisfy" the interface**. Unlike in other languages, ⚠️ **Go does not require you to explicitly declare that a type implements an interface**; if a type has all the methods the interface requires, it automatically satisfies the interface, e.g.:
    ```go
    type Vehicle interface {
        Drive() string
    }
    ```
    Here, `Vehicle` is an interface that requires a `Drive()` method returning a string.

3. **Implementing Interfaces with Structs**:

    A **struct can implement an interface by defining methods that match the interface’s method signatures**. In Go, there's no need to use a keyword like `implements`; ⚠️ **it is done implicitly by providing the correct method signatures**, e.g.:
    ```go
    func (c Car) Drive() string {
        return "Driving a " + c.brand
    }
    ```
    In this case, the `Car` struct implements the `Vehicle` interface by providing a `Drive` method. This allows `Car` to be treated as a `Vehicle`.

4. **Custom Instantiation**:

    You can create factory-like functions that return interfaces. **This allows the function to return different types that satisfy the interface, without the caller needing to know the specific type**.
    ```go
    // Custom instantiation function that returns a Vehicle interface
    func NewCar(brand string) Vehicle {
        return Car{brand: brand}
    }
    ```
    In this example, `NewCar()` creates an instance of the `Car` struct and returns it as a `Vehicle`. Since `Car` implements the `Vehicle` interface, the returned type can be used in contexts where a `Vehicle` is expected, providing polymorphism.

5. **Polymorphism**:

    **Polymorphism refers to the ability to treat different types as the same interface type**, allowing you to write more general and flexible code.
    ```go
    func printVehicleInfo(v Vehicle) {
        fmt.Println(v.Drive())
    }

    func main() {
        myCar := NewCar("Tesla")
        printVehicleInfo(myCar)  // Output: Driving a Tesla
    }
    ```
    Here, `printVehicleInfo` accepts any type that implements the `Vehicle` interface, so you can pass different struct types that implement `Vehicle` without changing the function.

⚠️ **A more complete and efficient example using Interface passed as value with Struct instantiation passed as reference** (used for larger data structs not exposed, i.e. not returned by functions):
```go
package domain

type BookingUser interface {
    FirstName() string
    LastName() string
    Email() string
    NumberOfTickets() uint8
    IsValidInput(remainingTickets uint8, conferenceName string) bool
}

type user struct {
    firstName       string
    lastName        string
    email           string
    numberOfTickets uint8
}

func NewBookingUser(firstName string, lastName string, email string, numberOfTickets uint8) BookingUser {
    return &user{
        firstName:       firstName,
        lastName:        lastName,
        email:           email,
        numberOfTickets: numberOfTickets,
    }
}

func (u *user) FirstName() string {
    return u.firstName
}
func (u *user) LastName() string {
    return u.lastName
}
func (u *user) Email() string {
    return u.email
}
func (u *user) NumberOfTickets() uint8 {
    return u.numberOfTickets
}

func (u *user) IsValidInput(remainingTickets uint8, conferenceName string) (vi bool) {
    isValidUserName := len(u.firstName) >= 2 && len(u.lastName) >= 2
    vi = true
    if !isValidUserName {
        fmt.Println("Please enter a valid name")
        vi = false
    }

    isValidEmail := len(u.email) > 3 && strings.Contains(u.email, "@")
    if !isValidEmail {
        fmt.Println("Please enter a valid email address")
        vi = false
    }

    isValidTicket := u.numberOfTickets > 0
    if !isValidTicket {
        fmt.Println("Please enter a valid number of tickets")
        vi = false
    }

    // validate avaibaility:
    if u.numberOfTickets > remainingTickets {
        fmt.Println("Sorry, we only have", remainingTickets, "tickets left for", conferenceName)
        vi = false
    }
    return vi
}
```
```go
package main

import (
    // GO "internal" module:
    "fmt"
    "sync"
    "time"

    // Own module/GO external package:
    "github.com/paguerre3/gocomp/common"
    "github.com/paguerre3/gocomp/domain"
)

// package level variables:
var (
    remainingTickets uint8 = 30
    bookings               = make([]domain.BookingUser, 0, totalTickets) // len=0 (initial size), cap=50
    wg               sync.WaitGroup
)

const (
    conferenceName       = "Go Conference"
    totalTickets   uint8 = 50
)

func main() {
    soldTickets := initConference()
    common.DisplayBookings(bookings)

    greetUsers(soldTickets)

    // infinite loop for ==> for "true", i.e. for { }
    // for remainingTickets > 0 { // loop with true false condition.
    for {
        // multiple returns are allowed in go
        //userFirstName, userLastName, userEmail, userTickets := getUserInputs()
        bookingUser := domain.NewBookingUser(getUserInputs())
        if vi := bookingUser.IsValidInput(remainingTickets, conferenceName); !vi {
            // continue to next iteration to try again:
            continue
        }
        // more code here
    }
}
```	

***Key Points:***

- **Structs** define concrete data types.
- **Interfaces define a set of behaviors (methods) without specifying the underlying data type**.
- Go interfaces are **implicit**: if a type implements the required methods, it automatically satisfies the interface.
- Using **custom instantiation** functions can abstract the creation of structs while returning interface types, allowing flexibility in the underlying implementation.

This allows Go to achieve polymorphism while maintaining simplicity and flexibility in its type system.


---
### Data types and recommendations on when to pass them by value or reference

| **Data Type**     | **Pass by Value**                               | **Pass by Reference (Pointer)**                          |
|-------------------|-------------------------------------------------|----------------------------------------------------------|
| **Basic types**   | (int, float, bool, string)                      | N/A (basic types are small and cheap to copy)             |
|                   | **When to use:** Always pass by value.          |                                                          |
| **Structs**       | **When the struct is small (few fields)**       | **When the struct is large or complex (many fields/nested structs)** |
|                   | Protects the original data from modification.   | More efficient for large structs, avoids copying overhead.|
| **Arrays**        | **When the array is small**                     | **For large arrays or if modifications are needed**       |
|                   | Arrays are copied by default when passed by value. | Pass by reference to avoid copying large arrays.         |
| **Slices**        | **Rarely (internally passed by reference)**     | **Modify elements or append to slice**                   |
|                   | Safe to pass by value, but generally passed as a reference due to underlying array. | Changes reflect in the original slice.                   |
| **Maps**          | **Rarely (maps are reference types)**           | **Always passed as a reference**                         |
|                   | Maps are internally passed by reference, even when passed by value. | Modifications reflect in the original map.               |
| **Channels**      | **Always passed by value**                      | **N/A (channels are reference types)**                   |
|                   | Channels are lightweight, passing by value does not copy the underlying data. |                                                          |
| **Pointers**      | **N/A**                                         | **Always passed as a reference**                         |
|                   | Pointers themselves are small values and can be passed directly. | Pass pointer if you want to modify the underlying value.  |
| **Interfaces**    | **Always passed by value (points to data)**     | **N/A (interfaces are reference types)**                 |
|                   | Interfaces store type information and a pointer to the underlying data. |                                                          |

***Key Considerations:***

- **Small types** like integers, booleans, and small structs can be passed by value, as copying them is cheap and protects the original data from being modified.
- **Large types** like large structs, arrays, and slices should generally be passed by reference to avoid the performance overhead of copying.
- **Reference types** like **maps, slices, channels, and interfaces are passed by value (descriptor is passed by value)**, but they internally reference the same underlying data, meaning modifications affect the original.


---
### Error handling

Go uses a unique approach to error handling compared to traditional exception-based languages. Instead of exceptions, **Go uses explicit return values to handle errors**, providing a clear and explicit mechanism for managing errors at each step.

1. **Error Type:**

    Go has a built-in **`error`** type, which is an interface that can be implemented to describe an error. It typically returns `nil` if there’s no error, or an instance of an error if something goes wrong, i.e.:
    ```go
    type error interface {
        Error() string
    }
    ```

2. **Returning Errors:**

    **Functions often return a value and an error**. The caller is responsible for checking if the error is `nil` to determine if the operation succeeded, e.g.:
    ```go
    func divide(a, b int) (int, error) {
        if b == 0 {
            return 0, fmt.Errorf("cannot divide by zero")
        }
        return a / b, nil
    }

    result, err := divide(10, 0)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Result:", result)
    }
    ```

3. **Error Wrapping (Go 1.13+):**

    Go 1.13 introduced **error wrapping** with the `%w` verb in the `fmt.Errorf` function, which allows errors to be wrapped and unwrapped for better context and debugging.
    ```go
    err := fmt.Errorf("an error occurred: %w", originalErr)
    ```
    The `errors.Unwrap` function can be used to retrieve the original error, and `errors.Is` checks if a specific error is in the error chain.

4. **Custom Errors:**

    You can create custom error types by implementing the `Error()` method for more specific error handling, e.g.:
    ```go
    type CustomError struct {
        Msg string
    }

    func (e *CustomError) Error() string {
        return e.Msg
    }

    func doSomething() error {
        return &CustomError{"something went wrong"}
    }

    err := doSomething()
    if err != nil {
        fmt.Println(err)
    }
    ```

5. **Panic and Recover:**

    Go provides the **`panic`** and **`recover`** functions **for handling unrecoverable errors or exceptional situations**. However, they are **not used for general error handling but rather for catastrophic failures or conditions that should halt the program**.

- **Panic**: Immediately stops the program and starts the process of unwinding the stack.
- **Recover**: Used within a deferred function to regain control of a panicking `goroutine`.
    ```go
    func riskyOperation() {
        defer func() {
            if r := recover(); r != nil {
                fmt.Println("Recovered from panic:", r)
            }
        }()
        panic("Something went terribly wrong!")
    }

    riskyOperation()
    ```

6. **Best Practices:**

- Always check for errors returned by functions.
- Return meaningful error messages for better debugging.
- Use `panic` and `recover` sparingly, typically for program crashes or system-level errors.

Go's explicit error handling promotes clear and concise error management while encouraging developers to think about error handling at every function call.


---
### Functional Programming

Go is primarily an imperative language but supports **functional programming** features, enabling a hybrid approach. While Go doesn’t have some of the native functional programming constructs found in languages like Haskell or Lisp, it does allow the use of first-class functions, higher-order functions, closures, and immutability to facilitate functional patterns.

**Key Functional Programming Concepts in Go:**

1. **First-class Functions**:

   In Go, functions are first-class citizens, meaning **they can be assigned to variables, passed as arguments, and returned from other functions**. 
   E.g:
   ```go
   func add(a, b int) int {
       return a + b
   }

   func applyOperation(a, b int, op func(int, int) int) int {
       return op(a, b)
   }

   result := applyOperation(2, 3, add)  // Passes `add` as a function argument
   fmt.Println(result)  // Output: 5
   ```

2. **Higher-order Functions**:

   **A function that takes another function as an argument or returns a function is called a higher-order function**. Go supports this concept, which is fundamental to functional programming. E.g.:
   ```go
   func multiplier(factor int) func(int) int {
       return func(x int) int {
           return x * factor
       }
   }

   double := multiplier(2)
   fmt.Println(double(5))  // Output: 10
   ```

3. **Closures**:

   **A closure is a function that references variables from its surrounding scope**. In Go, closures are easily implemented and commonly used.

   **Example**:
   ```go
   func incrementer() func() int {
       count := 0
       return func() int {
           count++
           return count
       }
   }

   inc := incrementer()
   fmt.Println(inc())  // Output: 1
   fmt.Println(inc())  // Output: 2
   ```

4. **Immutability**:

   Functional programming emphasizes immutability, though Go does not enforce immutability like some functional languages. Developers can achieve **immutability by avoiding modification of variables after they are created, passing values rather than references, and using constants where possible**. E.g.:
   ```go
   func sum(values []int) int {
       total := 0
       for _, v := range values {
           total += v
       }
       return total
   }

   result := sum([]int{1, 2, 3})  // Immutable usage of the input slice
   ```

5. **Anonymous Functions**:

   Go supports anonymous functions, which **can be used inline without needing to declare them separately**. These are useful in many functional programming scenarios such as in callbacks or small utility functions. E.g.:
   ```go
   result := func(a, b int) int {
       return a + b
   }(3, 4)  // Directly invoked anonymous function
   fmt.Println(result)  // Output: 7W
   ```

6. **Map, Filter, and Reduce**:
   While Go does not provide native implementations of common functional programming patterns like `map`, `filter`, and `reduce`, these can be implemented using higher-order functions and slices. Example of `map`:
   ```go
   func mapFunc(arr []int, f func(int) int) []int {
       result := make([]int, len(arr))
       for i, v := range arr {
           result[i] = f(v)
       }
       return result
   }

   doubled := mapFunc([]int{1, 2, 3}, func(x int) int { return x * 2 })
   fmt.Println(doubled)  // Output: [2 4 6]
   ```

**Limitations in Go for Functional Programming:**

- **No Tail-Call Optimization**: Unlike some functional languages, Go does not optimize tail-recursive functions.
- **No Native Support for Immutable Data Structures**: Go does not natively enforce immutability.
- **Error Handling**: Go relies on explicit error handling rather than functional constructs like monads, often seen in purely functional languages.

***Conclusion:***

While Go is not a purely functional language, it provides enough tools to enable functional programming techniques. Developers can use closures, first-class functions, and higher-order functions to implement functional patterns, making Go a flexible language that supports both imperative and functional styles.


---
### Generics
Generics in Go were introduced in Go 1.18, allowing developers to **write functions, types, and methods that can work with any type**.

**Key Concepts:**

1. **Type Parameters**: Generics are based on type parameters, which allow you to define a function, method, or struct that works with any type. The type parameter is specified in square brackets (`[]`). E.g.:
   ```go
   func Print[T any](x T) {
       fmt.Println(x)
   }
   ```
   Here, `T` is the type parameter, and `any` means it can be any type.

2. **Type Constraints**: You can specify constraints for type parameters to restrict the types that can be used. Go uses *interfaces* as type constraints. E.g:
   ```go
   func Add[T int | float64](a, b T) T {
       return a + b
   }
   ```
   In this case, `T` can only be an `int` or `float64`.

3. **Generic Types**: You can also define generic structs or interfaces. E.g.:
   ```go
   type Pair[T any] struct {
       First  T
       Second T
   }
   ```

4. **Type Inference**: Go can often infer the type parameter based on the arguments passed to a function, making generics more convenient to use. E.g.:
   ```go
   Print(10) // Go infers that T is int
   ```

**Benefits:**

- **Code Reusability**: You can write a single implementation for multiple types.
- **Type Safety**: Go's type system ensures that the correct types are used, reducing runtime errors.

**Limitations:**

- Generics in Go are designed to be simple, so they are not as flexible or powerful as in languages like Java or C++.
- No **higher-kinded types** (types that take type parameters themselves).

Generics in Go offer more flexibility while keeping the language's simplicity and performance.


---
### Variadic Functions

A variadic function **can take a variable number of arguments of the same type**. It uses `...` before the type of the last parameter to denote that it accepts multiple arguments. E.g.:
```go
func sum(nums ...int) int {
    total := 0
    for _, num := range nums {
        total += num
    }
    return total
}
```
You can call it with any number of integers: `sum(1, 2, 3)`.
- **Arguments** are treated as a slice inside the function.
- It can still have regular parameters before the variadic parameter.


---
### Runes

A **rune** is an alias for the `int32` type and **represents a Unicode code point**. Runes are used to handle characters beyond basic ASCII, **allowing Go to work with multilingual text and special symbols**.

**Key Points:**

- A rune represents a single Unicode character.
- Go strings are UTF-8 encoded, so indexing a string gives you a byte, not a rune.
- To work with individual characters, you convert a string into a slice of runes.

Example that demonstrates how to use **`utf8` package to work with UTF-8 encoded strings, particularly for counting runes and decoding individual characters**:
```go
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str := "Hello, 世界"

	// Get the number of runes (Unicode code points) in the string
	runeCount := utf8.RuneCountInString(str)
	fmt.Printf("The string '%s' has %d runes.\n", str, runeCount)

	// Decode each rune in the string and print its value
	fmt.Println("Runes in the string:")
	for i, w := 0, 0; i < len(str); i += w {
		r, width := utf8.DecodeRuneInString(str[i:])
		fmt.Printf("Rune: %c, starts at byte position %d, width: %d bytes\n", r, i, width)
		w = width
	}
}
```

***Explanation:***

- `utf8.RuneCountInString(str)` counts the number of runes in the UTF-8 encoded string.
- `utf8.DecodeRuneInString` reads one rune (character) at a time and returns the rune and its width in bytes. This is useful for iterating over multi-byte characters like "世界".

Output:
```
The string 'Hello, 世界' has 9 runes.
Runes in the string:
Rune: H, starts at byte position 0, width: 1 bytes
Rune: e, starts at byte position 1, width: 1 bytes
Rune: l, starts at byte position 2, width: 1 bytes
Rune: l, starts at byte position 3, width: 1 bytes
Rune: o, starts at byte position 4, width: 1 bytes
Rune: ,, starts at byte position 5, width: 1 bytes
Rune:  , starts at byte position 6, width: 1 bytes
Rune: 世, starts at byte position 7, width: 3 bytes
Rune: 界, starts at byte position 10, width: 3 bytes
```

Runes allow for easier manipulation of characters in a string.


---
### Testing

Testing is a built-in feature using the `testing` package. **Unit tests are written in files ending with `_test.go`** and **use functions starting with `Test`**. You can also **use assertion libraries like `testify` to simplify test validation**.

**Key Points:**

- **Test Function Naming**: Test functions must start with `Test`, followed by the function name.
- **Test Files**: Place test functions in files ending with `_test.go`.
- **Running Tests**: Use `go test` in the terminal to run all tests.
- **Assertions**: Go’s `testify` library provides cleaner assertions like `assert.Equal`.

*Example using `testify`*:

1. Install `testify`:
   ```bash
   go get github.com/stretchr/testify/assert
   ```

2. Write a test:
   ```go
   package main

   import (
       "testing"
       "github.com/stretchr/testify/assert"
   )

   func Add(a, b int) int {
       return a + b
   }

   func TestAdd(t *testing.T) {
       result := Add(2, 3)
       expected := 5
       assert.Equal(t, expected, result, "They should be equal")
   }
   ```

Run tests with:
```bash
go test
```

⚠️ *Note that a common practice in GO is adding unit tests (`_test.go`) in the same directory/package where the code is located. 
This allows you to test the package code independently of each other when needed or test everything at once without setting a common folder.* 

**Test method naming and best practices ensuring readability, clarity, and organization:**

1. **Prefix with `Test`**:
   - All test functions should start with `Test` followed by a descriptive name for the functionality being tested.
   - This is required by Go's testing framework to recognize and execute the function as a test. E.g.:
   ```go
   func TestCalculateTotal(t *testing.T) { ... }
   ```

2. **Descriptive Names**:
   - Use clear, descriptive names that explain what the test is checking. The name should reflect the function and the specific case being tested.
   - Consider the pattern `TestFunctionName_StateUnderTest_ExpectedBehavior`. E.g.:
   ```go
   func TestAdd_PositiveNumbers_ReturnsSum(t *testing.T) { ... }
   ```

3. **Separate Test Cases by Functionality**:
   - Split tests by functionality or behavior to avoid clutter and improve clarity. For example, if you're testing different edge cases, it's better to have multiple test functions instead of one large one. E.g.:
   ```go
   func TestAdd_Zero_ReturnsSameNumber(t *testing.T) { ... }
   func TestAdd_NegativeNumbers_ReturnsSum(t *testing.T) { ... }
   ```

4. **Table-Driven Tests**:
   - For **functions that need to be tested with "many" different inputs, use table-driven tests**. These are tests that loop over a set of test cases defined in a table (slice of structs), making the code more concise and maintainable. E.g.:
   ```go
   func TestAdd(t *testing.T) {
       tests := []struct {
           name     string
           a, b     int
           expected int
       }{
           {"Add two positives", 2, 3, 5},
           {"Add zero and number", 0, 4, 4},
           {"Add two negatives", -2, -3, -5},
       }

       for _, tt := range tests {
           t.Run(tt.name, func(t *testing.T) {
               result := Add(tt.a, tt.b)
               if result != tt.expected {
                   t.Errorf("expected %d, got %d", tt.expected, result)
               }
           })
       }
   }
   ```

5. **Use `t.Run` for Subtests**:
   - **When testing multiple scenarios for the same function, use `t.Run` to organize subtests**. This makes it easier to see which test cases passed or failed and helps in debugging. E.g.:
   ```go
   func TestMultiply(t *testing.T) {
       t.Run("Positive numbers", func(t *testing.T) {
           result := Multiply(2, 3)
           if result != 6 {
               t.Errorf("expected 6, got %d", result)
           }
       })

       t.Run("Negative numbers", func(t *testing.T) {
           result := Multiply(-2, 3)
           if result != -6 {
               t.Errorf("expected -6, got %d", result)
           }
       })
   }
   ```

By following these conventions, your tests will be clearer, easier to maintain, and more reliable.


---
### Templates

**Templates** are a powerful feature, **allow generating dynamic content (such as HTML or text files) by combining static text with dynamic data**. Go templates come in two main forms: **text templates** (used for generating plain text) and **HTML templates** (specifically designed for safe HTML generation). Both types are part of Go's `text/template` and `html/template` packages.

**Key Concepts:**

1. **Template Definition**:
   Templates are defined using `{{` and `}}` as delimiters. Inside these delimiters, you can add variables, expressions, and functions to manipulate the template data.

2. **Variables**:
   - Variables can be passed to templates as a `map` or a `struct`. You access them using dot notation. For example, if your data contains `Person.Name`, you can render it with `{{ .Name }}`.
   - You can declare variables within the template using `{{ $var := ... }}`.

3. **Actions**:
   Actions are the commands that modify or access data within templates:
   - `{{ . }}`: Refers to the root object passed to the template.
   - `{{ .Field }}`: Accesses a field or method of a struct.
   - `{{ if }}`, `{{ else }}`: Conditional logic.
   - `{{ range }}`: Loops over arrays, slices, or maps.
   - `{{ with }}`: Modifies the dot (.) to a new scope for a block.
   - `{{ block }}`: Defines a block that can be overridden by other templates.

4. **Functions**:
   Go templates have many built-in functions such as `len`, `index`, `eq`, `and`, `or`, etc. You can also register custom functions in your Go code and use them in templates.

5. **Pipelines**:
   Go templates support pipelines (`|`) to chain the output of one function into the input of another. Example: `{{ .Name | printf "%q" }}`.

6. **Template Inheritance (HTML Templates)**:
   HTML templates allow for inheritance by using `{{ template "name" . }}`. This enables you to define base templates and override specific sections in child templates.

7. **Security (HTML Templates)**:
   The `html/template` package automatically escapes content to prevent Cross-Site Scripting (XSS) attacks. You can mark content as safe using the `html` or `js` functions if necessary.

*Here’s a basic example of a Go HTML template:*
```go
package main

import (
	"html/template"
	"os"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	tmpl := template.Must(template.New("example").Parse(`
		<h1>Hello, {{.Name}}!</h1>
		<p>You are {{.Age}} years old.</p>
	`))

	data := Person{Name: "Alice", Age: 30}
	tmpl.Execute(os.Stdout, data)
}
```

This will produce:
```html
<h1>Hello, Alice!</h1>
<p>You are 30 years old.</p>
```

Go templates are **highly efficient for generating dynamic content in web applications, configuration files, or any text-based output**.


---
### Further Reading

**GO Language:**
- [GO by Example](https://gobyexample.com/)
- [Effective GO](https://go.dev/doc/effective_go)
- [From Beginner to Senior](https://www.bytesizego.com/blog/learning-golang-2024)

**Domain Driven Design:**
- [DDD example #1 in GO (complete)](https://github.com/paguerre3/go-ddd)
- [DDD example #2 in GO (simplified)](https://github.com/paguerre3/go-ddd-api)

**Library and Docker:**
- [Library client example](https://github.com/paguerre3/faccounts)

**Blockchain and Cloud Patterns/Interviews:**
- [Building a Blockchain in GO](https://www.youtube.com/watch?v=jzmIxoiFBW0&list=PL4pLiB9n8GE3UTECkQaXpGUd0kTaXVhAf&index=1)
- [Building a Blockchain in GO (simplified approach)](https://www.youtube.com/playlist?list=PLJbE2Yu2zumC5QE39TQHBLYJDB2gfFE5Q)
- [Building a Blockchain in GO (complete BC code plus Cloud design patterns summary/interviews)](https://github.com/paguerre3/gobc)

**Tools & Frameworks:**
- [GO Kit as a set of packages and best practices for building microservices](https://github.com/go-kit/kit)
- [Wire for dependency injection](https://github.com/google/wire)
- [Watermill for working efficiently with message streams](https://github.com/ThreeDotsLabs/watermill)
- [GO blueprint to spin up a Go project with the corresponding structure](https://github.com/Melkeydev/go-blueprint)

**Code Organization:**
- [GO Standards](https://github.com/golang-standards/project-layout)
