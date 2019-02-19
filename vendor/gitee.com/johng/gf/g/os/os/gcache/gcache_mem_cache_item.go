// Copyright 2018 gf Author(https://git.91ybt.com/lib/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://git.91ybt.com/lib/gf.

package gcache

import "git.91ybt.com/lib/gf/g/os/gtime"

// 判断缓存项是否已过期
func (item *memCacheItem) IsExpired() bool {
	// 注意这里应当包含等于，试想一下缓存时间只有最小粒度为1毫秒的情况
	if item.e >= gtime.Millisecond() {
		return false
	}
	return true
}
