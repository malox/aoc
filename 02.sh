#!/usr/bin/env python3

import sys

path = sys.argv[1]
input = [x.split() for x in open(path,'r').readlines()]

depth,newdepth,pos,aim = 0,0,0,0

for cmd in input:
    if cmd[0] == "forward":
        pos += int(cmd[1])
        newdepth += int(cmd[1])*aim
    elif cmd[0] == "up":
        depth -= int(cmd[1])
        aim -= int(cmd[1])
    elif cmd[0] == "down":
        depth += int(cmd[1])
        aim += int(cmd[1])
    #print("%s %s - %d %d %d %d" % (cmd[0], cmd[1], depth, newdepth, pos, aim))

print(depth*pos)
print(newdepth*pos)
