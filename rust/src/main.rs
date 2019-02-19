extern crate itertools;
extern crate dynasm;
extern crate dynasmrt;

use std::env;
use std::mem;

use itertools::multipeek;

use dynasm::dynasm;
use dynasmrt::{DynasmApi, DynasmLabelApi};

mod bf;
use bf::{BFInterpreter, BFInstr};

/// main interface for executing a parsed brainfuck program
/// through JIT using dynasm
fn execute_jit(prog: BFInterpreter) -> () {
    let mut ops = dynasmrt::x64::Assembler::new().unwrap();
    let mut loops = Vec::new();
    let mut code = multipeek(program.iter().cloned());

    // initialize prologue
    let start = prologue!(ops);

    for inst in prog.inst_buf {
        match inst {
            BFInstr::IncrementPtr => {
            },
            BFInstr::DecrementPtr => {
            },
            BFInstr::IncrementVal => {
            },
            BFInstr::DecrementVal => {
            },
            BFInstr::Write => {
            },
            BFInstr::Read => {
            },
            BFInstr::BeginLoop => {
            },
            BFInstr::ExitLoop => {
            },
            _ => {} 
        }
    }

    if loops.len() != 0 {
        return BFError::RuntimeEndOfLoop;
    }

    // finalize with epilogue
    dynasm!(ops
            ;; epilogue!(ops, 0)
            ;->overflow:
            ;; epilogue!(ops, 1)
            ;->io_failure:
            ;; epilogue!(ops, 2)
    );
    let code = ops.finalize().unwrap();

    ()
}


fn main() {
    
    // parse arguments
    let mut args: Vec<_> = env::args().collect();
    if args.len() != 2 {
        println!("Expected one argument (.bf file), got {}", args.len());
    }

    // read file into str
    let mut file = File::open(args[1]).expect("Unable to open file");
    let mut contents = String::new();
    file.read_to_string(&mut contents).expect("Unable to read file");
    
    // initialize interpreter and execute using JIT strategy
    let prog = BFInterpreter::new_program(contents.as_str());
    execute_jit(prog);
}
