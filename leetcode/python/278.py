# The isBadVersion API is already defined for you.
# @param version, an integer
# @return a bool
# def isBadVersion(version):


class Solution(object):
    def firstBadVersion(self, n):
        """
        :type n: int
        :rtype: int
        """
        j = n
        i = 0
        while i <= j:
            m = (i + j) / 2

            isfirst = isBadVersion(m) and isBadVersion(m - 1) == False

            if isfirst:
                return m
            elif isBadVersion(m):
                j = m - 1
            else:
                i = m + 1

		# not found
		return 0
