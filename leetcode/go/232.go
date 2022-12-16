type MyQueue struct {
    input []int
    output []int
}


func Constructor() MyQueue {
    return MyQueue{}
}


func (this *MyQueue) Push(x int)  {
    this.input = append(this.input, x)
}


func (this *MyQueue) Pop() int {
    var res int 
    if len(this.output) == 0 {
      for len(this.input) > 0 {
          v := this.input[len(this.input) - 1] 
          this.input = this.input[:len(this.input) - 1]
          this.output = append(this.output, v)
      } 
    } 

    if len(this.output) == 0 {
        return res
    }

    res = this.output[len(this.output) - 1]
    this.output = this.output[:len(this.output) - 1]

    return res
}


func (this *MyQueue) Peek() int {
    var res int

    if len(this.output) == 0 {
      for len(this.input) > 0 {
          v := this.input[len(this.input) - 1]
          this.input = this.input[:len(this.input) - 1]
          this.output = append(this.output, v)
      } 
    }

   
    res = this.output[len(this.output) - 1]

    return res    
}


func (this *MyQueue) Empty() bool {
    var res bool

    if len(this.input) == 0 && len(this.output) == 0 {
        res = true
    }

    return res
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
