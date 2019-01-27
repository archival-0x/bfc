#[macro_use]
extern crate dynasmrt;

use std::env;
use std::mem;

use dynasmrt::{DynasmApi, DynasmLabelApi};

mod bf;
use bf::BFInterpreter;

fn execute_jit(prog: BFInterpreter) -> () {
    let mut ops = dynasmrt::x64::Assembler::new().unwrap();
}


fn main() {
    
    // parse arguments
    let mut args: Vec<_> = env::args().collect();
    if args.len() != 2 {
        println!("Expected one argument (.bf file), got {}", args.len());
    }
   
    // read file
}
