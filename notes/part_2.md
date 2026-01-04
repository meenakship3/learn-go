# learn go fast: full tutorial (part 2)

## structs and interfaces

- structs - to create your own type
  - kind of like a class with just attributes (but go doesn't have classes)

```go
type gasEngine struct{
  mpg uint8
  gallons uint8
}

func (e gasEngine) milesLeft() uint8 {
  return e.gallons*e.mpg
}
// the (e gasEngine) part "links" the method to the struct

var myEngine gasEngine = gasEngine{mpg:25, gallons:15}
// access gasEngine.mpg, gasEngine.gallons
```

- interface - contains certain methods that each of its type should implement
  - if they implement it, then they become a member of the interface
  - reference: https://www.youtube.com/watch?v=lbW-KVdIXaY

```go
type shape interface {
  area() float64
  circumf() float64
}

func printShapeInfo (s shape) float64 {
  fmt.Printf("area of %T is: %0.2f\n", s, s.area())
  fmt.Printf("circumference of %T is: %0.2f\n", s, s.circumf())
} // notice how this makes the function reusable instead of having to create shapeinfo for each shape separately
```

## pointers

- reference: https://www.youtube.com/watch?v=4B2rwYvuiBo
- points to a memory location
  - access the pointer (i.e. the memory location) using &<variable name>
  - access the **value** that it points to by using \*<variable name> (this is called dereferencing)
  - you can store the pointer itself in a variable (like `m := &name`), m has its own memory address (separate from name obviously)
- confusing part

```go
func updateName(s *string) {
  *s = "SM"
}
```

- here *string means a pointer is accepted as input and *s means we're attempting to point to the value not the memory location
- in a function, what is updated is a copy of the parameter, not the parameter itself. if you want to update the parameter, use pointers.

## goroutines

- `go process(<variable>)` to use go routines
- concurrency vs parallelism
  - concurrency - work on multiple tasks together (not necessarily at the same time - but while there's I/O for one task, you can move onto the other task while waiting)
  - parallelism - execute 2 tasks at exactly the same time
  - parallelism occurs when you have multiple CPU cores
- mutex - sync.Mutex or sync.RWMutex
  - m.Lock() and m.Unlock() around the goroutine code
- "if a function is computationally expensive, the amount of improvement you get (as a result of parallelism) is a function of the cores in your CPU"

### channels

- carry information from goroutines, notify when data is changed

```go
var c = make(chan int) // notice the syntax
c <- 1 // creates an unbuffered channel, with only space for 1 element
var i = <- c // 1 "popped out" of the channel and stored in i
```

- running just this code will give us an error since after writing to an unbuffered channel we'll just wait there forever (can't move to read). instead, do this -

```go
func main(){
  var c = make(chan int)
  go process(c)
  fmt.Println(<-c)
}

func process(c chan int) {
  c <- 123
}
```

- buffer channel - store more than 1 value at once
- it's possible to create if statements with the value contained in a channel
  - eg.: if the channel contains a certain value, then do this

## generics

- for polymorphism-ish
- same function but for different types of inputs

```go
func sumSlice[T int | float32 | float64](slice []T){
  var sumT
  for _, v := range slice{
    sum += v
  }
  return sum
}
```
