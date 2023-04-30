use anyhow::Result;
use std::{char, fs, process::exit};

#[derive(Debug, Clone, Copy)]
struct Shape {
    _id: &'static str,
    mappings: [char; 2],
    power: u8,
    weakness: char,
    strength: char,
}

const ROCK: Shape = Shape {
    _id: "rock",
    power: 1,
    mappings: ['A', 'X'],
    weakness: 'B',
    strength: 'C',
};

const PAPER: Shape = Shape {
    _id: "paper",
    power: 2,
    mappings: ['B', 'Y'],
    weakness: 'C',
    strength: 'A',
};

const SCISSORS: Shape = Shape {
    _id: "scissors",
    power: 3,
    mappings: ['C', 'Z'],
    weakness: 'A',
    strength: 'B',
};

const OPTIONS: [Shape; 3] = [ROCK, PAPER, SCISSORS];

const WIN: u8 = 6;
const DRAW: u8 = 3;
const LOSS: u8 = 0;

pub fn run_2() -> Result<()> {
    let input = fs::read_to_string("static/files/input_2.txt")?;
    let lines = input.lines().into_iter();

    let mut part_one_score: u32 = 0;
    let mut part_two_score: u32 = 0;

    for line in lines {
        let (char_one, char_two) = parse_line(line);
        let opponent_shape = get_shape_by_mapping(&char_one).expect("opponent shape");

        if let Some(shape) = shape_to_play(&char_one, &char_two, 1) {
            let score_1 = play(&opponent_shape, &shape);
            part_one_score += score_1 as u32;
        };

        if let Some(shape) = shape_to_play(&char_one, &char_two, 2) {
            let score_2 = play(&opponent_shape, &shape);
            part_two_score += score_2 as u32;
        };
    }

    println!("part 1 answer: {}", part_one_score);
    println!("part 2 answer: {}", part_two_score);

    Ok(())
}

fn play(shape_1: &Shape, shape_2: &Shape) -> u8 {
    if shape_2.mappings[0].cmp(&&shape_1.weakness).is_eq() {
        // println!("{} beats {}", shape_2.id, shape_1.id);
        return WIN + shape_2.power;
    }

    if shape_1.mappings[0].cmp(&&shape_2.weakness).is_eq() {
        // println!("{} beats {}", shape_1.id, shape_2.id);
        return LOSS + shape_2.power;
    }

    // println!("both chose {}", shape_1.id);
    DRAW + shape_2.power
}

fn parse_line(line: &str) -> (char, char) {
    let inputs = line
        .split(" ")
        .map(|e| e.chars().next().expect("invalid line"))
        .collect::<Vec<char>>();

    (inputs[0], inputs[1])
}

fn shape_to_play(opponent_mapping: &char, outcome: &char, strategy: u8) -> Option<Shape> {
    match strategy {
        1 => Some(get_shape_by_mapping(&outcome)?),
        2 => match outcome {
            'Y' => Some(get_shape_by_mapping(opponent_mapping)?),
            'Z' => Some(win(opponent_mapping)?),
            'X' => Some(lose(opponent_mapping)?),
            _ => exit(1),
        },
        _ => exit(1),
    }
}

fn get_shape_by_mapping(mapping: &char) -> Option<Shape> {
    for shape in OPTIONS {
        if let Some(_) = shape
            .mappings
            .iter()
            .find(|entry| entry.cmp(&mapping).is_eq())
        {
            return Some(shape);
        }
    }

    None
}

fn win(against: &char) -> Option<Shape> {
    let against_shape: Shape = get_shape_by_mapping(against)?;
    Some(get_shape_by_mapping(&against_shape.weakness)?)
}

fn lose(against: &char) -> Option<Shape> {
    let against_shape: Shape = get_shape_by_mapping(against)?;
    Some(get_shape_by_mapping(&against_shape.strength)?)
}
