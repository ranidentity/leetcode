def is_happy(n:int) ->bool:
    """
    Determines if a given positive integer is a happy number.

    A happy number is a number defined by the following process:
    Starting with any positive integer, replace the number by the sum of the
    squares of its digits. Repeat the process until the number equals 1 (where it will stay),
    or it loops endlessly in a cycle which does not include 1.
    Those numbers for which this process ends in 1 are happy.

    Args:
        n: The positive integer to check.

    Returns:
        True if the number is happy, False otherwise.
    """
    def get_next_number(num):
        total_sum = 0
        while num>0:
            digit = num % 10
            total_sum += digit *digit
            num //= 10
        return total_sum
    
    seen_numbers = set()

    current_number = n
    while current_number != 1 and current_number not in seen_numbers:
        seen_numbers.add(current_number)
        current_number = get_next_number(current_number)

    return current_number == 1