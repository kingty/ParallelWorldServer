
--用户验证信息
create table user(

	user_id bigint(20) unsigned not null primary key auto_increment,
	email varchar(100) not null unique comment '邮箱',
	password varchar(50) not null comment '密码',
	register_time date not null comment '注册时间',
	

)engine=innod default charset=utf8 auto_increment=1;
--用户信息
--还差一个绑定微信号，不知道存什么字段
create table user_detail(
	user_detail_id bigint(20) unsigned not null primary key auto_increment,
	user_id bigint(20) unsigned not null comment '用户id',
	username varchar(100) not null comment '用户展示名',
	status int(1) default 0 comment '用户状态',
	gender int(1) default 0 comment '用户性别',
	url varchar(100)  comment '用户头像url',
	motto varchar(200) not null comment '用户格言，显示的一句话',
	stone_count int(11) default 0 comment '用户拥有的能量石数量'，
	get_count int(1) default 0 comment '每天抽取卡片的数量'，
	get_time date  comment '最后一次抽取卡片的时间'
)engine=innod default charset=utf8 auto_increment=1;



--卡片
create table card(
	card_id bigint(20) unsigned not null primary key auto_increment,
	user_id bigint(20) unsigned not null comment '用户id，这里存主要是为了让card直接能知道user,不再通过卡包去找',
	url varchar(100)  comment '卡片内容的图片url',
	title varchar(100) not null comment '卡片标题',
	content text not null comment '卡片正文',
	isprivate int(1) not null default 0 comment '公开或私密',
	city varchar(20) comment '发布时的定位城市',
	status int(1) default 0 comment '卡包状态',
	create_time date not null comment '创建时间'
)engine=innod default charset=utf8 auto_increment=1;


--卡包
create table package(
	package_id bigint(20) unsigned not null primary key auto_increment,
	user_id bigint(20) unsigned not null comment '用户id',
	url varchar(100)  comment '卡包封面URL',
	packagename varchar(100) not null comment '卡包名称',
	isprivate int(1) not null default 0 comment '是否为私密卡包'，
	status int(1) default 0 comment '卡包状态',
	create_time date not null comment '创建时间'
)engine=innod default charset=utf8 auto_increment=1;

--卡包和卡片关联表
create table package_card(
	package_card_id bigint(20) unsigned not null primary key auto_increment,
	card_id bigint(20) unsigned not null comment '卡片id',
	package_id bigint(20) unsigned not null comment '卡包id',
	category int(1) default 0 comment '收藏的还是自己建的'
)engine=innod default charset=utf8 auto_increment=1;




--主题卡包
create table theme(
	theme_id bigint(20) unsigned not null primary key auto_increment,
	url varchar(100)  comment '卡包封面URL',
	bgurl varchar(100)  comment '卡包背景URL',
	packagename varchar(100) not null comment '卡包名称',
	status int(1) default 0 comment '卡包状态',
	create_time date not null comment '创建时间'
)engine=innod default charset=utf8 auto_increment=1;


--主题卡包和卡片关联表
create table theme_card(
	them_card_id bigint(20) unsigned not null primary key auto_increment,
	card_id bigint(20) unsigned not null comment '卡片id',
	theme_package_id bigint(20) unsigned not null comment '主题卡包id',
)engine=innod default charset=utf8 auto_increment=1;


--关注关系表
create table follow(
	follow_id bigint(20) unsigned not null primary key auto_increment,
	follower_id bigint(20) unsigned not null comment '关注人ID',
	followeder_id bigint(20) unsigned not null comment '被关注人ID',
	isfollowed int(1) default 0 comment '是否也被它关注，这个字段只要是用于查询好友列表时，不需要再次查找是否被关注，提高查询效率，有待商榷',
	follow_time date not null comment '关注时间'
)engine=innod default charset=utf8 auto_increment=1;

--私信表
create table message(
	message_id bigint(20) unsigned not null primary key auto_increment,
	sender_id bigint(20) unsigned not null comment '发送者id',
	receiver_id bigint(20) unsigned not null comment '接收者id',
	card_id bigint(20) unsigned not null default 0 comment '卡片id,如果是送能量石自动产生的推送信息，保存id，默认为0',
	content text not null comment '消息内容', 
	send_time date not null comment '发送时间',
	status int(1) default 0 comment '消息状态，是否删除',
)engine=innod default charset=utf8 auto_increment=1;


--能量石记录表
create table stone_log(
	stone_log bigint(20) unsigned not null primary key auto_increment,
	user_id bigint(20) unsigned not null comment '用户id',
	log varchar(100) not null comment '日志信息，例如：送给  xxx 2颗能量石，xxx需要标记高亮' 
	log_time date not null comment '记录时间',
	stone_count int(11) default 0 comment '变动的能量石数'，
	status int(1) default 0 comment '状态'
)engine=innod default charset=utf8 auto_increment=1;



