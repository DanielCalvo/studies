Ownership! Rusts's most unique feature!

### There's the stack and the heap
- The stack stores value in the order it gets them, and removes the value in an oposite order. Last in, first out
    - All data stored on the stack must have a known and fixed size
- Data with an unknown size gets stored in the heap
    - The allocator has to look for a place on the heap for your data
    - The heap is slower than the stack
- When you call a function, the function and it's values get pushed into the stack 

### Ownershio
- Each value in Rust has a variable that is it's owner. There can only be one owner at a time
- When the owner goes out of scope, the value is dropped
- Assigning a value to another variable moves the ownership

### Variable scope
- Uhm...

### Memory and Allocation
- To support a mutable and growable piece of text, we need to allocate this memory on the heap
- In Rust, memory is returned once the variable that owns it goes out of scope

### Copy trait
- Rust has a `copy` trat that can be placed on types that are stored on the stack
- chars, int, floats, bools and tuples (if their elements have copy) have the copy trait

### Ownership and functions
- Passing a variable to a function will move or copy, just as an assignment does
- Returning values can also transfer ownership

### References and borrowing
- Using the & sign, you can have a function take a reference to an object instead of taking ownership of the value
- When you have a function pointing to an object it does not own, the value it points to will not be dropped when the reference goes out of scope
- Just as variables are immutable by default, so are references!
- You can only have one mutable reference to a particular piece of data in the scope
- Using scopes you can have multiple mutable data references, just not simultaneous ones

### Dangling references
- Rust catches invalid references (references to things that have been deallocated) at compile time

### Rules of Refernces
- At any given time, you can have one mutable reference, or any number of inmutable references
- References must always be valid!

### The slice type
- Hmm, the slice type does not have ownership!