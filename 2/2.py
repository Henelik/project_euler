fib = [1, 1]

# generate the list of fibonacci numbers
while True:
    fib.append(fib[-1] + fib[-2])
    if fib[-1] > 4000000:
        break
    if fib[-1] % 2 == 0:
        print(fib[-1])
