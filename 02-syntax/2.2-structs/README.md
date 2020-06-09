# B"H



### Quotes

_"Implicit conversion of types is the Halloween special of coding. Whoever thought of them deserves their own special hell." - Martin Thompson_

---

### Notes on following snippet.

```go
type example struct {
	flag    bool
	counter int16
	pi      float32
}

func main() {

	var e1 example

	fmt.Printf("%+v\n", e1)

	// Declare a variable of type example and init using
	// a struct literal.
	e2 := example{
		flag:    true,
		counter: 10,
		pi:      3.141592,
	}

	// Display the field values.
	fmt.Println("Flag", e2.flag)
	fmt.Println("Counter", e2.counter)
	fmt.Println("Pi", e2.pi)
}
```

See we're creating a **_value_** **of type** `example`, named `e1`, setting it to its zero value. 

**Note**, you won't hear me use the word **object** in Go. We create **values** in Go. 

Always remember we optimize for **correctness** first.

---

That being said, how much memory is allocated for `e1`?
- You might guess seven - not a bad guess - but wrong.
    
![](img/size-incorrect.png)

- Wrong because we have the concept of **alignments**. And they come from the hardware. Alignments are there to make reading and writing memory as efficient as possible. 

- The hardware has the ability, let's say within one instruction, to read and write one **word boundary**. 

- So, the idea around alignments is, there's no reason to allow a value to cross over word boundaries if it can fit in one. 

- Now, what that means is, is that we're going to end up with this other concept of **padding** inside of our structs. 

- Well, when you get to a 2-byte value, like `counter`, we could cross over these boundaries. So what we need to do is, make sure that any 2-byte value always aligns within a single word boundary. How do you do that? Well, you make sure that that value always falls within a 2-byte address scheme. In other words, if we look at the last digit of any address, it always falls on a multiple of two, in terms of its address. It falls within, address zero, address two, address four, address six, address eight, you get it. 

- And if we have a 4-byte value, then it has to fall on a 4-byte alignment. Address zero, address four, address eight, etc. 

- And if it's an 8-byte value, then it's gotta fall in that full boundary address zero, address eight. 

- So, what we're doing, is looking at the size of a value to determine what its alignment is, and then making sure that it falls properly in memory within those boundaries. 

- Now we can how `e1`is an 8-byte value (byte 2 is used for padding.)

![](img/size-correct.png)
    

- A struct must properly align based on its largest field.
    
- The only way padding could be a problem initially, is maybe we've got so much padding in the struct that it's allocating much more memory than we want it to. And therefore, maybe we've got a larger footprint in memory. 
    
- Now, if we truly want to micro-optimize the padding away, what we need to do is, order fields from largest to smallest. 

- But, of course, initially let's only group things together that belong together. Instead of going right into this idea of optimizing for performance.
    
---

- Any time you see a compiler message about a **literal value**, or **literal type**, what we're really talking about is an **unnamed type**.

- An anonymous struct:
```go
// Declare a variable of an anonymous type set to its zero value.
var e1 struct {
    flag    bool
    counter int16
    pi      float32
}
```

- Example 2
```go
// Declare a variable of an anonymous type and init
// using a struct literal.
e2 := struct {
    flag    bool
    counter int16
    pi      float32
}{
    flag:    true,
    counter: 10,
    pi:      3.141592,
}
```

- These **literal types** can come in really handy
- For example, you have to do some unmarshaling, you need some type information, but it's not necessarily good to name it. Because this only need to be used in this one place. And we don't want to name something we're only using in one place. That would be pollution. 

---

- When a **type** is **named**, there's **NOT** going to be **implicit conversion**. 
- But a value, isn't named, it's a **literal type**, now we have this flexibility on assignment to a named type.
 
 