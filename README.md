# 🖨️⌨️ letgo

🌟 **描述**： **letgo** 是一个 **力扣**/**leetcode** 本地刷题工具[目前仅支持go语言]。代码主要由 **go** 语言实现，让用户可以方便的选择自己喜欢的
**IDE** 进行刷题。

----

## 🚀快速开始

> 1 请先准备好:
> - go 环境
> - 一个你喜欢的IDE
> - 网络畅通
> - leetcode 账号
>
> 2 将代码备份到本地
>
> 3.1 将[conf-template替换成conf.yaml](system_code/conf/conf-template.yaml)
>
> 3.2 [准备好leetcode cookie并修改conf.yaml](system_code/doc/leetcode_cookie.md)
>
> 4 letgo down -i 题目ID
> 
> ![ss](system_code\doc\letgo%20down.gif "yes")

## 🔎项目结构

> 截至 2024年1月25日 的项目结构

```shell
(base) PS E:\code\letgoV2>
.
├─.idea                 # idea 配置文件 请忽略
├─system_code           # 系统文件 ！如果你只是刷题，请着重看一下your_code下面的文件结构！
│  ├─commands           # 控制台命令
│  ├─conf               # 配置文件
│  ├─middleware         # 未来可能的中间件，目前还没有
│  ├─models             # 未来可能的数据库交互层
│  ├─pkg                # 常用模块
│  │  ├─common					
│  │  ├─e
│  │  ├─func_operator	# 负责启动your_code代码，将string类型的参数转换成go可接受的类型
│  │  ├─logging			# 日志控制
│  │  │  └─http_logging
│  │  ├─setting
│  │  └─util			# 一些常用的工具方法
│  ├─runtime			# 系统的运行日志
│  │  ├─http_logs
│  │  └─logs
│  └─service            # 服务层
│      ├─code_handle_service
│      ├─down_service
│      ├─generate_service
│      │  └─generate_params
│      └─leetcode_api
│          ├─leetcode_bodys
│          └─leetcode_common
└─your_code			# 刷题文件，你的刷题代码会保存到这里
    │				# 每当你使用命令"down"下题目后，会生成一个名为IDxxx_题目的文件夹
    └─IDzzzz_longest_substring_without_repeating_characters # 这里的ID将根据创建次数从[zzzz迭代至aaaa]表示从[0-45_6975]
        └─logs 			# 调用 test.go 时自动生成运行记录
        code.go			# 需要你实现的代码
        code_test.go	# 这个文件是为了方便通过IDE进行debug
        meta_data.go	# 包含一些代码相关的信息，包括测试用例 !!测试用例的填写格式应和官方一致!!
        README-en.md	# 纯英的文档
        README-zh.md

```

## 📢📣 声明

本项目仅供学习参考之用途，请勿用于商业。