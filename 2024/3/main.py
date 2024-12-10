import re

with open('message.txt', 'r') as file:
    sum = 0
    linee = ""
    for line in file:
        line = line.strip()
        linee += line
    muls = re.sub(r"don't\(\).*?do\(\)", "", linee)
    muls1 = re.findall("mul\(([1-9]+[0-9]{0,2},[1-9]+[0-9]{0,2})\)", muls)
    for i in muls1:
        nums = i.split(",")
        sum += int(nums[0]) * int(nums[1])
    print(sum)
