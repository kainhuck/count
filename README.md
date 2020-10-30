# count

## count是一个统计项目代码行数的工具，它支持多种参数
1. 指定后缀
2. 指定是否忽略空行
3. 指定忽略后缀

# Useage
```bash
对代码项目进行行数的统计

Usage:
  count [flags]

Flags:
  -b, --but string      指定不统计的文件后缀 (default "exe")
  -e, --empty           是否忽略空行
  -h, --help            help for count
  -i, --hide            是否忽略隐藏文件 (default true)
  -p, --path string     指定项目路径 (default ".")
  -s, --suffix string   请指定后缀，省略查找全部
  -v, --view            是否显示过程
```
