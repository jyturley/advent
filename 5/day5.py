#!/bin/python

def is_nice(string):
    vowels = set(["a", "e", "i", "o", "u"])
    vowel_count = 0
    double_letter_pass = False
    previous_char = ""

    if "ab" in string: return False
    if "cd" in string: return False
    if "pq" in string: return False
    if "xy" in string: return False

    for char in string:
        if vowel_count < 3 and char in vowels:
            vowel_count += 1
        if double_letter_pass == False and char == previous_char:
            double_letter_pass = True

        previous_char = char

    return double_letter_pass and vowel_count >= 3

def test():
    assert(is_nice("ugknbfddgicrmopn") == True)
    assert(is_nice("aaa") == True)
    assert(is_nice("jchzalrnumimnmhp") == False)
    assert(is_nice("haegwjzuvuyypxyu") == False)
    assert(is_nice("dvszwmarrgswjxmb") == False)
    print "all tests pass!"

def partone():
    nice_count = 0
    with open('input.txt', 'r') as f:
        data = f.readlines()
        for string in data:
            nice_count += is_nice(string)

    print nice_count


def parttwo():
    pass

if __name__ == '__main__':
    test()
    partone()
    # parttwo()
