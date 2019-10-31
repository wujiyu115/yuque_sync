# 语雀文档同步

# 说明
`front matter`解析来自[hugo](https://github.com/gohugoio/hugo/tree/master/parser),去除了`GRO`和`CSV`的支持,现在支持`YAML`,`JSON`,`TOML`


配置支持`json`和`yaml`
`json`配置示例:

```json
{
    //绝对路径,文章下载存放路径(可选)
    "postPath": "yuque",
    //仓库文章信息缓存(可选)
    "cachePath": "yuque.json",
    //使用语雀文章哪个字段作为文章名(可选Title|Slug)
    "mdFormat": "Title",
    //语雀token(必填)
    "token": "token",
    //语雀登录账户别名(必填)
    "login": "u22579",
    //语雀仓库名(必填)
    "repo": "xcd0mr",
    //格式化文章(必填markdown|hexo|hugo)
    "adapter": "markdown",
    "concurrency": 5,
    //是否只拉取发布过的文章
    "onlyPub": true
}
```

# 使用

配置好`config.json`或者`config.yaml`,编译`yuque_sync`,执行`./yuque_sync`

# TODO

- [x] 同步markdown到文件夹
- [ ] hexo格式转换
- [ ] hugo格式转换
- [ ] 支持语雀多仓库同步到不同的分类
- [X] 解析front-matter

