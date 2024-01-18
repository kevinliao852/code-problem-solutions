class Solution {
  public int fib(int n) {
    int a = 0;
    int b = 1;

    for (int i = 1; i <= n; i++) {
      int t = a;
      a = a + b;
      b = t;
    }
    return a;
  }
}
