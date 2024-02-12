package main

import (
	"testing"
	"time"
)

// Tests creation of the root dir.
// Note we're sleeping between passing data in, as processor will loop until
// the channel closes. We could consider doing this differently to make testing
// easier, but this will suffice for now.
func TestCd(t *testing.T) {
	processor := NewProcessor()
	go processor.process()

	processor.in <- "$ cd somefolder"
	time.Sleep(100 * time.Microsecond)
	if processor.currentDir.pathKey != "/somefolder/" {
		t.Fatalf("Processor incorrectly handled cd. Expected '/somefolder/' for currentDir.pathkey, got: '%s'.", processor.currentDir.pathKey)
	}

	processor.in <- "$ cd someOtherFolder"
	time.Sleep(100 * time.Microsecond)
	if processor.currentDir.pathKey != "/somefolder/someOtherFolder/" {
		t.Fatalf("Processor incorrectly handled cd. Expected '/somefolder/someOtherFolder/' for currentDir.pathkey, got: '%s'.", processor.currentDir.pathKey)
	}

	processor.in <- "$ cd .."
	time.Sleep(100 * time.Microsecond)
	if processor.currentDir.pathKey != "/somefolder/" {
		t.Fatalf("Processor incorrectly handled cd ... Expected '/somefolder/' for currentDir.pathkey, got: '%s'.", processor.currentDir.pathKey)
	}

	processor.in <- "$ cd someOtherFolder"
	time.Sleep(100 * time.Microsecond)
	if processor.currentDir.pathKey != "/somefolder/someOtherFolder/" {
		t.Fatalf("Processor incorrectly handled cd. Expected '/somefolder/someOtherFolder/' for currentDir.pathkey, got: '%s'.", processor.currentDir.pathKey)
	}
	close(processor.in)

	// Test the map.
	if len(processor.dirIndex) != 3 {
		t.Fatalf("processor.dirIndex incorrect length, expected 3, got: %d.", len(processor.dirIndex))
	}
	_, ok := processor.dirIndex["/somefolder/"]
	if !ok {
		t.Fatalf("processor.dirIndex incorrect, expected to find /somefolder/.")
	}
	_, ok = processor.dirIndex["/somefolder/someOtherFolder/"]
	if !ok {
		t.Fatalf("processor.dirIndex incorrect, expected to find /somefolder//someOtherFolder/.")
	}

	// Test the pathList.
	if len(processor.currentDir.pathList) != 2 {
		t.Fatalf("Incorrect pathList length for processor.currentDir.pathList, expected 2, got: %d.", len(processor.currentDir.pathList))
	}
	if processor.currentDir.pathList[1].pathKey != "/somefolder/" {
		t.Fatalf("Incorrect pathList entry for processor.currentDir.pathList[1], expected '/somefolder/', got: '%s'.", processor.currentDir.pathList[2].pathKey)
	}
	if processor.currentDir.pathList[0].pathKey != "/" {
		t.Fatalf("Incorrect pathList entry for processor.currentDir.pathList[0], expected '/', got: '%s'.", processor.currentDir.pathList[2].pathKey)
	}
}

// Tests whether the ls function works correctly.
func TestLs(t *testing.T) {
	processor := NewProcessor()
	go processor.process()

	processor.in <- "$ cd dir1"
	processor.in <- "100 test"
	processor.in <- "dir test2"
	processor.in <- "2000 test3"
	processor.in <- "$ cd dir1_1"
	processor.in <- "55 bla"
	processor.in <- "74 bla"
	processor.in <- "$ cd .."
	processor.in <- "$ cd dir1_2"
	processor.in <- "$ cd dir1_2_3"
	processor.in <- "98 hi"
	processor.in <- "12345678 fini"
	close(processor.in)
	time.Sleep(100 * time.Microsecond)
	if processor.dirTree.totalSize != 12348005 {
		t.Fatalf("Processor incorrectly tracked dirsize. Expected 12348005 for root, got: %d.", processor.dirTree.totalSize)
	}
}
