Notes from official Go blog on [error handling in Go](https://blog.golang.org/error-handling-and-go).

* The `error` type is an interface type that represents any value that can describe itself as a string.

* Most commonly used error implementation is `package errors` unexported `errorString` type.

  Function `errors.New(string) error` constructs a value of `errorString` type and returns it as an error.

  The same can be constructed using `fmt.Errorf(....) error` to add more formatted context to describe the error.

  For example, "opening /a/b/c failed. permission denied"

* To present more context and allow callers to inspect details of an error, a custom error type can also be defined.

  Later using type assertion, it's details can be inspected. 

  Example `net` package [defines](https://golang.org/pkg/net/#Error) custom `net.Error` type.
    ```go
  type Error interface {
      error
      Timeout() bool
      Temporary() bool 
  }
    ```
    Temporary() could be used to distinguish transient errors from permanent ones and accordingly be retried.

    ```go
    if nerr, ok := err.(net.Error); ok && nerr.Temporary() {
        time.Sleep(1e9)
        continue
    }
    ```

    Note the above inline error check is an idiom in Go.
    
Nil error value not equal to nil  
(go lang FAQ)

* Interfaces are implemented as two elements, a type T and a value V.

* An interface value is nil only when both T = nil and V is not set.

* For example, storing a nil pointer of type T inside an interface value would give it the type *T and value nil and hence make the interface value non-nil.

* Therefore storing any concrete value in an interface makes it not nil.

* Pitfall

  ```go
    func foo() error {
        var e *MyError = nil // MyError is a struct type
        if !successful {
            e = &MyError{}
        }
        return e
    }

    foo() != nil 
    // Pitfall!
    // will always be true
    // since an interface value for error is created
    // with type = *MyError, value = nil

    foo() != (*MyError)(nil)
    // will only be true when there's an error
    // but this is not feasible as you'd have to know
    // actual type the interface value holds
    

    // Therefore always explicitly return nil
    // to ensure correct interface value construction
    func foo() error {
        if !successful {
            return &MyError{}
        }
        return nil
    }
  ```
* Recommendation is to always return a interface type and if value is "missing" then return nil explicitly.

Questions

* function types can also implement interfaces.

  What's the use of this?

  In the http handler example, why can't we do:
  
  ```go
  func viewRecordWrapped(w http.ResponseWriter, r *http.Request) {
      if err := viewRecord(w, r); err != nil {
          http.Error(w, err.Error(), 500)
      }
  }
  ```

* `json` package nowhere explicitly states  where it returns the `SyntaxError` type.

  How does one be sure about the caller's type assertion if-else branch required when calling `json.Decode` ?

  Similarly for the `net` package, is it obvious that only the `DNSError` could be returned when doing the associated function calls for example `LookupAddr` ?

* How to create custom error type with a constructor function that stores stack trace for debugging?