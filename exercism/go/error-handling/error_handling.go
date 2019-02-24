package erratum

import (
	"os"
	"time"
	"log"
)

var info = log.New(os.Stdout, "INFO: ", log.LstdFlags | log.Lshortfile)
const retry = 0 * time.Second

func Use(o ResourceOpener, input string) (err error) {
	var resource Resource

	// while resource fails to open
	for {
		resource, err = o()
		if _, ok := err.(TransientError); ok {

			info.Printf("A transient error has occurred. " +
				"Will sleep for %v seconds and retry.\n", retry.Seconds())

			time.Sleep(retry)

		} else if err != nil { // not a transient error
			return err
		} else {
			break
		}
	}

	defer func() {
		// try closing, but if close fails, let's not lose the very first
		// error we got
		// best way would be to form a chain of errors?
		// but close's failure wouldn't necessarily be caused-by the panic
		cerr := resource.Close()
		if err == nil {
			err = cerr
		}
	}()


	defer func() {
		if r := recover(); r != nil {

			switch e := r.(type) {

				case FrobError:
					resource.Defrob(e.defrobTag)
					err = e

				case error:
					err = e
			}
		}
	}()

	// start using the resource
	resource.Frob(input)

	return
}
