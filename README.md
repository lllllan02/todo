# Todo

## 项目说明

这是一个用于管理任务列表的命令行工具，旨在帮助用户更高效地组织和跟踪个人或工作中的待办事项。

## 使用说明

### 安装

```shell
go install github.com/lllllan02/todo
```

### 启动

```shell
todo
```

### 添加任务 `add`

```shell
todo> add   
? Title: new task
✔  
```

### 任务列表 `ls`

```shell
todo> ls   
+----+----------+-------+---------------------+
| ID |  TITLE   | DONE  |     CREATED AT      |
+----+----------+-------+---------------------+
| 25 | new task | false | 2024-08-14 17:19:46 |
+----+----------+-------+---------------------+
```

### 完成任务 `done`

```shell
todo> done
? Select tasks to done  [Use arrows to move, space to select, <right> to all, <left> to none, type to filter]
> [x]  25 - new task
✔  
```

### 删除任务 `del`

```shell
todo> del
? Select tasks to delete  [Use arrows to move, space to select, <right> to all, <left> to none, type to filter]
> [x]  25 - new task
✔ 
```

## 待实现功能

- [x] 添加任务
- [x] 完成任务
- [x] 删除任务
- [x] 任务列表
- [ ] 任务优先级设置
- [ ] 任务分类管理
- [ ] 创建子任务
- [ ] 添加任务描述
- [ ] 添加评论 / 任务进度说明
- [ ] 。。。
