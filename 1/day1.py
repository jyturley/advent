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

def entered_basement(floor):
    return floor == -1

def find_position_of_basement(input):
    current_floor = 0
    current_position = 0
    for paren in input:
        current_position += 1
        if paren == "(":
            current_floor += 1

        elif paren == ")": 
            current_floor -= 1
            if entered_basement(current_floor):
                return current_position
        else:
            continue

    return 0 # no basement reached


def main():
    with open('input.txt', 'r') as f:
        data = f.read()
        # print find_floor(data)
        print find_position_of_basement(data)

if __name__ == '__main__':
    main()
