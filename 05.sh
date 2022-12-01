#!/usr/bin/env python3

import sys
path = sys.argv[1]

# parse commands
cmds = []
for x in [x[:-1].split(' -> ') for x in open(path,'r').readlines()]:
    cmds.append([int(y) for y in x[0].split(',')] + [int(y) for y in x[1].split(',')])

#print(cmds)

# build matrix
def buildMatrix(cmds):
    x,y = 0,0
    for cc in cmds:
        x,y = max(x, cc[0], cc[2]), max(y,cc[1],cc[3])
    x,y = x+1, y+1
    #print("x,y=%d,%d" % (x,y))
    matrix = []
    for t in range(y):
        v = [0]*x
        matrix.append(v)
    return matrix

def printMatrix(m):
    print("\n" + "-"*len(m))
    for line in m:
        s = ""
        for cc in line:
            s += "." if cc == 0 else str(cc)
        print(s)
    print("-"*len(m) + "\n")

def countMatrix(matrix):
    count = 0
    for line in matrix:
        for val in line:
            if val > 1:
                count += 1
    print("count %d" % count) 

matrix = buildMatrix(cmds)
#printMatrix(matrix)

# --------------------------- part 1

for cc in cmds:
    #print("processing cmd %s" % str(cc))
    if cc[0] == cc[2]:
        x,y = min(cc[1], cc[3]), max(cc[1],cc[3])
        for idx in range(x, y+1):
            #print("x - %d %d" % (cc[0], idx))
            matrix[idx][cc[0]] += 1
    elif cc[1] == cc[3]:
        x,y = min(cc[0], cc[2]), max(cc[0],cc[2])
        for idx in range(x, y+1):
            #print("y - %d %d" % (idx, cc[1]))
            matrix[cc[1]][idx] += 1 
    #printMatrix(matrix)
#printMatrix(matrix)

countMatrix(matrix)

# --------------------------- part 2
dummy = 0
#printMatrix(matrix)

for cc in cmds:
    if cc[0] == cc[2] or cc[1] == cc[3]:
        dummy += 1# nothing to do
    else:
        #print("processing cmd %s" % str(cc))
        a,b = cc[0], cc[1]
        c,d = cc[2], cc[3]
        #print(" ab %d %d - b c %d %d" % (a,b,c,d))
        if a>c:
            a,b,c,d = c,d,a,b
            #print("swapped ab %d %d - b c %d %d" % (a,b,c,d))
        if b>d:
            while a <= c and b >= d:
                matrix[b][a] += 1
                a += 1
                b -= 1
        else:
            while a <= c and b <= d:
                matrix[b][a] += 1
                a += 1
                b += 1
        #printMatrix(matrix)
#printMatrix(matrix)

countMatrix(matrix)


# ---------------------------
