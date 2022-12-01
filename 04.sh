#!/usr/bin/env python3

import sys

path = sys.argv[1]
input = [x[:-1] for x in open(path,'r').readlines()]

numbers = [int(x) for x in input[:1][0].split(',')]
input = input[2:]

class Row:
    def __init__(self, input, parsed = False):
        self.founds = []
        self.values = input if parsed else [int(x) for x in input.split()]
    def __repr__(self):
        return str(self.values)

    def check(self, number):
        if number in self.values:
            self.founds.append(number)
        return len(self.founds) == len(self.values)

    def sum(self):
        sum = 0
        for number in self.values:
            sum += 0 if number in self.founds else number
        return sum

class Board:
    def __init__(self, input):
        self.score = 0
        self.rows = [Row(x) for x in input]
        self.cols = []
        for idx in range(len(self.rows)):
            cc = [x.values[idx] for x in self.rows]
            self.cols.append(Row(cc, True))

    def __repr__(self):
        ss = " [\n"
        for rr in self.rows:
            ss += "    " + str(rr) + "\n"
        return ss + " ]\n"

    def _sum(self, irows):
        self.score = 0
        for row in irows:
            self.score += row.sum()

    def _check(self, number, irows):
        for row in irows:
            if row.check(number):
                self._sum(irows)
                return True
        return False

    def check(self, number):
        if self._check(number, self.rows):
            return True
        else:
            return self._check(number, self.cols)

def createBoards(input):
    boards = []
    tmp = []
    for line in input:
        if line == '':
            boards.append(Board(tmp))
            tmp = []
        else:
            tmp.append(line)
    boards.append(Board(tmp))
    return boards

# part one
boards = createBoards(input)

result = -1
for val in numbers:
    for board in boards:
        if board.check(val):
            result = (val * board.score)
            
        if result > 0 : 
            break
    if result > 0 : break
print("one %d" % result)    

#part two
boards = createBoards(input)

for val in numbers:
    newboards = []
    for board in boards:
        if board.check(val):
            if len(boards) == 1:
                print("two %d" % (val * board.score))
        else:
            newboards.append(board)
    boards = newboards
