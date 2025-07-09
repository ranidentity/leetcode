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
    
    #first round will have a,d 
    # then combine to become ad, then pop d and go for next which is e
    def reCurve(current_combination: list[str], next_digit_index:int):
        print(current_combination)
        if len(current_combination) == len(digits): # done checking all numpads
            result.append("".join(current_combination)) #ad
            return
        digit = digits[next_digit_index]
        letters = numPads.get(digit,"") # get the letters e.g. 2= abc
        for letter in letters: #abc #def
            current_combination.append(letter) # store a, e  
            reCurve(current_combination,next_digit_index+1)
            current_combination.pop() # clear off d
    reCurve([],0)
    return result


# better outcome
def iterativeLetterCombinations(digits: str) -> list[str]:
    m = {'2':'abc','3':'def','4':'ghi','5':'jkl','6':'mno','7':'pqrs','8':'tuv','9':'wxyz'}

    if len(digits) == 0:
        return []

    res = ['']
    #23
    for i in range(len(digits)): #2
        n = len(res) # 1 # 3
        for j in range(n): #when 3, #0 #1 #2
            for c in m[digits[i]]: # abc # def
                res.append(res[j] + c) # '',a, b, c #a when j=0 #j=1, b #j=2, c 
        res = res[n:] # pop 1. '', 2. a,b,c 

    return res

print(iterativeLetterCombinations("23"))
