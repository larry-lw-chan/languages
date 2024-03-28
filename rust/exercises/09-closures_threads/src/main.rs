// Silence some warnings so they don't distract from the exercise.
#![allow(dead_code, unused_imports, unused_variables)]
use crossbeam::channel;
use std::thread;
use std::time::Duration;

fn pause_ms(ms: u64) {
    thread::sleep(Duration::from_millis(ms));
}
fn main() {
    let (tx, rx) = channel::unbounded();
    // Cloning a channel makes another variable connected to that end of the channel so that you can
    // send it to another thread.
    let rx2 = rx.clone();

    let handle_a = thread::spawn(move || {
        for msg in rx2 {
            println!("Thread A received: {}", msg);
        }
    });

    let handle_b = thread::spawn(move || {
        for msg in rx {
            println!("Thread B received: {}", msg);
        }
    });

    for i in 0..50 {
        tx.send(i).unwrap();
        pause_ms(20); // slight delay to make things more interesting
    }
    drop(tx); // close the channel so the receiving ends will break their loops and end

    // Join the child threads for good hygiene.
    handle_a.join().unwrap();
    handle_b.join().unwrap();

    println!("Main thread: Exiting.")
}
