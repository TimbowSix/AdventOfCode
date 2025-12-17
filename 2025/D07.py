import utils

DAY = 7

raw = utils.get_input_lines(DAY)
#raw = utils.get_input_lines_test(DAY)

grid = [list(line) for line in raw]

result = 0
for y in range(len(grid)):
    for x in range(len(grid[y])):
        symbol = grid[y][x]
        if not(symbol == "S" or symbol == "|"): continue
        if y+1 >= len(grid): break
        below_symbol = grid[y+1][x]
        if below_symbol == "^":
            result +=1
            grid[y+1][x-1] = "|"
            grid[y+1][x+1] = "|"
        else:
            grid[y+1][x] = "|"

print(result)


grid = [list(line) for line in raw]
for y in range(len(grid)):
    for x in range(len(grid[y])):
        if grid[y][x] == ".":
            grid[y][x] = 0
        elif grid[y][x] == "S":
            grid[y][x] = 1

for y in range(len(grid)):
    for x in range(len(grid[y])):
        symbol = grid[y][x]
        if symbol == 0 or symbol == "^": continue
        if y+1 >= len(grid): break
        below_symbol = grid[y+1][x]
        if below_symbol == "^":
            grid[y+1][x-1] += symbol
            grid[y+1][x+1] += symbol
        else:
            grid[y+1][x] += symbol

# for i in grid:
#     print(["." if s == 0 else str(s) for s in i])

print(sum(grid[-1]))
