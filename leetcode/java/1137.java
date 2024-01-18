class Solution {
  public int tribonacci(int n) {
    if (n == 0)
      return 0;
    if (n == 1)
      return 1;
    if (n == 2)
      return 1;

    var a = 1;
    var b = 1;
    var c = 0;

    for (int i = 2; i < n; i++) {
      var t1 = a;
      var t2 = b;
      a = a + b + c;
      b = t1;
      c = t2;
    }

    return a;
  }
}
