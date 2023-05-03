class Solution:
    def dayOfYear(self, date: str) -> int:
        t = [31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31]
        year, month, day = date.split("-")
        days = 0
        for i in range(int(month) - 1):
            days += t[i]
        if int(month) > 2 and self.isLeapYear(year):
            days += 1

        return days + int(day)

    def isLeapYear(self, year: str) -> bool:
        y = int(year)
        return y % 100 != 0 and y % 4 == 0 or y % 400 == 0