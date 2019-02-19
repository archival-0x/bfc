//! bf.rs
//!
//!     normal Brainfuck program interpreter, as the API can
//!     work as standalone BF implementation w/out dynamic
//!     compilation.


// defines the default number of cells within
// a Brainfuck tape.
const TAPE_SIZE: usize = 30000;


/// defines valid instructions for Brainfuck execution
pub enum BFInstr {
    IncrementPtr,
    DecrementPtr,
    IncrementVal,
    DecrementVal,
    Write,
    Read,
    BeginLoop,
    ExitLoop
}


pub struct BFInterpreter {
    pub inst_buf: Vec<BFInstr>,
    tape_size: usize,
}


/// defines errors encountered during BFInterpreter interaction
pub enum BFError {
    InvalidInstruction,
    BadTapeSize,
    RuntimeOverflow,
    RuntimeEndOfLoop
}


impl BFInterpreter {

    /// internal parser helper that converts raw instruction chars
    /// into valid Brainfuck instructions.
    fn parse_raw_instr(instruction: char) -> Result<BFInstr, BFError> {
        match instruction {
            '>' => Ok(BFInstr::IncrementPtr),
            '<' => Ok(BFInstr::DecrementPtr),
            '+' => Ok(BFInstr::IncrementVal),
            '-' => Ok(BFInstr::DecrementVal),
            '.' => Ok(BFInstr::Write),
            ',' => Ok(BFInstr::Read),
            '[' => Ok(BFInstr::BeginLoop),
            ']' => Ok(BFInstr::ExitLoop),
            _   => Err(BFError::InvalidInstruction)
        }
    }


    /// `new_program` generates a new Brainfuck struct for interaction by
    /// parsing a raw text of the program.
    pub fn new_program(&self, program: &str) -> Result<BFInterpreter, BFError> {

        // use filter_map in order to create instr vector iteratively with
        // instruction parser helper
        let mut inst_buf: Vec<BFInstr> = program.chars()
            .filter_map(|c| self.parse_raw_intr(c).is_ok())
            .collect::<Vec<BFInstr>>();

        Ok(BFInterpreter {
            inst_buf: inst_buf,
            tape_size: TAPE_SIZE
        })
    }


    /// `override_tape` changes the size of Brainfuck tape.
    /// (default is 30,000).
    pub fn override_tape(&mut self, size: usize) -> () {
        self.tape_size = size;
    }
}
