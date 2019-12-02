
def doit():
    try:
        sum = 0
        f=open("adv04i.txt", "r")
        l=f.readlines()
        for x in l:
            pwd = x.split()
            pwd2 = []
            for p in pwd:
                #print "test " + str(p) + " - " + str(''.join(sorted(p)))
                pwd2.append(''.join(sorted(p)))
            #print "test " + str(x) + " - " + str(len(pwd2)) + " - " + str(len(set(pwd2))) 
            if len(pwd2)==len(set(pwd2)):
                #print str(x) + " is valid"
                sum += 1
            #else:
                #print str(x) + " is not valid"
        print "SUM " + str(sum)
    finally:
        f.close()
                    
doit()
