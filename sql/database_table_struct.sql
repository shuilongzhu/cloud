create database cloud;
use cloud;
SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;
-- ----------------------------
-- Table structure for user_info
-- ----------------------------
create table user_info
(
    id                          int auto_increment comment '自增ID'
        primary key,
    user_id                     varchar(32)  not null comment '用户id',
    user_group_id               varchar(32)  not null comment '用户分组id',
    user_group_name             varchar(32)  null comment '用户组名称',
    user_name                   varchar(32)  null comment '用户姓名',
    nickname                    varchar(32)  null comment '用户昵称',
    role                        int          null comment '角色类型 0：普通用户；1：系统超级管理员；2:运维观察员；3:模型管理员',
    parent_node_id              varchar(32)  null comment '父节点ID',
    number                      varchar(32)  null comment '手机号',
    user_password               varchar(32)  null comment '用户密码 md5加密',
    emil                        varchar(32)  null comment '邮箱',
    note                        varchar(512) null comment '备注信息',
    created_time                datetime     null comment '创建时间',
    update_time                 datetime     null comment '更新时间',
    login_time                  datetime     null comment '登录时间',
    login_ip                    varchar(32)  null comment '登录IP',
    avatar_storage_address      varchar(128) null comment '头像存储地址',
    Id_photo_address            varchar(128) null comment '身份证照片存储地址',
    Id_number                   varchar(32)  null comment '身份证号码',
    Id_card_verification_status int          null comment '身份证验证状态 0：待提交；1：待审核；2：审核通过；3：审核驳回，需修改信息；4：审核拒绝;',
    user_state                  int          null comment '用户状态 1：正常；2：已禁用；',
    WeChat_ID                   varchar(32)  null comment '微信ID',
    qq_id                       varchar(32)  null comment 'QQ ID',
    alternate_field1            varchar(32)  null comment '备用字段1',
    alternate_field2            varchar(32)  null comment '备用字段2',
    alternate_field3            varchar(32)  null comment '备用字段3'
)comment '用户表 ' collate = utf8mb4_bin;

create index index_user_group_id
    on user_info (user_group_id);

create index index_user_id
    on user_info (user_id);
-- ----------------------------
-- Table structure for user_group
-- ----------------------------
create table user_group
(
    id               int auto_increment comment '自增ID'
        primary key,
    user_group_id    varchar(32)  null comment '用户组id',
    user_group_name  varchar(32)  null comment '用户组名称',
    user_group_note  varchar(128) null comment '用户组描述',
    created_time     datetime     null comment '创建时间',
    update_time      datetime     null comment '更新时间',
    delete_tag       int          null comment '删除标记 0:正常；1删除',
    alternate_field1 varchar(32)  null comment '备用字段1',
    alternate_field2 varchar(32)  null comment '备用字段2'
)comment '用户组表 ' collate = utf8mb4_bin;

create index index_user_group_id
    on user_group (user_group_id);
-- ----------------------------
-- Table structure for user_login_info
-- ----------------------------
create table user_login_info
(
    id            int auto_increment comment '自增ID'
        primary key,
    user_id       varchar(32) null comment '用户id',
    user_group_id varchar(32) null comment '用户分组id',
    login_time    datetime    null comment '登录时间',
    login_device  int         null comment '登录设备 1:pc端；2:android；3:iOS;',
    login_ip      varchar(32) null comment '登录IP',
    created_time  datetime    null comment '创建时间'
)comment '用户登录记录表 ' collate = utf8mb4_bin;

create index index_group_id
    on user_login_info (user_group_id);

create index index_user_id
    on user_login_info (user_id);
-- ----------------------------
-- Table structure for user_role
-- ----------------------------
create table user_role
(
    id          int auto_increment comment '自增ID'
        primary key,
    name        varchar(255) not null comment '角色组名称',
    `desc`      varchar(255) not null comment '角色组描述',
    create_time int          null comment '创建时间',
    modify_time int          null comment '修改时间'
)comment '用户角色表 ' collate = utf8mb4_bin;
-- ----------------------------
-- Table structure for channel_info
-- ----------------------------
create table channel_info
(
    id                  bigint       not null comment '自增ID'
        primary key,
    device_id           varchar(32)  null comment '设备ID',
    device_name         varchar(32)  null comment '设备名称',
    channel_id          varchar(32)  null comment '通道ID',
    channel_name        varchar(32)  null comment '通道名称',
    channel_ip          varchar(32)  null comment '通道IP地址',
    channel_username    varchar(32)  null comment '通道用户名',
    channel_password    varchar(32)  null comment '通道登录密码',
    channel_access_type varchar(32)  null comment '通道接入类型',
    take_way            varchar(32)  null comment '取流方式',
    channel_address     varchar(128) null comment '通道安装地址',
    channel_url         varchar(128) null comment '通道直播流地址',
    channel_model       varchar(32)  null comment '通道选择模型',
    setup_time          varchar(32)  null comment '配置时间',
    channel_status      varchar(32)  null comment '通道状态 1:在线2:离线3:关闭',
    access_time         datetime     null comment '接入时间',
    created_time        datetime     null comment '创建时间',
    delete_time         datetime     null comment '删除时间',
    delete_tag          int          null comment '删除标记 0:正常；1：已删除；',
    take_address        varchar(128) null comment '取流地址',
    is_false            int          null comment '是否是假通道（0：不是假通道；1：是假通道）',
    alternate_field2    varchar(32)  null comment '备用字段2',
    longitude           varchar(32)  null comment '通道地理经度',
    latitude            varchar(32)  null comment '通道地理纬度'
)comment '设备通道信息表' collate = utf8mb4_bin;

create index index_chan_id
    on channel_info (channel_id);

create index index_device_id
    on channel_info (device_id);
-- ----------------------------
-- Table structure for device_info
-- ----------------------------
create table device_info
(
    id                bigint                  not null comment '自增ID'
        primary key,
    device_sn         varchar(32)             null comment '设备ID',
    name              varchar(32)             null comment '设备名称',
    gb_id             varchar(32)             null comment '设备国标ID',
    group_id          varchar(32)             null comment '设备分组ID',
    user_group_id     varchar(32)             null comment '用户分组id',
    device_group_name varchar(32)             null comment '设备分组名称',
    product_type      varchar(32)             null comment '产品类型 1:计算盒子；2:布控球；3:AI枪机 4:普通摄像头 5：服务器',
    product_img_url   varchar(128)            null comment '产品图片存储地址',
    status            int                     null comment '设备状态 1:在线；2:不在线；3:未激活；4暂停服务',
    location          varchar(32)             null comment '设备安装位置',
    created_time      datetime                null comment '创建时间',
    update_time       datetime                null comment '更新时间',
    delete_tag        int         default 0   null comment '删除 0:正常；1:已删除；',
    alternate_field1  varchar(32) default '3' null comment '2:已激活;3:未激活；4暂停服务',
    type              int                     null comment '设备类型',
    uuid              varchar(64)             null comment '设备UUID',
    is_open           int         default 1   null comment '是否开启(1开启；2关闭)',
    licence           varchar(255)            null comment '许可证',
    airia_ship        varchar(16)             null comment '单兵中airiaShip版本号'
)comment '设备信息表 ' collate = utf8mb4_bin;

create index index_device_id
    on device_info (device_sn);

create index index_group_id
    on device_info (group_id);
-- ----------------------------
-- Table structure for device_group
-- ----------------------------
create table device_group
(
    id                bigint        not null comment '自增ID'
        primary key,
    device_group_id   varchar(32)   null comment '设备分组ID',
    user_group_id     varchar(32)   null comment '用户分组id',
    device_group_name varchar(32)   null comment '设备分组名称',
    device_group_note varchar(128)  null comment '设备分组描述',
    created_time      datetime      null comment '创建时间',
    update_time       datetime      null comment '更新时间',
    delete_tag        int default 0 null comment '删除 0:正常；1：已删除；'
)comment '设备分组表 ' collate = utf8mb4_bin;

create index index_device_group_id
    on device_group (device_group_id);
-- ----------------------------
-- Table structure for app_info
-- ----------------------------
create table app_info
(
    id           int auto_increment comment '自增ID'
        primary key,
    apk_name     varchar(64)       null comment 'apk文件名称',
    apk_url      varchar(128)      null comment 'apk文件url',
    app_img_url  varchar(128)      null comment 'app图标url',
    support_chip varchar(64)       null comment '支持设备芯片 天玑800等',
    note         varchar(512)      null comment '备注',
    manual_url   varchar(128)      null comment '功能手册url',
    created_man  varchar(128)      null comment '上传人',
    update_man   varchar(128)      null comment '更新人',
    delete_tag   int     default 0 null comment '删除标记 0:正常；1：已删除；',
    created_time datetime          null comment '创建时间',
    update_time  datetime          null comment '更新时间',
    version      varchar(64)       null,
    version_id   int               null comment 'apk版本号id',
    file_type    varchar(32)       null,
    category_id  tinyint default 0 null comment '应用分类 0未定义 1边缘应用 2移动单兵app 3物联管理平台 4物联管理app'
)comment '应用信息表' collate = utf8mb4_bin;

create index index_apk_name
    on app_info (apk_name);
-- ----------------------------
-- Table structure for device_event_info
-- ----------------------------
create table device_event_info
(
    id              int auto_increment comment '自增ID 用于标识事件'
        primary key,
    device_id       varchar(32)  not null comment '设备ID',
    device_name     varchar(32)  null comment '设备名称',
    event_type_name varchar(32)  null comment '事件类型',
    event_note      varchar(128) null comment '事件描述 事件名称',
    event_status    int          null comment '事件状态',
    dispose_time    varchar(32)  null comment '处理时间',
    created_time    datetime     not null comment '创建时间',
    update_time     datetime     null comment '更新时间',
    time_stamp      int          null comment '事件上传时间戳'
)comment '设备事件信息表 ' collate = utf8mb4_bin;

create index index_device_id
    on device_event_info (device_id);

create index index_event_note
    on device_event_info (event_note);

create index index_event_type
    on device_event_info (event_type_name);
-- ----------------------------
-- Table structure for request_record
-- ----------------------------
create table request_record
(
    id             bigint                                                                       not null comment 'ID'
        primary key,
    req_method     enum ('GET', 'POST', 'PUT', 'DELETE', 'HEAD', 'OPTIONS', 'TRACE', 'CONNECT') not null comment '请求方式',
    req_url        varchar(1024)                                                                not null comment '请求路由',
    req_param      varchar(4096)                                                                null comment '请求参数',
    code           int                                                                          not null comment '状态码',
    client_ip      varchar(200)                                                                 null comment '请求IP',
    user_name      varchar(200)                                                                 null comment '操作用户名字',
    start_time     datetime                                                                     null comment '请求开始时间',
    execution_time bigint                                                                       null comment '请求执行时间(ms)',
    delete_tag     int unsigned default '0'                                                     null comment '删除标记 0:正常；1：已删除',
    create_time    datetime                                                                     null comment '创建时间'
)comment '项目请求记录表' collate = utf8mb4_bin;

-- ----------------------------
-- Table structure for user_authorization
-- ----------------------------
create table user_authorization
(
    id              int auto_increment comment '自增ID'
        primary key,
    device_id       varchar(32)   null comment '设备ID',
    device_name     varchar(255)  null comment '设备名称',
    device_group_id varchar(64)   null comment '设备组ID',
    user_id         varchar(32)   null comment '用户id',
    created_time    datetime      null comment '创建授权时间',
    delete_time     datetime      null comment '删除授权时间',
    update_time     datetime      null comment '更新授权时间',
    delete_tag      int           null comment '删除标记 0:正常；1：已删除；',
    user_group_id   varbinary(64) null comment '用户组ID'
)comment '用户授权记录表 ' collate = utf8mb4_bin;

create index index_device_id
    on user_authorization (device_id);

create index index_user_id
    on user_authorization (user_id);
-- ----------------------------
-- Table structure for product_info
-- ----------------------------
create table product_info
(
    id             int auto_increment comment '自增ID'
        primary key,
    product_design varchar(32)   null comment '外观：手机、盒子、服务器',
    product_type   varchar(32)   null comment '产品型号：FWB01',
    product_logo   varchar(32)   null comment '图片存储地址',
    product_name   varchar(128)  null comment '产品名字：防爆手机',
    delete_tag     int default 0 null comment '删除标记 0:正常；1：已删除；',
    created_time   datetime      null comment '创建时间',
    update_time    datetime      null comment '更新时间',
    xpu_type       varchar(32)   null comment '芯片名称',
    xpu_version    varchar(255)  null comment '芯片版本',
    product_level  varchar(32)   null comment '产品类型(英文，用于映射前端所需字段)'
)comment '产品信息表 ' collate = utf8mb4_bin;