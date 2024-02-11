package main

// A node representing a dir in a file structure.
type Dir struct {
	// Name of the dir.
	name string
	// Size of each file in the dir, and any of it's child dirs (recursively).
	totalSize int
	// The path as a string for this dir. Staring with a / where each parent
	// dir's name is separated by a /. Includes the Dir's name.
	pathKey string
	// A slice holding all parents to this Dir, the root being the first
	// element, and the Dir's parent the last.
	pathList []*Dir
}

// Creates a new root dir.
func newRootDir() *Dir {
	result := new(Dir)
	result.name = ""
	result.pathKey = "/"
	result.pathList = []*Dir{}
	return result
}

// Creates a new dir and returns it.
// Only uses parent to determine pathKey and list.
func newDir(name string, parent *Dir) *Dir {
	result := new(Dir)
	result.name = name
	result.pathKey = parent.pathKey + name + "/"
	result.pathList = append(parent.pathList, parent)
	return result
}
