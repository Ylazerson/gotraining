# B"H





## Pointers

Pointers provide a way to share data across program boundaries. Having the ability to share and reference data with a pointer provides the benefit of efficiency. There is only one copy of the data and everyone can see it changing. The cost is that anyone can change the data which can cause side effects in running programs.


## Escape Analysis

* When a value could be referenced after the function that constructs the value returns.
* When the compiler determines a value is too large to fit on the stack.
* When the compiler doesn’t know the size of a value at compile time.
* When a value is decoupled through the use of function or interface values.



## Stack vs Heap

_"The stack is for data that needs to persist only for the lifetime of the function that constructs it, and is reclaimed without any cost when the function exits. The heap is for data that needs to persist after the function that constructs it exits, and is reclaimed by a sometimes costly garbage collection." - Ayan George

## Links

### Pointer Mechanics

   
   
[Using Pointers In Go](https://www.ardanlabs.com/blog/2014/12/using-pointers-in-go.html) - William Kennedy    
[Understanding Pointers and Memory Allocation](https://www.ardanlabs.com/blog/2013/07/understanding-pointers-and-memory.html) - William Kennedy    




### Escape Analysis and Inlining

[Go Escape Analysis Flaws](https://docs.google.com/document/d/1CxgUBPlx9iJzkz9JWkb6tIpTe5q32QDmz8l0BouG0Cw)  
[Compiler Optimizations](https://github.com/golang/go/wiki/CompilerOptimizations)

### Garbage Collection

[The Garbage Collection Handbook](http://gchandbook.org/)  
[Tracing Garbage Collection](https://en.wikipedia.org/wiki/Tracing_garbage_collection)  
[Go Blog - 1.5 GC](https://blog.golang.org/go15gc)  
[Go GC: Solving the Latency Problem](https://www.youtube.com/watch?v=aiv1JOfMjm0&index=16&list=PL2ntRZ1ySWBf-_z-gHCOR2N156Nw930Hm)  
[Concurrent garbage collection](http://rubinius.com/2013/06/22/concurrent-garbage-collection)  
[Go 1.5 concurrent garbage collector pacing](https://docs.google.com/document/d/1wmjrocXIWTr1JxU-3EQBI6BK6KgtiFArkG47XK73xIQ/edit)  
[Eliminating Stack Re-Scanning](https://github.com/golang/proposal/blob/master/design/17503-eliminate-rescan.md)  
[Why golang garbage-collector not implement Generational and Compact gc?](https://groups.google.com/forum/m/#!topic/golang-nuts/KJiyv2mV2pU) - Ian Lance Taylor  
[Getting to Go: The Journey of Go's Garbage Collector](https://blog.golang.org/ismmkeynote) - Rick Hudson  

### Static Single Assignment Optimizations

[GopherCon 2015: Ben Johnson - Static Code Analysis Using SSA](https://www.youtube.com/watch?v=D2-gaMvWfQY)  
[package ssa](https://godoc.org/golang.org/x/tools/go/ssa)    
[Understanding Compiler Optimization](https://www.youtube.com/watch?v=FnGCDLhaxKU)

### Debugging code generation

[Debugging code generation in Go](https://rakyll.org/codegen/) - JBD    

## Code Review

[Pass by Value](example1/example1.go) ([Go Playground](https://play.golang.org/p/9kxh18hd_BT))  
[Sharing data I](example2/example2.go) ([Go Playground](https://play.golang.org/p/mJz5RINaimn))  
[Sharing data II](example3/example3.go) ([Go Playground](https://play.golang.org/p/GpmPICMGMre))  
[Escape Analysis](example4/example4.go) ([Go Playground](https://play.golang.org/p/n9HijcdZ3pT))  
[Stack grow](example5/example5.go) ([Go Playground](https://play.golang.org/p/vBKF2hXvKBb))  

### Escape Analysis Flaws

[Indirect Assignment](flaws/example1/example1_test.go)  
[Indirection Execution](flaws/example2/example2_test.go)  
[Assignment Slices Maps](flaws/example3/example3_test.go)  
[Indirection Level Interfaces](flaws/example4/example4_test.go)  
[Unknown](flaws/example5/example5_test.go)  


--- --- --- --- --- --- --- --- ---

Remember, the lack of performance is gonna come from **four** places. 
1. **Latency** around networking, disk I/O, etc. 
2. Allocations and memory, garbage collection. 
3. How we access data .
4. Algorithm efficiency. 

---

Everything in Go is **pass by value**. 
- When I say pass by value, what I mean is WYSIWYG. 

---

#### Behind the Scenes - Part 1 

Remember
- **G**: Goroutine
- **P**: Logical Processor
- **M**: OS Thread
- **C**: Core

When your Go program starts up, it's gonna be given a **`P`** or a **logical processor** for every **core** (**`C`**) that's identified on the host machine. 

That **`P`** is given a _real live_ **operating system thread** that we call **`M`**.
 
That **`M`** is going be scheduled by the operating system scheduler on a particular core.

We get one more thing from the run time, and that is our **goroutine**, our **`G`**. 

---

#### Behind the Scenes - Part 2 

**Path of Execution**

**Threads** are our **path of execution** at the **operating system** level. 

All code at some point gets into **machine code**, and the operating system's job is to choose a path of execution, a thread to execute those instructions one after the other, starting from `main`. That's the job of the thread. 

---

There's **3** areas of memory that we may talk about throughout the class. 
1. The Data Segment 
    - usually reserved for global variables and read-only values
2. Stacks
3. Heaps 

We're going focus on stacks and heaps.

---

#### Behind the Scenes - Part 3

**Stacks**

A stack is a data structure that every thread is given. 

At the operating system level, your stack is a **contiguous** block of memory and usually it's allocated to be **1MG**. 

**1MG** of memory for every stack, and therefore for every thread. 

**`G`**s are very much like **`M`**s; we could almost say that they're the same but **`G`** above the operating system. 

A **`G`** has a stack of memory; **2K** in size (a lot lot smaller than **1MG**) 


#### Behind the Scenes - Part 4

By the time the goroutine that was created for this Go program wants to execute main, it's already executed a bunch of code from the run time. 

Any time a goroutine makes a function call, what it's going to do is take some memory off the stack. We call this a frame of memory. 

Remember, every line of code we write is either reading memory or writing memory. Obviously, also allocating at some point. 

The goroutine only has **direct** access to memory for the frame that it is operating on - called the **active frame**. 

The stack frame is serving a really important purpose. It's creating a sandbox, a layer of isolation. 

#### Behind the Scenes - Part 5

**Mechanics and Semantics** 
- Mechanics: how things work. 
- Semantics: how things behave. 

---

Every time you make a **function call**, what we're really doing is crossing over a **program boundary**. 

Program boundaries are important to identify, because it means that we're gonna be moving out of this sandbox or frame and we're gonna be moving into a new one. 

Every time we call a function, what's going to happen is we're going to slice a new frame off the stack. 

This is where the idea of **parameters** comes from. Parameters are really serving a mechanical purpose on top of our design purpose, and that is to get data inside of this new frame so the goroutine can operate this new data transformation in a level of isolation with a level of immutability. 

**_Everything_** in Go is **passed by value**, what that means is that we're gonna make a **copy** of the data as it goes across the boundary. 

There's two types of data that we operate with; the **value itself** or the **value's address**. Yes, addresses are data, and I want you to always understand that. 

When you pass the **data value itself** into a parameter, this is called **value semantics**. 

When you pass the **data value's address** into a parameter, this is called **pointer semantics**. 

---

**Pointer semantics** serve one purpose and that is to share a piece of data across a **program boundary**.

What happens when you pass **data value's address**: 
- Now you might think that this is what's called a pass by reference. 
- It really isn't. 
- Remember pass by value means WYSIWYG!! 
- It just so happens that the data we are copying and passing is not a value, it's an address. 

Pointer variables are not special. They aren't, they serve one purpose, and that is to store addresses. 

---

`*inc++`
- This is what we call an **indirect** memory read or write. 
- A read/modify/write operation through pointer indirection. 

![](img/indirect.png)

---

However, there's a huge cost to pointers. Remember now we've just walked away from our mutation isolation, our immutability. What we've done is set ourselves up for what we call **side effects** in our code. When we start mutating memory, we have to be very, very careful. 

---

Any memory below the **active frame** (**`AF`**) is not valid memory. 

---


#### Behind the Scenes - Part 6

We don't have **constructors** in Go. Yay, we don't want that - it hides cost. 

What we do have is what I call **factory functions**.
- Factory function is a function that creates a value, initializes it for use, and returns it back to the caller. This is great for readability, it doesn't hide cost, we can read it, and lends to simplicity in terms of construction. 

```go
type user struct {
	name  string
	email string
}

func main() {
	u1 := createUserV1()
	u2 := createUserV2()

	println("u1", &u1, "u2", u2)
}

/*
Factory function createUserV1
    - It returns a value of type user 
    - i.e. this function is using value semantics.
*/ 
//go:noinline
func createUserV1() user {
	u := user{
		name:  "Bill",
		email: "bill@ardanlabs.com",
	}

	println("V1", &u)

	return u
}

//look at the return, it's not using value semantics anymore. Version two is using pointer semantics.

/*
Factory function createUserV2
    - It returns an address/pointer 
    - i.e. this function is using pointer semantics.
*/ 
//go:noinline
func createUserV2() *user {
	u := user{
		name:  "Bill",
		email: "bill@ardanlabs.com",
	}

	println("V2", &u)

	return &u
}
```

Always ask yourself; what semantic is in play? This is critical.

When we we call `createUserV2()` above:
- The compiler is really really powerful. 
- The compiler is able to perform **static code analysis** 
- In this case what the compiler will do is perform what's called **escape analysis**. 
- This will determine whether a value gets to be placed on the **stack**, which is what we want, or it **escapes** to the **heap**.
- Our first priority is that a value stays on the stack. It's very very fast to leverage the stack. Also stacks are self-cleaning, which means that the garbage collector doesn't even get involved. 
- When a value **escapes** a stack and ends up on the heap, we call that an **allocation**. 

![](img/heap.png)

---

So memory is left alone on the stack as we go _up_ and cleaned on the way _down_, so the garbage collector doesn't have to get involved -it's self-cleaning. 

---

Anytime you mix semantics we're going to have a problem. 

So, here is a general guideline; never use **pointer semantics** during **construction** -  rather only use **value semantics** during **construction**. 

uggh:

![](img/ughh.png)

... more uggh:

![](img/ughh2.png)

---

#### Behind the Scenes - Part 7

If the compiler doesn't know the size of a value at compile time, it must immediately construct it on the heap. 

This is because the stack frames for every function are sized at compile time. Frames are not dynamic. 

What happens when you've got a Go routine that's making lots of function calls and eventually it runs out of stack space?  

We get a new stack - very unique from a programming language

It's going to do is allocate a larger stack, 25% larger than the original one, and then copy all the frames over.

But this isn't something that's going to happen all of the time. 2K is usually more than enough for our stacks, because you usually don't go more than even like 10 function calls deep. 

![](img/new-stack.png)

---

See [example5.go](example5/example5.go) for interesting example.

---

Since a value can move in memory that's on the stack, this actually creates an interesting constraint for us in Go - no stack can have a pointer to another stack. 

Imagine, we had all of these stacks all over the place, hundreds of thousands of Go routines with pointers to each other's stacks. If a stack had to grow, we would have to track every pointer that points to this stack and update it as well. You want to talk about _stop the world_ latency, that would be insane. 

---

#### Behind the Scenes - Part 8

**Deep Dive into GC**
- Not crucial for now to remember all the nitty-gritty, but it is cool stuff
- [2.3 Pointers - Part 5 (GC)](https://learning.oreilly.com/videos/ultimate-go-programming/9780135261651/9780135261651-UGP2_01_02_03_05)
- [Garbage Collection Semantics Part I](https://www.ardanlabs.com/blog/2018/12/garbage-collection-in-go-part1-semantics.html) - William Kennedy





dddddd

Now we're gonna talk about constants and constants are really a fascinating part of the language for me, because of the way they're implemented, and again remember Go is all about these costs and these trade offs so let's talk about the mechanics around constants I find them really interesting and I'll show you some of the cooler aspects of constants and how they work in the language. You'll end up using a few of them in almost every program that you write. Alright, so one of the really interesting things about constants, for me is that they only exist at compile time. Constants are something that really have a much different feel and flavor to than your common variables. Look at this right here on line 15 and 16. I am declaring two constants, but these are constants of a kind, and that's usually thing I've only seen it, seen it like this in this language. Most of the time we think about constants as being read-only variables, and it's absolutely not the case in Go. So, when we look at constants, they can come into this flavor that you see here, that is of a kind. Notice that there's no type information during the declaration of these constants on line 15 or 16. And again, it's based on the value on the right hand side of this assignment, whether they will be of kind integer or kind float. Right below there, just to show you the contrast, those are constants of a type, of type int and of type float64. Now, the big difference between constants of a kind and of a type, are that constants of a kind can be implicitly converted by the compiler. This gets really interesting, remember we talked about during the struct types, that there's no implicit conversion of that concrete data, but when we start talking about constants of a kind, that goes out the door. Now, constants of a type still hold true that once something is of a type, then that's it, the implicit conversion goes away. But constants of a kind, there we get some flexibility and it's a very powerful mechanism in Go, I'll show you from code readability. So there we go, ui, uf, two constants of a kind, kind int, kind float, and ti and tf, constants of type int and type float. Now, if we continue to talk about the idea of these kinds and what happens there, there's something very special in Go, and it's called Kind Promotion. And this Kind Promotion will tell you how things promote so floats promote over ints and types always promote over kind. So, on line 24, and I had to comment that out because once I made that constant of a type Uint8 it was bound by the laws of type and you could only put a value within the scope of a one biter 8-bit integer. But what's also really interesting about constants of a kind, because they're not type based, they're technically not limited by precision, and if you look at the specification, if you look at the specification around constants of a kind, you'll notice that the specification says that a constant has to be at least of 256 bits of precision. That makes constants of a kind almost like, or the compiler when it comes to constants of a kind like a high precision calculator. Now, if you look at on line 30 I want to show you some very interesting things here. What we're doing is we're multiplying the value of three by the value of 0.333. On the surface we see that we have an integer and we have a floating point. Remember, you can't have an implicit conversion between two different types. We could argue that these are two different types an int and a float, but they're constants, literal values in Go are constants, they're constants of a kind, they're technically unnamed constants of a kind. And so through this idea of Kind Promotion, that kind integer three will be promoted up to kind float. We now have like kind on both sides of that multiplication and var at this point is variable ends up being a variable of that float64 type. Very, very interesting, powerful stuff where we were able to work with literal values from a kind perspective, we don't have to worry about doing any explicit conversions and then this promotion takes care of making sure that everything is right. But remember that we're dealing with 256 of precision, 256 bits of precision when we're dealing with constants of a kind, and when we now convert us back to a variable where you see answer, we're moving that down to a 64 bit level of precision. There will be some precision loss, but again, remember that floating points already are already not precise, right, IEEE754 binary decimals. Look on line 33, this time we're doing division between basically a kind int1 and a kind float3.0. Again, we get the same thing, that integer promotes up to be of kind float, and we actually what we would say is have the exact representation of 1/3. Sometimes, in the old days we used to call what we considered constants of a kind to be exact, they were like these very exact numbers because we had such high levels of precision that they were exact. So, we would look at third truly as 1/3 even though, eventually it turns to 56 bits but there be some precision loss. But then on line 36 you could see that there's no promotion going on. One is of kind int, three is of kind int, we do the division, we end up with zero, because that's what's that's gonna be, everything stays within the kind integer. Now, I also told you that we can promote from kind to type and look at this on line 40. We're creating a constant of type int8 reside in the value of one, so it's still a constant, still only exists at compile time, but it's bound by the int8 type. And now this is where the power of kind and type comes in. Look on line 41. We're taking the literal value of kind int2 multiplying it by the constant of int8, and now, again, through promotion, that kind 2 value promotes up to int8, it now promotes to be a type constant. Again, we've got to have those like types across, you know, that expression, and two ends up being a constant of type int8. So these are the mechanics of constants of a kind, constants of a type, and the promotion, and I want to show you how powerful this stuff is. But before we do, I want to show you as well that this idea that constants only exist at compile time and they have these high precisions, like a minimum of 256 bits of precision is real. Now, it's hard to show you this is real when things only exist at compile time and we have to run a program at run time, but I want you to look at these constants as I've declared. This constant is called maxInt and it represents the maximum signed integer you could have for a 64 bit precision, there it is. But look on line 14, I'm storing a much larger number there, much larger than our 64 bit variables could ever hold. Now these are integers, but again, we have 256 bits of precision. If we didn't have these higher levels of precision, we really wouldn't be able to do this assignment at all. So what I'm really gonna show you is, as I run this, the compiler accepts line 14 even though this is a much larger number we could ever put in a variable. I mean look, if I take this and say let's create a constant of type int64 with the same number, this won't compile, that number is much larger than a 64 bit can hold and I've still only scratched the surface on what I could store in this constant. I mean I can keep going, I've got 256 bits of precision here. So, the compiler is really like this high precision calculator when it comes to constants. Let me show you more practical use of constants and how this kind and type promotion really makes our life a little bit better. Look at this constant tier of type Duration. This is a second way to declare a type here in Go, what I would say is, the name type Duration is based, based on int64. This is not an alias, we have really two distinct named types here. We're just using int64 as our base information or our base memory model for Duration. And I only want to do these types of things when the new type has its own representation and meaning, and it does here in the time package. Duration represents time. Doesn't represent an int64, it represents nanoseconds of time. I want you to look at these constants 'cause this is a real practical and clean way of how this idea of type and kind work together in constants. So we create a type constant named Nanosecond, which represents one nanosecond of time, it's based on a duration, and then if you go to line 16, look at what we're doing. We're taking the literal value of kind int1000, this is a constant, an unnamed constant of kind int, we're multiplying it by our type constant Nanosecond, this gets promoted to be of type Duration, and we end up with another constant of type Duration. How amazing is this? This looks simple on the surface, but the engineering behind this is really, really powerful. In fact, all of these constants end up being of type Duration, and represent the proper int64 value for these incremental units of time. And more gets into this. I mean look at this Add method. We will talk about methods later, but I want to focus on the parameter here. Notice how Add takes a value of type Duration, right. And we might do this, we might say I only want to pass Duration values into Add. Well, one of the interesting things about constants of a kind, again, is that they can be implicitly converted when there's compatibility. This is one of the first things that really messed with me when I started learning Go. I mean look at this, on line 39 I called the now function of the time package and I create a variable which gives me the current time. Now when I call Add on line 42, notice I'm passing the literal, unnamed constant of kind int minus five. You might say, whoa, whoa, whoa, whoa, whoa, that's not of type Duration, what's going on, why would the compiler accept that? This is one of those trade offs because values of a kind, right, constants of a kind, can be implicitly converted, we can pass that in, it would represent minus five nanoseconds, and this would really mess you up. This is one of the reasons why we can't have enumerations in Go. We don't want to create types as aliases to get compiler protection when they're based on let's say, those built-in types, which is where constants are allowed to be. Constants can only be based on the built-in types because again, they only exist at compile time. But look at this, this is something that we're gonna do a lot in Go, the timeout constant, right. Five is of kind int, there it is, time.Second is of type Duration, it's still a constant, but it's a constant of type Duration, that five int kind promotes up to Duration and timeout becomes a constant of type Duration, and that constant of type Duration supports the API very, very cleanly. This is really nice and great API design in leveraging constants of a type and constants of a kind. Very, very powerful stuff. The one thing you could never do is what we see here. Notice that I'm converting our integer, right, our constant minus five of kind int into a variable, or a value of type int64. I cannot pass minus five to Add because that is a value based on a name type int64 and it only wants Duration. So those constants of a kind can be implicitly converted through those calls, any value that's based on a type, we've got to have like types everywhere, everywhere. Very, very interesting and powerful stuff around constants. Now there's one last thing about constants that you'll probably want to use as you're writing production level code. This is the idea of the keyword iota, I-O-T-A. It's really interesting and it can be very, very confusing at first. If you notice that I'm using the concept of a block, we can do this with vars and types and we do it with imports all the time. Notice that we're using the parentheses here to block a bunch of, in this case, constants, so I don't have to write const again over and over again. And this iota stuff works very, very well when it comes to constants. Now, when we start a block of constants like we did here, iota starts at the value of zero. And then every time we use it inside the block, it will increment by one. The output of this right here, if I just run it, you'll see here that it's zero, one, two. Zero, one, two. We get this incrementing feature for free inside the constant block. Now, most likely, because we want the incrementing feature, we're gonna do it like this. We only have to assign iota one time to the very first constant in the block and we will automatically, for free, get the incremental. Again, on the output here, we see here is two, zero, one, two, same exact output, the thing is I didn't have to repeat myself and keep putting iota there. Now, this is something also that you'll see a lot of. Maybe we don't want to start at zero, we want to start at one. So the first constant, we take we know what the starting point for iota is zero, incremented by one initially on that first constant, and then we'll start at one. One, two, and three. And, we do this in the log package, and sometimes what you want is a set of constants that bit precision in some sort of flagging system, like we do at login. Notice this time what we're doing is we're using the iota value to shift bits to the left one and that gives us one, two, four, eight, 16, 32. You might see that a lot with constants as well. So, this iota is a very powerful mechanism if you're creating a set of constants that are gonna have some unique IDs and it just kind of let's the language set all that up for you. So constants are really, really powerful in Go. Remember now that there's two types of constants, constants of a kind and constants of a type. Your literal values in Go are constants of a kind, they're unnamed constants, constants of a kind can be implicitly converted by the compiler, which means that you can't really have enumerations in Go. You're not gonna get those compiler protections. Once a compiler is based on a type, then the full laws of type are gonna restrict its ability to be anything other than its particular precision. Remember, constants of a kind can have up to 256 bits of precision, we got really like a high precision calculator in Go. And then again, those literal values are all constants of a kind.
