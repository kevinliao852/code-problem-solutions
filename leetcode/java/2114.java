class Solution {
  public int mostWordsFound(String[] sentences) {
    int sum = 0;
    for (String s : sentences) {
      int count = 0;
      for (char w : s.toCharArray()) {

        if (w == ' ') {
          count++;
          sum = Math.max(count, sum);
        }
      }
    }
    return sum + 1;
  }
}
