import copy
with open("inputDay11.txt") as f:
    file = []
    for i in f:
        file.append(i.replace("\n", ""))
# print(file)


# [y][x]
def countOcc(x, y, layout):
    nearby = [[y, x + 1], [y, x - 1], [y + 1, x], [y - 1, x], [y - 1, x - 1], [y - 1, x + 1], [y + 1, x - 1],
              [y + 1, x + 1]]
    occs = 0
    for i in nearby:
        if i[0] < len(layout) and i[1] < len(layout[i[0]]):
            if layout[i[0]][i[1]] == "#":
                occs += 1
    return occs


before = []
working = copy.deepcopy(file)
while before != file:
    before = copy.deepcopy(working)
    y = 0
    while y < len(file):
        x = 0
        row = list(file[y])
        while x < len(row):
            occs = countOcc(x, y, file)
            if row[x] == "L" and occs == 0:
                row[x] = "#"
                working[y] = "".join(row)
            elif row[x] == "#" and occs >= 4:
                row[x] = "L"
                working[y] = "".join(row)
            x += 1
        y += 1
    file = copy.deepcopy(working)


seats = 0
for i in file:
    seats += i.count("#")

print(seats)