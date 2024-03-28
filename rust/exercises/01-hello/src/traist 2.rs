struct Duck {
    name: String,
}

struct Horse {
    name: String,
}

struct Pegasus {
    name: String,
}

trait Fly {
    fn fly(&self) -> &str;
}

trait Explode {
    fn explode(&self) -> &str;
}

trait Ride {
    fn ride(&self) -> &str;
}

trait Duckable: Fly + Explode {}

impl Explode for Duck {
    fn explode(&self) -> &str {
        "The duck is Boom"
    }
}

impl Fly for Duck {
    fn fly(&self) -> &str {
        "The duck is flying"
    }
}

impl Fly for Pegasus {
    fn fly(&self) -> &str {
        "The duck is flying"
    }
}

impl Ride for Pegasus {
    fn ride(&self) -> &str {
        "The pegasus is riding"
    }
}

impl Ride for Horse {
    fn ride(&self) -> &str {
        "The horse is riding"
    }
}

impl Explode for Horse {
    fn explode(&self) -> &str {
        "The horse is exploding"
    }
}

// Trait Default Method
trait Run {
    fn run(&self) {
        println!("I am running")
    }
}
struct Robot {}
impl Run for Robot {}

// fn main() {
//     let duck = Duck {
//         name: "Donald".to_string(),
//     };
//     let horse = Horse {
//         name: "Bucephalus".to_string(),
//     };
//     let pegasus = Pegasus {
//         name: "Peggy".to_string(),
//     };

//     println!("A duck named {}", duck.name,);
//     println!("A house named {}", horse.name,);
//     println!("A pegasus named {}", pegasus.name,);

//     let hello = "hello".to_string();
//     let world = "world";
//     let hello_world = hello + world;
//     println!("{}", hello_world);

//     let hello2 = "hello".to_string();
//     let world2 = "world".to_string();
//     let hello_world2 = hello2 + &world2;
//     println!("{}", hello_world2);

//     let robot = Robot {};
//     robot.run();
// }
