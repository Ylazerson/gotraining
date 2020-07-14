# B''H 


### Notes

* Slices are like dynamic arrays with special and built-in functionality.
* There is a difference between a slices length and capacity and they each service a purpose.
* Slices allow for multiple "views" of the same underlying array.
* Slices can grow through the use of the built-in function append.

### Links

[Go Slices: usage and internals](https://blog.golang.org/go-slices-usage-and-internals) - Andrew Gerrand    
[Strings, bytes, runes and characters in Go](https://blog.golang.org/strings) - Rob Pike    
[Arrays, slices (and strings): The mechanics of 'append'](https://blog.golang.org/slices) - Rob Pike    
[Understanding Slices in Go Programming](https://www.ardanlabs.com/blog/2013/08/understanding-slices-in-go-programming.html) - William Kennedy    
[Collections Of Unknown Length in Go](https://www.ardanlabs.com/blog/2013/08/collections-of-unknown-length-in-go.html) - William Kennedy    
[Iterating Over Slices In Go](https://www.ardanlabs.com/blog/2013/09/iterating-over-slices-in-go.html) - William Kennedy    
[Slices of Slices of Slices in Go](https://www.ardanlabs.com/blog/2013/09/slices-of-slices-of-slices-in-go.html) - William Kennedy    
[Three-Index Slices in Go 1.2](https://www.ardanlabs.com/blog/2013/12/three-index-slices-in-go-12.html) - William Kennedy    
[SliceTricks](https://github.com/golang/go/wiki/SliceTricks)    

### Code Examples

[Declare and Length](example1/example1.go) 
[Reference Types](example2/example2.go) 
[Appending slices](example4/example4.go) 
[Taking slices of slices](example3/example3.go) 
[Slices and References](example5/example5.go) 
[Strings and slices](example6/example6.go) 
[Variadic functions](example7/example7.go) 
[Range mechanics](example8/example8.go) 
[Three index slicing](advanced/example1/example1.go) 


<br>
---

### 3.3 Slices—Part 1 (Declare and Length and Reference Types)

**Slices** - this is something that you must learn, **you must master**, you can't cheat on because all of the data you'll be working with or at least the majority of it should be and probably will be stored in slices. This is your go-to data structure. 

<br>
---

##### `make` function 

It allows us to create **3** of the **reference types** that are also built into the language: 
- slice
- map
- channel
- interface values 
- functions


They're reference types because they are data structures that have a **pointer**. They're also reference types because when any of these types are set to their **zero value**, they're considered to be **`nil`** in this language. 

**A string is actually very close to being a reference type**. The problem is that when a string is set to its zero value, **it's not `nil`, it's empty** so I can't really put it in that class. 

---
<br>

For a slice, we're gonna use `make` when we already know ahead of time how much memory to allocate towards its backing data structure which is an array. 

```go
fruits := make([]string, 5)
```

![](img/slice.png)

A slice is a **three word** or 24 byte data structure on our AMD64 architectures and it's very similar to the string where you get a pointer and you have the length of bytes. In this case not bytes though, but the length here of five which means we're gonna have five strings. 

And when the length is only set on the make call, then the capacity matches the length. 

---
<br>

**The slice value just like the string is designed to be using value semantics.** As we're gonna learn our built in reference types are designed around value semantics. They're designed to be kept on the stack. 

![](img/slice-2.png)


I don't wanna see you sharing (with pointers) the slice value. It is designed to stay on your stack, it's designed to be using value semantics and we should be making copies of it. 

