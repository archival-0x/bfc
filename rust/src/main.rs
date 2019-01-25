#[cfg(target_arch="x86_64")]
#![feature(plugin)]
#![plugin(dynasm)]

#[macro_use]
extern crate dynasmrt;

use dynasmrt::{DynasmApi, DynasmLabelApi};


fn main() {
    let mut ops = dynasmrt::x64::Assembler::new().unwrap();
}
