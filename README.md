# Collection

> 数据集，作为其他项目的子模块









### git submodule



1. 添加

```
git submodule add URL PATH

cat .gitmodules

```

2. 查看

```

git submodule

```

3. 获取包含子模块的仓库

```

git clone URL

git submodule init
git submodule update


// OR

git clone --recursive URL 

```


4. 拉取更新过子模块

```
cd SUBMODULE_DATA

git status -s
git submodule update

```

5. 在包含子项目的仓库中更新子模块

```
cd SUBMODULE_DATA

git add .
git commit -m "MESSAGE"

git push origin master

// 即：和是一个托管项目操作一致

```


6. 移除子模块


