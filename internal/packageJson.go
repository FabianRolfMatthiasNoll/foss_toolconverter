package internal

import (
	"encoding/json"
	"os"
)

type Package struct {
	Dependencies map[string]string
}

func (pkg Package) GetTopLevelDependencies(pkgPath string) ([]string, error) {
	file, err := os.ReadFile(pkgPath)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(file, &pkg)
	if err != nil {
		return nil, err
	}

	topDeps := make([]string, len(pkg.Dependencies))

	i := 0
	for k := range pkg.Dependencies {
		topDeps[i] = k
		i++
	}

	return topDeps, nil
}
