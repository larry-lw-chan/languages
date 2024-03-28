#[derive(Debug)]
struct Rectangle {
    width: u32,
    height: u32,
}

impl Rectangle {
    fn new(width: u32, height: u32) -> Rectangle {
        Rectangle { width, height }
    }

    fn area(&self) -> u32 {
        self.width * self.height
    }
}

fn main() {
    let rectangle = Rectangle::new(30, 50);
    println!("rectangle is {:?}", rectangle);
    println!("rectangle is {:#?}", rectangle);
    println!("rectangle area is {}", rectangle.area());
}
