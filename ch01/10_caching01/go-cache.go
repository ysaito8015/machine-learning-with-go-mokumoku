package main

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"time"
)

func main() {
	// Create a cache with a default expiration time of 5 minutes, and which
	// purges expired items every 10 minutes
	c := cache.New(5*time.Minute, 10*time.Minute)

	// Set the value of the key "foo" to "bar", with the default expiration time
	// Key: foo, Value: bar
	c.Set("foo", "bar", cache.DefaultExpiration)

	// Set the value of the key "baz" to 42, with no expiration time
	// (the item won't be removed until it is re-set, or removed using
	// c.Delete("baz")
	c.Set("baz", 42, cache.NoExpiration)

	// Get the string associated with the key "foo" from the cache
	foo, found := c.Get("foo")
	if found {
		fmt.Printf("Key: foo, Value: %#v\n", foo)
	}

	// Since Go is statically typed, and cache values can be anything, type
	// assertion is needed when values are being passed to functions that don't
	// take arbitrary types, (i.e. interface{}). The simplest way to do this for
	// values which will only be used once--e.g. for passing to another
	// function--is:
	baz, found := c.Get("baz")
	if found {
		fmt.Println(baz.(int) * 2.0)
	}

	// This gets tedious if the value is used several times in the same function.
	// You might do either of the following instead:
	/*
		if x, found := c.Get("foo"); found {
			foo := x.(string)
			// ...
		}
		// or
		var foo string
		if x, found := c.Get("foo"); found {
			foo = x.(string)
		}
		// ...
		// foo can then be passed around freely as a string

		// Want performance? Store pointers!
		c.Set("foo", &MyStruct, cache.DefaultExpiration)
		if x, found := c.Get("foo"); found {
			foo := x.(*MyStruct)
			// ...
		}
	*/
}
