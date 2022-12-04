#!/usr/bin/python3

import re, sys

filename = sys.argv[1] if len(sys.argv) > 1 else "test"
doc = open(filename, "r").read()
#print(doc)

partone = 0
for value in re.split("\[|\]|{|}|,|:|\"", doc):
  if value.lstrip("-").isdigit():
    partone += int(value)
print("partone {}".format(partone))

import json

parsed = json.load(open(filename, "r"))
#print(parsed)

def check(node):
  if type(node) is str:
    return 0
  if type(node) is int:
    return node
  if type(node) is dict:
    total = 0
    items = node.values()
    if "red" not in items:
      for tt in items:
        total += check(tt)
    return total
  if type(node) is list:
    total = 0
    for tt in node:
      total += check(tt)
    return total

print("parttwo {}".format(check(parsed)))
