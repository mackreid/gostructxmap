### In development (documentation included)

## Go Struct x Map 
Will convert a struct to a map, with the ability to define your own rules (or use pre-made ones).

## Install
`go get github.com/mackreid/gostructxmap`

## Basic Useage
```
// define your struct with added `sxm` tags
type Book struct {
    ID string `sxm:"-"`
    Name string `sxm:"name"`
    Author string `sxm:"author,omitempty"`
    Pages int `sxm:"pages"`
}

testStruct := Book {
    ID: "123",
    Name: "Test Book Name",
    Pages: 250
}

testMap := gostructxmap.Make(book)
// testMap Output
{
    "name": "Test Book Name",
    author: "",
    pages: 250,
}
```

## Pre Defined Rules
To ignore a field in the struct you should ignore the tag completely or explicitly set the tag `sxm:"-"`
### WithOmit
```
// define your struct with added `sxm` tags
type Book struct {
    ID string `sxm:"-"`
    Name string `sxm:"name"`
    Author string `sxm:"author,omitempty"`
    Pages int `sxm:"pages"`
}

testStruct := Book {
    ID: "123",
    Name: "Test Book Name",
    Pages: 250
}

testMap := gostructxmap.Make(book)
// testMap Output - author has been omitted 
{
    "name": "Test Book Name",
    pages: 250,
}
```

### WithMask
... more to come ...