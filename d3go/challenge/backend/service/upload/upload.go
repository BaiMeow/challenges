package upload

import (
	"archive/zip"
	"d3go/utils"
	"io"
	"os"
	"path"
	"path/filepath"
)

func Unzip(zipPath string, outputPath string) (*utils.Tree, error) {
	r, err := zip.OpenReader(zipPath)
	defer r.Close()
	if err != nil {
		return nil, err
	}

	if err := os.MkdirAll(outputPath, 0750); err != nil {
		return nil, err
	}
	tree := utils.NewTree(filepath.Base(outputPath))
	for _, f := range r.File {
		utils.Paths2Tree(f.Name, path.Join(outputPath, f.Name), tree)
		p, _ := filepath.Abs(filepath.Join(outputPath, f.Name))
		content, err := f.Open()
		if err != nil {
			return nil, err
		}
		data, err := io.ReadAll(content)
		if err != nil {
			content.Close()
			return nil, err
		}
		if err := os.MkdirAll(filepath.Dir(p), 0750); err != nil {
			content.Close()
			return nil, err
		}
		if err := os.WriteFile(p, data, 0750); err != nil {
			content.Close()
			return nil, err
		}
		content.Close()

	}

	return tree, nil
}
