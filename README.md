Go In-memory cache
================================
Caching is a technique used to manage frequently accessed data in-memory. The method facilitates access to data while also reducing the workload on the main data stores. Not only can caching enhance performance; it can also improve the scalability and availability of an application.

## How install

```bash
go get -u github.com/DiasOrazbaev/cache
```

See it in action:
## Example
```go
package main

import (
  "fmt"
  "github.com/DiasOrazbaev/cache"
)

func main() {
  // create new instance of cache
  c := cache.New()

  // Set some value to cache
  c.Set("key", "value")

  // get some value by key
  fmt.Println(c.Get("key")) // output: value

  // delete from cache
  c.Delete("key")

  fmt.Println(c.Get("key")) // output: nil
}
```