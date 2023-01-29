# go_tiktok_project

## Background

a simple implementation of Tiktok.

## Install

This project uses [golang](https://github.com/golang) and [hertz](https://github.com/cloudwego/hertz). Go check them out if you don't have them locally installed.

## Usage

```bash
├── common
│   ├── middleware         # 中间件
│   ├── dal                # 数据访问
│   ├── config             # 配置
│   └── errlog             # 错误信息日志
├── handler                # 应用目录
├── go.mod                 # go 依赖管理
├── go.sum
├── idl                    # 接口定义
│   ├── comment.proto      # 评论接口
│   ├── favorite.proto     # 点赞接口
│   ├── follow.proto       # 关系接口
│   ├── user.proto         # 用户接口
│   └── video.proto        # 视频接口
├── util                   # 工具
└── service
    └── ...                
```

## Badge

[![standard-readme compliant](https://img.shields.io/badge/readme%20style-standard-brightgreen.svg?style=flat-square)](https://github.com/RichardLitt/standard-readme)

## Contributing

## License

