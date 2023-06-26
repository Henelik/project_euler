primes = [2, 3, 5, 7, 11, 13]

for i in range(13, 2000000):
    is_prime = True
    for p in primes:
        if i % p == 0:
            is_prime = False
            break
    if is_prime:
        primes.append(i)

print(sum(primes))