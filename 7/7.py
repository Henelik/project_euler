primes = [2, 3, 5, 7, 11, 13]

for i in range(13, 4294967295):
    is_prime = True
    for p in primes:
        if i % p == 0:
            is_prime = False
            break
    if is_prime:
        primes.append(i)
        if len(primes) >= 10001:
            print(primes[10000])
            break