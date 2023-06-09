### 文件目录

```go
energy                 
├─ api                 // api接口实现
│  ├─ analysis         
│  ├─ system           
│  └─ login.go         
├─ config              // 保存项目初始化所需数据
│  └─ config.ini       
├─ dataReceive         // 存放接受设备数据的接口
├─ log                 // 日志打印文件夹
│  ├─ log              
│  └─ log20220513.log  
├─ middleware          // 中间件
│  ├─ cors.go          // 跨域中间件
│  ├─ jwt.go           // jwt身份鉴权中间件
│  └─ logger.go        // 打印日志中间件
├─ model               // 存放子功能所需结构体和一些方法
│  ├─ analysis         
│  ├─ system           
│  ├─ db.go            // 连接数据库
│  └─ User.go          
├─ routes              // 存放各个子功能路由文件 
│  └─ login.go         
├─ utils               // 工具包
│  ├─ errmsg           
│  │  └─ errmsg.go     // 存放错误代码
│  └─ setting.go       // 读取config.ini中的数据并加载
├─ go.mod              
├─ go.sum              
├─ main.go             // main函数文件
├─ README.md           
└─ router.go           // 路由文件
```

### mongo 相关数据文档

#### calculation_result：存储计算结果
```json
{
    "time": "2022/05/01",
    "name": "energy_boiler_efficiency_day",//name见redis的table_name表格
    "value": [0.9, 0.8, 0, 0.88]
}
```

#### opc_data：按小时存储原始数据，每小时的列为一个数组

```json
{
    "itemid": "server.A%E6%B3%B5%E8%BF%90%E8%A1%8C1",
    "value": [false, false, true, true],
    "time": "2022/09/13 03"
}
```

#### loukong：楼控数据，在楼控原始数据基础上加上时间和表名

```json
{
	"time": "2022/05/01 08:05",
	"name": "heat",
    "info": {}
}
```

name目前有：heat：各组团热表；GA：太阳能热水，每分钟一个

