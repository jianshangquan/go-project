use std::time::{Instant, Duration};

fn main() {
    let start = Instant::now();
    let mut count = 0;

    while start.elapsed() < Duration::from_secs(10) {
        count += 1;
        // println!("Rust: looped {}", count);
    }

    println!("Looped {} times in 1 minute.", count);
}
