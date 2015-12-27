#!/bin/python
import re
import time
import sys
import random

def generate_replacements(conversions, molecule, step_two=False):
    replacements = []
    for key in conversions: 
        for match_index in [m.start() for m in re.finditer(key, molecule)]:
            for conversion in conversions[key]:
                replacement = molecule[:match_index] + conversion + molecule[match_index+len(key):]
                if not replacement in replacements:
                    replacements.append(replacement)
    return replacements

def parse_to_tuple(test=False):
    text_file = 'test_input.txt' if test else 'input.txt'
    with open(text_file, 'r') as f:
        data = f.readlines()

    conversions = []
    r = re.compile(r'(\w+) => (\w+)')
    for conversion_or_molecule in data:
        m = re.match(r, conversion_or_molecule)
        if m:
            key, value = m.groups()
            if not key in conversions:
                conversions[key] = []
            conversions.append((key, value))
        else:
            molecule = conversion_or_molecule.strip()

    return conversions, molecule

def parse(reverse=False, test=False):
    text_file = 'test_input.txt' if test else 'input.txt'
    with open(text_file, 'r') as f:
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
            conversions[key] = value
        else:
            molecule = conversion_or_molecule.strip()

    return conversions, molecule

def test():
    conversions, molecule = parse(reverse=True, test=True)
    # assert(steps_to_target_molecule(conversions, seed_molecule=molecule, target_molecule='e', seen=set()) == 3)
    # assert(steps_to_target_molecule(conversions, seed_molecule='HOHOHO', target_molecule='e', seen=set()) == 6)
    print steps_to_target_molecule_recurse(conversions, seed_molecule=molecule, target_molecule='e', seen=set())
    assert(steps_to_target_molecule_recurse(conversions, seed_molecule=molecule, target_molecule='e', seen=set()) == 3)
    assert(steps_to_target_molecule_recurse(conversions, seed_molecule='HOHOHO', target_molecule='e', seen=set()) == 6)
    print "all tests pass!"

def partone():
    conversions, molecule = parse()
    print conversions
    print molecule
    replacements = generate_replacements(conversions, molecule)
    print len(set(replacements))

def parttwo():
    conversions, molecule = parse_to_tuple()
    count = 0
    mol = molecule
    shuffle = 0
    print conversions
    while len(mol) > 1:
        start = mol
        for replacement, to_replace in conversions:
            while to_replace in mol:
                count += mol.count(to_replace)
                mol = mol.replace(to_replace, replacement)

        if start == mol:
            random.shuffle(conversions)
            mol = molecule
            count  = 0
            shuffle += 1
            print shuffle

    print count

if __name__ == '__main__':
    # test()
    # partone()
    parttwo()
