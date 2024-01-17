class Solution {
  public boolean uniqueOccurrences(int[] arr) {
    Map<Integer, Integer> d = new HashMap<>();
    Map<Integer, Boolean> c = new HashMap<>();

    for (int i = 0; i < arr.length; i++) {
      if (d.containsKey(arr[i])) {
        d.put(arr[i], d.get(arr[i]) + 1);
      } else {
        d.put(arr[i], 1);
      }
    }

    for (Integer v : d.values()) {
      if (c.containsKey(v)) {
        return false;
      } else {
        c.put(v, true);
      }
    }

    return true;
  }
}
