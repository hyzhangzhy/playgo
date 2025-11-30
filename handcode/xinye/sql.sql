/*
写一条delete sql语句，删除name多余的重复记录（保留重复项最小id的记录），表名是tb；
id     name
1      张三
2      王二
3      张三
4      张三
5      哈哈
6      王二
7      小小
结果：
id     name
1      张三
2      王二
5      哈哈
7      小小
*/

select id from tb group by name;
delete from tb where