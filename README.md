# gowebflow
Golang bindings for webflow api

# Install
```go get github.com/jancimertel/gowebflow```

# Usage
```
package main

import (
	"fmt"
	"github.com/jancimertel/gowebflow"
	"github.com/jancimertel/gowebflow/response"
)

func main() {
	// create client instance
	client, err := gowebflow.NewClient("<YOUR_API_KEY>")
	if err != nil {
	    panic(err)
	}
	
	// container for all items
	var allData []response.Item
	
	// get all items from one collection
	page := uint(0)
	for {
		var pageData []response.Item
		hasNextPage, err := client.PaginateItems("<COLLECTION_ID>", page, &pageData)
		if err != nil {
			panic(err)
		}
		
		for _, item := range pageData {
			allData = append(allData, item)
		}

        // no more items: (offset + page items) >= total items
		if !hasNextPage {
			break
		}
		page++
	}
	
	fmt.Println(len(allData))
}
```

# Notes
Instead of reflection, public method are using pointers to containers. It should always be a pointer to slice, be it 
default `&[]response.Item` or your custom one.