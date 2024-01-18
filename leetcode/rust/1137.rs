impl Solution {
    pub fn tribonacci(n: i32) -> i32 {
        if n == 2 {
            return 1;
        }
        if n == 1 {
            return 1;
        }
        if n == 0 {
            return 0;
        }

        let mut a = 1;
        let mut b = 1;
        let mut c = 0;

        for _ in 2..n {
            (a, b, c) = (a + b + c, a, b);
        }

        return a;
    }
}
