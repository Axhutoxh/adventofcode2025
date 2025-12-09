fresh_list = []
total_items = 0

with open("input.txt", "r") as file:
    for line in file:
        if len(line) == 1 and line[0] == "\n":
            break
        ranges = line.split("-")
        lower, upper = int(ranges[0]), int(ranges[1])
        fresh_list.append((lower, upper))

fresh_list = sorted(fresh_list, key=lambda x: x[0])

for idx in range(0, len(fresh_list) - 1):
    print("++++++++++++++++++++++++++++++++++++")
    (cmp_lower, cmp_upper) = fresh_list[idx]
    if cmp_lower == -1 or cmp_upper == -1:
        print("[SKIP] comparator")
        continue
    for inner_idx in range(idx + 1, len(fresh_list)):
        print("------------------------------------")
        (cur_lower, cur_upper) = fresh_list[inner_idx]
        mn = min(
            [
                fresh_list[idx][0],
                fresh_list[idx][1],
                fresh_list[inner_idx][0],
                fresh_list[inner_idx][1],
            ]
        )
        mx = max(
            [
                fresh_list[idx][0],
                fresh_list[idx][1],
                fresh_list[inner_idx][0],
                fresh_list[inner_idx][1],
            ]
        )
        print(
            f"idx: {idx}, cmp_lower: {cmp_lower}, cmp_upper: {cmp_upper}, inner_idx: {inner_idx}, cur_lower: {cur_lower}, cur_upper: {cur_upper} (mn: {mn}, mx: {mx})",
            end="",
        )
        if cur_lower == -1 or cur_upper == -1:
            print("[SKIP] current")
            continue
        #   Lcmp-----------Ucmp
        #       Lcur------------Ucur

        #   Lcur-----------Ucur
        #       Lcmp------------Ucmp
        # if (cmp_lower <= cur_lower and cur_lower >= cmp_upper) or (
        #    cur_lower <= cmp_lower and cmp_lower >= cur_upper
        # ):
        #    print("[COND1]")
        #    fresh_list[idx] = (mn, mx)
        #    fresh_list[inner_idx] = (-1, -1)
        #    continue

        if cmp_lower <= cur_lower and cur_lower <= cmp_upper:
            print(f" [COND1] ({cmp_lower}) <= ({cur_lower}) <= ({cmp_upper})")
            fresh_list[idx] = (mn, mx)
            cmp_lower, cmp_upper = fresh_list[idx]
            fresh_list[inner_idx] = (-1, -1)
            continue
        if cur_lower <= cmp_lower and cmp_lower <= cur_upper:
            print(f" [COND2] ({cur_lower}) <= ({cmp_lower}) <= ({cur_upper})")
            fresh_list[idx] = (mn, mx)
            cmp_lower, cmp_upper = fresh_list[idx]
            fresh_list[inner_idx] = (-1, -1)
            continue
        print("[NO COND]")

for lower, upper in fresh_list:
    print(f"===> ({lower}, {upper} => {upper - lower + 1})")
    if lower == -1 and upper == -1:
        continue
    total_items += upper - lower + 1


print(f"Sum: {total_items}")
