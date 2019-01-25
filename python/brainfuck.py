#!/usr/bin/env python3
"""
brainfuck.py

    a bf interpreter in Python but calls C through ffi

"""

import sys
import os.path
import cffi

def main():

    # read input

    # initialize ffi
    ffi = cffi.FFI();
    ffi.cdef("""


    """)
    lib = ffi.dlopen(obj_path)

    # call interpet (works on internal tape state)

    # call global tape helpers

if __name__ == "__main__":
    main()

