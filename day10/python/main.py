import numpy as np
from scipy.optimize import linprog


with open("input") as f:
    ls = f.read().strip().split("\n")

tasks = []
for l in ls:
    toggles, *buttons, counters = l.split()
    toggles = [x == "#" for x in toggles[1:-1]]
    moves = [set(map(int, b[1:-1].split(","))) for b in buttons]
    counters = list(map(int, counters[1:-1].split(",")))
    tasks.append((toggles, moves, counters))


def solve(goal, moves, part1):
    n, m = len(moves), len(goal)
    c = [1] * n
    A_eq = [[i in move for move in moves] for i in range(m)]
    bounds = [(0, None)] * n
    if part1:
        c += [0] * m
        A_eq = np.hstack([A_eq, -2 * np.eye(m)])
        bounds += [(None, None)] * m
    return linprog(c, A_eq=A_eq, b_eq=goal, bounds=bounds, integrality=True).fun


# Part 1
print(sum(solve(goal, moves, True) for goal, moves, _ in tasks))

# Part 2
print(sum(solve(goal, moves, False) for _, moves, goal in tasks))