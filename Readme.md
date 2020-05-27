# Golang React
Simple UI library done for fun and to experiment with golang.

### Install
```bash
$ go get -u github.com/alindesign/go-react
```

### Docs
A documentation about api it's available on [GoDoc](https://godoc.org/github.com/alindesign/go-react)

### Usage

#### Default example
```go
package main

import (
    "fmt"
    "github.com/alindesign/go-react"
)

func List(num int) *react.Element {
    var items []*react.Element
    for i := 0; i < num; i++ {
        items = append(items, react.CreateElement("li", react.Props{
            "className": fmt.Sprintf("list__item item-%d", i+1),
        }, react.Text("Item %d", i+1)))
    }

    return react.CreateElement("ul", react.Props{
        "className": "list__root",
    }, items...)
}

func App() *react.Element {
    return react.CreateElement("div", react.Props{
        "className": "root",
    }, List(4))
}

func main () {
    html := react.Render(App())
    
    fmt.Printf("Output: %s\n", html)
}

// Output: <div class="root"><ul class="list__root"><li class="list__item item-1">Item 1</li><li class="list__item item-2">Item 2</li><li class="list__item item-3">Item 3</li><li class="list__item item-4">Item 4</li></ul></div>
```
#### Elements
Built in shortcuts to HTML elements
```go
element := react.H1(nil, "child 2")
div := react.Div(nil, "child 1", element)
```
#### Fragments
```go
fragment := react.Fragment(
    react.H1(nil, "Title"),
    react.H3(nil, "text"),
)
```

#### Styles
Style properties will be converted to `kebab-case` so `backgroundColor` will become `background-color` 
```go
style := react.Style {
    "color": "#444",
    "backgroundColor": "#ecf0f1",
    "font-size": "14px",
}
```

#### Props
Properties will be converted to `kebab-case` so `DataString` will become `data-string`
```go
div := react.Div(react.Props{
    "className": "simple-div",
    "style": react.Style {
        "color": "#444",
        "backgroundColor": "#ecf0f1",
        "font-size": "14px",
    },
}, 
    react.Text("Hello world!"),
)
```

#### Document
Create a simple html boilerplate
```go
doc := react.Document(nil, nil)
// will render to: 
// <!doctype html><html lang="en"><head></head><body></body></html> 

customDoc := react.Document(&react.DocumentProps{
    BodyProps: react.Props{ "className": "body" },
    Body: react.H1(nil, react.Text("Hello!")),
}, nil) 
// will render to:
// <!doctype html><html lang="en"><head></head><body class="body"><h1>Hello!</h1></body></html> 
```

#### Components
```go
// Func components
func SimpleComponent (ctx *react.Context) *react.Element { 
    return react.Span(nil, react.Text("Hello!"))
}
custom := react.Component(SimpleComponent)
// will render to:
// <span>Hello!</span>

// Struct components
type SimpleComponent struct {}

func (s *SimpleComponent) Render (ctx *react.Context) *react.Element { 
    return react.Span(nil, react.Text("Hello!"))
}

custom := react.Component(&SimpleComponent{})
// will render to:
// <span>Hello!</span>
 
```

#### Render
Default renderers
```go
html := react.Render(react.H1(nil, react.Text("Hello!"))) // <h1>Hello!</h1>

html := react.RenderWithStats(react.H1(nil, react.Text("Hello!")))
// Print rendered nodes and time 
// <h1>Hello!</h1>

html := react.RenderBytes(react.H1(nil, react.Text("Hello!"))) // render []byte("<h1>Hello!</h1>")
```

Passing data to context
```go
func Greeter (ctx *react.Context) *react.Element {
    return react.H1(nil, react.Text("Hello %s!", ctx.GetString("name")))
}
html := react.Render(react.Component(Greeter), map[string]interface{}{
    "name": "Joe",
}) // <h1>Hello Joe!</h1>
```

Custom Renderer
```go
element := react.H1(nil, react.Text("Hi!"))
renderer := react.NewRenderer(element)
html := renderer.String()
```

#### Notes
Any text should be a react element instance:
```go
text := react.Text("Hi there!") // Hi there!
hello := react.Text("Hi there, %s!", "Joe") // Hi there, Joe!
h1 := react.H1(nil, react.Text("Hi there!")) // <h1>Hi there!</h1>
```

### Background
After a long night of no sleep and thinking on a better way to reuse HTML on golang, I started to search for something that has the possibility to create reusable components like on go lang, but without any success and tons of Javascript & React hours, I was thinking to adopt React component definition style, and implemented it in the way I was thinking they have done it.

### Contributing
If you have ideas or improvements you can contribute to this library.

### Disclaimer
I have done the project pure for fun and experimenting new things, I don't think that my implementation it's the best, so if you think that something can be improved, submit your PR.

### TODO
- Clean up
- Improve performance
