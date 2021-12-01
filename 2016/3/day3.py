#!/bin/python
from sets import Set

def track_houses_with_robo_santa(input):
    visited_houses = Set([])
    santa_location      = [0,0]
    robo_santa_location = [0,0]
    visited_houses.add(tuple(santa_location))
    santas_turn = True

    for house in input:
        if house == "^":
            if santas_turn:
                santa_location[1] += 1
            else:
                robo_santa_location[1] += 1

        if house == ">":
            if santas_turn:
                santa_location[0] += 1
            else:
                robo_santa_location[0] += 1
        if house == "v":
            if santas_turn:
                santa_location[1] -= 1
            else:
                robo_santa_location[1] -= 1
        if house == "<":
            if santas_turn:
                santa_location[0] -= 1
            else:
                robo_santa_location[0] -= 1

        if santas_turn:
            visited_houses.add(tuple(santa_location))
        else:
            visited_houses.add(tuple(robo_santa_location))

        santas_turn = not santas_turn

    return len(visited_houses)


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

    print track_houses_with_robo_santa("^>")
    print track_houses_with_robo_santa("^>v<")
    print track_houses_with_robo_santa("^v^v^v^v^v")

def partone():
    with open('input.txt', 'r') as f:
        data = f.read()
        print track_houses(data)

def parttwo():
    with open('input.txt', 'r') as f:
        data = f.read()
        print track_houses_with_robo_santa(data)

if __name__ == '__main__':
    test()
    parttwo()
