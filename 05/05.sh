#!/usr/bin/python

import sys

def readfile(filename):
    file = open(filename, "r")
    str = []
    for line in file:
        str.append(line.replace('\n', ''))
    return str

class PartOne:
    def __init__(self, input):
        self.input = input

    def isInvalid(self, current):
        forbidden = ["ab", "cd", "pq", "xy"]
        for str in forbidden:
            if str in current:
                #print "isInvalid " + current
                return True

        return False

    def isValid(self, current):
        gvowels = "aeiou"
        vowels = 0
        twice = 0
        for idx in range(len(current)):
            if current[idx] in gvowels:
                vowels += 1
            if idx > 0 and current[idx] == current[idx-1]:
                twice += 1
        #print "isValid " + current + " - vowels " + str(vowels) + " - twice " + str(twice)
        return True if vowels > 2 and twice > 0 else False

    def parse(self):
        nice = 0
        for curr in self.input:
            if self.isInvalid(curr):
                continue
            elif self.isValid(curr):
                nice += 1
        print "TOT " + str(nice)

class PartTwo:
    def __init__(self, input):
        self.input = input

    def middle(self, current):
        for idx in range(1, len(current) - 1):
            if current[idx-1] == current[idx+1] and current[idx] != current[idx+1]:
                return True
        return False

    def repeat(self, current):
        end = len(current) - 1
        for idx in range(end) :
            a = current[idx]
            b = current[idx+1]
            for jdx in range(idx + 2, end) :
                c = current[jdx]
                d = current[jdx+1]
                if a == c and b == d:
                    return True
        return False

    def parse(self):
        nice = 0
        for curr in self.input:
            if self.middle(curr) and self.repeat(curr):
                nice += 1
        print "TOT2 " + str(nice)

if __name__ == "__main__":
    filename = "test"
    if len(sys.argv) > 1 :
        filename = sys.argv[1]

    input = readfile(filename)

    pone = PartOne(input)
    pone.parse()

    ptwo = PartTwo(input)
    ptwo.parse()
