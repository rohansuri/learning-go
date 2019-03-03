package erratum

import (
	"log"
	"os"
)

var info = log.New(os.Stdout, "INFO: ", log.LstdFlags | log.Lshortfile)

func Use(o ResourceOpener, input string) (err error) {
	var resource Resource

	// open the resource for the first time
	// semantically should be out of the loop
	resource, err = o()

	// if resource fails to open, we have to retry
	for err != nil {

		if _, ok := err.(TransientError); !ok {
			return err
		}

		info.Printf("A transient error has occurred. " +
			"Will retry opening the resource.\n")

		resource, err = o()
	}

	defer func() {
		// try closing, but if close fails, let's not lose the very first
		// error we got.
		// best way would be to form a chain of errors?
		// but close's failure wouldn't necessarily be caused-by the panic
		// Java's equivalent of this is suppressed exceptions
		cerr := resource.Close()
		if err == nil {
			err = cerr
		}
	}()


	defer func() {
		if r := recover(); r != nil {

			if frobErr, ok := r.(FrobError); ok {
				resource.Defrob(frobErr.defrobTag)
			}

			err = r.(error)
		}
	}()

	// start using the resource
	resource.Frob(input)

	return
}

