
def doit():
    try:
        sum = 0
        f=open("adv04i.txt", "r")
        l=f.readlines()
        for x in l:
            pwd = x.split()
            #print "test " + str(x) + " - " + str(len(pwd)) + " - " + str(len(set(pwd))) 
            if len(pwd)==len(set(pwd)):
                #print str(x) + " is valid"
                sum += 1
            #else:
                #print str(x) + " is not valid"
        print "SUM " + str(sum)
    finally:
        f.close()
                    
doit()
