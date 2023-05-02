use anyhow::Result;
use std::fs;

pub fn run_three() -> Result<()> {
    let rucksacks = fs::read_to_string("static/files/input_3.txt")?;
    let mut rucksacks = rucksacks.lines();
    let mut part_one_score: i32 = 0;
    let mut part_two_score: i32 = 0;
    let mut cache: Vec<&str> = vec![];

    while let Some(item_list) = rucksacks.next() {
        let item_list = item_list.trim();
        let (pile_one, pile_two) = item_list.split_at(item_list.len() / 2);

        if let Some(common_one) = find_common_item(pile_one, pile_two, None) {
            let score = score_item(&common_one).expect("scoring item");
            part_one_score += score as i32;
        }

        cache.push(item_list);
        if cache.len() == 3 {
            if let Some(common_two) = find_common_item(cache[0], cache[1], Some(cache[2])) {
                let score = score_item(&common_two).expect("scoring item");
                part_two_score += score as i32;
            }

            cache = vec![];
        }
    }

    println!("part 1 answer: {}", part_one_score);
    println!("part 2 answer: {}", part_two_score);

    Ok(())
}

fn find_common_item(pile_one: &str, pile_two: &str, pile_three: Option<&str>) -> Option<char> {
    for char_one in pile_one.chars() {
        if pile_two.contains(char_one) {
            if let Some(pile_three) = pile_three {
                if pile_three.contains(char_one) {
                    return Some(char_one);
                }
            } else {
                return Some(char_one);
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
