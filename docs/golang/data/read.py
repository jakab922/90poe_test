import csv

with open("data.csv", "r") as f:
    reader = csv.reader(f)
    first = True
    coll = set()
    for i, line in enumerate(reader):
        if first:
            first = False
            continue
        phone = line[3]
        phone = ''.join(ch for ch in phone if ch in list("0123456789+"))
        if phone.startswith("0"):
            phone = "+44" + phone[1:]

        if phone in coll:
            print i, line[1], line[3]
            exit()

        coll.add(phone)

