package archivereader

import (
	systar "archive/tar"
	"io"
	"os"
)

type tar struct {
}

func (t *tar) Open(name string) ([]inforer, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	tr := systar.NewReader(f)

	list := make([]inforer, 1)
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		list = append(list, hdr)
	}
	return list, nil
}

func init() {
	Register("tar", new(tar))
}
