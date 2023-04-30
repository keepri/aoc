use anyhow::Result;
use std::fs;

pub fn run_1() -> Result<()> {
    let input = fs::read_to_string("static/files/input_1.txt")?;

    let found = input.split("\n\n").collect::<Vec<_>>();

    let mut totals: Vec<u32> = vec![];
    for segment in found.clone().iter() {
        let calories = segment
            .split("\n")
            .into_iter()
            .map(|calorie_count| calorie_count.parse::<u32>().unwrap_or(0));

        let total: u32 = calories.sum();

        totals.push(total);
    }

    totals.sort();

    let sum: u32 = get_last_three(&totals).iter().sum();

    println!("part 1 answer: {}", totals[totals.len() - 1]);
    println!("part 2 answer: {}", sum);

    Ok(())
}

fn get_last_three(vec: &Vec<u32>) -> Vec<u32> {
    let (index_1, index_2, index_3) = (vec[vec.len() - 3], vec[vec.len() - 2], vec[vec.len() - 1]);
    vec![index_1, index_2, index_3]
}
