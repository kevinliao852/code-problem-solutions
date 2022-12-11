impl Solution {
    pub fn convert_temperature(celsius: f64) -> Vec<f64> {
        let mut res = Vec::new();

        res.push(celsius + 273.15);
        res.push(celsius * 1.80 + 32.00);

        return res;
    }
}
