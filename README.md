## 项目介绍

> The code was not generated using GPT tools.

本项目提供了一个终端 todo 的编写示例，使用了 cobra 库包。

## 支持的功能

- [x] 添加 todo 项
  - [x] 支持 -d 参数直接添加
  - [x] 支持键盘手工输入
  - [x] 重复添加校验，当任务重复时，提示任务重复，拒绝添加
- [x] 支持将事项保存到 todo.json
- [x] 列出所有 todo 项
  - [x] 支持根据任务分组展示
- [x] 支持 任务划分 group, 默认为 default，支持创建任务时，通过 -g 传入自定义任务名称
- [x] 支持任务状态管理，创建后，任务状态为 Pending
- [x] 支持更新任务
  - [x] 更新任务名称
  - [x] 更新任务状态，并做合法值校验
  - [x] 更新任务分组
- [x] 标记 todo 项为完成， 完成状态为 Completed
- [x] 删除 todo 项

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
