# MySQL



|     列名      |                             描述                             |
| :-----------: | :----------------------------------------------------------: |
|      id       | 查询语句中每出现一个 select 关键字，MySQL 就会为它分配一个唯一的 id 值，某些子查询会被优化为 join 查询 |
|  select_type  |              SELECT 关键字对应的那个查询的类型               |
|     table     |                             表名                             |
|  partitions   |                        匹配的分区信息                        |
|     type      |             针对单表的查询方式（全表扫描、索引）             |
| possible_keys |                        可能用到的索引                        |
|      key      |                        实际用到的索引                        |
|    key_len    |                      实际用到的索引长度                      |
|      ref      |                          当使用索引                          |





## 海量数据下如何查找一条数据



1. 使用布隆过滤器，快速过滤掉不存在的记录。Redis 情况下可以用 bitmap 结构来实现布隆过滤器。

2. 在 Redis 中建立数据缓存，将我们对 Redis 使用场景的理解，尽量表达出来。可以以普通字符串的小时来存储

   例如 userId-->user.json，以一个hash来存储一条记录（userId -> username -> userAge）。以一个整的 hash来存储所有的数据，

   Userinfo -> field 就用 userId，value 就用 user.json。一个hash最多能支持 2^32-1（40多个亿）个键值对

3. 



## 索引



**普通索引**：允许被索引的数据列包含重复的值

**唯一索引**：可以保证数据记录的唯一性 （`UNIQUE` `KEY` catename (catid) ）

**主键**：是一种特殊的唯一索引，在一张表中只能定义一个主键索引，主键用于唯一标识一条记录，使用关键字 PRIMARY KEY 来创建

**联合索引**：索引可以覆盖多个数据列，如像 INDEX (column A, column B) 索引

**全文索引**：通过建立倒排索引，可以极大的提升检索效率，解决判断字段是否包含的问题，是目前搜索引擎使用的一种关键技术。可以通过 ALTER TABLE table_name ADD FULLTEXT (column) 来创建全文索引

**最左原则**：

```mysql
CREATE TABLE users (
  last_name varchar(50) not null,
  first_name varchar(50) not null,
  birthday date not null,
  gender enum('m', 'f') not null,
  key(last_name, first_name, birthday)
);
```

该表创建了三个组合索引

username，city，birthday

username，city

username

1. 查询必须从索引的最左边的列开始，否则无法使用到索引。例如下面这条语句就使用不到索引

   ```mysql
   SELECT * FROM users WHERE birthday = NOW();
   ```

2. 存储引擎不能使用索引中范围条件右边的列。例如以下语句，LIKE 是范围查询

   ```mysql
   SELECT * FROM users WHERE last_name = 'Simith' AND first_name LIKE '%Chuck%' AND birthday = NOW();
   ```

3. 不能跳过某一索引列



索引可以极大的提高数据的查询速度，通过索引可以在查询过程中，使用优化隐藏器，提高系统的性能。

但是会降低 insert、delete、update 的速度，因为在执行这些写操作时，还要操作索引文件。

索引需要占物理空间，除了数据表占数据空间之外，每一个索引还要占一定的物理空间，如果要建立聚簇索引，那么需要占据的空间就会更大，如果非聚簇索引很多，一旦聚集索引改变，那么所有肥聚集索引都会跟着改变。



**索引覆盖**：

索引覆盖就是一个 SQL 在执行时，可以利用索引来快速查找，并且此 SQL 所要查询的字段在当前索引对应的字段中都包括了，那么就表示此 SQL 走完索引之后就不用回表了，所需要的字段都在当前索引的叶子节点上存在，可以直接作为结果返回了。



## 数据一致性

1. 先更新 MySQL，再更新 Redis，如果更新 Redis 失败，可能仍然不一致
2. 先删除 Redis 缓存数据，再更新 MySQL，再次查询到的时候将数据添加到缓存中，这种方案能解决1方案的问题，但是在高并发下性能较低，而且仍然可能会出现数据不一致的问题，比如线程1删除了 Redis 缓存数据，正在更新 MySQL，此时另一个线程再查询，那么就会把 MySQL 中的老数据又查到 Redis 中
3. 延时双删，步骤是：先删除 Redis 缓存数据，再更新 MySQL，延迟几百毫秒再删除 Redis 缓存数据，这样就算在更新 MySQL 时，有其他线程读了 MySQL，把老数据读到了 Redis 中，那么也会被删除掉，从而保持数据一致













































