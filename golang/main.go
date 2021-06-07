package main

import (
	"fmt"
	"os/exec"
	"time"
)

func main() {
	// 设置开始时间
	startTime := time.Date(2021, 1, 1, 2, 3, 3, 0, time.Local)
	// 结束时间为当前日期
	endTime := time.Now().Local()
	// end < start 返回true
	for endTime.After(startTime) {
		commitDate := startTime.Format("Mon Jan 02 15:04:05 2006 +0800")
		// --allow-empty 允许空提交 --date 设置提交时间
		cmd := exec.Command("git", "commit", "--allow-empty", "--date", commitDate, "-m", "每天带点绿")
		if err := cmd.Run(); err != nil {
			fmt.Println(err)
		}
		// 增加一天
		startTime = startTime.AddDate(0, 0, 1)
	}
	fmt.Println("完成")
}
