# B"H

## Constants

- **Constants** are a way to create a _named identifier_ whose value can never change. 
- They also provide an incredible amount of flexibility to the language. 
- The way constants are implemented in Go is very unique.

## Notes

* Constants are **not** variables.
* They exist only at **compilation**.
* Untyped constants can be _**implicitly converted**_ where typed constants and variables can't.
* Think of untyped constants as having a **Kind**, not a **Type**.
* See the power of constants and their use in the standard library.


### Details


- Constants are really a fascinating part of the language.
- You'll end up using a few of them in almost every program that you write. 
- They only exist at compile time. 
- Much different feel and flavor than common variables. 
- The big difference between _constants of a kind_ and _of a type_, are that constants of a kind can be implicitly converted by the compiler. 
- **Kind Promotion**: 
    - Tells you how things promote, such as floats promote over ints and types always promote over kind. 
- A constant has to be at least of **256 bits** of precision (like a high precision calculator). 
- Remember now that there's **two** types of constants, constants of a kind and constants of a type. Literal values in Go are **constants of a kind**, they're **unnamed constants**, constants of a kind can be implicitly converted by the compiler. 