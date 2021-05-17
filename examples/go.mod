module github.com/oyekanmiayo/go-notion/examples

go 1.15

replace github.com/oyekanmiayo/go-notion/notion => ../notion

require (
	github.com/dghubble/sling v1.3.0 // indirect
	github.com/oyekanmiayo/go-notion/notion v0.0.0-00010101000000-000000000000
	golang.org/x/oauth2 v0.0.0-20210514164344-f6687ab2804c // indirect
)
