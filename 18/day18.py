#!/bin/python
import re

def is_on_for_next_frame(grid, x, y):
	num_neighbors_on = 0
	max_index = len(grid) - 1
	if x != 0:
		num_neighbors_on += grid[x-1][y]
		if y != 0:
			num_neighbors_on += grid[x-1][y-1]
		if y != max_index:
			num_neighbors_on += grid[x-1][y+1]

	if x != max_index:
		num_neighbors_on += grid[x+1][y]
		if y != 0:
			num_neighbors_on += grid[x+1][y-1]
		if y != max_index:	
			num_neighbors_on += grid[x+1][y+1]

	if y != 0:
		num_neighbors_on += grid[x][y-1]

	if y != max_index:
		num_neighbors_on += grid[x][y+1]

	if grid[x][y]:
		# true iff 2 or 3 neighbors on
		return num_neighbors_on == 2 or num_neighbors_on ==3
	else:
		# true iff 3 neighbors on
		print num_neighbors_on == 3

def get_next_frame(grid):
	next_frame = []
	for x in xrange(len(grid)):
		next_frame.append([])
		for y in xrange(len(grid[x])):
			next_frame[x].append(1 if is_on_for_next_frame(grid, x, y) else 0)

	return next_frame

def parse():
	grid = []
	with open('input.txt', 'r') as f:
		data = f.readlines()

	for x, row in enumerate(data):
		print row, x
		grid.append([])
		for y, cell in enumerate(row):
			grid[x].append(1 if cell is '#' else 0)
		print grid[x]

	return grid

def test():
	grid = [[0,1,0,1,0,1], [0,0,0,1,1,0], [1,0,0,0,0,1], [0,0,1,0,0,0], [1,0,1,0,0,1], [1,1,1,1,0,0]]
	print get_next_frame(grid)
	print "all tests pass!"

def partone():
	grid =  parse()
	pass

def parttwo():
	pass

if __name__ == '__main__':
    test()
    # partone()
    # parttwo()
