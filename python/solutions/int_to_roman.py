
def intToRoman2(num:int)-> str:
    if num == 0:
        return ''

    symbols_value = {
        'M' : 1000, 'CM' : 900, 'D' : 500, 'CD' : 400,
        'C' : 100, 'XC' : 90, 'L' : 50, 'XL' : 40, 'X' : 10,
        'IX' : 9, 'V' : 5, 'IV' : 4, 'I' : 1
    }
    
    res = []
    for symbol, value in symbols_value.items():
        cnt = num // value
        res.append(symbol * cnt)
        num -= cnt * value

    return ''.join(res)                

def intToRoman(num:int) -> str:    
    roman_num = ""
    # Thousands place
    count = num // 1000
    roman_num += "M" * count
    num%= 1000
    # Hundreds place
    if num >= 900:
        roman_num += "CM"
        num -= 900
    if num >= 500:
        roman_num += "D"
        num -= 500
    if num >= 400:
        roman_num += "CD"
        num -= 400
    if num >= 100:
        count = num // 100
        roman_num += "C" * count
        num %= 100
    # Tens place
    if num >= 90:
        roman_num += "XC"
        num -= 90
    if num >= 50:
        roman_num += "L"
        num -= 50
    if num >= 40:
        roman_num += "XL"
        num -= 40
    if num >= 10:
        count = num // 10
        roman_num += "X" * count
        num %= 10
    # Ones place
    if num >= 9:
        roman_num += "IX"
        num -= 9
    if num >= 5:
        roman_num += "V"
        num -= 5
    if num >= 4:
        roman_num += "IV"
        num -= 4
    if num > 0:
        roman_num += "I" * num
    return roman_num