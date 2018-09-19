package archivereader

import (
	syszip "archive/zip"
)

type zip struct {
}

func (z *zip) Open(name string) ([]inforer, error) {
	zr, err := syszip.OpenReader(name)
	if err != nil {
		return nil, err
	}
	defer zr.Close()

	list := make([]inforer, 1)
	for _, f := range zr.File {
		list = append(list, f)
	}
	return list, err
}

func init() {
	Register("zip", new(zip))
}
