import re
from functools import reduce
from operator import add, mul



def solve(lines):
    if not lines:
        return 0


    width = max(len(line) for line in lines)
    grid = [line.rstrip('\n').ljust(width) for line in lines]
    h = len(grid)

  
    sep = []
    for c in range(width):
        if all(grid[r][c] == ' ' for r in range(h)):
            sep.append(True)
        else:
            sep.append(False)

    problem_ranges = []
    in_problem = False
    start = 0
    for c in range(width):
        if not sep[c]:
            if not in_problem:
                in_problem = True
                start = c
        else:
            if in_problem:
                in_problem = False
                problem_ranges.append((start, c - 1))
    if in_problem:
        problem_ranges.append((start, width - 1))

    grand_total = 0

    for c_start, c_end in problem_ranges:
 
        op_row = None
        for r in range(h):
            segment = grid[r][c_start:c_end + 1]
            if '+' in segment or '*' in segment:
                op_row = r

        if op_row is None:
            continue


        op_segment = grid[op_row][c_start:c_end + 1]
        if '+' in op_segment and '*' in op_segment:
   
            plus_idx = op_segment.find('+') if '+' in op_segment else float('inf')
            mul_idx = op_segment.find('*') if '*' in op_segment else float('inf')
            op = add if plus_idx < mul_idx else mul
        elif '+' in op_segment:
            op = add
        elif '*' in op_segment:
            op = mul
        else:
            continue


        numbers = []
        for c in range(c_start, c_end + 1):
            digits = []
            for r in range(op_row):
                ch = grid[r][c]
                if ch.isdigit():
                    digits.append(ch)
            if digits:
                num = int(''.join(digits))
                numbers.append(num)

        if not numbers:
            continue

        result = reduce(op, numbers)
        grand_total += result

    return grand_total

lines = []
total_items = 0

with open("input.txt", "r") as file:
    for line in file:
        lines.append(line)

print(solve(lines))