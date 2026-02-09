class Solution:
    def countGoodIntegers(self, n: int, k: int) -> int:
        factorial = [1] * (n+1)
        for i in range(1, n+1):
            factorial[i] = factorial[i-1] * i

        def count_valid_permutations(digit_count: Counter) -> int:
            # Total permutations = n! / (d_0! * d_1! * ... d_9!)
            denom = 1
            for d in digit_count.values():
                denom *= factorial[d]
            total_permutations = factorial[n] // denom

            if digit_count[0] == 0:
                return total_permutations

            digit_count[0] -= 1
            denom2 = 1
            for d in digit_count.values():
                denom2 *= factorial[d]
            leading_zero_permutations = factorial[n-1] // denom2
            digit_count[0] += 1

            return total_permutations - leading_zero_permutations

        half_len = (n + 1) // 2 
        start = 10 ** (half_len - 1)
        end = 10 ** half_len
        seen_multisets = set()

        for left_num in range(start, end):
            left_str = str(left_num)
            if n % 2 == 0:
                right_str = left_str[::-1]
            else:
                right_str = left_str[-2::-1]

            palindrome_str = left_str + right_str
            palindrome_num = int(palindrome_str)

            if palindrome_num % k == 0:
                digit_count = frozenset(Counter(palindrome_str).items())
                seen_multisets.add(digit_count)
        result = 0
        for multiset in seen_multisets:
            c = Counter(dict(multiset))
            c_int = Counter({int(d): cnt for d, cnt in c.items()})
            result += count_valid_permutations(c_int)

        return result
            