def aaa(l):
    return all(abs(int(l[i]) - int(l[i + 1])) in range(1, 4) for i in range(len(l) - 1))

def asc(l):
    return all(int(l[i]) < int(l[i+1]) for i in range(len(l) - 1))

def desc(l):
    return all(int(l[i]) > int(l[i+1]) for i in range(len(l) - 1))

def generate_sublists(lst):
    """Generates sublists by removing one element at a time."""
    return [lst[:i] + lst[i+1:] for i in range(len(lst))]

with open('input.txt', 'r') as file:
    sum = 0
    for line in file:
        line = line.strip()
        x = line.split(" ")
        if(aaa(x)):
            if asc(x):
                sum += 1
                continue
            if desc(x):
                sum += 1
                continue
        sublists = generate_sublists(x)
        for lst in sublists:
            if(aaa(lst)):
                if asc(lst):
                    sum += 1
                    break
                if desc(lst):
                    sum += 1
                    break
    print(sum)
