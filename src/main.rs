#![allow(special_module_name)]

mod lib;

use crate::lib::twenty_two::{one::run_1, three::run_3, two::run_2};
use std::process::exit;

const RUN: u8 = 3;

fn main() -> Result<(), anyhow::Error> {
    match RUN {
        1 => run_1(),
        2 => run_2(),
        3 => run_3(),
        _ => {
            eprintln!("not done");
            exit(1)
        }
    }
}

// TODO:

// [ ] add by year selection when multiple years will exist
