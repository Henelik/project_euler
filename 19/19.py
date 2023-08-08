days_in_months = [31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31]

total_sundays = 0
current_weekday = 2 # January 1, 1901 was a Tuesday

for year in range(1901, 2001):
    for month in range(12):
        # track fist sundays
        if current_weekday == 0:
            total_sundays += 1

        days_in_month = days_in_months[month]
        # check February for leap years
        if month == 1 and (year % 400 == 0 or (year % 4 == 0 and year % 100 != 0)):
            days_in_month = 29

        current_weekday = (current_weekday+days_in_month)%7

print(total_sundays)
