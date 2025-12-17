import utils
from PIL import Image, ImageDraw
import av
from tqdm import tqdm


def get_color_gradient(value, max_value):
    ratio = value / max_value
    red = int(255 * ratio)
    green = int(255 * (1 - ratio))
    blue = 0
    return (red, green, blue)

def make_frame(grid):
    dimension = len(grid)
    scale = 4
    img = Image.new('RGB', (dimension*scale, dimension*scale), color = 'white')
    d = ImageDraw.Draw(img)
    for y, line in enumerate(grid):
        for x, symbol in enumerate(line):
            if isinstance(symbol, int) and symbol > 0:
                max_val = max([val for val in line if isinstance(val, int)] + [1])
                color = get_color_gradient(symbol, max_val)
                for dy in range(scale):
                    for dx in range(scale):
                        d.point((x*scale+dx, y*scale+dy), fill=color)
    return img

def m():
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
    return max(grid[-1])

if __name__ == "__main__":
    DAY = 7
    raw = utils.get_input_lines(DAY)
    frames = []
    max_val = m()
    grid = [list(line) for line in raw]
    for y in range(len(grid)):
        for x in range(len(grid[y])):
            if grid[y][x] == ".":
                grid[y][x] = 0
            elif grid[y][x] == "S":
                grid[y][x] = 1

    for y in range(len(grid)):
        frames.append(make_frame(grid))
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

    file_name = f"D{str(DAY).zfill(2)}.mp4"
    output = av.open(file_name, 'w')
    stream = output.add_stream('h264', 12)
    stream.width = frames[0].width
    stream.height = frames[0].height
    for img in tqdm(frames, desc="Making video"):
        frame = av.VideoFrame.from_image(img)
        packet = stream.encode(frame)
        output.mux(packet)
    packet = stream.encode(None)
    output.mux(packet)
    output.close()
    print(f"Video saved to {file_name}")
