## 常用命令

tail（尾巴）：从尾部开始展示文本

```bash
tail -n 100 test.log
tail -f test.log
```

grep：全局正则表达式搜索

```bash
grep '20:[0-5][0-9]:' *.log  #匹配当前目录下搜索log日志中，20点的日志
grep '20:[0-5][0-9]' 1.log 2.log 3.log  #指定在这三个文件中查找#grep规则是支持正则表达式的

ps -ef|grep java    #查找所有java进程,-c可以统计查找的个数  

grep '20:[1-5][0-9]:' *.log | grep -v '20:[3-4][0-9]:'   # -v反向选择，相当于过滤
grep 'ab|bc' *.log  #支持|语法，匹配含有ab或者bc的文本行
```

cut：从每个文件中剪出每行选定部分（由列表指定），并将它们写入标准输出

top：查看进程列表

diff：比较两个文件的不同

```bash
diff a.txt b.txt
```

tar：用来压缩和解压文件，tar 本身不具备压缩功能，只具有打包功能，压缩及解压是调用其他的功能来完成

```bash
tar -cvf test.tar test.txt # 打包 tar -cvf 包名 文件名
tar -xvf test.tar # 解包 tar -xvf 包名
tar -zcvf test.tgz test.txt # 压缩 tar -czvf 包名 文件名
tar -zxvf test.tgz # 解压 tar -xzvf 包名
```

man：查看命令的释义和用法

less/more：less 和 more 类似，但是用 less 可以随意浏览文件，而 more 仅能向前移动，却不能向后移动，而且less 在查看之前不会加载整个文件

```bash
less -m test.log
ps -a | less -N # -N 显示每行的行号 ps 查看进程信息并通过 less 分页显示
```

ls：查看文件夹

wc：统计指定文件中字节数、文字、行数，并将统计结果输出

- -l 统计行数
- -w 统计词数，一个字被定义为由空白、跳格或换行字符分隔的字符串
- -m 统计字符数
- -c 统计字节数

```bash
wc -l test.txt # 统计文件行数
```

查找文件

```bash
find / -name a.txt
whereis
which
locate
```











