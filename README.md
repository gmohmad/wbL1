### 1. What's the most efficient way to concatenate strings?
strings.Builder. Example:
```
builder := strings.Builder{}

builder.WriteString("some string...")
builder.WriteString("...some string ")

concatenatedString := builder.String()
```
### 2. What are interfaces and how are they used in Go?
Interfaces in Go are types that let us define sets of methods that object must impement, helping us achieve polymorphism. 
If an object impements all methods that we define in the interface, it automatically implements the interface, we 
don't have to explicitly say that an object defines the interface. This is possible because Go supports duck typing.
Example:
```
type I interface {
    SomeMethodThatObjectsMustImplement()
}
```

### 3. What's the difference between RWMutex and Mutex?
Mutex provides Lock and Unlock methods to lock access to a critical section so that only one goroutine at a time can
access it, doesn't matter if it's trying to write or read the data.
RWMutex, in addition to Lock and Unlock, provides RLock and RUnlock. Using these allow multiple goroutines to read 
the data, but only one to write. In short, Mutex lets us block for both read and write operations, while RWMutex also
lets us block either for read or write operations.

### 4. What's the difference between buffered and unbuffered channels?
Writing to an unbuffered channel blocks until another goroutine reads from it, and vice versa.
Buffered channels allow us to set a 'capacity' on the channel - so when we write to a channel, if the capacity isn't exceeded,
it'it will not block

### 5. What's the size of an empty struct (struct{}{})?
0 bytes.

### 6. Is there operator/method overloading in Go?
Nope.

### 7. In what order will map[int]int elements be displayed?
Elements in maps are unordered, so we can't know in what order they will be displayed. But when using some functions 
from the 'fmt' package, the elements might be sorted, like when using fmt.Println

### 8. What's the difference between new and make?
new is used to allocate memory for a new value of a specified type and return a pointer to it,
while make is used to create and initialize reference types like slices, maps and channels.

### 9. How many ways are there to define a slice or a map variable?
```
var slice []int
slice := new([]int)
slice := []int{}
slice := make([]int)

var m map[int]int
m := new(map[int]int)
m := map[int]int{}
m := make(map[int]int)
```

### 10. What will this program output and why?
```
func update(p *int) {
  b := 2
  p = &b
}

func main() {
  var (
     a = 1
     p = &a
  )
  fmt.Println(*p)
  update(p)
  fmt.Println(*p)
}
```
It will print
```
1
1
```
Thats because in 'update' function we're updating the value of the local 'p' variable, not the value of 'p' we created in
main function, because in Go, everything is passed by a copy, even pointers.

### 11. What will this program output and why?
```
func main() {
  wg := sync.WaitGroup{}
  for i := 0; i < 5; i++ {
     wg.Add(1)
     go func(wg sync.WaitGroup, i int) {
        fmt.Println(i)
        wg.Done()
     }(wg, i)
  }
  wg.Wait()
  fmt.Println("exit")
}
```
It will randomly print numbers from 0 to 4 and then deadlock because the WaitGroup is passed by a copy to the anonymous goroutine, 
so calling wg.Done() does not affect 'wg' we created in main function

### 12. What will this program output and why?
```
func main() {
  n := 0
  if true {
     n := 1
     n++
  }
  fmt.Println(n)
}
```
It will print 0, because variables created in a 'if' block are scoped within that 'if' block

### 13. What will this program output and why?
```
func someAction(v []int8, b int8) {
  v[0] = 100
  v = append(v, b)
}

func main() {
  var a = []int8{1, 2, 3, 4, 5}
  someAction(a, 6)
  fmt.Println(a)
}
```
It will print 
```
[100 2 3 4 5]
```
because the 'append' function allocates a new underlying array if the capacity of the slice
we're calling it on is exceeded, but before that we're operating on the same underlying array, that's why 1 became 100.

### 14. What will this program output and why?
```
func main() {
  slice := []string{"a", "a"}

  func(slice []string) {
     slice = append(slice, "a")
     slice[0] = "b"
     slice[1] = "b"
     fmt.Print(slice)
  }(slice)
  fmt.Print(slice)
}
```
It will print 
```
[b b a][a a]
```
because again, append allocates a new underlying array if we exceed the capacity, when we do 'append(slice, "a")',
we exceed the capacity of the 'slice' slice, which is 2, and because of that, the changes we make in the anonymous function
do not apply to the slice from the main function.
