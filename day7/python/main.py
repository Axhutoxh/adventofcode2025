grid = [list(line) for line in open("input.txt", "r").read().splitlines()]

rows = len(grid)
cols = len(grid[0])

start_col = grid[0].index("S")

visited = set()
count = 0

def traverse(row, col):
    global count

    if row < 0 or row >= rows or col < 0 or col >= cols:
        return

    if (row, col) in visited:
        return

    visited.add((row, col))

    current_cell = grid[row][col]

    if current_cell == "^":
        count += 1
        traverse(row + 1, col - 1)
        traverse(row + 1, col + 1)
    elif current_cell == ".":
        traverse(row + 1, col)
        
traverse(1, start_col) 

print(count)
