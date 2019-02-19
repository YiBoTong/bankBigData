// Copyright 2018 gf Author(https://git.91ybt.com/lib/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://git.91ybt.com/lib/gf.

// 内存锁.
package gmlock

var locker = New()

// 内存写锁，如果锁成功返回true，失败则返回false;过期时间单位为毫秒，默认为0表示不过期
func TryLock(key string, expire ...int) bool {
	return locker.TryLock(key, expire...)
}

// 内存写锁，锁成功返回true，失败时阻塞，当失败时表示有其他写锁存在;过期时间单位为毫秒，默认为0表示不过期
func Lock(key string, expire ...int) {
	locker.Lock(key, expire...)
}

// 解除基于内存锁的写锁
func Unlock(key string) {
	locker.Unlock(key)
}

// 内存读锁，如果锁成功返回true，失败则返回false;过期时间单位为毫秒，默认为0表示不过期
func TryRLock(key string, expire ...int) bool {
	return locker.TryRLock(key, expire...)
}

// 内存写锁，锁成功返回true，失败时阻塞，当失败时表示有写锁存在;过期时间单位为毫秒，默认为0表示不过期
func RLock(key string, expire ...int) {
	locker.RLock(key, expire...)
}

// 解除基于内存锁的读锁
func RUnlock(key string) {
	locker.RUnlock(key)
}
