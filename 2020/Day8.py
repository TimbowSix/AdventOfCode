import copy

with open("inputDay8.txt") as f:
    file = []
    for i in f:
        file.append(i.replace("\n", "").split())
# print(file)


def boot(file):
    acc = 0
    i = 0
    seenI = []
    broke = False
    while i < len(file):
        # print(i, file[i])
        if i in seenI:
            # print("Last index:", seenI[len(seenI)-1])
            # print("Last command:", file[seenI[len(seenI)-1]])
            # print("First double:", i)
            # print("First double:", file[i])
            # print("Double Index", seenI.index(i))
            # print("Double Index-1", seenI[seenI.index(i)-1])
            # print("Double Index-1", file[seenI[seenI.index(i)-1]])
            broke = True
            break
        elif file[i][0] == "acc":
            acc += int(file[i][1])
            seenI.append(i)
            i += 1
        elif file[i][0] == "nop":
            seenI.append(i)
            i += 1
        elif file[i][0] == "jmp":
            seenI.append(i)
            i += int(file[i][1])
    return acc, broke


print(boot(file)[0])


def getFix(file):
    for i in range(len(file)):
        newFile = copy.deepcopy(file)
        if newFile[i][0] == "jmp":
            newFile[i][0] = "nop"
        elif newFile[i][0] == "nop":
            newFile[i][0] = "jmp"

        # print(i, boot(newFile)[0], newFile[i], ":", file[i])
        test = boot(newFile)
        if not test[1]:
            print(i, test[0])


getFix(file)