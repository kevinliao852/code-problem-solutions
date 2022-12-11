func mostWordsFound(sentences []string) int {
	var sum int

	var wg sync.WaitGroup
	var ch = make(chan int, 2)
	var done = make(chan bool)

	go func() {
		for {
			data, ok := <-ch
			if !ok {
				done <- true
			}

			if sum < data {
				sum = data
			}
		}
	}()

	for i := 0; i < len(sentences); i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			cnt := 0
			for j := 0; j < len(sentences[i]); j++ {
				if string(sentences[i][j]) == " " {
					cnt++
					ch <- cnt
				}
			}
		}(i)
	}

	wg.Wait()
	close(ch)
	<-done

	return sum + 1
}
