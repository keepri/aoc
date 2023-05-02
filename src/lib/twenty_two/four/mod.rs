use anyhow::Result;
use std::fs;

pub fn run_four() -> Result<()> {
    let input = fs::read_to_string("static/files/input_4.txt")?;
    let mut input = input.lines();
    let mut part_one_score: i32 = 0;
    let mut part_two_score: i32 = 0;

    while let Some(pair) = input.next() {
        let pair = pair.trim();
    }

    println!("part 1 answer: {}", part_one_score);
    println!("part 2 answer: {}", part_two_score);

    Ok(())
}
