package hello

import (
	"fmt"
	"testing"

	"github.com/rohansuri/learning-go/stringutil"
)

func TestHello(t *testing.T) {
	fmt.Println("hello world")
}

func TestReverse(t *testing.T) {
	fmt.Println(stringutil.Reverse("hello"))
}
