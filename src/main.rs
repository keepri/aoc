#![allow(special_module_name)]

mod twenty_two;

use crate::twenty_two::{one::run_1, two::answer::run_2};
use std::process::exit;

const RUN: u8 = 2;

fn main() -> Result<(), anyhow::Error> {
    match RUN {
        1 => run_1(),
        2 => run_2(),
        _ => {
            eprintln!("not done");
            exit(1)
        }
    }
}

// TODO:

// [ ] add by year selection when multiple years will exist
