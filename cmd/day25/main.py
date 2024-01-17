import networkx as nx


graph = nx.Graph()
with open("./input.txt") as f:
    for line in f.readlines():
        vertex, neiString = line.split(": ")
        for nei in neiString.split():
            graph.add_edge(vertex, nei)
_, (l, r) = nx.stoer_wagner(graph)
ans = len(l) * len(r)
print(f"Part 1: {ans}")
