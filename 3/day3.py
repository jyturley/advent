#!/bin/python
from sets import Set

def track_houses(input):
    visited_houses = Set([])
    current_location = [0,0]
    visited_houses.add(tuple(current_location))

    for house in input:
        if house == "^":
            current_location[1] += 1
        if house == ">":
            current_location[0] += 1
        if house == "v":
            current_location[1] -= 1
        if house == "<":
            current_location[0] -= 1

        visited_houses.add(tuple(current_location))

    return len(visited_houses)
    

def test():
    print track_houses(">")
    print track_houses("^>v<")
    print track_houses("^v^v^v^v^v")


def partone():
    with open('input.txt', 'r') as f:
        data = f.read()
        print track_houses(data)

def parttwo():
    pass

if __name__ == '__main__':
    test()
    partone()
