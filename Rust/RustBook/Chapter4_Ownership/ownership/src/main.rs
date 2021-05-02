fn main() {
    let mut s = String::from("hello");
    s.push_str("asd");
    println!("{}",s);

    //Does what you expect, they're both on the stack and have a known size at compile time
    //There is no difference between deep and shallow clonning here
    let x = 5;
    let y = x;

    //The pointer for the string on the heap is copied, but not the data for the string in the heap
    let s1 = String::from("hello");
    let s2 = s1;
    //println!("{}, world!", s1); //Does not work, as s1 is an invalid reference after being shallow clonned by s2

    //Actually clones the string
    let s1 = String::from("hellooo");
    let s2 = s1.clone(); //Deep clone, not done by default in Rust!

    println!("{}",s2);

    //Integers have a known size at compile time and are stored on the stack!
    let x = 5;
    let y = x;
    println!("x = {}, y = {}", x, y);


    let mut mystr1 = String::from("asasdasd");
    change(&mut mystr1);
    println!("{}", mystr1);


    //Borrowing an inmutable value multiple times does not seem to be a problem
    let s = String::from("aaaa");
    let r1 = &s;
    let r2 = &s;
    println!("{}, {}", r1, r2);

    //Rust does not allow you to borrow a mutable value multiple times
    let mut ss = String::from("aaaa");
    //let r3 = &mut ss;
    //let r4 = &mut ss;
    //println!("{}, {}", r3, r4);

    let mut mystr2 = String::from("to meme is to dream");
    let word = first_word(&mystr2);
    let mut otherword= first_word(&mystr2);
    println!("First word: {}", word);
    println!("First word: {}", otherword);
}

fn first_word(s: &String) -> &str {
    let bytes = s.as_bytes();
    for (i, &item) in bytes.iter().enumerate(){
        println!("{} {}", i, item);
        if item == b' '{
            return &s[0..i];
        }
    }
    &s
}

fn calculate_length(s: &mut String) -> usize {
    s.len()
}

fn change(s: &mut String){
    s.push_str(" changed!")
}