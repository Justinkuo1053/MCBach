# 🎵 MCBach - 音樂評論神器 

> 一個讓音樂愛好者暢所欲言的超讚平台！🔥

[![Go Version](https://img.shields.io/badge/Go-1.23+-00ADD8?style=for-the-badge&logo=go)](https://golang.org/)
[![Gin Framework](https://img.shields.io/badge/Gin-Web%20Framework-00ADD8?style=for-the-badge)](https://gin-gonic.com/)
[![Spotify API](https://img.shields.io/badge/Spotify-API-1ED760?style=for-the-badge&logo=spotify&logoColor=white)](https://developer.spotify.com/)
[![MySQL](https://img.shields.io/badge/MySQL-Database-4479A1?style=for-the-badge&logo=mysql&logoColor=white)](https://www.mysql.com/)

## 🌟 這是什麼？

想像一下 **Instagram** 遇上了 **Spotify**，再加上一點 **Reddit** 的討論氛圍 - 這就是 MCBach！🎯

一個專為音樂發燒友打造的評論平台，讓你可以：
- 🔍 探索最新音樂專輯
- 💬 對喜愛的專輯發表熱辣評論
- 👍 給神評論按讚（還有專業讚！）
- 🏆 參與評論大賽，成為音樂評論界的傳奇
- 📊 用數據證明你的音樂品味有多讚
- 🎵 直接連結到 Spotify 聽音樂

## ✨ 超酷功能

### 🎸 音樂探索
- **新專輯發現**：即時同步 Spotify 最新發行
- **藝人資訊**：完整的藝人專輯收藏
- **智能推薦**：根據你的喜好推薦音樂

### 💭 社群互動
- **評論功能**：想說什麼就說什麼！
- **雙重按讚系統**：
  - 👍 普通讚：表達喜愛
  - 🏅 專業讚：值 2 倍分數的專業認可
- **編輯評論**：說錯話了？沒關係，可以編輯！
- **🏆 評論大賽系統**：
  - 📊 智能統計每位用戶的平均按讚數
  - 🥇 定期舉辦「音樂評論王」競賽
  - 📈 追蹤評論品質趨勢
  - 🎯 發掘最具影響力的音樂評論家

### 🔐 用戶系統
- **註冊/登入**：JWT 加密，安全可靠
- **會員制度**：只有會員才能按讚（防止灌水）
- **個人資料**：展示你的音樂品味

## 🚀 快速開始

### 前置需求
```bash
Go 1.23+
MySQL 8.0+
Spotify Developer Account
```

### 1️⃣ Clone 專案
```bash
git clone https://github.com/Justinkuo1053/MCBach.git
cd MCBach
```

### 2️⃣ 安裝依賴
```bash
go mod download
```

### 3️⃣ 環境配置
建立 `.env` 檔案：
```env
# 資料庫配置
DATABASE_URL=user:password@tcp(localhost:3306)/mcbach?charset=utf8mb4&parseTime=True&loc=Local

# Spotify API
SPOTIFY_CLIENT_ID=your_spotify_client_id
SPOTIFY_CLIENT_SECRET=your_spotify_client_secret

# JWT Secret
JWT_SECRET=your_super_secret_key
```

### 4️⃣ 資料庫設定
```bash
# 建立資料庫
mysql -u root -p
CREATE DATABASE mcbach;
```

### 5️⃣ 啟動服務
```bash
cd cmd
go run main.go
```

🎉 **搞定！** 服務現在運行在 `http://localhost:8080`

## 📡 API 端點

### 🎵 音樂相關
```
GET  /spotify/new-releases    # 獲取最新專輯
GET  /albums                  # 瀏覽專輯列表
GET  /artists/:spotifyId      # 獲取藝人資訊
POST /artists                 # 批量創建藝人
```

### 💬 評論系統
```
GET  /comments?albumId=123    # 獲取專輯評論
POST /comments                # 發表評論
PUT  /comments/:id            # 編輯評論
DELETE /comments/:id          # 刪除評論
POST /comments/:id/like       # 普通按讚
POST /comments/:id/pro-like   # 專業按讚
GET  /comments/top-users      # 月度排行榜
GET  /comments/contest-stats  # 評論大賽統計數據
```

### 👤 用戶管理
```
POST /auth/signup            # 註冊
POST /auth/signin            # 登入
POST /auth/signout           # 登出
GET  /user/me                # 個人資料
PUT  /user/me                # 更新資料
```

## 🏗️ 專案架構

```
MCBach/
├── cmd/                    # 應用程式入口
│   ├── main.go            # 主程式
│   └── main_test.go       # 測試
├── internal/              # 內部模組
│   ├── album/             # 專輯模組
│   ├── artist/            # 藝人模組
│   ├── auth/              # 認證模組
│   ├── comment/           # 評論模組
│   ├── config/            # 配置模組
│   ├── relations/         # 關聯模組
│   ├── spotify/           # Spotify 整合
│   └── user/              # 用戶模組
├── go.mod                 # Go 模組
├── go.sum                 # 依賴鎖定
└── README.md             # 你正在看的這個！
```

每個模組都採用 **Clean Architecture** 設計：
- 📁 `controllers/` - HTTP 處理層
- 📁 `services/` - 業務邏輯層  
- 📁 `repositories/` - 資料存取層
- 📁 `models/` - 資料模型

## 🛠️ 技術棧

| 類別 | 技術 | 為什麼選它？ |
|------|------|-------------|
| **後端框架** | Gin | 🚀 超快速、輕量級 |
| **資料庫** | MySQL + GORM | 🗄️ 可靠穩定、ORM 好用 |
| **認證** | JWT | 🔐 無狀態、安全 |
| **外部 API** | Spotify Web API | 🎵 最豐富的音樂資料 |
| **配置管理** | godotenv | ⚙️ 環境變數管理 |

## 🎨 使用範例

### 發表評論
```bash
curl -X POST http://localhost:8080/comments \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer your_jwt_token" \
  -d '{
    "albumId": 123,
    "content": "這張專輯太神了！每首歌都是經典 🔥"
  }'
```

### 專業按讚
```bash
curl -X POST http://localhost:8080/comments/456/pro-like \
  -H "Authorization: Bearer your_jwt_token"
```

### 查看評論大賽排名
```bash
curl -X GET http://localhost:8080/comments/contest-stats \
  -H "Authorization: Bearer your_jwt_token"
```

## 🤝 貢獻指南

想要讓 MCBach 變得更棒嗎？歡迎你的貢獻！

1. 🍴 Fork 這個專案
2. 🌿 建立你的功能分支 (`git checkout -b feature/AmazingFeature`)
3. 💫 提交你的改動 (`git commit -m 'Add some AmazingFeature'`)
4. 📤 推送到分支 (`git push origin feature/AmazingFeature`)
5. 🎯 開啟一個 Pull Request

## 📋 TODO List

- [ ] 🏆 **評論大賽進階功能**
  - [ ] 平均按讚數計算演算法
  - [ ] 季度/年度評論王評選
  - [ ] 評論品質分析儀表板
  - [ ] 大賽獎勵機制
- [ ] 🔔 即時通知系統
- [ ] 📱 RESTful API 文檔
- [ ] 🎭 用戶頭像上傳
- [ ] 🏷️ 音樂標籤系統
- [ ] 📊 評論情感分析
- [ ] 🌐 國際化支援
- [ ] 🐳 Docker 容器化
- [ ] ☁️ 雲端部署指南

## 🐛 遇到問題？

1. 檢查 `.env` 檔案配置
2. 確認 MySQL 服務運行正常
3. 驗證 Spotify API 憑證
2. 查看 [Issues](https://github.com/Justinkuo1053/MCBach/issues) 頁面
5. 還是不行？開個新 Issue 吧！

## 📜 授權條款

本專案採用 MIT 授權條款 - 詳見 [LICENSE](LICENSE) 檔案

---

<div align="center">

**用 ❤️ 和 ☕ 製作**

如果這個專案對你有幫助，給個 ⭐ 吧！

[🐛 回報問題](https://github.com/Justinkuo1053/MCBach/issues) • [💡 功能建議](https://github.com/Justinkuo1053/MCBach/issues) • [📖 文檔](https://github.com/Justinkuo1053/MCBach/wiki)

</div>
