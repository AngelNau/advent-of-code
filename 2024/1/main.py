
with open('inputsmall.txt', 'r') as file:
    # Read each line in the file
    sum = 0
    listL = []
    listR = []
    for line in file:
        # Print each line
        line = line.strip()
        x = line.split("   ")
        listL.append(int(x[0]))
        listR.append(int(x[1]))
    # listL.sort()
    # listR.sort()
    # for i, el in enumerate(listL):
    #     sum += abs(listL[i] - listR[i])
    for i, el in enumerate(listL):
        sum += el * listR.count(el)
    print(sum)
