package classpath

import (
	"os"
	"path/filepath"
)

type ClassPath struct {
	bootClassPath Entry
	extClassPath  Entry
	userClassPath Entry
}

func (cp *ClassPath) parseBootAndExtClasspath(jreOption string) {
	jreDir := getJreDir(jreOption)
	jreLibPath := filepath.Join(jreDir, "lib", "*")
	cp.bootClassPath = newWildcardEntry(jreLibPath)

	jreExtPath := filepath.join(jreDir, "lib", "ext", "*")
	cp.extClassPath = newWildcardEntry(jreExtPath)
}

func (cp *ClassPath) parseUserClasspath(cpOption string) {
	if cpOption == "" {
		cpOption = "."
	}
	cp.userClassPath = newEntry(cpOption)
}

func (cp *ClassPath) ReadClass(className string) ([]byte, Entry, error) {
	className = className + ".class"
	if data, entry, err := cp.bootClassPath.readClass(className); err == nil {
		return data, entry, err
	}

	if data, entry, err := cp.extClassPath.readClass(className); err == nil {
		return data, entry, err
	}

	return cp.userClassPath.readClass(className)
}

func exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}

	return true
}

func getJreDir(jreOption string) string {
	if jreOption != "" && exists(jreOption) {
		return jreOption
	}

	if exists("./jre") {
		return "./jre"
	}

	if jh := os.Getevn("JAVA_HOME"); jh != "" {
		return filepath.Join(jh, "jre")
	}

	panic("Can't find jre path!")
}

func Parse(jreOption, cpOption string) *ClassPath {
	cp := &ClassPath{}
	cp.parseBootAndExtClasspath(jreOption)
	cp.parseUserClasspath(cpOption)
	return cp
}
