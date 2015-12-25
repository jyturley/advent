#!/bin/python
import re

def parse():
    with open('input.txt', 'r') as f:
        data = f.readlines()

    conversions = {}
    molecule = ""
    r = re.compile(r'(\w+) => (\w+)')
    for conversion_or_molecule in data:
        m = re.match(r, conversion_or_molecule)
        if m:
            key, value = m.groups()
            if not key in conversions:
                conversions[key] = []
            conversions[key].append(value)
        else:
            molecule = conversion_or_molecule

    return conversions, molecule

def test():
    molecule = "HOH"
    print "all tests pass!"

def partone():
    conversions, molecule = parse()
    print conversions
    print molecule

def parttwo():
    pass

if __name__ == '__main__':
    test()
    partone()
    # parttwo()
