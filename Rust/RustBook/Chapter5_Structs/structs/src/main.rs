
//You can scope user inside the main function or outside of it... as a global struct apparently?
struct User {
    username: String,
    email: String,
    sign_in_count: u64,
    active: bool,
}

fn main() {
    println!("Hello, world!");


    //Looks like you have to populate all fields or it won't compile!
    let mut user1 = User {
      username: String::from("John Doe"),
        email: "".to_string(),
        sign_in_count: 0,
        active: false
    };

    //The entire instance has to be mutable, you can't have a single field as mutable
    user1.sign_in_count = 1;

    println!("{}", user1.sign_in_count)
}
