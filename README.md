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
    "name": "boiler_efficiency_day",//name见redis的table_name表格
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



### redis文档

使用db2，存储展示数据，命名为"2022/05/01 08 table_name"，前面表示本地时间的年月日时，table_name为表格名字

| 表格                  | table_name                   | time精确到              | 值格式                                                  |
| --------------------- | ---------------------------- | ----------------------- | ------------------------------------------------------- |
| 电锅炉热效率          | boiler_efficiency_day        | 日（"2022/05/01"）      | 长度为0到23的数组（根据记录值变化），如[0.9, 0.8, 0...] |
| 蓄热水箱效率          | watertank_efficiency_day     | 日                      | 长度为0到23的数组（根据记录值变化），如[0.9, 0.8, 0...] |
| 能源站效率            | energystation_efficiency_day | 日                      | 长度为0到23的数组（根据记录值变化），如[0.9, 0.8, 0...] |
| 设备在线率（能源站）  | device_online_rate_hour      | 小时（"2022/05/01 08"） | 长度为0到60的数组（根据记录值变化），如[0.9, 0.8, 0...] |
| 锅炉总功率            | boiler_power_hour            | 小时                    | 长度为0到60的数组（根据记录值变化），如[100, 250, 0...] |
| 能源站碳排放（当日）  | energystation_carbon_day     | 日                      | 长度为0到23的数组                                       |
| 能源站碳排放（近7天） | energystation_carbon_week    | 近7天（"2022/05/01"）   | 长度为7的数组（mongo中只存当天的值）                    |
| 能源站负载率          | energy_pay_load              | 日                      | 长度为0到23的数组                                       |
|                       |                              |                         |                                                         |

