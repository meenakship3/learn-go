# learn go fast: full tutorial (part 1)

source - https://www.youtube.com/watch?v=8uiZC0l4Ajw

## contents

- about
- getting started
- data types
- arrays
- map
- loops
- :=
- strings

## about go

- statically and strongly typed
- compiled (creates machine code/.exe file) - makes it much much (120x) faster than python (interpreted)
- built in concurrency with gorunner or something like that

## getting started

- download from go.dev if not already
- create a go.mod file and set module path (typically a repository). module paths are used to -
  - Import paths - how other code will import packages from your module
  - Module identity - uniquely identifying your module in the Go ecosystem

```bash
go mod init <repo link>
```

- entry point for the application - main.go within a cmd directory

```go
// identify package right at the beginning
package main
import "fmt" // can print w/o this, just recommended

// you HAVE to create a function with the package name (if it's main)
func main() {
  fmt.Println("Hello World!")
}
```

<img src="/Users/meenu/Documents/Screenshots/Screenshot 2025-12-21 at 11.12.00_PM.png">

## data types

- int, int16 (upto ~37k), int32, int64
  - if you try to assign the val to a greater number (i.e. beyond the limit of int16), you'll get an overflow error
- float32, float64 (no plain float)
- uint (only positive values)
- can only do arithmetic with the same data types
- int division results in an int that's rounded down
- strings

  - can be defined with double quotes or back quotes (if you want them to be printed exactly as you've formatted them)
  - len(string) doesn't give you number of characters but number of bytes
  - ```go
    import "unicode/utf8"

    fmt.Println(utf8.RuneCountInString("hello"))
    ```

  - rune for a single character

- bool

```go
var myString string = "hello"
var myBool bool
// shorthand
myVar := "text"
```

## functions and if

```go
import "errors"

func intDivision(numerator int, denominator int) (int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("Cannot divide by zero")
		return 0, err
	}
	var result int = numerator / denominator
	return result, nil
}

// calling it
var result, err = intDivision(numerator, denominator)
```

- or - ||, and - &&
  - only evaluates what is necessary

```go
switch remainder{
  case 0:
    something
  case 1, 2:
    something
}
// no break statement required
```

## arrays

```go
func main() {
	var intArr [3]int32 // array of size 3
	intArr[1] = 123
	fmt.Println(intArr[0])
	fmt.Println(intArr[1:3])
}
```

- if no values are assigned, array is [0,0,0] by default
- get memory location of elements using & `fmt.Println(&intArr[1]`

```go
// other ways of creating an array
var intArr [3]int32 = [3]int32{1,2,3}
intArr := [...]int32{1,2,3}
```

```go
// appending to an array
var intSlice []int32 = []int32{1, 2, 3}
intSlice = append(intSlice, 7)
// append array to array
var intSlice2 []int32 = []int32{8,9}
intSlice = append(intSlice, intSlice2...)
```

- use cap(intSlice) to see how the capacity expands after append
- `var intSlice3 []int32 = make(int32[], 3, 8)` - to set length and capacity (3,8)

```go
for i, v := range intArr {
  // gets index and value
}
```

## map

- dict
- ```go
  var myMap map[string]uint8 = make(map[string]unint8)
  // or
  var myMap = map[string]unint8{"Adam":23, "Sarah":45}
  ```
- can access value using square brackets like python
- if the key doesn't exist, default value of the value (lols) type will be returned (eg.: 0 in the case of unit8)

```go
var age, ok = myMap2["Jason"] // ok is a bool, tells you if the key was found in map
if ok{
  fmt.Printf("The age is %v", age)
} else {
  fmt.Println("Invalid name")
}
```

- `delete(myMap2, "Adam")` to delete key-value from map

```go
for name, age := range myMap2 {
  // something
}
```

## loops

```go
// while
for i<10 {
  fmt.Println(i)
  i = i + 1
}

// for
for i:=0; i<10; i++{
  fmt.Println(i)
}
```

## :=

- declares and assigns variables in one step
  - type of variable is inferred
- can only be used within functions, not at package level
- can be used to declare multiple variables (of different types) at the same time. atleast one variable has to be new though

## strings

- can be indexed using [] but what we get an array of bytes (utf-8) not characters lol
- instead represent a string as an array of runes
  - gets character count

```go
var myString = []rune("resume")
```

- strings are immutable, can be concatenated using + but it creates a new string every time
  - instead import strings and use string builder

```go
var strSlice = []string{"s", "e", "n", "t", "i", "m", "e", "n", "t"}
var strBuilder strings.Builder
	for i := range strSlice {
		strBuilder.WriteString(strSlice[i]) // each char added to the builder but string created later
	}
	var catStr = strBuilder.String()
	fmt.Printf("\n%v", catStr)
```
