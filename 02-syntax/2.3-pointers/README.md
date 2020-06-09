# B"H





## Pointers

Pointers provide a way to share data across program boundaries. Having the ability to share and reference data with a pointer provides the benefit of efficiency. There is only one copy of the data and everyone can see it changing. The cost is that anyone can change the data which can cause side effects in running programs.

## Notes

* Use pointers to share data.
* Values in Go are always pass by value.
* "Value of", what's in the box. "Address of" ( **&** ), where is the box.
* The (*) operator declares a pointer variable and the "Value that the pointer points to".

## Escape Analysis

* When a value could be referenced after the function that constructs the value returns.
* When the compiler determines a value is too large to fit on the stack.
* When the compiler doesn’t know the size of a value at compile time.
* When a value is decoupled through the use of function or interface values.

## Garbage Collection History

The design of the Go GC has changed over the years:
* Go 1.0, Stop the world mark sweep collector based heavily on tcmalloc.
* Go 1.2, Precise collector, wouldn't mistake big numbers (or big strings of text) for pointers.
* Go 1.3, Fully precise tracking of all stack values.
* Go 1.4, Mark and sweep now parallel, but still stop the world.
* Go 1.5, New GC design, focusing on latency over throughput.
* Go 1.6, GC improvements, handling larger heaps with lower latency.
* Go 1.7, GC improvements, handling larger number of idle goroutines, substantial stack size fluctuation, or large package-level variables.
* Go 1.8, GC improvements, collection pauses should be significantly shorter than they were in Go 1.7, usually under 100 microseconds and often as low as 10 microseconds.
* Go 1.9, Large object allocation performance is significantly improved in applications using large (>50GB) heaps containing many large objects.
* Go 1.10, Many applications should experience significantly lower allocation latency and overall performance overhead when the garbage collector is active.

## Garbage Collection Semantics

[Garbage Collection Semantics Part I](https://www.ardanlabs.com/blog/2018/12/garbage-collection-in-go-part1-semantics.html) - William Kennedy

## Stack vs Heap

_"The stack is for data that needs to persist only for the lifetime of the function that constructs it, and is reclaimed without any cost when the function exits. The heap is for data that needs to persist after the function that constructs it exits, and is reclaimed by a sometimes costly garbage collection." - Ayan George

## Links

### Pointer Mechanics

[Pointers vs. Values](https://golang.org/doc/effective_go.html#pointers_vs_values)    
[Language Mechanics On Stacks And Pointers](https://www.ardanlabs.com/blog/2017/05/language-mechanics-on-stacks-and-pointers.html) - William Kennedy    
[Using Pointers In Go](https://www.ardanlabs.com/blog/2014/12/using-pointers-in-go.html) - William Kennedy    
[Understanding Pointers and Memory Allocation](https://www.ardanlabs.com/blog/2013/07/understanding-pointers-and-memory.html) - William Kennedy    

### Stacks

[Contiguous Stack Proposal](https://docs.google.com/document/d/1wAaf1rYoM4S4gtnPh0zOlGzWtrZFQ5suE8qr2sD8uWQ/pub)  

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

When your Go program starts up, it's gonna be given a **`P`** or a **logical processor** for every **core** that's identified on the host machine. 

That **`P`** is given a _real live_ **operating system thread** that we call **`M`**.
 
That **`M`** is going be scheduled by the operating system scheduler on a particular core.

We get one more thing from the run time, and that is our **goroutine**, our **`G`**. 

---

#### Path of Execution

**Threads** are our **path of execution** at the **operating system** level. 

All of the code you're writing at some point gets into **machine code**, and the operating system's job is to choose a path of execution, a thread to execute those instructions one after the other, starting from `main`. That's the job of the thread. 


The operating system schedules threads. But, what's important here is the data structures now. There's three areas of memory that we may talk about throughout the class. There's the data segment, there's stacks, and there are heaps. The data segment's usually reserved for your global variables, your read-only values, we don't care about that so much today. What we're gonna focus on is stacks and heaps. A stack is a data structure that every thread is given. At the operating system level, your stack is a contiguous block of memory and usually it's allocated to be one meg. One meg of memory for every stack, and therefore for every thread. If you had 10,000 threads, you can add that up. That's 10,000 megs of memory immediately being used out of the box. This concept of a stack is coming to us all the way through the hardware. The hardware wants this stuff. It really helps simplify our programming models down there, and it's really gonna help us, too, in our programming models, as we're gonna see again to understand cost and what's happening. Since an M, or an operating system thread, is a path of execution, and it has a stack, and it needs that stack in order to do its job, our G, which is our path of execution at Go's level. Gs are very much like Ms, we could almost, at this point, say that they're the same but this is above the operating system. A G has a stack of memory, too. Except in Go today, that stack is 2K in size. 2K. It was 4K for a long time. Today it's 2K. Notice how much smaller, and I mean significantly smaller, the stack is for a goroutine as it relates to an operating system. That's important, because we wanna be able to have a lot of goroutines in our program. Lots of paths of execution running. Remember, Go is focused both on integrity and minimizing the use of resources throughout the running program. This is one of those areas where we're seeing that. When our Go program starts up, we get a G. Our G is our path of execution. Its job is to execute every single instruction that we've written. The M is gonna host itself on top of the M to actually get that happening at the hardware level. What we have is this 2K stack, and this is what's important here, this is what we're gonna be talking about. By the time the goroutine that was created for this Go program wants to execute main, it's already executed a bunch of code from the run time. Eventually is says hey, I wanna execute the main function now, we're ready. This is what is going to happen. As the goroutine executes code and starts jumping between functions, this stack becomes critical to make all that happen. Eventually, the Go routine says I want to execute main, and any time a goroutine makes a function call, what it's going to do is take some memory off the stack. We call this a frame of memory. It's gonna slice a frame of memory off the stack so it can now execute the code that you see here inside of main. What I want you to understand right now is that in order to execute this data transformation, we have to be able to read and write memory. Remember, every line of code we write is either reading memory or writing memory. Obviously, also allocating at some point. It's doing those three things, and what's important here is that the goroutine only has direct access to memory. Let's think about this for a second. The goroutine only has direct access to memory for the frame that it is operating on. It wants to execute main. This is now our active frame. Here it is. The goroutine is now executing within this context, and this is, for our purposes right now, the only memory the goroutine can read and write to directly. This is it. What does that mean to us? What it means is if this data transformation has to be executed by the goroutine, and it can only operate within the scope of memory within this frame, it means all of the data that the goroutine needs to perform this data transformation has to be in here. This frame. If we look on line 10, what we see is a variable declaration of count assigned to the value of 10. We basically are now gonna be allocating our four bytes of memory right here inside this frame. It has to be inside this frame, because if it's not, the goroutine can't access it. Understand that this frame is serving a really important purpose. It's creating a sandbox, a layer of isolation. It's giving us a sense of immutability that the goroutine can only mutate or cause problems here and nowhere else in our code. This is very, very powerful constructs that we're gonna wanna leverage and it starts to allow us to talk about things like semantics. You're gonna hear me use the words mechanics and semantics. When I talk about mechanics, I'm talking about how things work. When I talk about semantics, I'm talking about how things behave. The semantics are gonna let us understand what the behavior is, the impact we're gonna have. The mechanics are gonna allow us to visualize and see how all that works. We need both. The semantics, though, are very, very important and powerful. What we're looking at so far here is that we've created a value of type Int, its value is 10, it has been allocated or created within the frame here because this is the only place the goroutine can operate in terms of memory. On line 13, we're doing two things. I wanna bring everything back to English. On line 13, what you're gonna see is that I'm asking us to display the value of count. Any time I use the variable's name I want us to think value of. What's in the box? Over here, you can see I'm using the ampersand operator. Ampersand operator isn't unique to Go. Many programming languages have used it, and it usually means the same thing. Address of, where is the box? Hey, if you got a box in memory, it's got to be somewhere. We're gonna be using hexadecimal numbers for this. We only have to worry about looking at the last four. I could pretend that this is an address F1BC. Again, we're using hexadecimal numbers, cause it's just more efficient when you have these very large numbers which are addresses. Value of, what's in the box? That's the variable and only the variable. Ampersand variable, address of, where is the box? It's gotta be located somewhere in memory. It's gonna be a memory address. There it is. I want you to look on line 16. In line 16, we're about to make another function call. I want you to think about this for a second. Every time you make a function call, what we're really doing is crossing over a program boundary. Program boundaries are important to identify, because in this case, this program boundary, this function call, means that we're gonna be moving out of this sandbox or frame and we're gonna be moving into a new one. Every time we call a function, what's going to happen is we're going to slice a new frame off the stack. This is now going to be the active frame. Our goroutine is now going to be operating within this sandbox, this level of isolation. Which once again means that we're gonna execute this code inside of increment. We're gonna perform this new data transformation, then the data the goroutine needs to perform this transformation better be inside of this frame cause this is the only place we can operate in. This is where the idea of parameters comes from. We have been using parameters our whole programming lives. We learned about API design. We've done all of these things, but I wanna show you the mechanics behind it because without the mechanics of parameters, we wouldn't be able to encapsulate at all. These parameters are really serving a mechanical purpose on top of our design purpose, and that is to get data inside of this new frame so the goroutine can operate this new data transformation in a level of isolation with a level of immutability. What we're doing on line 16 is passing the value of count across this programming boundary. Because everything in Go is passed by value, what that means is that we're gonna make a copy of the data as it goes across the boundary. You're gonna hear me start using three different terms. You're gonna hear me use the word data, and data's what we're working with. Concrete data, if you don't understand the data, you don't understand the problem. There's two types of data that we operate with. It is the value itself, like this integer value 10, or it's the value's address. Yes, addresses are data, and I want you to always understand that. In this case, what we're passing across the program boundary is the value itself, which means we're making a copy of the value, the four bytes. We're gonna throw it over this program boundary, which means that those four bytes have to end up, now, inside this frame as well. This is where your parameters come from. This parameter we're declaring inside of inc, increment, sorry. Inc int, what that is there for is to capture or hold that value we're throwing over the program boundary so the goroutine can operate this data transformation. This becomes inc. Passed by value means we make copies and we store copies. There it is. Now, the goroutine can operate within this function right here, which it is, and on line 27 what you see is a read-modify-write operation. Line 27, inc plus plus, which means that we are now mutating memory, but we're mutating it right here. What's really important is that this frame is allowing the goroutine to mutate memory without any cause of side effects throughout the program. The memory mutation is happening in isolation within our sandbox with a level of mutability in the sense that we don't affect anything else outside of this execution context. This is a very, very important and powerful mechanism. What we're really seeing here is what is called values semantics. We're gonna focus a lot in this class on the value and pointer semantics behavior that the language gives you. If you wanna write code in Go that is optimized for correctness, that you can read and understand the impact of things, then your value and your pointer semantics are everything. What we're looking at here is an example of value semantics. There's a value here in main for count. Now, we operate a different transformation around that data, and that means that every data transformation, every piece of code that's operating on it gets its own copy of the value. This is value semantics. It's very, very important. Value semantics has the benefit of again this idea of isolation and mutability. It's going to also, potentially, in many cases give us some performance. We're gonna talk about that. But, I told you engineering is not about just hacking code, it's about knowing the cost of things. Everything has a cost, nothing is free. So the question now is what is the cost of value semantics? One of the costs of value semantics is that we have multiple copies of the data throughout the program. There is no efficiency with the value semantics, and sometimes it can be very complicated to get a piece of data that's changed and to get that updated everywhere it needs to be. Our value semantics are very powerful semantics because it's gonna reduce things like side effects. We'll talk about that. It's giving us isolation, it's giving us these levels of immutability that are so important towards integrity. Sometimes, the inefficiency of value semantics might cause more complexity in the code. It might even cause some performance problems, and performance does matter. One of the things we've gotta learn is how to balance our value and our pointer semantics. Right now, all we're looking at here is the value semantics. When the increment function returns and we're back up in main, now what's going to happen is this goroutine is no longer operating in this active frame. This is now our active frame. The goroutine is now operating here, and if I were to run this program, what we're gonna see is that mutation that we made got isolated here, did not carry up or forward. Now, this piece of code is still operating on its own copy. This was operating on its own copy. We got the benefit of this mutation not affecting anything else. This is value semantics and this is a big part, again, of the beauty of the pass by value. What we see is what you get.