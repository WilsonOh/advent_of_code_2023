with open("./input.txt") as fin:
    with open("./out.dot", "w") as fout:
        fout.write("graph {\n")
        for line in fin.readlines():
            vertex, neiString = line.split(": ")
            neiList = neiString.split()
            fout.write(f"\t{vertex} -- {{{', '.join(neiList)}}};\n")
        fout.write("}")
