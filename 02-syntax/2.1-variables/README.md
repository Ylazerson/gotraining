# B"H


## Key Notes

- **Type is Life!**

- **Type** provides us two pieces of information:
    1. size
    2. representation (what it represents)

- In general, just use `int` 
    - Its size will be based on the architecture that we are building for.
    - For example, on an AMD64 (a 64-bit architecture): 
        - our **pointer** size or the **address** size is gonna be 64 bits or eight bytes. 
        - Go will do the following: the generic **word size** will be 8 bytes (a Go **word** refers to a value that can change size depending on the architecture)  
        - In other words, Go says, if your address and your word size is 8 bytes, let's just make the integer follow. 
        - ... **mechanical sympathies** at its finest!. 

- **Zero Values** 
    - ... very very important - it's for integrity 
    - In other words, any time we allocate a byte of memory, if we don't pre-initialize it with some value, and we're gonna run electrons through the machine and make sure that we set it all to zero. 
    - This is going to help with a huge number of bugs that we found in the past ... 


- Follow this rule: if you want to declare a variable and it's gonna be set to its zero value state, use `var` approach.
    - `var` gives us zero value 100% of the time
```go
var a int
var b string
var c float64 
var d bool
```

![](img/str.png)

- **`string`**
    - And `string` is really kind of made up datatype for any language
    - Go has a unique way
    - Now, when we look at on line 16 and we say this var b string, what we're seeing is a string being declared and created setting it to its zero value which is what we call an empty string, and so an empty string in Go is going to look something like this. B is going to be a two word data structure, strings are two words. Notice I'm using the term word because the size of a string will change depending on the architecture, on the Playground the word is gonna be two, four byte words, so that's eight bytes, on our 64 bit, it'll be 16 two eight byte words, there it is, where the first word is a pointer, we're gonna talk about pointers and the next word is the number of bytes. Since this is an empty string, we don't have any pointer to a backing array yet, and there are no bytes. So a word is a two word data value, eight to 16 bytes of depending upon architecture. Again, I want to talk about cost, engineering cost, so we can predict reasonably well how things are going to run, how things are gonna happen in memory, so type is important, the size of things can be important, we really want understand the cost of the things we're taking. Now, you're probably not gonna always want to create, construct, and initialize to zero value. There are gonna be times where you wanna pre-initialize something, and this is where the short variable declaration operator comes from, that is the colon equal that you see here on line 27. A lot of people think that this syntax from the language comes from C, but actually, the Pascal programming language had a large influence also in the syntax of this language. This is one of those places where Pascal comes from. Now, this operator is really a productivity operator, it's allowing us to declare and initialize at the same time, I don't want to use the concept that that Go has duck typing, it doesn't. It's really struct-based typing, but in this particular case, I want us to always look at the short variable declaration operator as a productivity operator, so aa is 10, bb is string, cc is a float, and dd is bool, and it's based on the fact that the compiler knows what type of value is sitting on the other side of that short variable declaration, and if we go back to hello for a second, what hello would look like from a string now, is an array of five bytes, right, H-E-L-L-O. Our pointer would point to that, and we would say we have five bytes right there. We're gonna go deeper into this stuff as the class goes on, but I really believe in visuals, being able to visualize code and early on in my career, I used to keep a piece of paper at my desk with a pen and I would draw these things out. So what you're seeing me draw on the board is really my mental models, my visual mental models of code and how things go and you're gonna see a lot of that throughout this class and I really think it helps and it's something that it might help you as well. So we've got the short variable declaration operator when it's something that's gonna not be zero value. Now, you might see code like this, and I could see this. I'm using the short variable declaration operator to assign or declare an integer and set it a zero value. There's nothing necessarily wrong with this. A big part of what we are gonna learn about writing code and mental models is consistency. If you wanna do this, then be consistent. I wouldn't. I would use var for zero value, because I think the readability marker of var is just way too strong to walk away, and if I was doing a code review, I would ask us to switch that to var, but it's not for any other reason for readability and consistency and readability. So I don't want to do that, you won't see that in this code base, but that's up to you. As long as everybody on your team is doing it the same way, then it's going to be okay. Now, there's one more concept here that I want to talk about, and it's called conversion. Go doesn't have casting, it has conversion. And what conversion means really is that we may be taking a cost, a memory cost, as we convert values from one type to the other. If you've never heard of casting before, what casting has done traditionally and it's really a part of helping with performance, is saying this following thing. Let's say that we allocated a one byte integer. There it is, there is my one byte integer. Let's say for some silly crazy reason, I decide that I really want a, not to represent a one byte integer, but a four byte integer. Casting will allow me to do something like this, where I could tell the compiler, look, you know and I know that a is an int eight or one byte integer, but casting let's us pretend that what that memory really is is a four byte integer, and the compiler trusting us will just say, okay, and suddenly now, if I'm casting a from one bytes to four bytes, I have the ability to read and write memory across those four bytes from this particular location. I could be potentially corrupting a lot of memory here. Now, this is a silly example because really, where casting traditionally comes in is kind of two places. One, if you're dealing with data coming over the wire and you've got large number of bytes, you probably would like blindly copy those bytes somewhere in memory and then you would cast or overlay a structure on top of it, right? And that's gonna be very very very fast. You're just gonna say, hey, those bytes over there, those 20 bytes starting at that address location, they really represent this structure, and then that's gonna let you, because without type, you can't read and write to memory, it's gonna let us do that, but if you're off by one byte, then we've got real problems, so casting comes with the idea that if you're off, or seeing that one byte is you overlay that struct, now you are reading and writing the bytes you didn't. And that's a real real problem with casting, so Go says, look, integrity is number one. We care about integrity as our highest priority, so you can use the unsafe package to do some casting, what we'd rather have is conversion, what Go would say here is we don't wanna even set up this scenario, so if you really want a to be four bytes, then we're gonna convert a into a new value that will be four bytes and maybe, and in this particular case, even have to give it a new variable name, in this case in the sample, we used aaa. But the idea of conversion over casting is again, is an integrated play. There could be a cost of new memory but we always rather be safe than sorry. So this is what I wanna share with you in the variable section, you already know what a variable is, but to sum up here, okay, type is everything. Type is life. It gives us two pieces of information, size of the memory that we're gonna be allocating and reading and writing and what that memory represents, and without the type information, we will have chaos. It gives us the levels of integrity we need, right? We have the idea of zero value in Go, and I wanna use the keyword var for that. All memory's initialized at least to a zero value state. We can use the short variable declaration operator when we are initializing something to it other than a zero value state. There is exceptions to everything, and part of engineering is knowing when to take that exception, but these are the guidelines we're gonna be following, and we have conversion over casting, again, it's an integrity play to keep our software, again, and our data and our memory safe.


## Variables

Variables are at the heart of the language and provide the ability to read from and write to memory. In Go, access to memory is type safe. This means the compiler takes type seriously and will not allow us to use variables outside the scope of how they are declared.

## Notes

* The purpose of all programs and all parts of those programs is to transform data from one form to the other.
* Code primarily allocates, reads and writes to memory.
* Understanding type is crucial to writing good code and understanding code.
* If you don't understand the data, you don't understand the problem.
* You understand the problem better by understanding the data.
* When variables are being declared to their zero value, use the keyword var.
* When variables are being declared and initialized, use the short variable declaration operator.

## Links

[Built-In Types](http://golang.org/ref/spec#Boolean_types)    
[Variables](https://golang.org/doc/effective_go.html#variables)    
[Gustavo's IEEE-754 Brain Teaser](https://www.ardanlabs.com/blog/2013/08/gustavos-ieee-754-brain-teaser.html) - William Kennedy    
[What's in a name](https://www.youtube.com/watch?v=sFUSP8Au_PE)    
[A brief history of “type”](http://arcanesentiment.blogspot.com/2015/01/a-brief-history-of-type.html) - Arcane Sentiment    

## Code Review

[Declare and initialize variables](example1/example1.go) ([Go Playground](https://play.golang.org/p/xD_6ghgB7wm))

## Exercises

### Exercise 1 

**Part A:** Declare three variables that are initialized to their zero value and three declared with a literal value. Declare variables of type string, int and bool. Display the values of those variables.

**Part B:** Declare a new variable of type float32 and initialize the variable by converting the literal value of Pi (3.14).

[Template](exercises/template1/template1.go) ([Go Playground](https://play.golang.org/p/mQiNGaMaiAa)) | 
[Answer](exercises/exercise1/exercise1.go) ([Go Playground](https://play.golang.org/p/Ygxt9kW_WAV))
___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
