# skills-validate-go

skills-validate-go 是一個用 Go 語言實作的 Agent Skills 驗證工具。

> **注意：** 本程式庫僅供展示用途，並不建議用於生產環境。

## 功能特色

- ✅ 驗證技能（skill）目錄
- ✅ 讀取並解析技能屬性
- ✅ 產生 agent prompt XML

## 安裝方式

### 使用 go install

```bash
go install github.com/jhaoheng/skills-validate-go/cmd/skills-validate@latest
```

### 從原始碼安裝

```bash
git clone https://github.com/jhaoheng/skills-validate-go.git
cd skills-validate-go
make install
```

### 下載執行檔

請至 [releases page](https://github.com/jhaoheng/skills-validate-go/releases) 下載最新版本。

## 使用說明

### CLI 指令

```bash
# 驗證一個技能
skills-validate validate path/to/skill

# 讀取技能屬性（輸出 JSON）
skills-validate read-properties path/to/skill

# 產生 <available_skills> XML 給 agent prompt 使用
skills-validate to-prompt path/to/skill-a path/to/skill-b
```

### Go API 範例

```go
package main

import (
    "fmt"
    "github.com/jhaoheng/skills-validate-go/pkg/skillsref"
)

func main() {
    // 驗證技能目錄
    problems, err := skillsref.Validate("my-skill")
    if err != nil {
        panic(err)
    }
    if len(problems) > 0 {
        fmt.Println("驗證錯誤:", problems)
    }

    // 讀取技能屬性
    props, err := skillsref.ReadProperties("my-skill")
    if err != nil {
        panic(err)
    }
    fmt.Printf("技能: %s - %s\n", props.Name, props.Description)

    // 產生可用技能的 prompt
    prompt, err := skillsref.ToPrompt([]string{"skill-a", "skill-b"})
    if err != nil {
        panic(err)
    }
    fmt.Println(prompt)
}
```

## 開發相關

### 先決條件

- Go 1.21 或以上版本

### 建置指令

```bash
# 編譯執行檔
make build

# 執行測試
make test

# 執行測試並產生覆蓋率報告
make test-coverage

# 執行程式碼靜態檢查
make lint

# 格式化程式碼
make fmt
```

### 專案結構

```
skills-validate-go/
├── cmd/
│   └── skills-validate/    # CLI 進入點
├── internal/
│   ├── parser/            # SKILL.md 解析器
│   ├── validator/         # 驗證邏輯
│   ├── prompt/            # Prompt 產生
│   ├── models/            # 資料模型
│   └── errors/            # 錯誤定義
├── pkg/
│   └── skillsref/         # 公開 API
└── testdata/              # 測試用資料
```

## 本地測試

你可以在專案根目錄執行下列指令進行本地測試：

```bash
go test -v -race -coverprofile=coverage.out ./...
go tool cover -func=coverage.out
go tool cover -html=coverage.out
```

Makefile：

```bash
make test
make test-coverage
```

## 授權

Apache 2.0

## 貢獻

歡迎貢獻！請隨時提交 Pull Request。
