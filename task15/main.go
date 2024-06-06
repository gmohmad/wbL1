package main

/*
var justString string
func someFunc() {
  v := createHugeString(1 << 10)
  justString = v[:100]
}

func main() {
  someFunc()
}

So in this code we initialize a huge string and assign it to 'v' variable. And in Go, strings are just immutable
slices of bytes, and a slice is just a struct that has len, cap, and pointer to an element of the
underlying array fields. Next, we create a new string (slice of bytes) using 'v[:100]' and assing it to 'justString', which
is a global variable. So because the slice we have in 'justString' is referencing the first 100 values from the underlying
array, and 'justString' is a global variable (which means it will not be cleaned up by the GC), the GC will not be able to
reclaim the memory taken by the underlying array of the huge string we created.

To fix this, we can just move 'justString' into the 'someFunc' function, which will allow the GC to clean it up after
'someFunc' function returns:

func someFunc() {
  v := createHugeString(1 << 10)
  justString := v[:100]
}

func main() {
  someFunc()
}
*/
