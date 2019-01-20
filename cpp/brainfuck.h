#ifndef BRAINFUCK_H
#define BRAINFUCK_H

#include <fstream>
#include <string>
#include <iostream>
#include <stdexcept>

#define TAPE_SIZE 30000

class Interpreter;

std::string checkFile(const std::string &filename); 

#endif
