# 2

## 設定檔案

### 尋找路徑
1. 命令列指定路徑
2. 當前工作目錄底下的 `filelist-generator.json`
3. 使用默認設定

### 忽略路徑
使用 [path.Match](https://pkg.go.dev/path#Match) 匹配路徑，如果匹配就會被忽略。