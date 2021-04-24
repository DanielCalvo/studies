## Section 1: Getting started

### 2. What is rust
- Memory safe systems language!
- Rust is a modern systems programming language
- Memory safety helps a lot to reduce vulnerabilities!
- Rust has no garbage collector
- No null types or pointers
- No exceptions in rust.
- Modern package manager! cargo is similar to npm
- No data races!

### 3. Installing rust
- https://doc.rust-lang.org/book/
- Just follow the docs. I'm trying to use apt-get, let's see if it works out...

### 4. Setting up the development environments
- https://www.rust-lang.org/tools
- Ended up installing the rust plugin on pycharm, let's see how that goes.

### 5. Cargo
- `cargo new example` <- Creates a project named example
- `cargo new --help`
- cargo.toml contains dependencies for the project, as well as metadata and compiler settings
- `cargo new` also creates a new git repo, interesting
- https://crates.io/ <- Package registry
- `cargo install cargo-expand`
- There are 3 release channels for rust: Stable, beta, nightly
- `rustup toolchain list`

## Section 2: Manual Memory Management
- Code for this course is at https://github.com/gavadinov/Learn-Rust-by-Building-Real-Applications under "memory management"

### 7. Introduction
- The stack, the heap, pointers and smart pointers

### 8. The stack
- What is the stack?
    - The stack is a special region of the memory that stores variables created by each function
- For every function called, a new stack frame is allocated oh top of the current one. Only the function that created the stack frame has access to it!
- The size of every variable on the stack has to be known at compile time
- When a function exits, it's stack frame is released
- Each function has a limited stack space. If you reach the end of the stack, you get a stack overflow error, ha!

### 9. The heap
- The heap is a region of the memory that is not automatically managed!
- You have to manually alocate memory there and manually free memory there
- It has no size restrictions. Only limited by the limits of the system
- Accessible by any function, anywhere in the program
- Heap allocations are expensive and should be avoided when possible
- On the stack you only store the address for the memory on the heap that has the value you want
- You always need to deallocate manually the memory that you have allocated on the heap!

### 10. Smart pointers
- A smart pointer is just a wrapper to a raw pointer adding additional capabilities to it
- There are many types of smart pointers but the main one makes sure to free the memory it points to when this memory goes out of scope

### 11. Explore the memory layout in GDB
- That was really cool! GDB was used to inspect the stack and heap as a small program ran. I really should try out more things with GDB...

## Section 3: Building a command line application

### 14. Basic data types
- Four basic data types: Booleans, characters, integers and floats
- Integer types: u8, i8, u16, i16, ... same with 32, 64 and 128
- u: unsigned, i: signed integer
- unsgined: can only be positive
- signed: can be positive or negative. Number indicates the size of the integer. 8 = 8 bit integer
- usize and isize: Architecture dependent integers. 32 bit on a 32 bit arch, 64bit on a 64bit arch
- floating points: f32, f64
- bool: true or false, one byte in size
- char: holds a single unicode value. It is always 4 bytes in size

### 15. Functions
- Yay let's finally write some code!
- `cargo new mars_calc`
- `cargo run` <- Yay got a hello world!

### 16. Macros
- Every time you seem something with an exclamation mark it is a macro and not a function call!
- Like `println!`
- Macros are used for meta programming, this means code that writes more code (?)
- A macro can be called with a variable number of parameters and different types
- Downside: Macro definitions are more complex than functions
- Println is a macro since it can receive a variable number or arguments 
- To expand macros: 
    - `rustup toolchain install nightly`
    - `cargo expand` 

### 17. Mutability
- All variables in rust are immutable by default :o
- Once you assign a value to them, they can never change
- You have to explicitly declare a variable as mutable in order to be able to mutate it
- Keyword: `mut`

### 18. The standard library
- The standard library is an external crate, available to all projects by default
- had to pass &mut myvar to the io::stdin read line function, why?

### 19. Ownership rules in Rust
1. Each value in rust is owned by a variable
2. When the owner goes out of scope, the variable will be deallocated
3. There can only be one owner at a time
- This lecture was really heavy on the theory. Rewatch it later or read the docs if you have to

### 20. References and borrowing
- References allow you infer to a value without taking ownership of it
- Add an ampersand before the type!
- In rust: Passing refernces as parameters == borrowing
- There also are mutable references!
- You cannot have a mutable borrow and an immutable borrow at the same time!
- Also heavy on the theory. Maybe read about ownership, references and borrowing on the official docs!

### 21. Explore ownership and borrowing in GDB
- A reference is just a pointer
- You really need to play with the debugger later!

### 22. Finishing touches
- Rust has a result type... that functions seem to return. More on this later!