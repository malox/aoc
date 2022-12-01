#!/usr/bin/env python3

import sys
import copy

path = sys.argv[1]
input = [x[:-1] for x in open(path,'r').readlines()]

depth,newdepth,pos,aim = 0,0,0,0

rates = []
pos = len(input[0])


for idx in range(pos):
    zero, one = 0,0
    for signal in input:
        if signal[idx] == "0":
            zero += 1
        else:
            one += 1
    rates.append((zero,one))

#print(rates)

gamma, epsilon = "", ""
for rate in rates:
    gamma += "0" if rate[0] > rate[1] else "1"
    epsilon += "1" if rate[0] > rate[1] else "0"

print( "%s %s - %d" % (gamma, epsilon, int(gamma,2) * int(epsilon,2)) )

# part two

oxygen, co2 = "", ""

tmp = copy.deepcopy(input)
for idx in range(pos):
    zero, one = 0,0
    for signal in tmp:
        if signal[idx] == "0":
            zero += 1
        else:
            one += 1

    aux = []
    for signal in tmp:
        if zero > one and signal[idx] == "0" :
            aux.append(signal)
        elif zero <= one and signal[idx] == "1" :
            aux.append(signal)

    tmp = aux
    if len(tmp) == 1 : break

oxygen = tmp[0]

tmp = copy.deepcopy(input)
for idx in range(pos):
    zero, one = 0,0
    for signal in tmp:
        if signal[idx] == "0":
            zero += 1
        else:
            one += 1

    aux = []
    for signal in tmp:
        if zero <= one and signal[idx] == "0" :
            aux.append(signal)
        elif zero > one and signal[idx] == "1" :
            aux.append(signal)

    tmp = aux
    if len(tmp) == 1 : break

co2 = tmp[0]

print("%s %s - %d" % (oxygen, co2, int(oxygen,2)*int(co2,2)))

