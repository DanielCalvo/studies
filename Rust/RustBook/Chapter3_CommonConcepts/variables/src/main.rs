fn main() {
    //Immutable variable (default in Rust)
    let x = 5;
    println!("{}",x);

    //Constant
    const MY_CONST :u32 = 100;
    println!("{}",MY_CONST);

    //Variable shadowing
    let a = 5;
    let a = a + 1;
    println!("{}", a);

    //Shadowing the variable recreates it
    let spaces = "    ";
    let spaces = spaces.len();

    //This will not compile, you're not allowed to change the type of a variable
    // let mut spaces = " ";
    // spaces = spaces.len()

    //When a type is ambiguous, annotate it:
    let guess: u32 = "42".parse().expect("Not a number!");
    println!("{}", guess);

    //Rust has 4 scalar types: integers, floating-point numbers, booleans and characters
    let f= true;
    let c = 'a';

    //Rust has two primitive compound types: Tuples and arrays
    let tup: (i32, f64, u8) = (500, 6.4, 1);
    println!("{}", tup.0);

    //Every element of an array must have the same type
    //Arrays in rust have a fixed lenght :o
    let arr = [1,2,3,4,5];
    println!("{}", arr)
    //Arrays are useful when you want data on the stack and not on the heap (?)
    //There is a vector data type that can grow in size -- discussed on chapter 8
    //Trying to access an index on the array that doesn't exist causes a panic

}
