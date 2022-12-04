/*
// Definition for a Node.
class Node {
    public int val;
    public List<Node> children;

    public Node() {}

    public Node(int _val) {
        val = _val;
    }

    public Node(int _val, List<Node> _children) {
        val = _val;
        children = _children;
    }
};
*/

class Solution {
  public List<Integer> preorder(Node root) {
    List<Integer> list = new ArrayList<Integer>();

    if (root == null) {
      return list;
    }

    Collections.addAll(list, root.val);
    this.traverse(root, list);

    return list;
  }
  public void traverse(Node node, List<Integer> list) {
    if (node != null) {
      for (int i = 0; i < node.children.size(); i++) {
        list.add(node.children.get(i).val);
        traverse(node.children.get(i), list);
      }
    }
  }
}
