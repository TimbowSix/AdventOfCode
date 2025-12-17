import utils

DAY = 6

raw = utils.get_input_lines(DAY)
#raw = utils.get_input_lines_test(DAY)

grid = [line.split() for line in raw]

problem_count = len(grid[0])

total = 0
for x in range(problem_count):
    numbers = []
    operation = None
    for line in grid:
        if line[x].isdigit():
            numbers.append(int(line[x]))
        else:
            operation = line[x]
    if operation is None:
        raise ValueError("Hier stimmt was nicht Kollege")

    if operation == "+":
        total += sum(numbers)
    elif operation == "*":
        n = 1
        for i in numbers:
            n *= i
        total += n

print(total)

total = 0
numbers = []
operation = "+"
for x in range(len(raw[-1])):
    if raw[-1][x] in ["+", "*"]: # new number col
        #print(numbers)
        if operation == "+":
            total += sum(numbers)
        if operation == "*":
            n = 1
            for i in numbers:
                n *= i
            total += n
        operation = raw[-1][x]
        numbers = []
    num = ""
    for y in range(len(grid)-1):
        if x >= len(raw[y]): continue
        num += raw[y][x]
    num = num.strip()
    if num: numbers.append(int(num))

#print(numbers)
if operation == "+":
    total += sum(numbers)
if operation == "*":
    n = 1
    for i in numbers:
        n *= i
    total += n
print(total)
