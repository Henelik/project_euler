def calculate():
    for c in range(1000, 1, -1):
        for b in range(1000-c, 1, -1):
            for a in range(1000-c-b, 1, -1):
                #print(f"a: {a}, b: {b}, c: {c}")
                if a+b+c == 1000 and a*a + b*b == c*c:
                    return a*b*c

print(calculate())