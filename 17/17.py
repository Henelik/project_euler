digits = {
    1: "one",
    2: "two",
    3: "three",
    4: "four",
    5: "five",
    6: "six",
    7: "seven",
    8: "eight",
    9: "nine",
    10: "ten",
    11: "eleven",
    12: "twelve",
    13: "thirteen",
    14: "fourteen",
    15: "fifteen",
    16: "sixteen",
    17: "seventeen",
    18: "eighteen",
    19: "nineteen",
}

tens = {
    20: "twenty",
    30: "thirty",
    40: "forty",
    50: "fifty",
    60: "sixty",
    70: "seventy",
    80: "eighty",
    90: "ninety",
}

def number_to_english(n):
    result = ""
    if n >= 100:
        if n % 100 == 0:
            return digits[n//100] + " hundred"
        result += digits[n//100] + " hundred and "
        n %= 100
    if n in list(digits.keys()):
        return result + digits[n]
    result += tens[(n//10)*10]
    n %= 10
    if n in list(digits.keys()):
        result += "-" + digits[n]
    return result

print(len("".join(["one thousand"] + [number_to_english(n) for n in range(1, 1000)]).replace(" ", "").replace("-", "")))