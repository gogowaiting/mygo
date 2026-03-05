package interview

import "sort"

type Question struct {
	ID       string
	Topic    string
	Level    int
	Prompt   string
	Answer   string
	Keywords []string
}

var questions = []Question{
	{ID: "go-memory-escape", Topic: "memory", Level: 2, Prompt: "什么是逃逸分析？", Answer: "编译器分析变量是否逃出当前栈帧，逃逸则分配到堆上。常见触发场景：返回局部变量地址、闭包捕获变量、接口装箱。", Keywords: []string{"逃逸分析", "栈", "堆"}},
	{ID: "go-gmp", Topic: "runtime", Level: 3, Prompt: "GMP 模型分别代表什么？", Answer: "G 是 goroutine，M 是内核线程，P 是调度上下文。调度器将 G 放到 P 的本地队列，由 M 绑定 P 执行。", Keywords: []string{"GMP", "调度"}},
	{ID: "go-channel-close", Topic: "concurrency", Level: 2, Prompt: "channel 关闭后读取和写入行为？", Answer: "关闭后可继续读，读到零值且 ok=false；向已关闭 channel 写会 panic。通常由发送方负责关闭。", Keywords: []string{"channel", "close"}},
	{ID: "go-context", Topic: "concurrency", Level: 2, Prompt: "context 的核心用途是什么？", Answer: "跨 goroutine 传递取消信号、超时截止时间和请求级元数据，常用于链路控制和资源回收。", Keywords: []string{"context", "cancel", "timeout"}},
	{ID: "go-interface-nil", Topic: "type-system", Level: 3, Prompt: "为什么 interface 看起来是 nil 但比较不等于 nil？", Answer: "interface 由动态类型和值组成。只有类型和值都为 nil 时接口才等于 nil。", Keywords: []string{"interface", "nil"}},
	{ID: "go-map-concurrency", Topic: "concurrency", Level: 2, Prompt: "原生 map 为什么并发不安全？", Answer: "map 在扩容和写入时会修改内部结构。并发读写会触发数据竞争甚至 panic。可用 mutex/sync.Map 保护。", Keywords: []string{"map", "并发"}},
	{ID: "go-gc", Topic: "runtime", Level: 2, Prompt: "Go GC 的基本特点？", Answer: "并发三色标记+写屏障，目标是降低 STW 时间；通过 GOGC 调整触发频率与内存占用权衡。", Keywords: []string{"GC", "三色标记"}},
	{ID: "go-defer", Topic: "language", Level: 1, Prompt: "defer 的执行时机和参数求值时机？", Answer: "defer 在函数返回前按后进先出执行；defer 调用的参数在注册 defer 时立即求值。", Keywords: []string{"defer", "LIFO"}},
}

func List() []Question {
	out := make([]Question, len(questions))
	copy(out, questions)
	sort.Slice(out, func(i, j int) bool {
		if out[i].Topic == out[j].Topic {
			if out[i].Level == out[j].Level {
				return out[i].ID < out[j].ID
			}
			return out[i].Level < out[j].Level
		}
		return out[i].Topic < out[j].Topic
	})
	return out
}

func Find(id string) (Question, bool) {
	for _, q := range questions {
		if q.ID == id {
			return q, true
		}
	}
	return Question{}, false
}

func Topics() []string {
	seen := map[string]struct{}{}
	for _, q := range questions {
		seen[q.Topic] = struct{}{}
	}
	out := make([]string, 0, len(seen))
	for t := range seen {
		out = append(out, t)
	}
	sort.Strings(out)
	return out
}
