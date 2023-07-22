# completes in 1h57m53.87s

class TriangleNumbers:
    def __iter__(self):
        self.index = 0
        self.cum_sum = 0
        return self
    
    def __next__(self):
        self.index += 1
        self.cum_sum += self.index
        return self.cum_sum
    
def count_divisors(n):
    divisors = 1
    for i in range(1, n//2):
        if n%i == 0:
            divisors += 1
    return divisors

for n in iter(TriangleNumbers()):
    if count_divisors(n) > 500:
        print("found solution:", n)
        exit()