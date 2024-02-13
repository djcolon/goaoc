package main

import "testing"

// Tests creation of the root dir.
func TestNewRootDir(t *testing.T) {
	rootDir := newRootDir()
	if rootDir.name != "" {
		t.Errorf("Root dir should have empty name.")
	}
	if rootDir.pathKey != "/" {
		t.Errorf("Root dir should have pathKey '/'.")
	}
	if len(rootDir.pathList) != 0 {
		t.Errorf("Root dir should have empty path list.")
	}
}

// Tests creation of a child dir.
func TestNewDir(t *testing.T) {
	currentNode := newRootDir()
	dir := newDir("someName", currentNode)
	if dir.name != "someName" {
		t.Errorf("Returned dir has incorrect name: %s, expected: %s.", dir.name, "someName")
	}
	if dir.pathKey != "/someName/" {
		t.Errorf("Returned dir has incorrect pathKey: %s, expected: %s.", dir.name, "/someName/")
	}
	if len(dir.pathList) != 1 {
		t.Errorf("Returned dir has incorrect pathList length: %d, expected: 1.", len(dir.pathList))
	}
	if dir.pathList[0] != currentNode {
		t.Errorf("Returned dir has incorrect pathList, expected root node, got: %v.", dir.pathList[0])
	}
}

// Tests creation of multiple child dirs.
func TestNewDirToDepth(t *testing.T) {
	root := newRootDir()
	level1 := newDir("level1", root)
	level2 := newDir("level2", level1)
	level3 := newDir("level3", level2)

	if level3.name != "level3" {
		t.Errorf("Level3 dir has incorrect name: %s, expected: level3.", level3.name)
	}
	if level3.pathKey != "/level1/level2/level3/" {
		t.Errorf("Returned dir has incorrect name: %s, expected: '/level1/level2/level3/'.", level3.name)
	}
	if len(level3.pathList) != 3 {
		t.Errorf("Returned dir has incorrect pathKey length: %d, expected: 3.", len(level3.pathList))
	}
	if level3.pathList[0] != root {
		t.Errorf("Returned dir has incorrect pathList, expected root node in [0], got: %v.", level3.pathList[0])
	}
	if level3.pathList[1] != level1 {
		t.Errorf("Returned dir has incorrect pathList, expected level1 node in [1], got: %v.", level3.pathList[1])
	}
	if level3.pathList[2] != level2 {
		t.Errorf("Returned dir has incorrect pathList, expected level2 node in [2], got: %v.", level3.pathList[2])
	}
}
