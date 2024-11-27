use anyhow::Result;
use std::{collections::VecDeque, fs};

pub fn run_five() -> Result<()> {
    let input = fs::read_to_string("static/files/input_5.txt")?;
    let mut part_one_answer: String = "".into();
    let mut part_two_answer: String = "".into();
    let mut input = input.lines();
    let mut queue_arr = vec![
        VecDeque::<char>::new(),
        VecDeque::<char>::new(),
        VecDeque::<char>::new(),
        VecDeque::<char>::new(),
        VecDeque::<char>::new(),
        VecDeque::<char>::new(),
        VecDeque::<char>::new(),
        VecDeque::<char>::new(),
        VecDeque::<char>::new(),
    ];

    while let Some(line) = input.next() {
        if line.len() == 0 {
            break;
        }

        let line_chars: Vec<char> = line.chars().collect();
        for i in (1..line.len()).step_by(4) {
            let elem = line_chars[i];
            if elem as u8 != 32 && !elem.is_numeric() {
                queue_arr[i / 4].push_front(elem);
            }
        }
    }

    let mut actions: Vec<(usize, usize, usize)> = vec![];
    while let Some(line) = input.next() {
        let line_split: Vec<&str> = line.split(" ").collect();
        actions.push((
            line_split[1].parse::<usize>().unwrap(),
            line_split[3].parse::<usize>().unwrap() - 1,
            line_split[5].parse::<usize>().unwrap() - 1,
        ));
    }

    for action in actions {
        for _ in 0..action.0 {
            if let Some(elem) = queue_arr[action.1].pop_back().take() {
                queue_arr[action.2].push_back(elem);
            }
        }
    }

    for list in queue_arr {
        part_one_answer = format!("{}{}", part_one_answer, list.clone().pop_back().unwrap());
    }

    println!("part 1 answer: {}", part_one_answer);
    println!("part 2 answer: {}", part_two_answer);

    Ok(())
}
