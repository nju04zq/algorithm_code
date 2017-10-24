package main

import "fmt"
import "strings"

func split(path string) []string {
	skip := func(path string, i int) int {
		for ; i < len(path); i++ {
			if path[i] != '/' {
				break
			}
		}
		return i
	}
	get := func(path string, i int) int {
		for ; i < len(path); i++ {
			if path[i] == '/' {
				break
			}
		}
		return i
	}
	elements := make([]string, 0)
	for i := 0; i < len(path); {
		i = skip(path, i)
		j := get(path, i)
		if j > i {
			elements = append(elements, path[i:j])
		}
		i = j
	}
	return elements
}

func simplifyPath(path string) string {
	final := make([]string, 0)
	for _, element := range split(path) {
		if element == "." {
			continue
		} else if element == ".." {
			if len(final) > 0 {
				final = final[:len(final)-1]
			}
		} else {
			final = append(final, element)
		}
	}
	if len(final) == 0 {
		return "/"
	} else {
		return "/" + strings.Join(final, "/")
	}
}

func testSimplify(path string) {
	fmt.Printf("Before %q, after %q\n", path, simplifyPath(path))
}

func main() {
	testSimplify("/")
	testSimplify("//")
	testSimplify("/home/")
	testSimplify("/a/./b/../../c/")
	testSimplify("/home/../../")
}
