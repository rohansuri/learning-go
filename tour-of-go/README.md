Important points extracted out from "A Tour of Go"

## Table of contents

* [Basics](#basics)
	* [Declaring types and identifiers](#declaring-types-and-identifiers)
	* [Basic types](#basic-types)
    * [Type conversion](#type-conversion)
    * [Constant declaration](#constant-declaration)
* [Flow control ](#flow-control)
* [More types](#more-types)
	* [Pointers](#pointers)
	* [Structs](#structs)
	* [Arrays](#arrays)
	* [Slices](#slices)
	* [Range](#range)
	* [Map](#map)
	* [Functions](#functions)
	* [Closures](#closures)
* [Methods and Interface](#methods-and-interfaces)
	* [Methods](#methods)
	* [Interfaces](#interfaces)
	* [Empty interface](#empty-interface)
	* [Type assertion](#type-assertion)
* [Questions](#questions)


## Basics

#### Declaring types and identifiers

* A name is exported if it begins with a capital letter.

* Type is specified after identifier name.  
	`func add(x int, y int)`

* Two or more consecutive identifiers have the same type then you are allowed to state the type only once.  
	`func add(x, y int)`

* Function can return multiple results

* Function's return values can be named (to document results in a meaningful way)  
	`func add(x, y int) (result int)`
	
	Thereby no need to name them explicitly in the return statement (but the return keyword is still needed)


* var declarations can also initialise variables then and there thereby not requiring explicit type declaration.  
	`var sum = 100`

* var declaration is required for variables declared outside functions.

* Inside functions, var declarations with initialisations can use a short hand operator `:=`
```go
func add(x, y int) (sum int) {
	sum := x + y
	return
}
```

#### Basic types

* `int`, `uint` are 32bits on 32bit systems and 64bit on 64bit systems.  
	For system independent integers we have `int32`, `int64`, etc.

* Amongst the usual basic types, `string` and `rune` are basic types too in Go.  
	`rune` is an alias for `int32`, representing a Unicode code point (code point is a numeric value from the Unicode space)

#### Type conversion

* Expressions `T(v)` are type conversion expressions which convert v to type T.  
	These are always explicitly needed.

	```go
	x, y := 3, 4
	div := float64(x)/float64(y) // not type converting either of x or y is an error
	fmt.Println(div)
	```

#### Constant declaration
* `const Pi = 3.14`

## Flow control 

* Syntax for if conditionals and for loops do not require parentheses but do require the body always be enclosed in `{}`

* Switch statements don't fall-through and the cases do not need to be constants i.e. they could be expressions.

* Deferred function calls are executed after the surrounding function returns. Multiple defers are stacked.

## More types

#### Pointers

* Go has pointers. Using `*` to dereference and `&` address-of.
	Pointers hold memory address of a value.
	
* Zero value is `nil`

* No pointer arithmetic is permitted.
	
#### Structs

* Structs hold many fields (POJOs essentially) 
```go
type Vertex struct {
	X int
	Y int
}
```

* Members could be accessed directly using dot operator


* Structs are constructed using struct literals  
	```go
	Vertex{1,2}   
	Vertex{X:1} // Y is 0 implicitly
	```

* Accessing struct fields using a pointer can also be done using the dot operator (rather than the cubersome `(*p).X`)
	
#### Arrays
* Array declaration type is of the form `[n]T`

* Arrays are of fixed size, zero indexed.

* `oddNos := [4]int{1, 3, 5, 7}` is an array

#### Slices
* Slice declaration type is of the form `[ ]T`

* Slices are dynamically-sized view of arrays (references to a portion of the array)

* Obtain a slice using `a[low: high]`, high isn't included.  
	You may omit low or high or both and defaults (0, length) will be used.

* `var firstTwo []int = oddNos[0:3]` // [1, 3]

* Modifying a slice modifies the underlying array and all other slices that cover the same modified range. 

* `naturalNos := []int{1, 2, 3}`  
  creates an array and returns a slice reference to it  

	`naturalNos = naturalNos[:2]` // [1, 2]  
	`naturalNos = naturalNos [1:]` // [2]  

* Slice has a length as well as a capacity.  
	* length is number of elements that come under the slice
	* capacity is number of elements present in the underlying array starting from current slice
	
	```go
	s := []int{1, 2, 3, 4}
	len(s) is 4
	cap(s) is 4
	
	s = s[:2]
	len(s) is 2
	cap(s) is 4
	
	s = s[:4] // extend the slice's length
	len(s) is 4
	cap(s) is 4
	
	s = s[2:] // drop first two elements, therefore capacity reduced by 2
	len(s) is 2
	cap(s) is 2
	```
* To create a dynamically-sized array we use built-in *make* function.  
	It allocates a zeroed array and a slice reference to it.   
	`darr := make(int[], 5)` // len = 5
	
* We got dynamic arrays, now how do we add elements to it?  
	Using built-in *append* function.  
	```go
	var s []int
	s = append(s, 1)
	s = append(s, 2, 3, 4)
	```
	It takes the slice and list of elements to be added, returning a slice.  
	Returns a slice to a bigger array if required.
	
* If time to spare, go through [Slice internals blog](https://blog.golang.org/go-slices-usage-and-internals)


#### Range

* Ranging for loop over slices returns two values -- index and copy of element at that index
```go
	for index, value := range slice {
		// do something
	}
```

* If index not required then an underscore can be placed instead of it.

#### Map

* Creating a map is done using built-in `make(map[keyType]valueType)` function

* Similar to struct literals, we have map literals.  
```go
idToPersonMap := map[int]Person {
		1: Person {name: "Leo"},
		2: {name: "Cristiano"} // it is ok to omit the struct type name
	}
```

* Insert in a map `m[key] = value`

* Retrieve a value from the map using `value = m[key]`

* Delete a key in a map using function `delete(map, key)`

* Test if a key is present using two-value assignment  
  `value, isPresent := map[key]`   
	// if absent then value is the appropriate zero-value

#### Functions

* Functions are values too i.e. they can be passed as arguments or returned from functions.


#### Closures
* Closures are functions in Go which access, assign, reference variables from outside it's own body.

## Methods and Interfaces

#### Methods

* Methods are functions defined on a type.

* Their signature contains a special receiver argument of the type on which they're defined.

```go
// Vertex is the special receiver argument of type struct
func (v Vertex) Abs() float64 {
	return (v.X * v.X + v.Y * v.Y)
}
```
* Methods can be defined on non-struct types too.
```go
type MyInt int

func (i *MyInt) Increment() {
	*i++
}
```

* Methods can only be defined on types declared in the same package.

* Methods by default pass by value i.e. create a copy of recevier.   
  To modify the receiver you use pointer receivers.  
	(as in the MyInt example above)

* Pointer receivers are of the form `*T` where T itself cannot be a pointer. 

* For convenience, Go lets you invoke methods with pointer receivers on a pointer type as well as value type.  
  The vice versa is true as well.

	The MyInt example above can be invoked as

```go
 i := MyInt(1)
 i.Increment() // go treats this as (&i).Increment()
```
* You should use pointer receivers as they avoid copying the receiver and you'd mostly require modifying the receiver.

#### Interfaces

* An interface type is defined as a set of method signatures.

* Value could be any value that implements those methods.

* Interface values can be thought of as being a tuple of (value, concrete type)

* A type implements an interface implicitly by implementing the methods.  
 There is no explicit keyword declaration.

 ```go
	type I interface {
		Foo()
	}

	type IImpl struct { }

	func (i IImpl) Foo() {
		// do something
	}

	var i I
	impl := IImpl{}
	i = impl // Impl implements I
	i = &impl
 ```

* Pointer types `*T` *method set* includes the non-pointer receiver methods as well but the reverse is not true.

	Method sets are methods you can call on the type and that determine the interfaces a type implements.

	See method sets in Go lang spec.

```go
type Impl2 struct{}

func (i *Impl2) Foo(){}

var i I
impl := Impl2{}
i = impl // doesn't work

// error raised:
// Impl2 does not implement I (Foo method has pointer receiver)

i = &impl // works
```

* Interface values with nil concrete values invoke the methods with a nil receiver and not throw null pointer exception.  
```go
var i I
var impl *IImpl // zero value is nil
i = impl
i.Foo() // allowed
```  

* However invocations on nil interface values are runtime errors.
```go
var i I
i.Foo() // runtime error
```

#### Empty interface

* An interface with no methods is known as an empty interface  
` interface{} `

* It can hold values of any type.

* Used by code that handles values of unknown type.

#### Type assertion

* Access to an interface value's underlying concrete value.

* The form of type assertion is:  
  ```t := i.(T)```
  
	This statement asserts that i holds a value of type T and assigns the underlying value to t, else triggers a panic.  

	(analogous to Java's type safety instanceof check followed by a type cast)  

	For example,
	```go
	var i interface{} = "hello"
	t := i.(string)
	t, ok = i.(float64) // t will be zero value for float64, ok is false
	t = i.(float64) // assertion without ok, panic!
	``` 

* Type switches let you do a series of type assertions.  
  The cases specify a number of types to compare to an interface value to.
  ```go
	switch v := i.(type){ // type is a keyword here
		case int:
		case string:
		default:
	}
	```


## Questions

* Do we have access specifiers in Go?

* Go passes arguments to functions as pass-by-value?

* Does Go compile to machine-specific assembly?

* Examples of using the empty interface

* Are panics the equivalents of run time exceptions?

* How does multiple assignment work without overwriting each other?  
	For example in fibonacci, this works  
	 *a, b = b, a + b*  
	
  Go inspects this statement involving multiple assignments, where an assignment includes accessing a variable that in turn got assigned before in the same statement.  
Seeing this, it generates the required temporary variables.

```
        0x0008 00008 (exercise-fibonacci-closure.go:10) MOVQ    (CX), DX
        0x000b 00011 (exercise-fibonacci-closure.go:10) MOVQ    (AX), BX
        0x000e 00014 (exercise-fibonacci-closure.go:10) MOVQ    DX, (AX)
        0x0011 00017 (exercise-fibonacci-closure.go:10) ADDQ    BX, DX
        0x0014 00020 (exercise-fibonacci-closure.go:10) PCDATA  $2, $4
        0x0014 00020 (exercise-fibonacci-closure.go:10) MOVQ    DX, (CX)
        
        where line 10 is:
        a, b = b, a + b
        
        (generated using go tool compile -S exercise-fibonacci-closure.go)
```