import utils

raw = utils.get_input(4)
#raw = utils.get_input_test(4)

grid = [list(line) for line in raw.splitlines()]

def get_grid(grid, x, y):
    if y >= len(grid) or y < 0:
        return None
    line = grid[y]
    if x >= len(line) or x < 0:
        return None
    return line[x]

def check_adjacent(grid, x,y):
    r = [
            get_grid(grid, x-1, y) == ROLL, # left
            get_grid(grid, x+1, y) == ROLL, # right
            get_grid(grid, x, y-1) == ROLL, # up
            get_grid(grid, x, y+1) == ROLL, # down
            get_grid(grid, x-1, y-1) == ROLL, #up left
            get_grid(grid, x+1, y-1) == ROLL, #up right
            get_grid(grid, x-1, y+1) == ROLL, #down left
            get_grid(grid, x+1, y+1) == ROLL, #down right
    ]
    return r

ROLL = "@"

result = 0
for y, line in enumerate(grid):
    for x, symbol in enumerate(line):
        if symbol != ROLL: continue
        r = check_adjacent(grid, x,y)
        if sum(r) < 4:
            result += 1
print(result)

changed = True
removed = 0
while changed:
    changed = False
    for y, line in enumerate(grid):
        for x, symbol in enumerate(line):
            if symbol != ROLL: continue
            r = check_adjacent(grid, x,y)
            if sum(r) < 4:
                removed += 1
                changed = True
                grid[y][x] = "x"

print(removed)
