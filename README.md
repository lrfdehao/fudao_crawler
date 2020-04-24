# FUDAO crawler

使用 `golang` 实现腾讯课堂 `https://www.fudao.com` 下的课程抓取小工具。

## 技术选型

golang + MySQL

使用cronjob每天定时启动一次。

## 思路
1. 根据 **grade_subject** 接口抓取所有的 `年级(grade)` 以及对应的 `科目(subject)`。
2. 根据 `年级(grade)` 和 `科目(subject)` 调用 **subject_course** 接口抓取所有的课程与课程包 并整合所有 `课程id(course_id)`。
3. 根据所有的 `课程id(course_id)` 调用 **course_detail** 接口抓取所有的 `目录(course_directory)` 和 `教师(teacher)` 信息。

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
``` sql
CREATE TABLE `course_directory` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `csid` int(11) DEFAULT NULL,
  `title` varchar(45) DEFAULT NULL,
  `begin_time` timestamp NULL DEFAULT NULL,
  `end_time` timestamp NULL DEFAULT NULL,
  `tid` bigint(20) DEFAULT NULL,
  `grade` int(11) DEFAULT NULL,
  `subject` int(11) DEFAULT NULL,
  `cs_type` int(11) DEFAULT NULL,
  `exam_duration` int(11) DEFAULT NULL,
  `cid` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `csid_idx` (`csid`),
  KEY `cid_idx` (`cid`)
) ENGINE=InnoDB AUTO_INCREMENT=10045 DEFAULT CHARSET=utf8mb4;

CREATE TABLE `course_info` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `cid` int(11) unsigned NOT NULL,
  `course_name` varchar(45) DEFAULT NULL,
  `begin_time` timestamp NULL DEFAULT NULL,
  `end_time` timestamp NULL DEFAULT NULL,
  `grade` int(11) DEFAULT NULL,
  `subject` int(11) DEFAULT NULL,
  `apply_num` int(11) DEFAULT NULL,
  `pre_amount` int(11) DEFAULT NULL,
  `af_amount` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `cid_idx` (`cid`)
) ENGINE=InnoDB AUTO_INCREMENT=922 DEFAULT CHARSET=utf8mb4;

CREATE TABLE `teacher` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `tid` bigint(20) DEFAULT NULL,
  `name` varchar(45) DEFAULT NULL,
  `desc` varchar(1024) DEFAULT NULL,
  `pic` varchar(1024) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `tid_idx` (`tid`)
) ENGINE=InnoDB AUTO_INCREMENT=187 DEFAULT CHARSET=utf8mb4;
```