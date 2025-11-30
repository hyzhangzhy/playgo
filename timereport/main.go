package main

import (
	"fmt"
	"os/exec"
	"time"
)

func main() {
	fmt.Println("程序已启动，将每半小时（整点和半点）报时一次...")

	// 启动时先报一次
	speakCurrentTime()

	// 使用 Ticker 每秒检查一次时间
	// 这种方式不依赖于计算长时间的 Sleep，避免了系统休眠导致的计时偏差
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for t := range ticker.C {
		// 检查是否是整点或半点 (0分或30分)，且秒数为0
		if (t.Minute() == 0 || t.Minute() == 30) && t.Second() == 0 {
			speakCurrentTime()
		}
	}
}

func speakCurrentTime() {
	now := time.Now()
	timeStr := now.Format("15:04")
	fmt.Printf("现在时间是: %s\n", timeStr)

	// 构建语音文本
	hour := now.Hour()
	minute := now.Minute()

	var period string
	if hour < 12 {
		period = "上午"
	} else if hour == 12 {
		period = "中午"
	} else {
		period = "下午"
	}

	speakHour := hour
	if hour > 12 {
		speakHour = hour - 12
	}
	if speakHour == 0 {
		speakHour = 12
	}

	var text string
	if minute == 0 {
		text = fmt.Sprintf("现在时间是 %s %d点整", period, speakHour)
	} else {
		text = fmt.Sprintf("现在时间是 %s %d点%d分", period, speakHour, minute)
	} // 调用 Windows PowerShell 进行语音合成
	// 注意：这需要 Windows 环境
	cmd := exec.Command("powershell", "-Command", fmt.Sprintf(`Add-Type -AssemblyName System.Speech; (New-Object System.Speech.Synthesis.SpeechSynthesizer).Speak('%s')`, text))
	err := cmd.Run()
	if err != nil {
		fmt.Println("语音报时失败:", err)
	}
}
