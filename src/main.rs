#![allow(special_module_name)]

mod lib;

use crate::lib::twenty_two::{one::run_1, three::run_3, two::run_2};
use std::{env, process};

fn main() -> Result<(), anyhow::Error> {
    if let true = env::args().count() < 3 {
        eprintln!("please provide year and day");
        process::exit(1)
    };

    let args = env::args().collect::<Vec<String>>();
    let year = args[1].parse::<u16>().unwrap_or(0);
    let day = args[2].parse::<u16>().unwrap_or(0);

    match year {
        2022 => match day {
            1 => run_1(),
            2 => run_2(),
            3 => run_3(),
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
