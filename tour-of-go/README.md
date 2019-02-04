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
* [Questions](#questions)


## Basics

#### Declaring types and identifiers

* A name is exported if it begins with a capital letter.

* Type is specified after identifier name.  
	*func add(x int, y int)*

* Two or more consecutive identifiers have the same type then you are allowed to state the type only once.  
	*func add(x, y int)*

* Function can return multiple results

* Function's return values can be named (to document results in a meaningful way)  
	*func add(x, y int) (result int)*
	
	Thereby no need to name them explicitly in the return statement (but the return keyword is still needed)


* var declarations can also initialise variables then and there thereby not requiring explicit type declaration.  
	*var sum = 100*

* var declaration is required for variables declared outside functions.

* Inside functions, var declarations with initialisations can use a short hand operator :=
```
func add(x, y int) (sum int) {
	sum := x + y
	return
}
```

#### Basic types

* int, uint are 32bits on 32bit systems and 64bit on 64bit systems.  
	For system independent integers we have int32, int64, etc.

* Amongst the usual basic types, string and runes are basic types too in Go.  
	Rune is an alias for int32, representing a Unicode code point (code point is a numeric value from the Unicode space)

#### Type conversion

* Expressions T(v) are type conversion expressions which convert v to type T.  
	These are always explicitly needed.

	```
	x, y := 3, 4
	div := float64(x)/float64(y) // not type converting either of x or y is an error
	fmt.Println(div)
	```

#### Constant declaration
* const Pi = 3.14

## Flow control 

* Syntax for if conditionals and for loops do not require parentheses but do require the body always be enclosed in {}

* Switch statements don't fall-through and the cases do not need to be constants i.e. they could be expressions.

* Deferred function calls are executed after the surrounding function returns. Multiple defers are stacked.

## More types

#### Pointers

* Go has pointers. Using * to dereference and & address-of.
	Pointers hold memory address of a value.
	
* Zero value is *nil*

* No pointer arithmetic is permitted.
	
#### Structs

* Structs hold many fields (POJOs essentially) 
```
type Vertex struct {
	X int
	Y int
}
```

* Members could be accessed directly using dot operator


* Structs are constructed using struct literals  
	*Vertex{1,2}*  
	*Vertex{X:1}* // Y is 0 implicitly
	
* Accessing struct fields using a pointer can also be done using the dot operator (rather than the cubersome (*p).X)
	
#### Arrays
* Array declaration type is of the form *[n]T*

* Arrays are of fixed size, zero indexed.

* oddNos := [4]int{1, 3, 5, 7} is an array

#### Slices
* Slice declaration type is of the form *[ ]T*

* Slices are dynamically-sized view of arrays (references to a portion of the array)

* Obtain a slice using a[low: high], high isn't included.  
	You may omit low or high or both and defaults (0, length) will be used.

* var firstTwo []int = oddNos[0:3] // [1, 3]

* Modifying a slice modifies the underlying array and all other slices that cover the same modified range. 

* *naturalNos := []int{1, 2, 3}* creates an array and returns a slice reference to it  
	*naturalNos = naturalNos[:2]* // [1, 2]  
	*naturalNos = naturalNos [1:]* // [2]  

* Slice has a length as well as a capacity.  
	* length is number of elements that come under the slice
	* capacity is number of elements present in the underlying array starting from current slice
	
	```
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
	*darr := make(int[], 5)* // len = 5
	
* We got dynamic arrays, now how do we add elements to it?  
	Using built-in *append* function.  
	```
	var s []int
	s = s.append(s, 1)
	s = s.append(s, 2, 3, 4)
	```
	It takes the slice and list of elements to be added, returning a slice.  
	Returns a slice to a bigger array if required.
	
* If time to spare, go through [Slice internals blog](https://blog.golang.org/go-slices-usage-and-internals)


#### Range

* Ranging for loop over slices returns two values -- index and copy of element at that index
```
	for index, value := range slice {
		// do something
	}
```

* If index not required then an underscore can be placed instead of it.

#### Map

* Creating a map is done using built-in *make(map[keyType]valueType)* function

* Similar to struct literals, we have map literals.  
```
idToPersonMap := map[int]Person {
		1: Person {name: "Leo"},
		2: {name: "Cristiano"} // it is ok to omit the struct type name
	}
```

* Insert in a map *m[key] = value*

* Retrieve a value from the map using *value = m[key]*

* Delete a key in a map using function *delete(map, key)*

* Test if a key is present using two-value assignment  
  *value, isPresent := map[key]* // if absent then value is the appropriate zero-value

#### Functions

* Functions are values too i.e. they can be passed as arguments or returned from functions.


#### Closures
* Closures are functions in Go which access, assign, reference variables from outside it's own body.



## Questions

* Do we have access specifiers in Go?

* Does Go compile to machine-specific assembly?

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