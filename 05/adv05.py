import sys

def parse(line):
    #print("input   = " + line)
    modified = False
    skip = False
    newline = ""
    for it in range(0, len(line)-1):
        #print(str(it) + " - " + line[it] + " " + line[it+1])
        if skip == True:
            #print( " skip " + line[it])
            skip = False
        else:
            if line[it] == line[it+1] or line[it].lower() != line[it+1].lower():
                #print( " append " + line[it])
                newline += line[it]
            else:
                #print( " discard " + line[it])
                skip = True
                modified = True
    if skip != True:
        #print( " append last char " + line[len(line)-1])
        newline += line[len(line)-1]
    #else:
        #print( " skip last char " + line[len(line)-1])

    if modified == True:
        return parse(str(newline))
    #print(str(len(newline)) + " output = " + newline + "\n")

    return len(newline)

def parsetwo(line):
    modified = True
    tmpline = line
    while modified == True:
        modified = False
        skip = False
        newline = ""
        for it in range(0, len(tmpline)-1):
            if skip == True:
                skip = False
            else:
                if tmpline[it] == tmpline[it+1] or tmpline[it].lower() != tmpline[it+1].lower():
                    newline += tmpline[it]
                else:
                    skip = True
                    modified = True

        if skip != True:
            newline += tmpline[len(tmpline)-1]

        tmpline = newline

    return len(tmpline)

def partwo(line):
    for ch in range(ord('a'),ord('z')+1):
        unit = chr(ch)
        newline = line.replace(unit, '').replace(unit.upper(),'')
        print (unit + " " + str(parsetwo(newline)))

for line in sys.stdin:
    print(parse(line.strip()))
    partwo(line.strip())
    
