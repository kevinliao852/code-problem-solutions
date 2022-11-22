// The API isBadVersion is defined for you.
// bool isBadVersion(int version);

class Solution {
public:
  int firstBadVersion(int n) {
    unsigned long int i = 0;
    unsigned long int j = n;
    while (i <= j) {
      // 0 ~ 4294967295
      unsigned long int m = (i + j) / 2;

      if (isBadVersion(m) == true && isBadVersion(m - 1) == false) {
        return m;
      }

      if (isBadVersion(m) == true) {
        j = m - 1;
      } else {
        i = m + 1;
      }
    }
    return 0;
  }
};
