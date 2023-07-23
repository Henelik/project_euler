# This problem can be rephrased as 
# "how many combinations of 0 and 1 of length 40 are there, with exactly 20 0s"

# Find all combinations that start with a 0, then multiply by 2
#for i in range(1, 20):
#    solution = [0] * i

# count_permutations only counts the permutations starting with the given mode - HALF of all available
def count_permutations(num_current, num_other):
    count = 0
    if num_other == 0:
        return 1
    for i in range(1, num_current+1):
        count += count_permutations(num_other, num_current-i)
    return count

print(count_permutations(20, 20)*2)