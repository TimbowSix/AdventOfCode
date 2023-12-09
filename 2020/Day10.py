with open("inputDay10.txt") as f:
    file = []
    for i in f:
        file.append(int(i.replace("\n", "")))

file.sort()


# print(file)


def getDiffs(file):
    file.sort()
    final = file[len(file) - 1] + 3
    start = 0
    difs = {1: 0, 2: 0, 3: 0}
    idx = 0
    dif = file[0] - start
    difs[dif] = difs[dif] + 1
    while idx < len(file):
        if idx == len(file) - 1:
            dif = final - file[idx]
        else:
            dif = file[idx + 1] - file[idx]
        idx += 1
        difs[dif] = difs[dif] + 1
    return difs


difs = getDiffs(file)
# print(difs)
print("Solution part 1:", difs[1] * difs[3])


# leider nicht meine eigene lösung :(

file = [0]+file
file.append(file[-1]+3)

poss = [1]+[0]*file[-1]
for i in file[1:]:
    poss[i] = poss[i-1] + poss[i-2] + poss[i-3]

print("Solution part 2:", poss[-1])
