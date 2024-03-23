package node

type Node struct {
	Name            string
	Ip              string
	Cores           int
	Memory          int
	MemoryAllocated int
	Disk            int
	DisAllocated    int
	Role            string
	TaskCount       int
}
