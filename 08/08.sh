#!/usr/bin/python

import sys

filename = sys.argv[1] if len(sys.argv) > 1 else "test"
lines = [x[:len(x)-1] for x in open(filename, "r").readlines()]
#print(lines)

partone = 0
parttwo = 0
for l in lines:
  data = 0
  size = len(l)
  escape = False
  hex = 0
  extra=6 # base add for "\"\""
  for it in range(1, size-1):
    if hex > 0:
      hex -= 1
      if hex == 0:
        data += 1
        extra += 5
    elif escape:
      escape = False
      if l[it] == 'x':
        hex = 2
      else:
        data += 1
        extra += 4 if l[it] == "\"" or l[it] == "\\" else 3
    elif l[it] == "\\":
      escape = True
    else:
      data += 1
      extra += 1

  #print("line={} data={} size={} extra={}".format(l, data, size, extra))
  partone += (size-data)
  parttwo += (extra - size)

print("part one {}".format(partone))
print("part two {}".format(parttwo))
