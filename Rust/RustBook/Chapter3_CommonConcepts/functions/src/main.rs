fn main() {
    hello_world();

    let a = 10;
    print_int(a);

    //Statement
    let b = 3;

    //Statements do not return values
    //let x = (let asd = 123); <- Does not compile

    //Expressions evaluate to something
    //Expressions do not include ending semicolons
    let y = {
        let bbb = 4;
        bbb+2+51; //warns that this is unused
        567 //returns this
    };
    print_int(y);

    print_int(multiply_by_two(42))
}

//Rust doesn't care if you define this below or above main
fn hello_world(){
    println!("Hello, world!");
}

fn print_int(x: i32){
    println!("{}", x);
}

fn multiply_by_two (x: i32) -> i32 {
    x * 2 //This is an expression. Adding a semicolon would change it to a statement
}