#!/usr/bin/env python3

import sys

input = [
    "D2FE28",
    "38006F45291200",
    "EE00D40C823060",
    "8A004A801A8002F478",
    "620080001611562C8802118E34",
    "C0015000016115A2E0802F182340",
    "A0016C880162017C3686B18A3D4780",
    "00569F4A0488043262D30B333FCE6938EC5E5228F2C78A017CD78C269921249F2C69256C559CC01083BA00A4C5730FF12A56B1C49A480283C0055A532CF2996197653005FC01093BC4CE6F5AE49E27A7532200AB25A653800A8CAE5DE572EC40080CD26CA01CAD578803CBB004E67C573F000958CAF5FC6D59BC8803D1967E0953C68401034A24CB3ACD934E311004C5A00A4AB9CAE99E52648401F5CC4E91B6C76801F59DA63C1F3B4C78298014F91BCA1BAA9CBA99006093BFF916802923D8CC7A7A09CA010CD62DF8C2439332A58BA1E495A5B8FA846C00814A511A0B9004C52F9EF41EC0128BF306E4021FD005CD23E8D7F393F48FA35FCE4F53191920096674F66D1215C98C49850803A600D4468790748010F8430A60E1002150B20C4273005F8012D95EC09E2A4E4AF7041004A7F2FB3FCDFA93E4578C0099C52201166C01600042E1444F8FA00087C178AF15E179802F377EC695C6B7213F005267E3D33F189ABD2B46B30042655F0035300042A0F47B87A200EC1E84306C801819B45917F9B29700AA66BDC7656A0C49DB7CAEF726C9CEC71EC5F8BB2F2F37C9C743A600A442B004A7D2279125B73127009218C97A73C4D1E6EF64A9EFDE5AF4241F3FA94278E0D9005A32D9C0DD002AB2B7C69B23CCF5B6C280094CE12CDD4D0803CF9F96D1F4012929DA895290FF6F5E2A9009F33D796063803551006E3941A8340008743B8D90ACC015C00DDC0010B873052320002130563A4359CF968000B10258024C8DF2783F9AD6356FB6280312EBB394AC6FE9014AF2F8C381008CB600880021B0AA28463100762FC1983122D2A005CBD11A4F7B9DADFD110805B2E012B1F4249129DA184768912D90B2013A4001098391661E8803D05612C731007216C768566007280126005101656E0062013D64049F10111E6006100E90E004100C1620048009900020E0006DA0015C000418000AF80015B3D938"
    ]

# --------------------------- common

conv = lambda s : "".join([format(int(x, 16), '04b') for x in s])

verbose = False
def Print(packets):
    global verbose
    if verbose:
        for pckt in packets: 
            pckt.dump("-- ")

def logdebug(s) :
    global verbose
    if verbose:
        print(s)

class Packet:
    def __init__(self, bits):
        self.version = int(bits[:3], 2)
        self.type = int(bits[3:6], 2)
        self.number = ""
        self.length = 0
        self.read_bytes = 5 if self.type == 4 else 1
        self.packets = []
        self.packets_size = 0
        self.packets_number = 0
        self.current = None
        logdebug("+++    debug {} - {}".format(bits, self)) 

    def literal(self, bits):
        self.number += bits[1:]
        if bits[0] == '0':
            self.read_bytes = 0
            self.number = int(self.number, 2)

    def computeLength(self, bits):
        self.length = 15 if bits == '0' else 11
        self.read_bytes = self.length

    def subPacket(self, bits):
        if not self.current:
            self.current = Packet(bits)
        else:
            self.current.read(bits)
        if self.current.done():
            self.packets.append(self.current)
            self.current = None
            self.read_bytes = 0 if len(self.packets) == self.packets_number else 6
        else:
            self.read_bytes = self.current.read_bytes

    def read(self, bits):
        logdebug("  >>>  debug {} - {}".format(bits, self))
        if self.type == 4:
            self.literal(bits)
        elif self.packets_size > 0:
            self.packets = Parse(bits)
            self.read_bytes = 0
        elif self.packets_number > 0:
            self.subPacket(bits)
        elif self.length == 0:
            self.computeLength(bits)
        elif self.length == 15:
            self.packets_size = int(bits, 2)
            self.read_bytes = self.packets_size
        elif self.length == 11:
            self.packets_number = int(bits, 2)
            self.read_bytes = 0 if len(self.packets) == self.packets_number else 6
        logdebug("  <<<  debug {} - {}".format(bits, self))

    def done(self) :
        return self.read_bytes == 0

    def __repr__(self):
        if self.type == 4:
            return "Packet v {}, t {}, number {}, read {}".format(self.version, self.type, self.number, self.read_bytes)
        else:
            return "Packet v {}, t {}, length {}, size {}, number {}, read {}".format(self.version, self.type, self.length, self.packets_size, self.packets_number, self.read_bytes)

    def dump(self, tag):
        logdebug("{} {}".format(tag, self))
        for psub in self.packets :
            psub.dump(tag*2)

    def eval(self):
        if self.type == 0:
            return sum([pckt.eval() for pckt in self.packets])
        elif self.type == 1:
            result = 1
            for pckt in self.packets:
                result = result * pckt.eval()
            return result
        elif self.type == 2:
            return min([pckt.eval() for pckt in self.packets])
        elif self.type == 3:
            return max([pckt.eval() for pckt in self.packets])
        elif self.type == 4:
            return self.number
        elif self.type == 5:
            return 1 if self.packets[0].eval() > self.packets[1].eval() else 0
        elif self.type == 6:
            return 1 if self.packets[0].eval() < self.packets[1].eval() else 0
        elif self.type == 7:
            return 1 if self.packets[0].eval() == self.packets[1].eval() else 0
        else:
            print("WARN this should never happen")
            return 1
        

def Parse(sequence):
    idx = 0
    packets = []
    current = None
    while idx < len(sequence):
        if current == None:
            if idx+6 >= len(sequence): break
            current = Packet(sequence[idx:idx+6])
            idx += 6
        else :
            read = current.read_bytes
            current.read(sequence[idx:idx+read])
            idx += read
        if current.done():
            packets.append(current)
            current = None
    return packets

def Versions(packets):
    return sum([pkct.version for pkct in packets]) + sum([ Versions(pkct.packets) for pkct in packets])

# --------------------------- part 1 & 2

for input_str in input:
    #print("hex : {}\nbin : {}".format(input_str, conv(input_str)))
    packets = Parse(conv(input_str))
    Print(packets)
    print("part one {} two {} - input {}".format(Versions(packets), packets[0].eval(), input_str[:50]))

# --------------------------- 

