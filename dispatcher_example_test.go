package dispatcher_test

import (
	"fmt"
	"strconv"

	"github.com/kamiazya/go-dispatcher"
)

func ExampleDispatcher() {

	// greet task generator
	greetTaskGen := func(j int) dispatcher.Task {
		return func() error {
			fmt.Println("hello", strconv.Itoa(j))
			return nil
		}
	}

	// create dispacher
	d, err := dispatcher.New()
	if err != nil {
		fmt.Println(err.Error())
	}

	// start dispacher
	d.Start()

	for i := 1; 10 >= i; i++ {
		// dispach task
		d.Dispatche(greetTaskGen(i))
	}

	d.Wait()

	// Unorderd output:
	// hello 1
	// hello 2
	// hello 3
	// hello 4
	// hello 5
	// hello 6
	// hello 7
	// hello 8
	// hello 9
	// hello 10
}
