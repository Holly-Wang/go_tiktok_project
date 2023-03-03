# go_tiktok_project

## About

a simple implementation of Tiktok.

```bash
.
├── build.sh               # 编译脚本
├── common                 # 通用组件
│   ├── authenticate       # 鉴权
│   ├── config             # 配置
│   ├── const.go           
│   ├── dal                # 数据访问
│   │   ├── mysql
│   │   └── redis
│   ├── errlog             # 错误信息日志
│   └── middlewares        # 中间件
├── docker-compose.sh
├── docker-compose.yml
├── handler                # 接口
├── idl                    # 接口定义
│   ├── *.proto
│   └── biz/model
├── main.go
├── output
│   |-- bin
│   └── conf
├── README.md
├── router.go
├── service
├── start.sh
├── util                   # 工具
└── video_data             # 存储文件
```

## Usage

``` bash
git clone https://github.com/Holly-Wang/go_tiktok_project

cd go_tiktok_project

go build -o tiktok && ./tiktok
```


## Generate pb

proto files must add `option go_package = "pb/";` and are included in `./idl`

so you can use this cmd to generate *.pb.go files 

``` bash
cd idl

hz model -idl service.proto
```

## Contributing

## License

