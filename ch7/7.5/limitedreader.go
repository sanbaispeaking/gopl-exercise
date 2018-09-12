package limitreader

import "io"

type limitReader struct {
	r     io.Reader
	limit int64
}

func (l *limitReader) Read(p []byte) (int, error) {
	n, err := l.r.Read(p)
	if err != nil {
		if err != io.EOF {
			return n, err
		}
	}
	if int64(n) > l.limit {
		n = int(l.limit)
		err = io.EOF
	}
	return n, err
}

// LimitReader returns a wrapped reader that report EOF after reading n bytes
// from the original reader r
func LimitReader(r io.Reader, n int64) io.Reader {
	lr := limitReader{r, n}
	return &lr
}
