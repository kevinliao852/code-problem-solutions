type Data struct {
	i int
	p int
}

type StockSpanner struct {
	stack []Data
	count int
}

func Constructor() StockSpanner {
	return StockSpanner{
		stack: make([]Data, 0, 0),
		count: 0,
	}
}

func (this *StockSpanner) push(price int) {
	this.stack = append(this.stack, Data{this.count, price})
}

func (this *StockSpanner) pop() {
	this.stack = this.stack[:len(this.stack)-1]
}

func (this *StockSpanner) top() Data {
	return this.stack[len(this.stack)-1]
}

func (this *StockSpanner) size() int {
	return len(this.stack)
}

func (this *StockSpanner) Next(price int) int {
	this.count++

	if this.count == 1 {
		this.push(price)
		return 1
	}

	var data int = this.count

	for this.size() > 0 {
		if this.top().p > price {
			data = this.count - this.top().i
			break
		}

		this.pop()
	}

	this.push(price)
	return data
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */
