use rand::Rng;
use std::cmp::Ordering;
use std::io; //Load the IO library from the standard library so we can handle input/output

fn main() {
    println!("Guess the number!");
    println!("Input your guess");

    //Let creates a variable
    //Mut signals that we want this variable to be mutable. Variables are inmutable by default in Rust
    //Assign an empty string to this variable
    //New is an associated function of the the string type

    let secret_number = rand::thread_rng().gen_range(1, 101);

    loop {
        let mut guess = String::new();

        io::stdin()
            .read_line(&mut guess) //This argument is a reference!
            .expect("Failed to read line"); //This is related to error handling. You'll learn more about error handling later!

        let guess: u32 = match guess.trim().parse() {
            Ok(num) => num,
            Err(_) => continue,
        };

        println!("You guessed {}", guess);

        match guess.cmp(&secret_number) {
            Ordering::Less => println!("Too small!"),
            Ordering::Greater => println!("Too big!"),
            Ordering::Equal => {
                println!("You win!");
                break;
            }
        }
    }
}
