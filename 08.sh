#!/usr/bin/env python3

import sys
path = sys.argv[1]

# parse commands
signals = [ x[:-1].split(' | ') for x in open(path,'r').readlines() ]
#print("signals {} ".format(signals))

# --------------------------- part 1

tot = 0
for line in signals:
    for digit in line[1].split():
        val = len(digit)
        if val == 2 or val == 3 or val == 4 or val == 7:
            tot += 1

print("part one %d" % tot)

# --------------------------- part 2

tot = []

def put(cc, i, s):
    #print("    adding %d / %s" % (i,s))
    cc[i] = s
    cc[s] = i

def mySort(s):
    return "".join(sorted(s))

for line in signals:
    patterns = [mySort(x) for x in line[0].split()]
    digits = [mySort(x) for x in line[1].split()]
    conv, tmp = dict(), dict()
    #print("patterns {} - digits {}".format(patterns, digits))

    for patt in patterns:
        val = len(patt)
        tmp[val] = [patt] + tmp.get(val,[])
    #print(" grouped by lenght = {}".format(tmp))

    while len(conv) < len(patterns)*2:
        #print("----------\nnew iteration - conv = ".format(conv)) 
        for patt in patterns:
            #print ("parsing %s" % patt)
            val = len(patt)
            if val == 2:
                put(conv, 1, patt)
                for three in tmp[5]:
                    if patt[0] in three and patt[1] in three:
                        put(conv,3,three)
            elif val == 3:
                put(conv, 7, patt)
            elif val == 4:
                put(conv, 4, patt)
                bd = ""
                for char in patt:
                    if char not in tmp[2][0]:
                        bd += char
                for zero in tmp[6]:
                    if bd[0] not in zero or bd[1] not in zero:
                        put(conv, 0, zero)
                e, beg, zero, seven = "", "", conv[0], tmp[3][0]
                for char in zero:
                    beg += "" if char in seven else char
                #print("e {} - beg {} - zero {} - seven {}".format(e,beg,zero,seven))
                for sixornine in tmp[6]:
                    if sixornine != zero:
                        for char in beg:
                            if char not in sixornine:
                                e = char
                                put(conv, 9, sixornine)
                for six in tmp[6]:
                    if six != zero and six != conv[9]:  
                        put(conv, 6, six)
                for two in tmp[5]:
                    if e in two:
                        put(conv, 2, two)
            elif val == 7:
                put(conv, 8, patt)
            elif len(conv) == 18:
                for five in tmp[5]:
                    if five not in conv:
                        put(conv, 5, five)
    #print("mapped conversion table (len={}) => {} \n  digits {}".format(len(conv), conv, digits)) 
    tot.append(conv[digits[0]]*1000 + conv[digits[1]]*100 + conv[digits[2]]*10 + conv[digits[3]])
    
print("part two {}".format(sum(tot)))

# --------------------------- 


