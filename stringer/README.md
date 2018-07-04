# go generate sample

Try "go generate" Go1.4~

.go file
```
//go:generate stringer -type=Fruit
package main
```

```
# a tool called stringer
go get golang.org/x/tools/cmd/stringer
```

stringer generate source
	String() method of type of "Fruit"

So, you can use var name for fmt.Println()!!!

# how to run
```
go generate
go run *.go
# ---> new file called fruit_string.go
```

FYI:
http://mattn.kaoriya.net/software/lang/go/20141208093852.htm
https://blog.golang.org/generate

Enjoy!
