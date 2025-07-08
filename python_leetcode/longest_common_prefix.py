def longestCommonPrefix(strs:list[str])->str:
    if not strs:
        return ""
    if len(strs) == 1:
        return strs[0]

    strs.sort()

    first_str = strs[0]
    last_str = strs[len(strs) - 1]
    common_prefix = []
    for i in range(min(len(first_str), len(last_str))):
        if first_str[i] == last_str[i]:
            common_prefix.append(first_str[i])
        else:
            break
    return "".join(common_prefix)


print(longestCommonPrefix(["flower","owfl","fluck"]))
