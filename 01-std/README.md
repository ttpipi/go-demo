## slog
- 真正的使用应该是 slog.New() 得到一个 logger slog.Logger, 然后通过 logger调用Log, Debug, Info等public方法
- 通过slog.Log, slog.Info等包函数直接使用, 其实是为了用户使用方法, 封装了一个defaultLogger

## interface