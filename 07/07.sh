#!/usr/bin/python

import sys

class Input:
    def __init__(self, line):
        arr = line.split()
        self.cmd = arr[:len(arr)-2]
        self.target = arr[len(arr)-1]

    def __repr__(self):
        return str(self.cmd) + ' -> ' + self.target

class Node:
    def __init__(self, name):
        self.name = name
        self.value = 0
        self.valid = False

    def __repr__(self):
        return "node " + str(self.name) + ':' + str(self.value) + ':' + str(self.valid)

class Wire:
    def __init__(self):
        self.one = None
        self.two = None
        self.out = None
        self.ops = "FWD"
        self.val = 0
        self.pending = True

    def __repr__(self):
        return "wire - one[" + str(self.one) + "] ops[" + self.ops + "] two[" + str(self.two) + "] " \
                    + "out[" + str(self.out) + "] val[" + str(self.val) + "] pending[" + str(self.pending) + "]"
    
    def run(self):
        #print "wire - before run : " + str(self)
        if self.pending == False :
            return
        
        if self.ops == "NOT" :
            if self.one.valid == False:
                return
            self.out.value = 65535 - self.one.value

        elif self.ops == "FWD" :
            if self.one != None and self.one.valid == True:
                self.out.value = self.one.value
            elif self.val > 0 :
                self.out.value = self.val
            else :
                return

        elif self.ops == "AND" :
            if self.one.valid == False or self.two.valid == False :
                return
            self.out.value = self.one.value & self.two.value
           
        elif self.ops == "OR" :
            if self.one.valid == False or self.two.valid == False :
                return
            self.out.value = int(self.one.value | self.two.value)

        elif self.ops == "LSHIFT" :
            if self.one.valid == False :
                return
            self.out.value = int(self.one.value << self.val)

        elif self.ops == "RSHIFT" :
            if self.one.valid == False:
                return
            self.out.value = int(self.one.value >> self.val)
        
        self.out.valid = True
        self.pending = False
        #print "wire -  after run : " + str(self)


class Parser:
    def __init__(self, filename):
        self.input = []
        self.nodes = dict()
        self.wires = []
        file = open(filename, "r")
        for line in file:
            self.input.append(Input(line))

    def getNode(self, name):
        x = self.nodes.get(name)
        if x == None:
            self.nodes.update({name : Node(name)})
        return self.nodes[name]

    def parse(self):
        #print self.input
        for instr in self.input:
            wire = Wire()
            wire.out = self.getNode(instr.target)
            if len(instr.cmd) == 1 :
                if instr.cmd[0].isdigit() :
                    wire.val = int(instr.cmd[0])
                else :
                    wire.one = self.getNode(instr.cmd[0]) 
            elif len(instr.cmd) == 2 :
                wire.ops = instr.cmd[0] # "NOT"
                wire.one = self.getNode(instr.cmd[1]) 
            elif len(instr.cmd) == 3 :
                wire.one = self.getNode(instr.cmd[0]) 
                wire.ops = instr.cmd[1] # and/or/[rl]shift
                if instr.cmd[2].isdigit() :
                    wire.val = int(instr.cmd[2]) 
                else :
                    wire.two = self.getNode(instr.cmd[2]) 

            self.wires.append(wire)
        
        #self.dump()

        while True :
            run = 0
            for wire in self.wires:
                if wire.pending == True :
                    wire.run()
                    if wire.pending == True :
                        print "pending wire " + str(wire)
                    run += 1
            if run == 0:
                break
        
        #self.dump()
        for k in self.nodes :
            #print k
            if k == "a" :
                print " - " + str(self.nodes[k])

    def dump(self):
        print "nodes :"
        for k in self.nodes :
            print " - " + str(self.nodes[k])

        print "wires :" 
        for wire in self.wires :
            print " - " + str(wire)


if __name__ == "__main__":
    filename = "test"
    if len(sys.argv) > 1 :
        filename = sys.argv[1]

    parser = Parser(filename)
    parser.parse()