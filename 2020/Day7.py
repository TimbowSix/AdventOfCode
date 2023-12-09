with open("inputDay7.txt") as f:
    file = []
    for i in f:
        file.append(i.replace("\n", "").split())
# print(file)
bags = {}
for i in file:
    numberOfContainingBags = int(len(i) / 4 - 1)
    # print(numberOfContainingBags)
    containingBags = []
    # 5:6, 9:10, 13:14
    pos = 5
    for n in range(numberOfContainingBags):
        containingBags.append(i[pos] + i[pos + 1])
        pos += 4

    bags[i[0] + i[1]] = containingBags


# print(bags)
# own: shinygold


def isPossible(bag):
    if "shinygold" in bags[bag]:
        return True
    else:
        for j in bags[bag]:
            if isPossible(j):
                return True


possibleBags = 0
for i in bags:
    if isPossible(i):
        possibleBags += 1

print("number of possible bags:", possibleBags)
bags_w_numbers = {}
for i in file:
    numberOfContainingBags = int(len(i) / 4 - 1)
    # print(numberOfContainingBags)
    containingBags = {}
    # 5:6, 9:10, 13:14
    pos = 5
    for n in range(numberOfContainingBags):
        containingBags[i[pos] + i[pos + 1]] = i[pos - 1]
        pos += 4

    bags_w_numbers[i[0] + i[1]] = containingBags
# print(bags_w_numbers)


def getBags(bag):
    totals = 0
    for x in bags_w_numbers[bag]:
        totals += int(bags_w_numbers[bag][x])
        totals += getBags(x)*int(bags_w_numbers[bag][x])
    return totals


print("total bags:", getBags("shinygold"))
