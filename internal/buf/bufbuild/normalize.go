package bufbuild

import (
	"fmt"
	"sort"
	"strings"

	"github.com/bufbuild/buf/internal/pkg/storage/storagepath"
	"github.com/bufbuild/buf/internal/pkg/util/utilstring"
)

func normalizeAndValidateRoots(inputRoots []string) ([]string, error) {
	if len(inputRoots) == 0 {
		inputRoots = []string{"."}
	}
	return normalizeAndValidateFileList(inputRoots, "root")
}

func normalizeAndValidateRootsExcludes(inputRoots []string, inputExcludes []string) ([]string, []string, error) {
	roots, err := normalizeAndValidateRoots(inputRoots)
	if err != nil {
		return nil, nil, err
	}

	if len(inputExcludes) == 0 {
		return roots, nil, nil
	}

	excludes, err := normalizeAndValidateFileList(inputExcludes, "exclude")
	if err != nil {
		return nil, nil, err
	}

	rootMap := utilstring.SliceToMap(roots)
	excludeMap := utilstring.SliceToMap(excludes)

	// verify that no exclude equals a root directly
	for exclude := range excludeMap {
		if _, ok := rootMap[exclude]; ok {
			return nil, nil, fmt.Errorf("%s is both a root and exclude, which means the entire root is excluded, which is not valid", exclude)
		}
	}
	// verify that all excludes are within a root
	for exclude := range excludeMap {
		if !storagepath.MapContainsMatch(rootMap, exclude) {
			return nil, nil, fmt.Errorf("exclude %s is not contained in any root, which is not valid", exclude)
		}
		if storagepath.Ext(exclude) == ".proto" {
			return nil, nil, fmt.Errorf("excludes can only be directories but file %s discovered", exclude)
		}
	}
	return roots, excludes, nil
}

func normalizeAndValidateFileList(inputs []string, name string) ([]string, error) {
	if len(inputs) == 0 {
		return inputs, nil
	}

	var outputs []string
	for _, input := range inputs {
		if input == "" {
			return nil, fmt.Errorf("%s value is empty", name)
		}
		output, err := storagepath.NormalizeAndValidate(input)
		if err != nil {
			// user error
			return nil, err
		}
		outputs = append(outputs, output)
	}
	sort.Strings(outputs)

	for i := 0; i < len(outputs); i++ {
		for j := i + 1; j < len(outputs); j++ {
			output1 := outputs[i]
			output2 := outputs[j]

			if output1 == output2 {
				return nil, fmt.Errorf("duplicate %s %s", name, output1)
			}
			if strings.HasPrefix(output1, output2) {
				return nil, fmt.Errorf("%s %s is within %s %s which is not allowed", name, output1, name, output2)
			}
			if strings.HasPrefix(output2, output1) {
				return nil, fmt.Errorf("%s %s is within %s %s which is not allowed", name, output2, name, output1)
			}
		}
	}

	// already checked duplicates, but if there are multiple directories and we have ".", then the other
	// directories are within the output directory "."
	var notDotDir []string
	hasDotDir := false
	for _, output := range outputs {
		if output != "." {
			notDotDir = append(notDotDir, output)
		} else {
			hasDotDir = true
		}
	}
	if hasDotDir {
		if len(notDotDir) == 1 {
			return nil, fmt.Errorf("%s %s is within %s . which is not allowed", name, notDotDir[0], name)
		}
		if len(notDotDir) > 1 {
			return nil, fmt.Errorf("%ss %v are within %s . which is not allowed", name, notDotDir, name)
		}
	}

	return outputs, nil
}
