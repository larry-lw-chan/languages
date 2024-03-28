enum Colors {
    Red,
    Green,
    Blue,
}

enum DispenserItem {
    Empty,
    Ammo(u8),
    Thing(String, i32),
    Place { x: i32, y: i32 },
}

// fn main() {
//     let red = Colors::Red;
//     let green = Colors::Green;
//     let blue = Colors::Blue;

//     let color = Colors::Red;

//     let item = DispenserItem::Empty;
//     let ammo = DispenserItem::Ammo(5);
//     let thing = DispenserItem::Thing("Hat".to_string(), 5);

//     let place = DispenserItem::Place { x: 5, y: 10 };

//     // How to match a single variant... not super intuitive
//     if let Colors::Red = color {
//         println!("Red");
//     }

//     // The match statement is more intuitive
//     match red {
//         Colors::Red => println!("Red"),
//         _ => println!("Not red"),
//     }

//     match green {
//         Colors::Green => println!("Green"),
//         _ => println!("Not green"),
//     }

//     match blue {
//         Colors::Blue => println!("Blue"),
//         _ => println!("Not blue"),
    }

    let message = match item {
        DispenserItem::Empty => "Dispenser is empty",
        _ => "Dispenser is not empty",
    };

    println!("{}", message);

    match ammo {
        DispenserItem::Empty => println!("Dispenser is empty"),
        DispenserItem::Ammo(_) => println!("Dispenser has ammo"),
        DispenserItem::Thing(_, _) => println!("Dispenser has a thing"),
        DispenserItem::Place { .. } => println!("Dispenser is at a place"),
    }

    match thing {
        DispenserItem::Empty => println!("Dispenser is empty"),
        DispenserItem::Ammo(_) => println!("Dispenser has ammo"),
        DispenserItem::Thing(_, _) => println!("Dispenser has a thing"),
        DispenserItem::Place { .. } => println!("Dispenser is at a place"),
    }

    match place {
        DispenserItem::Empty => println!("Dispenser is empty"),
        DispenserItem::Ammo(_) => println!("Dispenser has ammo"),
        DispenserItem::Thing(_, _) => println!("Dispenser has a thing"),
        DispenserItem::Place { .. } => println!("Dispenser is at a place"),
    }
}
