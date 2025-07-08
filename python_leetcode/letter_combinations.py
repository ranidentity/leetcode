def letterCombinations(digits: str) -> list[str]:
    numPads = {
        '2': "abc",
        '3': "def",
        '4': "ghi",
        '5': "jkl",
        '6': "mno",
        '7': "pqrs",
        '8': "tuv",
        '9': "wxyz"
    }
    result = []

    if not digits:
        return result
    
    def reCurve(current_combination: list[str], next_digit_index:int):
        if len(current_combination) == len(digits):
            result.append("".join(current_combination))
            return
        