use std::io;

fn main() {
    let mut input = String::new();
    io::stdin().read_line(&mut input);
    println!("Hello, world!");
    println!("Number: {} String: {}", 100, "something!");
    println!("Weight on mars: {}kg ", calculate_weight_on_mars(80.0));
    let mut mars_weight = calculate_weight_on_mars(10.0);
}

fn calculate_weight_on_mars(weight: f32) -> f32 {
    weight / 9.81 * 3.711
}