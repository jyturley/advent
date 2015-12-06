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

def test():
    print find_dimensions("2x3x4")
    print find_dimensions("1x1x10")


def main():
    total_wrapping_paper = 0
    with open('input.txt', 'r') as f:
        data = f.readlines()
        for present in data: 
            total_wrapping_paper += find_dimensions(present)

        print total_wrapping_paper



if __name__ == '__main__':
    # test()
    main()
