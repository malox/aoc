#!/usr/bin/python

import hashlib

def parseImpl(input, idx, prefix):
    hexhash = hashlib.md5(input + idx).hexdigest()
    if hexhash.startswith(prefix):
        print "input " + input + " " + idx + " - " + hexhash
        return True
    return False

def parse(input, prefix = "00000"):
    idx = 0
    while not parseImpl(input, str(idx), prefix):
        idx += 1

if __name__ == "__main__":
    parse("abcdef")
    parse("pqrstuv")
    parse("bgvyzdsv")
    parse("bgvyzdsv", prefix = "000000")
