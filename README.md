## 项目介绍

> The code was not generated using GPT tools.

本项目提供了一个 todo 的编写示例，使用了 cobra 库包。

## 支持的功能

第一阶段，主要是学习 golang 几个主要的开发场景和主流库包。

- [x] 提供 CLI 终端命令 `cobra`
  - [x] 添加 todo 项
    - [x] 支持 -d 参数直接添加
    - [x] 支持键盘手工输入
    - [x] 重复添加校验，当任务重复时，提示任务重复，拒绝添加
  - [x] 支持将事项保存到 todo.json
  - [x] 列出所有 todo 项
    - [x] 支持根据任务分组展示
    - [x] 支持使用 `done` 子命令快速查看已完成的任务
  - [x] 支持 任务划分 group, 默认为 default，支持创建任务时，通过 -g 传入自定义任务名称
  - [x] 支持任务状态管理，创建后，任务状态为 Pending
  - [x] 支持更新任务
    - [x] 更新任务名称
    - [x] 更新任务状态，并做合法值校验
    - [x] 更新任务分组
  - [x] 标记 todo 项为完成， 完成状态为 Completed
  - [x] 删除 todo 项
  - [x] 支持搜索功能
  - [x] 支持命名错误时，给出建议命令
- [ ] 提供 HTTP 服务 `net/http`
  - [ ] 定义 API: `/api/v1/todos`
  - [ ] 支持 GET 查询全部的 todo
    - [ ] 支持默认返回全部
    - [ ] 支持根据状态返回
    - [ ] 支持指定搜索条件
      - [ ] 任务名称 模糊匹配 不区分大小写
      - [ ] 任务名称 + 分组
  - [ ] 支持使用 POST 创建任务
    - [ ] 支持重复检测
    - [ ] 支持指定组和任务状态
  - [ ] 支持使用 PUT 更新一个 todo
  - [ ] 支持使用 DELETE 删除一个 todo
- [ ] 提供 GRPC 服务
  - [ ]

## 代码理解增强

请逐行为我的 Go 代码添加注释。作为一名专业的 Go 工程师，请在注释中详细说明每行代码的数据类型、语法格式、在程序中的作用，
以及所使用的库包。所有解释都通过注释的形式添加到代码中。

Please add comments to my Go code line by line. As a professional Go engineer,
provide detailed explanations for each line, including data types, syntax structure,
purpose within the program, and the libraries used.
All explanations should be added as comments within the code.

## 使用说明

```bash
(base) samzonglu in ~/Git/goprojs/go-study-todo on main λ ./mytodo -h
A longer description that spans multiple lines and likely contains examples and usage of using your application.

Usage:
  todo [command]

Available Commands:
  add         Add a new task to the list
  completion  Generate the autocompletion script for the specified shell
  delete      Delete a task from the list
  help        Help about any command
  list        list all tasks
  search      A brief description of your command

Flags:
  -c, --config string   config file
  -h, --help            help for todo
  -r, --region string   AWS region (default "us-west-2")

Use "todo [command] --help" for more information about a command.
```

### list and add task

```bash
(base) samzonglu in ~/Git/goprojs/go-study-todo on main λ ./mytodo ls
No tasks to list
(base) samzonglu in ~/Git/goprojs/go-study-todo on main λ ./mytodo add -d "learn cobra"
Added task: learn cobra, Group is: default
(base) samzonglu in ~/Git/goprojs/go-study-todo on main λ ./mytodo add -d "learn golang" -g golang
Added task: learn golang, Group is: golang
(base) samzonglu in ~/Git/goprojs/go-study-todo on main λ ./mytodo add -d "write todo cli with cobra"
Added task: write todo cli with cobra, Group is: default
```

### delete task

```bash
(base) samzonglu in ~/Git/goprojs/go-study-todo on main λ ./mytodo ls
List Tasks:
>>> Group : default
1. learn cobra   status: Pending
2. write todo cli with cobra     status: Pending
>>> Group : golang
1. learn golang  status: Pending
(base) samzonglu in ~/Git/goprojs/go-study-todo on main λ ./mytodo delete -d "write todo cli with cobra"
Deleted task: write todo cli with cobra
(base) samzonglu in ~/Git/goprojs/go-study-todo on main λ ./mytodo ls
List Tasks:
>>> Group : default
1. learn cobra   status: Pending
>>> Group : golang
1. learn golang  status: Pending
```

### update task:

```bash
(base) samzonglu in ~/Git/goprojs/go-study-todo on main λ ./mytodo ls
List Tasks:
>>> Group : default
1. learn cobra   status: Pending
2. jack  status: Completed
>>> Group : golang
1. learn golang  status: Completed
(base) samzonglu in ~/Git/goprojs/go-study-todo on main λ ./mytodo update -d "learn cobra" -s "completed"
Task updated successfully
(base) samzonglu in ~/Git/goprojs/go-study-todo on main λ ./mytodo ls
List Tasks:
>>> Group : golang
1. learn golang  status: Completed
>>> Group : default
1. learn cobra   status: Completed
2. jack  status: Completed
```

### search task

```bash
(base) samzonglu in ~/Git/goprojs/go-study-todo on main ● ● λ ./mytodo search -h
A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.

Usage:
  todo search [flags]

Flags:
  -g, --group string      Group tasks by group
  -h, --help              help for search
  -k, --keywords string   Search tasks by keywords <required>

Global Flags:
      --config string   config file (default is $HOME/.todo.yaml)
(base) samzonglu in ~/Git/goprojs/go-study-todo on main ● ● λ ./mytodo search -k "learn"
>>> Group: default
Task: learn cobra	Status: completed
>>> Group: golang
Task: learn golang	Status: completed
(base) samzonglu in ~/Git/goprojs/go-study-todo on main ● ● λ ./mytodo search -k "learn" -g golang
>>> Group: golang
Task: learn golang	Status: completed
```

### suggestion cmd

```bash
(base) samzonglu in ~/Git/goprojs/go-study-todo on main ● λ ./mytodo adc
Error: unknown command "adc" for "todo"

Did you mean this?
	add

Run 'todo --help' for usage.
```

### run as server

```bash
 go run main.go server
2024/12/05 23:36:31 Starting server on :8080
2024/12/05 23:36:34 POST%!(EXTRA string=/api/v1/todos)
2024/12/05 23:36:34 Received todo: {Description:jack433 Status:pendding Group:default}
2024/12/05 23:36:34 Task already exists
2024/12/05 23:37:19 POST%!(EXTRA string=/api/v1/todos)
2024/12/05 23:37:19 Received todo: {Description:jack193 Status:pendding Group:default}
2024/12/05 23:37:19 Added task: jack193, Group is: default
2024/12/05 23:37:21 POST%!(EXTRA string=/api/v1/todos)
2024/12/05 23:37:21 Received todo: {Description:jack146 Status:pendding Group:default}
2024/12/05 23:37:21 Added task: jack146, Group is: default
2024/12/05 23:37:22 POST%!(EXTRA string=/api/v1/todos)
2024/12/05 23:37:22 Received todo: {Description:jack103 Status:pendding Group:default}
2024/12/05 23:37:22 Added task: jack103, Group is: default
2024/12/05 23:37:23 POST%!(EXTRA string=/api/v1/todos)
```
