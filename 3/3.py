def prime_factors(n):
    for i in range(2, n):
        if n % i == 0:
            return [i] + prime_factors(n//i)
    return [n]

print(max(prime_factors(600851475143)))
