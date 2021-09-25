package util

import (
	"io"
	"net/url"
	"path/filepath"
	"strings"
)

// GenFullRequestUrl 根据路径信息，生成全路径
func GenFullRequestUrl(protocal string, domain string, basepath string, path string, query map[string]string) string {
	ks := make([]string, 0)
	var rawQuery string
	if query != nil && len(query) > 0 {
		for k, v := range query {
			ks = append(ks, k, "=", v, "&")
		}
		rawQuery = strings.Join(ks[:len(ks)-1], "")
	}

	u := &url.URL{
		Scheme:   protocal,
		Host:     domain,
		Path:     filepath.Join(basepath, path),
		RawQuery: rawQuery,
	}
	return u.String()
}

// FastClose 快速关闭Closer，并忽略err
func FastClose(closer io.Closer) {
	_ = closer.Close()
}

// PageReaderIterator 分页读取迭代器
type PageReaderIterator interface {
	// Next 读取下面一条
	Next() ([]byte, error)
	// NextMulti 读取下面多条
	NextMulti(count int) ([]byte, error)
	// HasNext 是否有下一条
	HasNext() bool
}