package mlib

import "fmt"

func seqRun() {
	c1 := make(chan struct{})
	c2 := make(chan struct{})
	c3 := make(chan struct{})
	done := make(chan struct{})
	go func() {
		<-c1
		fmt.Println(1)
		c2 <- struct{}{}
	}()
	go func() {
		<-c2
		fmt.Println(2)
		c3 <- struct{}{}
	}()
	go func() {
		<-c3
		fmt.Println(3)
		done <- struct{}{}
	}()
	c1 <- struct{}{}
	<-done
	close(c1)
	close(c2)
	close(c3)
	close(done)
}
