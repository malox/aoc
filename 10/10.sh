#!/usr/bin/python

def parse(s):
  digits = [int(x) for x in s]
  res, curr, found = "", digits[0], 1
  for it in range(1, len(digits)):
    if curr == digits[it]: 
      found += 1
    else:
      res += "{}{}".format(found, curr)
      curr, found = digits[it], 1
  res += "{}{}".format(found, curr)
  return res
  
#for pp in ["1", "11", "211", "21", "1211", "111221"]:
#  print("{} => {}".format(pp, parse(pp)))

start = "3113322113"
for it in range(40):
  start = parse(start)
print("partone {}".format(len(start)))

start = "3113322113"
for it in range(50):
  start = parse(start)
print("parttwo {}".format(len(start)))
