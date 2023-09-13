# ordering-system

https://drawsql.app/teams/pipis-team/diagrams/ordering-db

## 商家

- 註冊
- 登入
- 商家資料(RU)
- 菜單管理(CRUD)
- 歷史紀錄(U)
- 報表統計

## 客戶

- 註冊
- 登入
- 獲取商家列表(R)
- 獲取商家資料(R)
- 獲取菜單列表(R)
- 訂單管理(CRUD)
- 歷史紀錄(R)

## 專案情境

<!-- - As a [type of user], I want [an action] so that [a benefit/a value] -->
<!-- - 作為一名註冊用戶，我希望可以重設我的密碼，這樣當我忘記密碼時可以重新訪問我的賬戶。 -->

### Sprint 1
- 作為一名商家用戶，我希望可以註冊帳號與透過驗證，並且填寫我的註冊資訊（Email、Password）、基本資訊（ID、UID、Name、Phone、Address），才能讓客戶方便認識我。
後續可填（Description、OpeningHours）
<!-- - 作為一名顧客用戶，我希望可以註冊帳號與透過驗證，並且填寫我的註冊資訊（Email、Password）、基本資訊（ID、UID、Name），才能進行點餐。 -->
```
UI順序
Email（按下登入，確認Firebase帳號有沒有註冊過、沒有註冊的話要發驗證6碼OTP信） -> 驗證碼 -> Password + 基本資訊（Name、Phone、Address）
```
- 作為一名商家用戶，我希望可以填寫Email、Password登入帳號，才能使用這個App
```
UI順序
Email（按下登入，確認Firebase帳號有沒有註冊過、有註冊過才會進到下一個頁面） -> Password
```

### Sprint 2
- 作為一名商家用戶，我希望可以新增一筆 menu，包含菜單的標題與描述、有哪些多少商品、價格多少，這樣我才能讓使用者看見我的菜單並執行後續的操作
- 作為一名商家用戶，我希望可以獲取自己的 menu 列表，這樣我才能清楚的看見我目前擁有的所有菜單
- 作為一名商家用戶，我希望可以在我新增完 menu 後點進去看到細節描述，包含這個菜單的標題與描述、裡面有多少商品、價格多少，這樣我才能確認資訊是否正確
- 作為一名商家用戶，我希望我的 menu 可以修改，這樣我才能在有錯誤發生時或有活動時更新我的菜單
- 作為一名商家用戶，我希望我的 menu 可以刪除，這樣我才能避免存在過多已經沒在使用的菜單
<!-- - 作為一名顧客用戶，我希望我可以看見所有的店家列表，好讓我挑選我要的餐廳
- 作為一名顧客用戶，我希望我可以看見指定的店家資訊，好讓我挑選我要的餐點 -->

### Sprint 3

<!-- - 作為一名顧客用戶，我希望我可以對喜歡的餐點進行點餐，這樣商家才能收到我的訂購資訊
- 作為一名顧客用戶，我希望我可以刪除我的訂單，這樣在我誤觸送單的時候才可以挽回 -->
- 作為一名商家用戶，我希望在顧客下訂單後，我可以即時收到顧客下訂的資訊，這樣我才能開始製作我的餐點
- 作為一名商家用戶，我希望在開始製作餐點前，可以將製作狀態更改為製作中，這樣顧客才能看見自己餐點的進度，並且避免顧客取消訂單
