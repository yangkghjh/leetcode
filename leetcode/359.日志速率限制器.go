package leetcode

/*
 * @lc app=leetcode.cn id=357 lang=golang
 *
 * [359] 日志速率限制器
 */

// @lc code=start

// Logger 日志
type Logger struct {
	records map[string]int
}

/** Initialize your data structure here. */

// NewLogger 创建
func NewLogger() Logger {
	// func Constructor() Logger {
	return Logger{
		records: map[string]int{},
	}
}

/** Returns true if the message should be printed in the given timestamp, otherwise returns false.
  If this method returns false, the message will not be printed.
  The timestamp is in seconds granularity. */

// ShouldPrintMessage 返回该条日志是否应该被输出
func (this *Logger) ShouldPrintMessage(timestamp int, message string) bool {
	t, ok := this.records[message]

	if !ok || (ok && timestamp-t >= 10) {
		this.records[message] = timestamp
		return true
	}

	return false
}

/**
 * Your Logger object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.ShouldPrintMessage(timestamp,message);
 */

// @lc code=end
