/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
	var m int

	for i, j := 1, n; i <= j; {
		m := (i + j) / 2

		if isBadVersion(m) == true && isBadVersion(m-1) == false {
			return m
		}

		if isBadVersion(m) == false {
			i = m + 1
		} else {
			j = m - 1
		}

	}

	return m
}
