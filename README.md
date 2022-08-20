# simple-cache
Golang simple memory cache tool. 

## How to use:
```go
package main

import (
	"log"

	cache "github.com/Viquad/simple-cache"
)

func main() {
	cache := cache.NewMemoryCache()

	err := cache.Set("userId", 42)
	if err != nil {
		log.Println(err.Error())
		return
	}

	val, err := cache.Get("userId")
	if err != nil {
		log.Println(err.Error())
		return
	}

	log.Printf("Get() returned %v\n", val)

	err = cache.Delete("userId")
	if err != nil {
		log.Println(err.Error())
		return
	}
}
``` 