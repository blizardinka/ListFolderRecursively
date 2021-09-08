package main

import (
	"bytes"
	"fmt"
	"io"
	"sort"
	"strings"

	//pkg represents i/o interfaces. And wrap implementation of such primitives in some abstract form to efficient work with it

	//package helps get iformation from operating system. Provides primitives
	//Q: primitives?
	//A: base data structure for representing specific value

	"os"
	// os helps to interract with system
	//"path/filepath"
	//helps to work with system pathes
	//"strings"
	//provide functions to interact with UTF-8 text
)

func main() {
	// //TODO: 19.06.21
	// //Q1: Which purpose to assign out := os.Stdout?
	// //A1: assign "os.Stout" var into "out"

	// out := os.Stdout

	// if !(len(os.Args) == 2 || len(os.Args) == 3) {
	// 	panic("usage go run main.go . [-f]")
	// }
	// // os.args holds command-line arguments
	// //path uese like entry points in our directory
	// path := os.Args[1]
	// printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	// err := dirTree(out, path, printFiles)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Fprint(out)
}

func dirTree(outWriter io.Writer, pathToFile string, inPrintFiles bool) error {
	buf := new(bytes.Buffer)
	var prefix string
	var isDir bool = false
	isDir, err := helpDirTree(buf, pathToFile, inPrintFiles, prefix, isDir)
	if buf.Len() == 0 {
		fmt.Print("empty")
	}
	outWriter.Write(buf.Bytes())
	return err
}

func helpDirTree(outBuffer *bytes.Buffer, pathToFile string, inPrintFiles bool, prefix string, isDir bool) (bool, error) {
	readingFile, err := os.Open(pathToFile)
	if err != nil {
		fmt.Println(err)
	}
	fileStorage, err := readingFile.Readdir(0)
	pathStorage := make([]string, len(fileStorage))
	for n, file := range fileStorage {
		pathStorage[n] = file.Name()
	}
	sort.Strings(pathStorage)
	for n, path := range pathStorage {
		if strings.Compare("", path) != 0 {
			newPath := pathToFile + "/" + path
			newOpenFile, err := os.Stat(newPath)
			if err != nil {
				fmt.Println(err)
			}
			if len(pathStorage)-1 != n {
				if newOpenFile.IsDir() {
					if strings.Count(prefix, "\t") == 0 {
						if isDir || strings.Count(prefix, "\t") == 1 {
							isDir = true
							prefix = "│	├───"
							outBuffer.WriteString(prefix + path + "\n")
							isDir, err = helpDirTree(outBuffer, newPath, inPrintFiles, prefix, isDir)
						}
						isDir = true
						prefix = "├───"
						outBuffer.WriteString(prefix + path + "\n")
						isDir, err = helpDirTree(outBuffer, newPath, inPrintFiles, prefix, isDir)
					}
				} else {
					if isDir || strings.Count(prefix, "\t") == 1 {
						isDir = false
						prefix = "│	├───"
						outBuffer.WriteString(prefix + path + "\n")
						continue
					}
					// if strings.Count(prefix, "\t")==1 {}
					isDir = false
					prefix = "├───" + path + "\n"
					outBuffer.WriteString(prefix)
				}
			} else {
				if newOpenFile.IsDir() {
					isDir = true
					prefix = "└───"
					outBuffer.WriteString(prefix + path + "\n")
				} else {
					if strings.Count(prefix, "\t") == 1 {
						isDir = false
						prefix = "│	└───"
						outBuffer.WriteString(prefix + path + "\n")
						continue
					}
					isDir = false
					prefix = "└───"
					outBuffer.WriteString(prefix + path + "\n")
				}

			}
		}
	}
	return isDir, err
}
