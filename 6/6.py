def sum_of_squares(limit):
    sum = 0
    for i in range(1, limit+1):
        sum += i*i
    return sum

def square_of_sums(limit):
    sum = 0
    for i in range(1, limit+1):
        sum += i
    return sum**2

print(square_of_sums(100) - sum_of_squares(100))

