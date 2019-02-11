
Important go cli invocations.

#### Run tests in all subdirectories

For example, to run all tests written as part of Tour of Go  

$ pwd  
/home/rohansuri/go/src/github.com/rohansuri/learning-go/tour-of-go  

$ go test ./...  
?   	github.com/rohansuri/learning-go/tour-of-go/closures	[no test files]  
?   	github.com/rohansuri/learning-go/tour-of-go/interfaces/reader	[no test files]  
ok  	github.com/rohansuri/learning-go/tour-of-go/interfaces/stringer	(cached)  
?   	github.com/rohansuri/learning-go/tour-of-go/maps	[no test files]  
ok  	github.com/rohansuri/learning-go/tour-of-go/slices	(cached)

TODO: why does this work? (refer to some documentation)

