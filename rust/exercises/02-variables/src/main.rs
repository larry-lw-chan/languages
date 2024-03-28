const STARTING_MISSILES: i8 = 8;
const READY_AMOUNT: i8 = 2;

fn main() {
    let (mut missles, ready) = (STARTING_MISSILES, READY_AMOUNT);

    println!("Firing {} of my {} missles!", ready, missles);
    missles = missles - 1;
    println!("Firing {} missles left!", missles);
}
