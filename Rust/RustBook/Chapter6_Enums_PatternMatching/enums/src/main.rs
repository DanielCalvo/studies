//Using an enum, each variant can have different types and amount of associated data
enum IpAddr {
    V4(u8, u8, u8, u8),
    V6(String),
}

enum Coin {
    Penny,
    Nickel,
    Dime,
    Quarter,
}

//This is similar to defining several different structs
enum Message {
    Quit, //no data associated with it
    Move {x: i32, y: i32}, //anonymous struct
    Write(String), //string
    ChangeColor(i32, i32, i32), //three i32 values
}

//You can also define methods on enums
impl Message{
    fn call(&self){
    }
}

fn main() {

    let home = IpAddr::V4(127,0,0,1);
    let loopback = IpAddr::V6(String::from("::1"));

   let m = Message::Write(String::from("hello"));
    m.call();

    //Rust doesn't have null feature!
    //But it does have an enum that can encode the concept of a value being present or absent
    //<T> is a generic parameter (type?) it's talked about later
    //<T> means the Some variant of the Option enum can hold one piece of data of any type
    enum Option<T>{
        Some(T),
        None,
    }

    let some_number = Some(5);
    let some_string = Option::Some("String!");
    let absent_number: Option<i32> = Option::None; //Compile can't infer type here, so you need to tell it explicitly

    let value = value_in_cents(Coin::Penny);
    println!("{}", value)

    //Stopped here: https://doc.rust-lang.org/book/ch06-02-match.html#patterns-that-bind-to-values

}

fn value_in_cents(coin: Coin) -> u8 {
    //Every pattern must be handled exhaustively
    //if expressions return a boolean, but a match can return anything
    match coin {
        Coin::Penny => {
            println!("Lucky penny!");
            1
        }
        Coin::Nickel => 5,
        Coin::Dime => 10,
        Coin::Quarter => 25,
    }
}


fn route(ip_kind: IpAddr){

}