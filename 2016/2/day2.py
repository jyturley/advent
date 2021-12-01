#!/bin/python
import os

def find_dimensions(present):
    dimensions = present.split("x")
    dimensions = [int(d) for d in dimensions]
    assert(len(dimensions) == 3)
    l, w, h = dimensions

    surface_area = (2 * l * w) + (2 * w * h) + (2 * h * l)
    dimensions_copy = list(dimensions)
    dimensions.remove(min(dimensions))
    slack = min(dimensions) * min(dimensions_copy)

    return surface_area + slack

def find_ribbon_length(present):
    dimensions = present.split("x")
    dimensions = [int(d) for d in dimensions]
    assert(len(dimensions) == 3)
    l, w, h = dimensions

    dimensions_copy = list(dimensions)
    dimensions.remove(min(dimensions))

    length_for_present = 2 * min(dimensions) + 2 * min(dimensions_copy)
    length_for_bow = l * w * h

    return length_for_present + length_for_bow


def test():
    print find_ribbon_length("2x3x4")
    print find_ribbon_length("1x1x10")


def partone():
    total_wrapping_paper = 0
    with open('input.txt', 'r') as f:
        data = f.readlines()
        for present in data: 
            total_wrapping_paper += find_dimensions(present)

        print total_wrapping_paper

def parttwo():
    total_ribbon = 0
    with open('input.txt', 'r') as f:
        data = f.readlines()
        for present in data:
            total_ribbon += find_ribbon_length(present)

        print total_ribbon

if __name__ == '__main__':
    test()
    parttwo()
