#!/usr/bin/env python3

import sys
path = sys.argv[1]

rounds = [x.split() for x in open(path,'r').readlines()]
#print(rounds)

rock = 1
paper = 2
scissor = 3

draw = 0
win = 1
lost = 2

def score(me, result):
    total = me
    if result == draw :
        total += 3 #draw
    elif result == win :
        total += 6 #win
    elif result == lost :
        total += 0 #lost
    return total

# ====== PART ONE

rules = dict()
rules[(rock, rock)] = draw
rules[(rock, paper)] = lost
rules[(rock, scissor)] = win
rules[(paper, rock)] = win
rules[(paper, paper)] = draw
rules[(paper, scissor)] = lost
rules[(scissor, rock)] = lost
rules[(scissor, paper)] = win
rules[(scissor, scissor)] = draw


def value(shape):
    if shape == "A" or shape == "X":
        return rock
    elif shape == "B" or shape == "Y":
        return paper
    elif shape == "C" or shape == "Z":
        return scissor

total = 0
for round in rounds:
    me = value(round[1])
    other = value(round[0])
    result = rules[(me, other)]
    total += score(me, result)
print("part one: {}".format(total))

# ====== PART TWO

newrules = dict()
newrules[(rock, draw)] = rock
newrules[(rock, win)] = paper
newrules[(rock, lost)] = scissor
newrules[(paper, lost)] = rock
newrules[(paper, draw)] = paper
newrules[(paper, win)] = scissor
newrules[(scissor, win)] = rock
newrules[(scissor, lost)] = paper
newrules[(scissor, draw)] = scissor

def value_two(shape):
    if shape == "X":
        return lost
    elif shape == "Y":
        return draw
    elif shape == "Z":
        return win

total = 0
for round in rounds:
    other = value(round[0])
    result = value_two(round[1])
    me = newrules[(other, result)]
    total += score(me, result)
print("part two: {}".format(total))

