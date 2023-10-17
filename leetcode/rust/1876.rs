impl Solution {
    pub fn count_good_substrings(s: String) -> i32 {
        let mut count: i32 = 0;
        let chars: Vec<char> = s.chars().collect();
        
        if chars.len() <= 2 {
            return 0;
        }

        for i in 0..chars.len() -2 {
            if chars[i] != chars[i + 1] && chars[i + 1] != chars[i + 2] &&  chars[i] != chars[i + 2] {
                count += 1;
            }
        }

        return count;
    }
}
