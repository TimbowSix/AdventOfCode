def switchRowsCols(arr: list) -> list:
    """Takes any 2x2 List, switches Columns to Rows"""
    cols = []
    for i in range(len(arr[0])):
        col = []
        for row in arr:
            col.append(row[i])
        cols.append(col)
    return cols

def binaryToDecimal(binList) -> int:
    """Takes an iterable of 0 and 1 and converts it into decimal Number"""
    sum = 0
    exp = len(binList)-1
    for bin in binList:
        bin = int(bin)
        sum += bin*2**exp
        exp -= 1
    return sum

def getFile(path: str) -> list:
    """Opens and reads a .txt or similar file and splits its lines"""
    with open(path, "r") as f:
        file = f.read().splitlines()
    return file
