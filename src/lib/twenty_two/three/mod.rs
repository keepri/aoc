use anyhow::Result;
use std::fs;

pub fn run_3() -> Result<()> {
    let rucksacks = fs::read_to_string("static/files/input_3.txt")?;
    let rucksacks = rucksacks.lines().collect::<Vec<_>>();

    let mut part_one_score: u32 = 0;
    let part_two_score: u32 = 0;

    for item_list in rucksacks {
        let item_list = item_list.trim();
        let (compartment_one, compartment_two) = item_list.split_at(item_list.len() / 2);
        let pile_one = compartment_one.chars().collect::<Vec<_>>();
        let pile_two = compartment_two.chars().collect::<Vec<_>>();

        match find_common_item(&pile_one, &pile_two) {
            Some(common_item) => {
                let score = score_item(&common_item).expect("scoring item");
                part_one_score += score as u32;
            }
            None => println!("could not find common item for {}", item_list),
        }
    }

    println!("part 1 answer: {}", part_one_score);
    println!("part 2 answer: {} - not done", part_two_score);

    Ok(())
}

fn find_common_item<'a, 'b>(pile_one: &'a Vec<char>, pile_two: &'b Vec<char>) -> Option<&'a char> {
    for char_one in pile_one {
        for char_two in pile_two {
            if char_one == char_two {
                return Some(&char_one);
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
