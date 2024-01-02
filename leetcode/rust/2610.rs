use std::collections::HashMap;

impl Solution {
    pub fn find_matrix(nums: Vec<i32>) -> Vec<Vec<i32>> {
        let mut res = Vec::new();
        let mut hash_map = HashMap::new();

        for &n in &nums {
            if hash_map.get(&n).unwrap_or(&0) == &0 {
                hash_map.insert(n, 1);
            } else {
                hash_map.insert(n, hash_map[&n] + 1);
            }

            if res.len() < hash_map[&n] {
                res.push(vec![n]);
            } else {
                res[hash_map.get(&n).unwrap() - 1].push(n);
            }
        }

        return res;
    }
}
