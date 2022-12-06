use anyhow::Result;
use std::{fs, process::exit};

pub fn run_3() -> Result<()> {
    let rucksacks = fs::read_to_string("static/files/input_3.txt")?;
    let rucksacks = rucksacks.lines().collect::<Vec<_>>();

    let mut final_score: u32 = 0;
    for item_list in rucksacks {
        let (compartment1, compartment2) = item_list.split_at(item_list.len() / 2);
        let compartment1_items = compartment1.trim().split("").collect::<Vec<_>>();
        let compartment2_items = compartment2.trim().split("").collect::<Vec<_>>();

        let common_item = find_common_item(&compartment1_items, &compartment2_items);
        let common_item = match common_item.chars().next() {
            Some(c) => c,
            _ => {
                eprintln!("common item to chars");
                exit(1)
            }
        };

        let score = match score_item(&common_item) {
            Some(s) => s,
            _ => {
                eprintln!("error scoring item");
                exit(1)
            }
        };

        let score = score as u32;
        final_score += score;
        // println!("common item {common_item} with score {score}");
    }

    println!("final score is {final_score}");

    Ok(())
}

fn find_common_item<'a, 'b>(pile1: &Vec<&'a str>, pile2: &Vec<&'b str>) -> &'a str {
    for item1 in pile1 {
        if item1.len() == 0 {
            // println!("skipped empty string");
            continue;
        }

        for item2 in pile2 {
            if item2.len() == 0 {
                // println!("skipped empty string");
                continue;
            }

            // println!("checking {item1} against {item2}");
            if item1.cmp(item2).is_eq() {
                return item1;
            }
        }
    }

    panic!("could not find common item");
}

fn score_item(item: &char) -> Option<u8> {
    let code = *item as u8;

    // uppercase
    if code.cmp(&64).is_gt() & code.cmp(&91).is_lt() {
        return Some(code - 38);
    }

    // lowercase
    if code.cmp(&96).is_gt() & code.cmp(&123).is_lt() {
        return Some(code - 96);
    }

    None
}
