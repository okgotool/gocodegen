# gocodegen

Generate CURD for golang, generate model、biz、api、swagger api, used Gin, GORM, Gen, go-swagger.

根据数据库 DML 和表生成 model、biz、api、swagger api 的工具

web 框架使用 Gin：https://github.com/gin-gonic/gin

数据库访问使用 GORM：https://gorm.io/

API 文档使用 go-swagger：https://github.com/go-swagger/go-swagger

底层代码生成采用：GORM Gen：https://github.com/go-gorm/gen

## 工程目录结构

```
- config # 工具配置文件
- db_dml # 数据库DML创建表sql文件
- example # 生成的示例代码
```

## 生成的代码目录结构

```
- api # api层代码，包括swagger，gin，api参数校验
- biz # 业务逻辑层
- dal # 数据访问层
  - model # 数据库model
  - query # Gen DAO查询逻辑
- model # 公共model层
  - response # 公共model，如request、response
- db_migrate # 数据库相关代码，如菜单
- ruoyi_vue3 # ruoyi vue3代码
- vue_element_admin # vue_element_admin代码

```

生成的 swagger API：

![swagger API example](https://raw.githubusercontent.com/okgotool/gocodegen/main/image/swagger_api.png)

## 使用方法

### 配置 gen.xml

配置：./config/gen.yaml

```
mysql: # 连接数据库来生成代码
  host: 127.0.0.1
  port: 3306
  user: user
  password: passwd
  database: dbname
  conn: charset=utf8mb4&parseTime=True
gen:
  dmlFolder: ./db_dml # 创建数据库表sql的目录
  genRuoyiVue3: # 生成Ruoyi vue3代码的配置
    enable: true  # 是否开启
    projectRoot: ./example/ruoyi_vue3 # Ruoyi vue3代码生成的根目录，将按结构生成代码
    topMenus: # 一级菜单
      - name: 管理首页
        path: dashboard
        icon: smile
        orderNum: 1
      - name: 系统设置
        path: system
        icon: component
        orderNum: 2
  genVueElementAdmin: # 生成 vue element admin代码配置
    enable: true # 是否开启
    projectRoot: ./example/vue_element_admin # 代码生成根目录，将按结构生成代码
  genApi:  # 生成后端API
    enable: true
    overWrite: true # 如果文件存在，是否覆盖
    apiGroup: v1 
    excludeColumnsForQueryParameters: # 查询API参数不包含的字段（驼峰格式对象属性）
      - deleted
      - lastmodified
    excludeColumnsForResponse:
      - deleted
  genBiz: # 业务层代码
    enable: true
    overWrite: true
  genDataSource:  # 数据源代码
    enable: true
    overWrite: false
  genDal: # dal代码
    enable: true
    overWrite: true
  dataSources: # 获取数据源代码
    - dataSourceName: db1 # 数据库连接名字，下面每个表可以指定不同的数据库
      dsn: root:passwd@(127.0.0.1:3306)/ai_iotman?charset=utf8mb4&parseTime=True # 直接使用连接字符串，测试时可使用
      # code: app.DBCoon # 使用代码片段获取
  genTables: # 生成哪些表，不配置时生成数据库中所有表，配置了则只生成下面列表的表的代码
    - tableName: sys_role # 需要生成代码的表名
      parentMenu: 设备管理 # 生成前端放到哪个一级菜单下
      menuName: 角色管理 # 生成前端的菜单名
    - tableName: test_cat
      parentMenu: 实时监控
      menuName: 动力系统
      dataSourceName: db1 # 使用的数据源
      fields: # 表字段的生成配置，配置了则使用配置的，否则使用默认配置
        - columnName: id # 表字段名称
          fieldName: Id # 生成的model的属性名
          fieldType: int64 # 生成的model的属性类型名
        - columnName: cat_name
          fieldName: NewFieldName
        - columnName: lastmodified
          isIgnore: false
  outputRoot: ./example # 后端代码生成的位置
  packageRoot: github.com/okgotool/gocodegen/example # 生成的代码的包，import使用


```

### 运行

#### 命令行

- 下载代码生成工具：https://github.com/okgotool/gocodegen/releases
- 解压到目录，如：./cmd/gen 下面
- 创建代码生成配置文件，如： ./gen_example.yaml
- 执行：

```
cd ./cmd/gen
./gocodegen ./gen_example.yaml
```

- 代码只需要在表变更时重新生成
- 代码生成工具不要上传到代码库，只需要上传代码生成配置文件

## 其它

### GORM Conventions 约定

- ID as Primary Key， ID 为主键
- struct name to snake_cases as table name，model 名格式约定
- Column db name uses the field’s name’s snake_case by convention，model 属性名约定
- For models having CreatedAt field, the field will be set to the current time when the record is first created if its value is zero，有 CreatedAt 属性（字段 created_at)时，记录创建时自动设成当前时间
- For models having UpdatedAt field, the field will be set to the current time when the record is updated or created if its value is zero，有 UpdatedAt 属性（字段 updated_at)时，记录更新时自动设成当前时间

### Gocodegen column Conventions Gocodegen 表字段约定

- 必须有 ID 字段，且为主键，字段类别推荐为 bigint(20)，model 属性类型为 int64
- 建议创建 created_at，updated_at 字段
- 建议创建 deleted 字段，字段类别推荐为 tinyint(2)，采用软删除方式删除记录，0 位未删除，1 为删除
- 其它常用字段：
  - updated_by - 最后更新人
  - created_by - 创建者
- 当字段为数字类型时，默认支持多个使用 in 查询，多个时逗号隔开

### Gocodegen API Conventions Gocodegen API 约定

- id 为主键
- 分页参数约定：
  - page - 页码，默认 1
  - pageSize - 每页记录数，默认 10
  - orderBy - 排序，格式："表字段名 asc|desc"，多个逗号隔开，默认 id 倒序： id desc
