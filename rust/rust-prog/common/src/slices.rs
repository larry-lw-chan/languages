use core::str;

fn slice() {
    let string = String::from("hello world");
    let hello = &string[0..5];
    let world = &string[6..];

    println!("{hello} {world}");
    println!("{string}");
}

fn first_world(s: &String) -> &str {
    let bytes = s.as_bytes();

    for (i, &item) in bytes.iter().enumerate() {
        if item == b' ' {
            return &s[0..i];
        }
    }
    &s[..]
}
