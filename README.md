# FUDAO crawler

使用 `golang` 实现腾讯课堂 `https://www.fudao.com` 下的课程抓取小工具。

## 技术选型

golang + MySQL

使用cronjob每天定时启动一次。

## 思路
1. 根据 **grade_subject** 接口抓取所有的 `年级(grade)` 以及对应的 `科目(subject)`。

![grade_subject](https://imgkr.cn-bj.ufileos.com/edef032b-b172-4182-a7ef-8daca8a5a4f5.png)

2. 根据 `年级(grade)` 和 `科目(subject)` 调用 **discover_subject** 接口抓取所有的课程与课程包 并整合所有 `课程id(course_id)`。

![discover_subject](https://imgkr.cn-bj.ufileos.com/75d7afae-7065-418a-8b0d-3ae29d8fb79d.png)

3. 根据所有的 `课程id(course_id)` 调用 **courseStaticDetail** 接口抓取所有的 `目录(course_directory)` 和 `教师(teacher)` 信息。

![courseStaticDetail](https://imgkr.cn-bj.ufileos.com/c4f1d3c3-4014-47e2-a969-3e795bd0d033.png)

- 其中抓取时 分页 `size=0` 为所有数据。
- `grade=0` 为测试数据。

## 目录结构

```
crawler
├── README.md
├── constant
│   └── constant.go
├── crawler
├── go.mod
├── go.sum
├── main.go
├── model
│   ├── db.go
│   └── resp.go
├── util
│   └── util.go
└── worker
    ├── course_detail.go
    ├── grade_subj.go
    ├── package_course.go
    └── subject_course.go
```

## SQL Table
[SqlTable](#sql_data) 
