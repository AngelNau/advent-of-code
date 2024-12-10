import re

def rek()

with open('message.txt', 'r') as file:
    sum = 0
    maze = ""
    for line in file:
        line = line.strip()
        maze += line
    for char in maze:
        if char == 'X':
            sum += rek()
    print(sum)
