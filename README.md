# Ordering System (Backend)

---

## 作品介紹

### 簡介

我與另一名前端工程師合作並設計並開發的一款餐廳點餐系統，分為 **商店管理系統（Store Ease 商店輕鬆理）** 與 **快速點餐系統（Order Ease 餐點輕鬆訂）** 兩部份；而這個 Project 是整個系統的後端架構。

- **商店輕鬆理**
  - 功能特色：
    - 帳戶管理： 以 E-Mail 註冊，並於後續更改使用者資訊或刪除用戶
    - 商店管理： 輕鬆管理多個商店資訊、菜單和座位配置。
    - 訂單管理： 生成 QRCode 提供給客戶，讓他們可以快速點餐，並實時接收訂單信息。
    - 即時通訊： 透過 FCM，在客戶點餐時即時接收訂單資訊，
  - 未來展望：
    整合理財系統，進行營業額統計與視覺化呈現。
- **餐點輕鬆訂**
  - 功能特色：
    - 帳戶管理： 可以 E-Mail 實名或匿名註冊，並於後續更改使用者資訊或刪除用戶
    - 快速點餐： 掃描商店生成的 QRCode，即時獲取商店菜單並快速點餐。
    - 訂單即時更新： 透過 FCM，在商店有任何訂單狀態更新時及時回饋給客戶。
    - 訂單紀錄： 紀錄所有的訂單歷史
  - 未來展望： 整合理財系統，實現預算控管與視覺化呈現。也可將訂單歷史的商家地理位置做成地圖，讓點餐也能紀錄用戶的行徑軌跡，增添趣味性。

### 開發時長

- 後端系統（:file_folder: `app`）包含設計規劃與溝通，從無到建立，開發總時長**為 8 週**。

## 負責項目

![Ordering System v1.0 Database Structure](screenshots/ordering_system_database_structure.png)

- Scrum 專案管理
- 具彈性化的資料庫與系統架構設計
- Gmail SMTP 設置
- 與前端工程師溝通，設計並製作所有 RESTful APIs
  - 多國語言資料管理切換
  - 圖片上傳
  - OTP 寄送、驗證
  - 新增/讀取/編輯/刪除 **用戶**資訊（搭配 Firebase Auth）
  - 新增/讀取/編輯/刪除 **商店**、**菜單**、**類別**、**座位**、**訂單**資訊
  - 串接 FCM，即時發送與回傳**訂單狀態資訊**
- 雲端專案佈署
  - Google Cloud Run
  - MySQL (Railway)

## 技術清單

### Language

- **Golang** `v1.20` |

  - **Main Dependencies**

    - Gin-Gonic (`github.com/gin-gonic/gin` v1.9.1)
    - MySQL & GORM (`gorm.io/driver/mysql` v1.5.1, `gorm.io/datatypes` v1.2.0, `gorm.io/gorm` v1.25.4)

  - **All Dependencies**

    ![All dependencies in backend system](screenshots/dependencies.png)

### Cloud Services

- **Server**
  - Google Cloud Run
- **Database**
  - MySQL (Deploy on **Railway**)
- **Others**
  - Gmail SMTP (OTP mail sender)
  - Firebase Auth
  - Firebase Cloud Messaging (FCM)

### Other Tools

- **Git** (Version control)
- **Docker** (Container)
- **Postman** (API tool)
- **Trello** (Scrum pattern)
- **DrawSQL** (Database design tool)
