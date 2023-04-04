首先在终端中运行go tool pprof profile.pprof命令，其中profile.pprof是已经生成的CPU profile文件的路径。
然后，go tool pprof会启动一个交互式命令行界面。
在命令行中输入top可以查看CPU使用率最高的函数列表。