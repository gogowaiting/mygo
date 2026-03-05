package learn

import "sort"

type Demo struct {
	ID          string
	Topic       string
	Title       string
	Description string
	Path        string
	Level       int
}

var demos = []Demo{
	{ID: "basic-00-pre", Topic: "basic", Title: "预定义标识符", Description: "Go 关键字和内置标识符基础", Path: "./cmd/101-pre", Level: 1},
	{ID: "basic-01-hello", Topic: "basic", Title: "Hello World", Description: "最小可运行程序", Path: "./cmd/102-hello", Level: 1},
	{ID: "basic-02-var-const", Topic: "basic", Title: "变量与常量", Description: "var/const 与类型推导", Path: "./cmd/103-var-const", Level: 1},
	{ID: "basic-03-for", Topic: "basic", Title: "循环", Description: "for 的常见写法", Path: "./cmd/104-for", Level: 1},
	{ID: "basic-04-if", Topic: "basic", Title: "条件分支", Description: "if/else 语法", Path: "./cmd/105-if", Level: 1},
	{ID: "basic-05-switch", Topic: "basic", Title: "Switch", Description: "值 switch 与类型 switch", Path: "./cmd/106-switch", Level: 1},
	{ID: "basic-06-array", Topic: "basic", Title: "数组", Description: "数组声明与访问", Path: "./cmd/107-array", Level: 1},
	{ID: "basic-07-slice", Topic: "basic", Title: "切片", Description: "切片初始化与截取", Path: "./cmd/108-slice", Level: 1},
	{ID: "basic-08-map", Topic: "basic", Title: "Map", Description: "键值结构与遍历", Path: "./cmd/109-map", Level: 1},
	{ID: "basic-09-range", Topic: "basic", Title: "Range", Description: "range 在集合中的行为", Path: "./cmd/110-range", Level: 1},
	{ID: "basic-10-func", Topic: "basic", Title: "函数", Description: "函数签名、返回值与闭包", Path: "./cmd/111-func", Level: 1},
	{ID: "basic-11-point", Topic: "basic", Title: "指针", Description: "值传递与指针传递", Path: "./cmd/112-pointer", Level: 2},
	{ID: "basic-12-struct", Topic: "basic", Title: "结构体", Description: "结构体声明与初始化", Path: "./cmd/113-struct", Level: 2},
	{ID: "basic-13-method", Topic: "basic", Title: "方法", Description: "接收者与方法集合", Path: "./cmd/114-method", Level: 2},
	{ID: "basic-14-interface", Topic: "basic", Title: "接口", Description: "接口定义与实现", Path: "./cmd/115-interface", Level: 2},
	{ID: "basic-15-channel", Topic: "concurrency", Title: "Channel", Description: "无缓冲/有缓冲与关闭语义", Path: "./cmd/201-channel", Level: 2},
	{ID: "basic-16-goroutine", Topic: "concurrency", Title: "Goroutine", Description: "并发执行与生命周期", Path: "./cmd/202-goroutine", Level: 2},
	{ID: "basic-range-chan", Topic: "concurrency", Title: "Range Channel", Description: "for-range 消费 channel", Path: "./cmd/203-range-chan", Level: 2},
	{ID: "basic-test-timer", Topic: "concurrency", Title: "定时与超时", Description: "timer、after、超时控制", Path: "./cmd/204-timer", Level: 3},
	{ID: "basic-slice-append", Topic: "basic", Title: "切片扩容", Description: "append 行为观察", Path: "./cmd/116-slice-append", Level: 2},
	{ID: "basic-nil-struct", Topic: "basic", Title: "nil 与结构体", Description: "nil 判定和零值语义", Path: "./cmd/117-nil-struct", Level: 2},
	{ID: "algo-bubble", Topic: "algorithm", Title: "冒泡排序", Description: "O(n^2) 稳定排序", Path: "./cmd/301-bubble-sort", Level: 1},
	{ID: "algo-select", Topic: "algorithm", Title: "选择排序", Description: "O(n^2) 原地排序", Path: "./cmd/302-select-sort", Level: 1},
	{ID: "algo-insert", Topic: "algorithm", Title: "插入排序", Description: "局部有序场景友好", Path: "./cmd/303-insert-sort", Level: 1},
	{ID: "algo-quick", Topic: "algorithm", Title: "快速排序", Description: "分治思想", Path: "./cmd/304-quick-sort", Level: 2},
	{ID: "adv-context", Topic: "advanced", Title: "Context", Description: "上下文取消与传播", Path: "./cmd/402-context", Level: 3},
	{ID: "adv-traffic-control", Topic: "advanced", Title: "流量控制", Description: "并发控制与节流思路", Path: "./cmd/403-traffic-control", Level: 3},
	{ID: "adv-guess-number", Topic: "advanced", Title: "猜数字", Description: "综合控制流练习", Path: "./cmd/401-guess-number", Level: 2},
	{ID: "adv-worker-pool", Topic: "advanced", Title: "Worker Pool", Description: "生产者-消费者与并发 worker", Path: "./cmd/404-worker-pool", Level: 3},
	{ID: "adv-sync-pool", Topic: "advanced", Title: "sync.Pool", Description: "对象复用降低分配压力", Path: "./cmd/405-sync-pool", Level: 3},
	{ID: "adv-atomic", Topic: "advanced", Title: "Atomic", Description: "无锁计数与原子操作", Path: "./cmd/406-atomic", Level: 3},
	{ID: "adv-generic-set", Topic: "advanced", Title: "Generic Set", Description: "泛型容器设计", Path: "./cmd/407-generic-set", Level: 3},
	// Web 开发 - Gin (5xx)
	{ID: "gin-basic", Topic: "web", Title: "Gin 基础", Description: "引擎初始化、路由、启动服务", Path: "./cmd/501-gin-basic", Level: 1},
	{ID: "gin-route", Topic: "web", Title: "路由进阶", Description: "路由分组、路径参数、查询参数", Path: "./cmd/502-gin-route", Level: 1},
	{ID: "gin-bind", Topic: "web", Title: "请求绑定", Description: "JSON/表单/Query 绑定与 validator", Path: "./cmd/503-gin-bind", Level: 2},
	{ID: "gin-response", Topic: "web", Title: "响应处理", Description: "JSON/XML/重定向响应", Path: "./cmd/504-gin-response", Level: 2},
	{ID: "gin-middleware", Topic: "web", Title: "中间件", Description: "自定义中间件、鉴权、中间件链", Path: "./cmd/505-gin-middleware", Level: 2},
	{ID: "gin-upload", Topic: "web", Title: "文件上传", Description: "文件上传与静态文件服务", Path: "./cmd/506-gin-upload", Level: 2},
	{ID: "gin-crud", Topic: "web", Title: "CRUD 综合", Description: "RESTful 风格完整 CRUD", Path: "./cmd/507-gin-crud", Level: 3},
	{ID: "interview-cli", Topic: "interview", Title: "面试题模块", Description: "面试题库 list/random/show", Path: "./cmd/interview", Level: 2},
}

func List() []Demo {
	out := make([]Demo, len(demos))
	copy(out, demos)
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

func Find(id string) (Demo, bool) {
	for _, d := range demos {
		if d.ID == id {
			return d, true
		}
	}
	return Demo{}, false
}

func Topics() []string {
	seen := map[string]struct{}{}
	for _, d := range demos {
		seen[d.Topic] = struct{}{}
	}
	out := make([]string, 0, len(seen))
	for t := range seen {
		out = append(out, t)
	}
	sort.Strings(out)
	return out
}
