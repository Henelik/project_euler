nums = [i*j for i in range(100,1000) for j in range(100,1000)]


def is_palindrome(s):
    return s[::-1] == s


for n in reversed(sorted(nums)):
    if is_palindrome(str(n)):
        print(n)
        break