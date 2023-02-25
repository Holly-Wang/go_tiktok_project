# go_tiktok_project

## About

a simple implementation of Tiktok.

```bash
├── common                 # 通用组件
│   ├── middlewares        # 中间件
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
├── storage                # 存储本地文件    
├── go.mod
└── go.sum     
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

