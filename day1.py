#!/bin/python

import os

def find_floor(input):
    current_floor = 0
    for paren in input:
        if paren == "(":
            current_floor += 1
        elif paren == ")": 
            current_floor -= 1
        else:
            continue

    return current_floor

def main():
    print "day 1 "
    with open('input.txt', 'r') as f:
        data = f.read()
        print data
        print find_floor(data)

if __name__ == '__main__':
    main()
