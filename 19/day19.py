#!/bin/python
import re

def generate_replacements(conversions, molecule):
    replacements = []
    for key in conversions:
         for match_index in [m.start() for m in re.finditer(key, molecule)]:
            for conversion in conversions[key]:
                # print match_index, key, conversion, len(conversion)
                # print molecule
                # print replacement
                replacement = molecule[:match_index] + conversion + molecule[match_index+len(key):]
                replacements.append(replacement)
    return replacements

def parse(reverse=False):
    with open('input.txt', 'r') as f:
        data = f.readlines()

    conversions = {}
    molecule = ""
    r = re.compile(r'(\w+) => (\w+)')
    for conversion_or_molecule in data:
        m = re.match(r, conversion_or_molecule)
        if m:
            if reverse:
                value, key = m.groups()
            else:
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
    replacements = generate_replacements(conversions, molecule)
    print len(set(replacements))

def parttwo():
    conversions, molecule = parse(reverse=True)
    print sorted(conversions, cmp=lambda x,y: 1 if len(x) < len(y) else -1)

if __name__ == '__main__':
    test()
    # partone()
    parttwo()
