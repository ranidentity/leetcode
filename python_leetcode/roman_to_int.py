symbol_values = {
    'M':1000,
    'CM':900,
    'D':500,
    'CD':400,
    'C':100,
    'L':50,
    'XL':40,
    'X':10,
    'V':5,
    'IV':4,
    'I':1
}

def romanToInt(s:str)->int:
    symbol_values = {
        'I': 1,
        'V': 5,
        'X': 10,
        'L': 50,
        'C': 100,
        'D': 500,
        'M': 1000
    }
    num = 0
    i = 0
    n = len(s)
    
    while i < n:
        # Check for subtractive combinations (e.g., IV, IX, XL, XC, CD, CM)
        if i < n - 1 and symbol_values[s[i]] < symbol_values[s[i+1]]:
            num += symbol_values[s[i+1]] - symbol_values[s[i]]
            i += 2  # Skip the next character since it's already processed
        else:
            num += symbol_values[s[i]]
            i += 1
    return num




print(romanToInt("MCMDCDIIII"))