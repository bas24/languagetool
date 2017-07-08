Languagetool API wrapper in Golang

Wrapper for opensource text proofreading api provided by <a href="www.languagetool.org">Languagetool.org</a> 

Link to api documentation <a href="www.languagetool.org/http-api/swagger-ui/#/">here.</a>
Note: Api has limit of 20 calls per minute per IP.

Example usage:

```go
	package main

	import (
		"fmt"
		lt "github.com/bas24/languagetool"
	)

	func main(){
		text := `Text with eror.`
		result, err := lt.Check(text, "en-US")
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(result)
	}
```

Better not very nice code, than no code.