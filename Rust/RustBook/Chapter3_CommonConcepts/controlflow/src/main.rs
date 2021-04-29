fn main() {
    println!("Hello, world!");

    let number = 5;

    if number < 10 {
        println!("condition was true!");
    } else {
        println!("condition was false");
    }

    let condition = true;

    //Rust need to know at compile time what type a variable is, so havint this var be an int else a string would not compile
    let num = if condition {
        5
    } else {
        6
    };
    println!("{}", num);

    // Sample loop
    let mut counter = 0;
    loop {
        println!("yay!");
        counter = counter+1;
        if counter > 5 {
            break
        }
    }

    //Returning values from a loop
    let result = loop {
        break 2 //Pass whatever you want to return from the loop to the break expression
    };
    println!("{}", result);

    let mut numm = 0;

    while numm < 10 {
        numm = numm+1;
    }
    println!("{}", numm);


    let a = [4,5,6,7,8,9];

    for element in a.iter() {
        println!("{}", element);
    }

    for number in (1..10) {
      print!("{}", number);
    }




}
