import utils

lines = utils.get_input_lines(1)
#lines = utils.get_input_lines_test(1)

instructions: list[int] = []

for line in lines:
    if line.startswith("R"):
        value = int(line[1:])
        instructions.append(value)
    elif line.startswith("L"):
        value = -int(line[1:]) 
        instructions.append(value)

dial = 50
MAX_DIAL = 99

zeros = 0
for instruction in instructions:
    #print(dial, instruction)
    instruction = instruction % (MAX_DIAL+1)

    dial = dial + instruction
    #print(dial)
    if dial < 0:
        dial = (MAX_DIAL+1) + dial
    elif dial > MAX_DIAL:
        dial = dial % (MAX_DIAL+1)
    if dial == 0:
        zeros += 1
    #print(dial, instruction)

print(zeros)


dial = 50
zeros = 0
for instruction in instructions:
    n = instruction
    while n != 0:
        if n < 0:
            dial -= 1
            n += 1
        else:
            dial += 1
            n -= 1
        if dial > MAX_DIAL:
            dial = 0
        elif dial < 0:
            dial = MAX_DIAL
        if dial == 0:
            zeros += 1

print(zeros)

