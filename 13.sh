#!/usr/bin/env python3

import sys, copy
path = sys.argv[1]

# --------------------------- common

points, folds = open(path,'r').read().split('\n\n')
points = [(int(x), int(y)) for x, y in [z.split(',') for z in points.split()]]
folds  = [(f, int(g)) for f,g in [tuple(x.split('=')) for x in folds.split() if '=' in x]]

#print("points : {}".format(points))
#print("folds : {}".format(folds))

# build board
board = []
max_row, max_col = max([x[1] for x in points])+1, max([x[0] for x in points])+1
for idx in range(max_row): board.append(['.']*max_col)
for pp in points: board[pp[1]][pp[0]] = '#'

def dump(board):
    print('-'*len(board[0]))
    for line in board:
        print("".join(line))
    print('-'*len(board[0]))
#dump(board)

# --------------------------- part 1 & 2

for cmd in folds:
    newboard = []
    row_num, col_num = len(board), len(board[0])
    if cmd[0] == 'x':
        for x in range(len(board)):
            newboard.append([])
        for jdx in range(cmd[1]):
            for idx in range(row_num):
                newboard[idx].append('#' if board[idx][jdx] == '#' or board[idx][col_num - jdx - 1] == '#' else '.')
    elif cmd[0] == 'y':
        for idx in range(cmd[1]):
            newline = []
            for jdx in range(col_num):
                newline.append('#' if board[idx][jdx] == '#' or board[row_num - idx - 1][jdx] == '#' else '.')
            newboard.append(newline)
    else: continue # it should never happen
    board = newboard
    count = 0
    for line in board:
        count += line.count('#')
    print("cmd {} - count {}".format(cmd, count))
    #dump(board)
dump(board)

# --------------------------- 

