package cpu

import "runtime"

/**
 *
 * @Author: zhangxiaohu
 * @File: manager.go.go
 * @Version: 1.0.0
 * @Time: 2020/1/14
 */

type CpuEntry struct {
	CoresCount int
}

func GetCpuInfomarion() CpuEntry{
	return CpuEntry{CoresCount:runtime.NumCPU()}
}

