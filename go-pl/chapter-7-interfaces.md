
#### Introduction

* Interfaces are abstractions about behaviours of other types.

* In Go, interfaces are satisfied implicitly.

  No need for explicit declaration on the concrete type for it to satisfy an interface.

  Possessing the required methods is enough.

* Therefore types not under your control can also be made to satisfy an interface. 

#### Interfaces as contracts

* The freedom to substitute one type for another that satisfies the same interface is called _substitutability_.

  It is the key to design adaptable APIs.

  And in Go, interfaces give you substitutability.

#### Interface types

* `io.Writer` is most widely used interface that provides an abstraction over all types to which bytes can be written.

  It could be a file, in-memory buffer, network connection, archiver, hashers, etc.

* Naming convention for interfaces with only a single method is to suffix the method with a _er_.

  method Write, interface Writer  
  method Close, interface Closer

* Just like struct embedding, you can embed other interfaces in an interface.

    ```go
    type ReadWriter interface {
        Reader
        Writer
    }
    ```

#### Interface satisfaction

* A value of type T does not possess all the methods that a *T pointer does and therefore they satisfy a different set of interfaces.

  Althought the invocation of *T methods on a T value is allowed, it is only a syntactic sugar.

  ```go
    type IntSet struct {...}
    func (set *IntSet) String() string {...}
    
    var s IntSet
    s.String() // OK, compiler converts it to invoking string method on the pointer we get by taking the address of variable s.
    (&s).String() // OK

    var _ fmt.Stringer = &s // OK since *IntSet has a method String() defined
    var _ fmt.Stringer = s // compile error

  ```

* `godoc -analysis=type <package-name>`  
   does static analysis for types defined in the specified package and lists the interfaces a concrete type implements.

* `interface{}` is the empty interface.

  No demands on the methods types should have to satisfy it.

* If the method requires to mutate the internal state contained in a concrete type, it'll always take a pointer receiver.

* We can define new groupings over concrete types and create an interface for even types not defined by us.  

  Since the concrete types would have those commonalities in methods and those would be in the public API, we can depend on it futuristically not to expect the satisfiability of the interface to change in future.

* A program that deals with digitized artifacts like music, book, film could have following concrete types.

  Magazine  
  Movie  
  Book  
  Album  
  TVPodcast  

  Some properties are common to all artifacts which can be expressed as an interface.

  ```go
  type Artifact interface {
      Title() string
      Creators() []string
      Created () time.Time
  }
  ```

  Some properties would be common to only certain type of artifacts like only printed artifacts have sense of pages.

    ```go
  type Text interface {
      Pages() int
  }

  type Streamer interface {
      Stream() (io.ReadCloser, error)
      RunningTime() time.Duration
      Format () string // MP3, MP4
  }

  type Audio interface {
      Streamer
  }

  type Video interface {
      Streamer
      Resolution() (x, y int)
  }
  ```


#### Interface values

* Interface values have two components a dynamic type and a dynamic value.

* An interface assignment involves an implicit conversion from a concrete type to an interface type.
   
   The dynamic type of interface is set to the concrete type and dynamic value is a copy of the type's value.

* Calling a method on a nil interface causes panic.

* If interface's dynamic values are not comparable for example in case of slices, maps, functions, then a panic occurs.

  Therefore care must be taken in using them in comparable situations like switch statements or map keys.

* You can embed interfaces inside structs.

#### http.Handler interface

Function types may also satisfy interfaces and become adapters to allow a multiple methods from a single type to satisfy the same interface.

For example,

* The Handler interface requires the following method for interface satisfaction:

  `ServeHTTP(ResponseWriter, *Request)`

* A single type wants to setup multiple of its own methods as http handlers.

  ```go
    type Foo struct { ... }
    func (f *Foo) foo(ResponseWriter, *Request){...} // on /foo
    func (f *Foo) fooz(ResponseWriter, *Request){...} // on /fooz
    func (f *Foo) baz(ResponseWriter, *Request){...} // on /baz
  ```
  Making Foo satisfy the Handler interface would only let us have one ServeHTTP method defined.  

  We'd then need to match the URL in an if-else branch.  
  Imagine when there are multiple types, each wanting to register their own handlers - every type would have one and only one ServerHTTP to satisfy the Handler interface, which would match the URL on a case by case basis.
  
  This becomes inconvenient.
  
  To plug in multiple methods of the same type to ServeHTTP, we use an adapter function type that itself satisfies the Handler interface:
  ```go
    type HandlerFunc func(ResponseWriter, *Request)

    func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
        f(w, r)
    }
   ```
    Now we simply use http.HandlerFunc(f.foo), http.HandlerFunc(f.fooz) and so on to setup multiple handlers.

#### error interface

* In the `errors` package we have an unexported type that implements the error interface `*errorString`.

  The reason `*errorString` implements the interface and not `errorString` itself is because the latter's dynamic interface value would be `string` whose equality comparison would be true for two strings having the same characters.

  With errors that is not desirable.  
  And therefore with `*errorString`, dynamic interface value returned by `errors.New` is a pointer value to the `errorString` instance which would be unique every time.

#### Type assertion

* Type assertion is an operation applied to interface values.

* x.(T)

  x is an expression of interface type.

  If T is a concrete type, type assertion checks if x's dynamic type is T.  
  If it is then the dynamic value of type T is extracted.

  ```go
    var x io.Writer
	x = os.Stdout
	stdout := x.(*os.File)
	fmt.Printf("%T\n", stdout) // extracted os.Stdout value of type *os.File
	fmt.Println(stdout.Name()) // /dev/stdout
  ```  

  If T is an interface type, then type assertion checks if x's dynamic type satisfies that interface.  
  If it does an interface value of type T is returned but with it's dynamic type, value being same as before.

  ```go
    var x io.Writer
	x = os.Stdout
	rw := x.(io.ReadWriteCloser)
	if err := rw.Close(); err != nil {
		t.Fatal(err)
	}
  ```

* Type assertion check without the ok return value panics if the check fails.

#### Type switches

* Multi way type assertion.

* No fall through of cases.

* Multiple cases can share their case.

* Convention is to receive type check result in the same variable name which has it's own lexical scope of visibility.
   
   Inside the cases the new variable can be used which will have the same type as the case.

   In other cases it will have the type that it already had.

  ```go
    switch x := x.(type) {
        case int, unint:
        case bool:
        case nil:
        default:
    }
  ```

#### A few words of advice 

* Interfaces are needed when there are two or more concrete types which have to be dealt with in a uniform way. 

* An exception for using interfaces even though having only one concrete type is when they need to be in different packages because of the dependencies needed by concrete type. Decoupling packages.

* A good rule of thumb for interface design is _ask only for what you need_.

#### Questions

* What values in Go are comparable? What are not?
* In what all combinations can you embed structs inside interfaces, etc?