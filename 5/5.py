for i in range(20, 20**20):
    valid = True
    for j in range(2, 21):
        if i % j != 0:
            valid = False
            break
    if valid:
        print(i)
        break