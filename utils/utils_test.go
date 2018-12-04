package utils

import "testing"

func TestRandomStr(t *testing.T) {
	str := RandomStr(10)
	InitLogs()
	LogDebug("a9a9a9a9a9a9a9aa9a")
	t.Error(str)
}