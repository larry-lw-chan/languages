use std::thread;

fn main() {
    let handle = thread::spawn(move || "Hello from a thread!");

    // Do stuff in main thread

    // wait until thread has exited
    handle.join().unwrap();
}
