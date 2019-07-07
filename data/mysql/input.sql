drop database if exists AliAuction;
create database AliAuction;
use AliAuction;


#创建缓存表
create table t_cache(
	cacheId integer not null auto_increment,
	name varchar(128) not null,
	value mediumtext not null,
	createTime datetime not null default CURRENT_TIMESTAMP,
	modifyTime datetime not null default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP, 
	primary key( cacheId )
)engine=innodb default charset=utf8mb4 auto_increment = 10001;

alter table t_cache add unique index nameIndex(name);