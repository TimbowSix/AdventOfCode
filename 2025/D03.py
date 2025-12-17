import utils

raw = utils.get_input_lines(3)
#raw = utils.get_input_lines_test(3)

joltages = []
for bank in raw:
    biggest_first = 0
    biggest_ind = 0
    for ind, bat in enumerate(bank):
        bat = int(bat)
        if bat > biggest_first and ind < len(bank)-1:
            biggest_first = bat
            biggest_ind = ind

    biggest_second = 0
    for bat in bank[biggest_ind+1:]:
        bat = int(bat)
        if bat > biggest_second:
            biggest_second = bat
    joltages.append(int(f"{biggest_first}{biggest_second}"))

print(sum(joltages))

def find_biggest(lst: list[int]) -> tuple[int, int]:
    biggest = 0
    biggest_ind = 0
    for ind, bat in enumerate(lst):
        bat = int(bat)
        if bat > biggest:
            biggest = bat
            biggest_ind = ind
    return biggest, biggest_ind

joltages = []
for bank in raw:
    batts = []
    ind = 0
    for i in range(12):
        upper = -(12-i)+1
        bank_part = bank[ind: upper] if upper < 0 else bank[ind:]
        bat, biggest_ind = find_biggest(bank_part)
        batts.append(bat)
        #print(bat, biggest_ind, ind, upper)
        ind += biggest_ind+1
    joltages.append(int("".join([str(b) for b in batts])))

#print(joltages)
print(sum(joltages))

