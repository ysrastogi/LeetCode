class Solution:
    def concatHex36(self, n: int) -> str:
        if n == 0:
            return "00"
        result = ""
        trigesimalCharacters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
        decimalCharacters = "0123456789ABCDEF"
        trigesimalString =""
        trigesimalDigit = n**3
        while trigesimalDigit > 0:
            remainder = trigesimalDigit%36
            trigesimalString = trigesimalCharacters[remainder] + trigesimalString
            trigesimalDigit //= 36

        decimalString =""
        decimalDigit = n**2
        while(decimalDigit>0):
            remainder = decimalDigit%16
            decimalString = decimalCharacters[remainder] +decimalString
            decimalDigit //=16

        result = decimalString+trigesimalString
        return result