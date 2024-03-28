struct RedFox {
    enemy: bool,
    life: u8,
}

trait Noisy {
    fn get_noise(&self) -> &str;
}

impl RedFox {
    fn new() -> Self {
        Self {
            enemy: false,
            life: 100,
        }
    }
}

impl Noisy for RedFox {
    fn get_noise(&self) -> &str {
        "Ring-ding-ding-ding-dingeringeding!"
    }
}

fn print_noise<T: Noisy>(item: T) {
    println!("{}", item.get_noise());
}

impl Noisy for u8 {
    fn get_noise(&self) -> &str {
        "BYTE!"
    }
}
// fn main() {
//     let fox = RedFox::new();
//     println!("The fox is an enemy: {}", fox.enemy);
//     println!("The fox life is: {}", fox.life);

//     let test: u8 = 5;
//     print_noise(test);
// }
