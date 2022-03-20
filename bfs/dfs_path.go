package bfs

type DFSPath[T any] interface {
	Path() []T
	Len() int
	Contains(string) bool
}

type dfsPath[T any] struct {
	path []T
	// Change value to int so we can keep count of instances?
	set map[string]bool
}

func (dp *dfsPath[T]) pop(s string) {
	dp.path = dp.path[:len(dp.path)-1]
	delete(dp.set, s)
}

func (dp *dfsPath[T]) push(t T, s string) {
	dp.path = append(dp.path, t)
	dp.set[s] = true
}

func (dp *dfsPath[T]) Path() []T {
	return dp.path
}

func (dp *dfsPath[T]) Len() int {
	return len(dp.path)
}

func (dp *dfsPath[T]) Contains(s string) bool {
	return dp.set[s]
}
