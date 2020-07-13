# B"H


- We're gonna learn why Go only has **arrays**, **slices**, and **maps**

---

### Data-Oriented Design

"Data dominates. If you've chosen the right data structures and organized things well, the algorithms will almost always be self-evident. Data structures, not algorthims, are central to programming." - Rob Pike

**Design Philosophy:**

* If you don't understand the data, you don't understand the problem.
* All problems are unique and specific to the data you are working with.
* Data transformations are at the heart of solving problems. Each function, method and work-flow must focus on implementing the specific data transformations required to solve the problems.
* If your data is changing, your problems are changing. When your problems are changing, the data transformations needs to change with it.
* Uncertainty about the data is not a license to guess but a directive to STOP and learn more.
* Solving problems you don't have, creates more problems you now do.
* If performance matters, you must have mechanical sympathy for how the hardware and operating systems work.
* Minimize, simplify and REDUCE the amount of code required to solve each problem. Do less work by not wasting effort.
* Code that can be reasoned about and **does not hide execution costs** can be better understood, debugged and performance tuned.
* Coupling data together and writing code that produces predictable access patterns to the data will be the most performant.
* Changing data layouts can yield more significant performance improvements than changing just the algorithms.
* Efficiency is obtained through algorithms but performance is obtained through data structures and layouts.


---

**Data-oriented Design** 
- Switch your brain away from object-oriented design. 
- Data-oriented design: understanding that every problem you solve is a **data** problem. - We are all data scientists at the end of the day. 
- **Integrity** comes from the data, our **performance** is going to be coming from the data, everything we do, our mental models, everything's going to be coming from the data. 
- If you don't understand the data you're working with you don't understand the problem you're trying to solve, because all problems are specific and unique to the data that you are working with, and data transformations are at the heart of everything we do. 

---

- Our problems are solved in the **concrete** data, our manipulations, our memory mutations, everything is in the concrete. 

---

- When we start looking at changes to our data/program etc. later in the course, we will start focusing on how to **decouple** the code from these data changes so these cascading changes are minimized. 
- But here's the thing, if you're abstracting in a general way, if you're building abstractions on top of abstractions, you're really walking away from all the things we've talked about so far, that the idea of optimizing for correctness and readability. 
- What we need is this balance of decoupling but thin layers of decoupling to deal with change. 
- If you are solving problems you don't have, you're now creating more problems that you do. 
- Everything we must do, everything we must do must be focused around minimizing, simplifying, and reducing the amount of code we need to solve every problem. 
- Data-oriented design. Think data, data, data, data, data. 


