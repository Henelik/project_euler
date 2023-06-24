nums = [i*j for i in range(1000,100) for j in range(1000,100)]


def is_palindrome(s):
    return s[::-1] == s


for n in reversed(sorted(nums)):
    print("checking", n)
    if is_palindrome(str(n)):
        print(n)
        break