#include "brainfuck.h"

class Interpreter {

    public:

        /* initialize tape with appropriate size */
        char tape[TAPE_SIZE];
        char *d = tape;

        const char *c;
        size_t loop;

        Interpreter(const char* code) {
            c = code;
            while (*c) {
                switch (*c) {
                    case '>': d++;             break;
                    case '<': d--;             break;
                    case '+': (*d)++;          break;
                    case '-': (*d)--;          break;
                    case '.': std::cout << *d; break;
                    case ',': std::cin >> *d;  break;
                    case '[':
                        loop = 1;
                        if (*d == '\0') {
                            do {
                                c++;
                                if     (*c == '[') loop++;
                                else if (*c == ']') loop--;
                            } while ( loop != 0 );
                        }
                        break;
                    case ']':
                        loop = 0;
                        do {
                            if      (*c == '[') loop++;
                            else if (*c == ']') loop--;
                            c--;
                        } while ( loop != 0 );
                        break;
                }
                c++;
            }
        }
};


std::string checkFile(const std::string &filename) {

    /* check file extension */
    if (filename.substr(filename.find_last_of(".") + 1) != "bf")
        throw std::runtime_error("ERROR: File is not of .bf extension");

    /* open file for reading */
    std::ifstream file(filename);
    if (!file.is_open())
        throw std::runtime_error("ERROR: Failed to open " + filename + "!");

    std::string input {
        std::istreambuf_iterator<char>(file),
        std::istreambuf_iterator<char>()
    };
    return input;
}


int main(int argc, char **argv) {

    /* arg check */
    if (argc != 2) {
        std::cerr << "Usage: " << argv[0] << " file.bf" << std::endl;
        return 1;
    }

    /* file check and run interpreter */
    Interpreter bf = Interpreter(checkFile(argv[1]).c_str());
    return 0;
}
