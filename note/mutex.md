### sync.Map
sync.Map其实就是维护两个map，一个read、一个dirty，他们的entry都是指向同一数据的指针，但是read map数据没dirty全。dirty的访问需要mu锁，而read的结构不会改变，所以访问不需要mu锁。read上的key可以安全的并发读写（不包括删除）。当新增key时，会通过mu锁写到dirty map。当删除key时，dirty map会直接删除， read map 则是变为expunge。当删除的key再次写入时，会先写dirty map，再写read map。如果读写的数据不在read map就会去dirty map找，并增加miss计数，miss到一定的数量会直接把dirty map 提升为 read map。

### sync.Mutex
如果只有一个goroutine，那么直接获取锁，否则进入慢路径  
慢路径：如果当前不是饥饿模式，那么自旋一段时间，还是没有获取到就进入等待队列。等到唤醒时还要跟新到的goroutine竞争，竞争不到插回队头。
饥饿模式：如果获取不到锁的时间超过1ms，就会进入饥饿模式，这时候新到的goroutine不再抢锁，而是乖乖排队。直到只剩最后一个waiter或者小于1ms时解除饥饿模式。

### MESI协议
数据首次加载到缓存是独占状态，加入独占状态的好处是避免无意义的BusRd请求，因为其他核心没有该数据。  
