#![allow(special_module_name)]

mod lib;

use crate::lib::twenty_two::{
    five::run_five, four::run_four, one::run_one, three::run_three, two::run_two,
};
use std::{env, process};

fn main() -> Result<(), anyhow::Error> {
    if let true = env::args().count() < 3 {
        eprintln!();
        eprintln!("please provide year and day");
        eprintln!();
        process::exit(1)
    };

    let args = env::args().collect::<Vec<String>>();
    let year = args[1].parse::<u16>().unwrap_or(0);
    let day = args[2].parse::<u16>().unwrap_or(0);

    match year {
        2022 => match day {
            1 => run_one(),
            2 => run_two(),
            3 => run_three(),
            4 => run_four(),
            5 => run_five(),
            _ => {
                eprintln!("{year} {day} not done");
                process::exit(1)
            }
        },
        _ => {
            eprintln!("year not done");
            process::exit(1)
        }
    }
}
