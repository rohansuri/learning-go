Notes from Go official blog on [errors are values](https://blog.golang.org/errors-are-values)

One might feel there's a lot of repetitive error handling and clutter of if err != null checks, for example:


```go
, err = fd.Write(p0[a:b])
if err != nil {
    return err
}
_, err = fd.Write(p1[c:d])
if err != nil {
    return err
}
// and so on
```

But there are techniques for graceful error handling.

One idea is to avoid repetitive error checking and do it only once after the entire operation is complete.

* `Scanner.Scan()` returns bool indicative of a successful scan of the next token or an unsuccessful scan in case of reaching end of input or an error.

  Encountering an error can be checked after finishing a looping scan by `Scanner.Err()` 
  
* The idea is to maintain an error state across multiple invocations and access it later after all invocations.

* Could be done using closures that maintains the state.

    ```go
    var err error
    write := func(buf []byte) {
        if err != nil {
            return // no-op
        }
        _, err = w.Write(buf)
    }
    write(p0[a:b])
    write(p1[c:d])
    // and so on
    // finally
    if err != nil {
        return err
    }

    ```

* If the same call pattern arises at multiple sites the closure would be duplicated.

  To associate independent error state with methods from multiple independent call sites, we create a struct.

    ```go
    type errWriter struct {
        err error
        w io.Writer
    }

    func (ew *errWriter) write(buf []byte){
        if err == nil {
            _, ew.err = ew.w.write(buf)
        }
    }

    ew := &errWriter{w: fd}
    ew.write(p0[a:b])
    ew.write(p1[c:d])
    // and so on
    if ew.err != nil {
        return ew.err
    }

    ```
* Caveat: what we have is simply an all or nothing check at the end and there's no way to know how many invocations succeeded.