// Package lru lru cache by golang
package lru

import "github.com/karlseguin/ccache/v2"

// NewCache 新增cache配置
func NewCache() *ccache.Cache {
	return ccache.New(ccache.Configure())
}
