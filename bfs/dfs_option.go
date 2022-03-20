package bfs

type DFSOption func(o *dfsOption)

type dfsOption struct {
	allowCycles     bool
	allowDuplicates bool
}

func AllowDFSCycles() DFSOption {
	return func(o *dfsOption) { o.allowCycles = true }
}

func AllowDFSDuplicates() DFSOption {
	return func(o *dfsOption) { o.allowDuplicates = true }
}
