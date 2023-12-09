with open("inputDay9.txt") as f:
    file = []
    for i in f:
        file.append(int(i.replace("\n", "")))


# print(file)
def checkValid():
    for i in range(len(file)):
        valid = False
        if i > 25:
            for n in file[i - 25:i]:
                for x in file[i - 25:i]:
                    if n + x == file[i]:
                        valid = True
            if not valid:
                return file[i]


def func2(number):
    i = 0
    start = 0
    range = 0
    while i < len(file):
        range += file[i]
        if range > number:
            range = 0
            start += 1
            i = start
        elif range == number:
            return file[start:i+1]
        else:
            i += 1


invalid = checkValid()
print("Invalid number:", invalid)
invalid_range = func2(invalid)
print("Min+Max of invalid Range:", max(invalid_range)+min(invalid_range))
