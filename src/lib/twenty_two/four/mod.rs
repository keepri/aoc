use anyhow::Result;
use std::{fs, process};

pub fn run_four() -> Result<()> {
    let input = fs::read_to_string("static/files/input_4.txt")?;
    let mut input = input.lines();
    let mut part_one_score: i32 = 0;
    let mut part_two_score: i32 = 0;

    while let Some(section_id_range_pair) = input.next() {
        let section_id_range_pair = section_id_range_pair.trim();

        match section_id_range_pair.split_once(",") {
            Some((range_one, range_two)) => {
                match (range_one.split_once("-"), range_two.split_once("-")) {
                    (Some((start_one, end_one)), Some((start_two, end_two))) => {
                        let start_one = start_one.parse::<i32>()?;
                        let end_one = end_one.parse::<i32>()?;
                        let start_two = start_two.parse::<i32>()?;
                        let end_two = end_two.parse::<i32>()?;

                        if !(start_one < start_two || end_one > end_two)
                            || !(start_two < start_one || end_two > end_one)
                        {
                            part_one_score += 1;
                        }

                        if !(end_one < start_two || start_one > end_two) {
                            part_two_score += 1;
                        }
                    }
                    _ => {
                        eprintln!("can't split entry {} or {}", range_one, range_two);
                        process::exit(1);
                    }
                }
            }
            _ => {
                eprintln!("can't split entry {}", section_id_range_pair);
                process::exit(1);
            }
        };
    }

    println!("part 1 answer: {}", part_one_score);
    println!("part 2 answer: {}", part_two_score);

    Ok(())
}
