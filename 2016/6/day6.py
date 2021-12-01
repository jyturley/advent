#!/bin/python
import re

def set_grid_pt2(start_x, start_y, end_x, end_y, value, grid):
    for x in xrange(end_x - start_x + 1):
        for y in xrange(end_y - start_y + 1):
            current_x = start_x + x
            current_y = start_y + y
            grid[current_x][current_y] += value
            if grid[current_x][current_y] < 0:
                grid[current_x][current_y] = 0


def parse_and_execute_instruction_pt2(instruction, grid):
    m = re.search("(\d+,\d+) through (\d+,\d+)", instruction)
    assert(m and len(m.groups()) == 2)

    start_x, start_y = [int(val) for val in m.group(1).split(',')]
    end_x,   end_y   = [int(val) for val in m.group(2).split(',')]  

    if instruction.startswith("turn on"):
        set_grid_pt2(start_x, start_y, end_x, end_y, 1, grid)

    if instruction.startswith("turn off"):
        set_grid_pt2(start_x, start_y, end_x, end_y, -1, grid)

    if instruction.startswith("toggle"):
        set_grid_pt2(start_x, start_y, end_x, end_y, 2, grid)

    return

def count_on_lights(grid):
    sofar = 0
    for row in grid:
        for light in row:
            sofar += light
    return sofar

def set_grid_pt1(start_x, start_y, end_x, end_y, value, grid):
    for x in xrange(end_x - start_x + 1):
        for y in xrange(end_y - start_y + 1):
            grid[start_x + x][start_y + y] = value

def toggle_grid_pt1(start_x, start_y, end_x, end_y, grid):
    for x in xrange(end_x - start_x + 1):
        for y in xrange(end_y - start_y + 1):
            current_x = start_x + x
            current_y = start_y + y
            grid[current_x][current_y] = 0 if grid[current_x][current_y] == 1 else 1

def parse_and_execute_instruction_pt1(instruction, grid):
    m = re.search("(\d+,\d+) through (\d+,\d+)", instruction)
    assert(m and len(m.groups()) == 2)

    start_x, start_y = [int(val) for val in m.group(1).split(',')]
    end_x,   end_y   = [int(val) for val in m.group(2).split(',')]  

    if instruction.startswith("turn on"):
        set_grid_pt1(start_x, start_y, end_x, end_y, 1, grid)

    if instruction.startswith("turn off"):
        set_grid_pt1(start_x, start_y, end_x, end_y, 0, grid)

    if instruction.startswith("toggle"):
        toggle_grid_pt1(start_x, start_y, end_x, end_y, grid)

    return

def test():
    grid = [[0 for x in xrange(10)] for x in xrange(10)]
    parse_and_execute_instruction_pt1("turn on 0,0 through 2,2", grid)
    assert( count_on_lights(grid) == 9 )

    grid = [[0 for x in xrange(1000)] for x in xrange(1000)]
    parse_and_execute_instruction_pt1("toggle 0,0 through 999,0", grid)
    assert( count_on_lights(grid) == 1000 )
    parse_and_execute_instruction_pt1("toggle 0,0 through 999,0", grid)
    assert( count_on_lights(grid) == 0 )

    grid = [[0 for x in xrange(1000)] for x in xrange(1000)]
    parse_and_execute_instruction_pt1("turn on 499,499 through 500,500", grid)
    assert( count_on_lights(grid) == 4 )
    print "all tests pass!"


def partone():
    grid = [[0 for x in xrange(1000)] for x in xrange(1000)]
    with open('input.txt', 'r') as f:
        for instruction in f.readlines():
            parse_and_execute_instruction_pt1(instruction, grid)
    print count_on_lights(grid)


def parttwo():
    grid = [[0 for x in xrange(1000)] for x in xrange(1000)]
    with open('input.txt', 'r') as f:
        for instruction in f.readlines():
            parse_and_execute_instruction_pt2(instruction, grid)
    print count_on_lights(grid)

if __name__ == '__main__':
    # test()
    # partone()
    parttwo()
