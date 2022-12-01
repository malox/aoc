#!/usr/bin/env python3

import sys
path = sys.argv[1]

# parse commands
commands = [x[:-1] for x in open(path,'r').readlines()]
#print("commands {} ".format(commands))

syntax = {'(':')', '[':']', '<':'>', '{':'}'}
#syntax.update({(v,k) for k,v in syntax.items()})
#print("syntax {} ".format(syntax))

prices = {
    ')': 3,
    ']': 57,
    '}': 1197,
    '>': 25137,
    '(': 1,
    '[': 2,
    '{': 3,
    '<': 4
    }

# --------------------------- common

def parseCommand(cmd):
    stack = []
    for op in cmd:
        if op in '([{<':
            stack.append(op)
        elif op != syntax[stack.pop()]:
            return (prices[op],0)

    score = 0
    while len(stack) > 0:
        op = stack.pop()
        score = (score*5) + prices[op]
    return (0,score)
    

parsed = [parseCommand(cmd) for cmd in commands]

# --------------------------- part 1

print("part one {} ".format(sum([pp[0] for pp in parsed])))

# --------------------------- part 2

scores = []
for val in [pp[1] for pp in parsed]:
    if val == 0 : continue
    scores.append(val)
print("part two {} ".format(sorted(scores)[int(len(scores)/2)]))

# --------------------------- 


