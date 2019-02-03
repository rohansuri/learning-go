Important points extracted out from "A Tour of Go"

## Basics

#### Declaring types and identifiers

* A name is exported if it begins with a capital letter.

* Type is specified after identifier name.  
	*func add(x int, y int)*

* Two or more consecutive identifiers have the same type then you are allowed to state the type only once.  
	*func add(x, y int)*

* Function can return multiple results

* Function's return values can be named (to document results in a meaningful way), then requiring no explicit return statement.  
	*func add(x, y int) (result int)*


* var declarations can also initialise variables then and there thereby not requiring explicit type declaration.  
	*var sum = 100*

* var declaration is required for variables declared outside functions.

* Inside functions, var declarations with initialisations can use a short hand operator :=
```
func add(x, y int) (sum int) {
	sum := x + y
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

#### Constant declaration:  
* const Pi = 3.14

## Flow control 

* Syntax for if conditionals and for loops do not require parentheses but do require the body always be enclosed in {}

* Switch statements don't fall-through and the cases do not need to be constants i.e. they could be expressions.

* Deferred function calls are executed after the surrounding function returns. Multiple defers are stacked.

## Questions

* Why do I have to type var to declare variables?
