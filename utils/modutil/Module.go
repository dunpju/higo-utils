package modutil

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/dunpju/higo-utils/utils/dirutil"
	"os/exec"
	"path/filepath"
	"runtime/debug"
)

type Mod struct{}

func (this *Mod) GetModInfo() *GoMod {
	return GetModInfo()
}

func (this *Mod) GetModule() (string, error) {
	return GetModule()
}

func (this *Mod) GetGoModChildPath(targetPath string) []string {
	return GetGoModChildPath(targetPath)
}

func GetModInfo() *GoMod {
	cmd := exec.Command("go", "mod", "edit", "-json")
	buffer := bytes.NewBufferString("")
	cmd.Stdout = buffer
	cmd.Stderr = buffer

	if err := cmd.Run(); err != nil {
		panic(err)
	}
	goMod := &GoMod{}
	err := json.Unmarshal(buffer.Bytes(), &goMod)
	if err != nil {
		panic(err)
	}
	return goMod
}

type GoMod struct {
	Module  Module
	Go      string
	Require []Require
	Exclude []Module
}

type Module struct {
	Path    string
	Version string
}

type Require struct {
	Path     string
	Version  string
	Indirect bool
}

// GetModule 获取模块名称
func GetModule() (string, error) {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return "", fmt.Errorf("not read build info")
	}
	for _, dep := range info.Deps {
		if dep.Version == "(devel)" {
			return dep.Path, nil
		}
	}
	return "", fmt.Errorf("not found devel module")
}

// GetGoModChildPath 获取子路径
func GetGoModChildPath(targetPath string) []string {
	childPath := make([]string, 0)
begin:
	abovePath := dirutil.Dirname(targetPath)
	files, err := filepath.Glob(targetPath + "/go.mod")
	if err != nil {
		panic(err)
	}
	if len(files) == 0 {
		path := []string{dirutil.Basename(targetPath)}
		childPath = append(path, childPath...)
		targetPath = abovePath
		goto begin
	}
	return childPath
}
