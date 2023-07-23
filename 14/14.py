def collatz_length(n):
    if n == 1:
        return 1
    if n%2 == 0:
        return collatz_length(n/2)+1
    else:
        return collatz_length(3*n+1)+1

longest_length = 0
num_with_longest_length = 0

for i in range(1, 1000000):
    length = collatz_length(i)
    if longest_length < length:
        longest_length = length
        num_with_longest_length = i

print(num_with_longest_length)