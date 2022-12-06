use anyhow::Result;
use std::fs;

pub fn run_3() -> Result<()> {
    let input = fs::read_to_string("static/files/input_3.txt")?;
    Ok(())
}
