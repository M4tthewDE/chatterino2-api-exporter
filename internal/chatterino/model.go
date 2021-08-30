package chatterino

type AllStats struct {
	memory *Memory
}

type Memory struct {
	alloc       int64
	totalAlloc  int64
	systemAlloc int64
	numGC       int
}
