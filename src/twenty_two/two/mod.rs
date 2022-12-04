use anyhow::Result;
use std::fs;

pub mod answer {
    use std::{collections::HashMap, process::exit};

    use super::*;

    #[derive(Debug, Clone, Copy)]
    struct Shape {
        id: &'static str,
        mapping: char,
        power: u8,
        weakness: char,
        strength: char,
    }

    const ROCK: Shape = Shape {
        id: "rock",
        power: 1,
        mapping: 'A',
        weakness: 'B',
        strength: 'C',
    };

    const PAPER: Shape = Shape {
        id: "paper",
        power: 2,
        mapping: 'B',
        weakness: 'C',
        strength: 'A',
    };

    const SCISSORS: Shape = Shape {
        id: "scissors",
        power: 3,
        mapping: 'C',
        weakness: 'A',
        strength: 'B',
    };

    const OPTIONS: [Shape; 3] = [ROCK, PAPER, SCISSORS];

    const WIN: u8 = 6;
    const DRAW: u8 = 3;
    const LOSS: u8 = 0;

    pub fn run_2() -> Result<()> {
        let input = fs::read_to_string("./src/lib/two/input.txt")?;
        let lines = input.lines().into_iter();

        let mut total_score: u32 = 0;
        for line in lines {
            let (shape_1, shape_2) = parse(line);
            let score = [play(&shape_1, &shape_2)].map(Into::<u32>::into)[0];
            total_score += score;
        }

        println!("final score {total_score}");

        Ok(())
    }

    fn play(shape_1: &Shape, shape_2: &Shape) -> u8 {
        let shape_2_power = shape_2.power;

        if shape_2.mapping.cmp(&shape_1.weakness).is_eq() {
            // println!("{} beats {}", shape_2.id, shape_1.id);
            return WIN + shape_2_power;
        }

        if shape_1.mapping.cmp(&shape_2.weakness).is_eq() {
            // println!("{} beats {}", shape_1.id, shape_2.id);
            return LOSS + shape_2_power;
        }

        // println!("both chose {}", shape_1.id);
        DRAW + shape_2_power
    }

    fn parse(line: &str) -> (Shape, Shape) {
        let inputs = line
            .split(" ")
            .map(|e| e.chars().next().expect("invalid line"))
            .collect::<Vec<_>>();

        let opponent_shape = inputs[0];
        let outcome_secret = inputs[1];
        let (shape_1, shape_2) =
            identify((opponent_shape, outcome_secret)).expect("identifying shape");

        (shape_1, shape_2)
    }

    // continue here: add param whole line so we can have pairs
    fn identify((opponent_move, outcome): (char, char)) -> Option<(Shape, Shape)> {
        let indications: HashMap<char, i8> = HashMap::from([('X', -1), ('Y', 0), ('Z', 1)]);

        let opponent_shape: Shape = get_shape_by_mapping(&opponent_move)?;

        let outcome = indications.get(&outcome)?;
        let outcome_shape = match outcome {
            0 => opponent_shape,
            1 => win(&opponent_shape)?,
            -1 => lose(&opponent_shape)?,
            _ => exit(1),
        };

        Some((opponent_shape, outcome_shape))
    }

    fn get_shape_by_mapping(mapping: &char) -> Option<Shape> {
        for option in OPTIONS {
            if option.mapping.cmp(mapping).is_eq() {
                return Some(option);
            }
        }

        None
    }

    fn win(against: &Shape) -> Option<Shape> {
        let shape: Shape = get_shape_by_mapping(&against.weakness)?;
        Some(shape)
    }

    fn lose(against: &Shape) -> Option<Shape> {
        let shape: Shape = get_shape_by_mapping(&against.strength)?;
        Some(shape)
    }
}
