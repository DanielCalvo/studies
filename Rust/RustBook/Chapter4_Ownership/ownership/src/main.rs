fn main() {
    let s = String::from("hello");
    println!("{}",s);

    let s1 = String::from("hellooo");
    let s2 = s1.clone(); //Deep clone, not done by default in Rust!

    println!("{}",s2);

    //Integers have a known size at compile time and are stored on the stack!
    let x = 5;
    let y = x;

    println!("x = {}, y = {}", x, y);

}
