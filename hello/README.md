Packages hello and stringutil created as followed from the "How to Write Go Code" guide

### Learnings:

*	One workspace as defined by $GOPATH (default is $HOME/go)  
	Inside which you have /src and /bin.  
	Inside of src you have multiple repos, which in turn contain many packages.  
	There isn't any restriction on the go file name, but package name should be last directory folder in the import path. 
	
* 	*go cli tool* is a
	*  build tool
	* test tool (go test, runs files suffixed by _test.go)
	* dependency management tool (go get)

* 	1 folder = 1 executable  
	i.e. you cannot have more than one main entry points in a single import path (also [see](https://www.reddit.com/r/golang/comments/35ntmp/how_to_sort_your_projects_in_one_folder_without/))


	Don't fight the tool, create one executable per package.

	On the other hand it makes sense to have only one entry point per package,  
	if you need more to just "run it around", you actually are wanting to test it around. 
	 
	So rather do the invocations from a _test.go. Each test case is a new entry point.

* rune type (alias for int32)  
	To cover entire Unicode space  
	Character literals are runes by default [see](https://stackoverflow.com/a/51611567/3804127)

	s := "hello你好" as string it's length is 11  
	5 single byte chars + 2*3 Chinese chars

	but as rune it's length is 7
	
	Therefore iterating over runes makes more sense since runes represent one logical character in the Unicode space, whereas strings are made up of single-single bytes.  
	(Think about reversing a Unicode string byte by byte and the error one would commit)


### TODO:
*	Why does the swap assignment work in Go (question on understanding the language construct)?