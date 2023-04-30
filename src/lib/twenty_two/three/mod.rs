use anyhow::Result;
use std::fs;

pub fn run_3() -> Result<()> {
    let rucksacks = fs::read_to_string("static/files/input_3.txt")?;
    let rucksacks = rucksacks.lines().collect::<Vec<_>>();

    let mut part_one_score: u32 = 0;
    let part_two_score: u32 = 0;

    for list in rucksacks {
        let (compartment_one, compartment_two) = list.split_at(list.len() / 2);
        let compartment_one_items = compartment_one.trim().split("").collect::<Vec<_>>();
        let compartment_two_items = compartment_two.trim().split("").collect::<Vec<_>>();

        let common_item =
            find_common_item([compartment_one_items, compartment_two_items]).expect("common item");
        let common_item = common_item.chars().next().expect("common item to chars");

        let score = score_item(&common_item).expect("scoring item");
        let score = score as u32;
        part_one_score += score;
    }

    println!("part 1 answer: {}", part_one_score);
    println!("part 2 answer: {} - not done", part_two_score);

    Ok(())
}

fn find_common_item(piles: [Vec<&str>; 2]) -> Option<&str> {
    let pile1 = &piles[0];
    let pile2 = &piles[1];

    for item1 in pile1 {
        if item1.len() == 0 {
            continue;
        }

        for item2 in pile2 {
            if item2.len() == 0 {
                continue;
            }

            if item1.cmp(item2).is_eq() {
                return Some(item1);
            }
        }
    }

    None
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
