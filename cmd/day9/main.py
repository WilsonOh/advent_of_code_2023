with open("cmd/day9/input.txt") as f:
    input = [[int(n) for n in line.split()] for line in f.readlines()]

    def extrapolate(seq: list[int], isPart2: bool) -> int:
        if all([num == 0 for num in seq]):
            return 0
        newSeq = [seq[i + 1] - seq[i] for i in range(len(seq) - 1)]
        nextNum = extrapolate(newSeq, isPart2)
        return seq[0] - nextNum if isPart2 else seq[-1] + nextNum

    p1 = sum((extrapolate(seq, isPart2=False) for seq in input))
    p2 = sum((extrapolate(seq, isPart2=True) for seq in input))
    print(f"Part 1: {p1}")
    print(f"Part 2: {p2}")
