Notes from 

* Go lang official blog on "Defer, Panic, Recover" [here](https://blog.golang.org/defer-panic-and-recover)

* Go lang [spec](https://golang.org/ref/spec#Run_time_panics) on errors, runtime panics

* Effective Go's error [section](https://golang.org/doc/effective_go.html#errors)


#### Panic

* Built-in function that stops ordinary flow of control.

* Panic caller's deferred statements are executed and control returns to its caller.

  To the caller, the call yet again behaves as a panic and therefore caller's deferred statement start executing.

  This goes up the call stack until someone recovers and if no one does, the program crashes.

* Panics are initiated by calling the built-in panic directly or by runtime panics like array's out of bounds access.

* In a runtime panic, a value of interface type `runtime.Error` is passed to panic.

#### Recover

* A call to recover regains control of a panicking goroutine and resumes normal execution.

  Note:  
  The function that called recover would be in it's deferred block, which is always executed after the function returns and therefore execution resumes normally from the calling function (see example below).

* returns nil if the goroutine isn't panicking, else returns the value passed to panic.

* Since only deferred functions are executed when panicking, using recover only makes sense to be used inside a defer function.

#### Defer

* Another control flow mechanism.

* Defer statement pushes a function call onto a list.

* List of saved calls is executed after the surrounding function returns but before the control returns to its caller.

* List of saved calls is executed in LIFO order.

* Used to perform clean-up actions.
 
  defer closing a file right after opening it.  
  defer unlocking a mutex after acquiring it.

* Function arguments in the deferred call are evaluated when the statement is seen.

* Return values from a deferred function are discarded.

* In case of returns without named parameters, deferred functions do not have access to internal result parameters and they are only set by the explicit return statement.

  ```go
  func c() (int) { // returns 1
    i := 0
    defer func() { i++ }()
    return 1
  }

  ```

* In case of returns with named parameters, deferred functions can read/assign to them.

  ```go
  func c() (i int) { // returns 2
    defer func() { i++ }()
    return 1 // i set to 1 by explicit return, then incremented by deferred func call
  }
  ```

  Useful in propagating errors after recovering from panic.


#### Defer, Panic, Recover in action
```go

func main(){
    fmt.Println(f())
    fmt.Println("Back to normal execution")
}

func f() (err error) {
	defer func() {
		if r := recover(); r != nil {
			if rerr, ok := r.(error); ok {
				err = rerr
			}
		}
	}()
	g()
	fmt.Println("I'm not executed, since g() always panics") // normal execution resumes in caller main, not in the recoverer
	return
}

func g() {
	panic(fmt.Errorf("I always panic"))
}
  ```

#### Simplifying error handling



#### Review

Defer, panic, recover seems different from try-catch-finally.

In try-catch-finally execution resumes right where you catch the exception.

In defer-panic-recover, execution resumes in the caller of the function where panic was recovered from.

