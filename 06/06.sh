#!/usr/bin/python

import sys

class Parser:
    def __init__(self, filename):
        self.input = []
        file = open(filename, "r")
        for line in file:
            parsed = line.replace('\n', '').replace(",", " ").replace("turn ", "").replace("through ", "")
            self.input.append(parsed.split())
        self.reset()

    def reset(self):
        self.matrix = []
        self.size = 1000
        for idx in range(self.size):
            row = []
            for jdx in range(self.size):
                row.append(0)
            self.matrix.append(row)

    def count(self):
        count = 0
        for idx in range(self.size):
            for jdx in range(self.size):
                count += self.matrix[idx][jdx]
                
        print "count " + str(count)
    
    def parse(self):
        for opt in self.input:
            action = opt[0]
            startrow = int(opt[1])
            startcol = int(opt[2])
            endrow = int(opt[3])
            endcol = int(opt[4])
            for idx in range(startrow, endrow + 1):
                for jdx in range(startcol, endcol + 1):
                    if action == "on":
                        self.matrix[idx][jdx] = 1
                    elif action == "off":
                        self.matrix[idx][jdx] = 0
                    elif action == "toggle":
                        if self.matrix[idx][jdx] == 1:
                            self.matrix[idx][jdx] = 0
                        elif self.matrix[idx][jdx] == 0:
                            self.matrix[idx][jdx] = 1
        self.count()

    def parseTwo(self):
        self.reset()
        for opt in self.input:
            action = opt[0]
            startrow = int(opt[1])
            startcol = int(opt[2])
            endrow = int(opt[3])
            endcol = int(opt[4])
            for idx in range(startrow, endrow + 1):
                for jdx in range(startcol, endcol + 1):
                    if action == "on":
                        self.matrix[idx][jdx] += 1
                    elif action == "off":
                        if self.matrix[idx][jdx] > 0:
                            self.matrix[idx][jdx] -= 1
                    elif action == "toggle":
                        self.matrix[idx][jdx] += 2
        self.count()
        

if __name__ == "__main__":
    filename = "test"
    if len(sys.argv) > 1 :
        filename = sys.argv[1]

    parser = Parser(filename)
    parser.parse()
    parser.parseTwo()
