# go_tiktok_project

## About

a simple implementation of Tiktok.

```bash
├── common                 # 通用组件
│   ├── middleware         # 中间件
│   ├── dal                # 数据访问
│   │   ├── mysql          
│   │   └── redis          
│   ├── config             # 配置
│   └── errlog             # 错误信息日志
├── handler                # 应用目录
├── idl                    # 接口定义
├── util                   # 工具
├── service
│   └── ...           
├── go.mod
└── go.sum     
```

## Usage

``` bash
git clone https://github.com/Holly-Wang/go_tiktok_project

cd go_tiktok_projec

go build -o tiktok && ./tiktok
```


## Generate pb

proto 文件需要包含 `option go_package = "pb/";`

``` bash
protoc idl/*.proto --go_out=./idl
```

## 

hertz-examples: https://github.com/cloudwego/hertz-examples

## Badge

[![standard-readme compliant](https://img.shields.io/badge/readme%20style-standard-brightgreen.svg?style=flat-square)](https://github.com/RichardLitt/standard-readme)

## Contributing

## License

