import utils
from PIL import Image, ImageDraw
import av
from tqdm import tqdm
import multiprocessing
import tempfile


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

def make_frame(grid):
    dimension = len(grid)
    scale = 4
    img = Image.new('RGB', (dimension*scale, dimension*scale), color = 'white')
    d = ImageDraw.Draw(img)
    for y, line in enumerate(grid):
        for x, symbol in enumerate(line):
            if symbol == ROLL:
                for dy in range(scale):
                    for dx in range(scale):
                        d.point((x*scale+dx, y*scale+dy), fill='black')
    return img

if __name__ == "__main__":
    raw = utils.get_input(4)
    #raw = utils.get_input_test(4)

    grid = [list(line) for line in raw.splitlines()]
    changed = True
    removed = 0
    grids = []
    while changed:
        changed = False
        for y, line in enumerate(grid):
            grids.append([row.copy() for row in grid])
            for x, symbol in enumerate(line):
                if symbol != ROLL: continue
                r = check_adjacent(grid, x,y)
                if sum(r) < 4:
                    removed += 1
                    changed = True
                    grid[y][x] = "x"

    print(removed)


    with multiprocessing.Pool() as pool:
        frames = list(tqdm(pool.imap(make_frame, grids), total=len(grids), desc="Generating frames"))


    output = av.open("D04.mp4", 'w')
    stream = output.add_stream('h264', 180)
    stream.width = frames[0].width
    stream.height = frames[0].height
    for img in tqdm(frames, desc="Making video"):
        frame = av.VideoFrame.from_image(img)
        packet = stream.encode(frame)
        output.mux(packet)
    packet = stream.encode(None)
    output.mux(packet)
    output.close()
