package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

// Responsible for the processing of instructions.
type processor struct {
	// Mapping of Dir pathKey to dir. For easy retrieval.
	dirIndex map[string]*Dir
	// Channel to pass in instructions from cli to processor.
	in chan string
	// The directory tree.
	dirTree *Dir
	// Current directory.
	currentDir *Dir
}

// Creates and initialises a processor.
func NewProcessor() *processor {
	result := new(processor)
	result.in = make(chan string, 10)
	result.dirTree = newRootDir()
	result.dirIndex = make(map[string]*Dir)
	result.dirIndex[result.dirTree.pathKey] = result.dirTree
	result.currentDir = result.dirTree
	return result
}

// Switch current direct to parent.
func (p *processor) cdUp() {
	if len(p.currentDir.pathList) == 0 {
		log.Fatal("Attempt to cd up out of root!")
	}
	p.currentDir = p.currentDir.pathList[len(p.currentDir.pathList)-1]
}

// Switch current dir into a named dir.
func (p *processor) cdIn(dirname string) {
	// See if the dir already exists.
	tempPath := fmt.Sprintf("%s%s/", p.currentDir.pathKey, dirname)
	targetDir, ok := p.dirIndex[tempPath]
	if ok {
		p.currentDir = targetDir
	} else {
		p.currentDir = newDir(dirname, p.currentDir)
		p.dirIndex[p.currentDir.pathKey] = p.currentDir
	}
}

// Will increase the current dir and all it's parents' file sizes by the given
// size.
// It would be quicker if we did this only once for the sum of the folder, but
// would make the program less easy to read.
func (p *processor) registerFileSize(size int) {
	//log.Printf("Adding %d to %s", size, p.currentDir.pathKey)
	// Add the size to the currentDir, and all of its parents.
	p.currentDir.totalSize += size
	for _, parent := range p.currentDir.pathList {
		//log.Printf("Adding %d to %s", size, parent.pathKey)
		parent.totalSize += size
	}
}

// Processes commands coming into the in channel until it closes.
func (p *processor) process() {
	for line := range p.in {
		// First parse the command. One of:
		words := strings.Fields(line)
		// Check if we received a command?
		if words[0] == "$" {
			// We can ignore ls, as any command not starting with $ is a result
			// from ls.
			if words[1] == "cd" {
				if len(words) < 3 {
					log.Println("Received cd without directory. Skipping.")
					continue
				}
				if words[2] == ".." {
					p.cdUp()
				} else {
					p.cdIn(words[2])
				}
			}
		} else {
			// This must be output from ls.
			// Ignore any dirs (we'll create them as we move in)
			if words[0] != "dir" {
				// Grab the file size and register it.
				size, err := strconv.Atoi(words[0])
				if err != nil {
					log.Fatalf("Received incorrectly formatted ls line: '%s'.", line)
				}
				p.registerFileSize(size)
			}
		}
	}
}

// Sums up the size of any directory of at most the given limit and returns the
// total.
func (p *processor) SumDirSizesUpTo(limit int) int {
	result := 0
	for _, dir := range p.dirIndex {
		if dir.totalSize <= limit {
			result += dir.totalSize
		}
	}
	return result
}
