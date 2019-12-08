用來實作驗證API相關通用程式碼
- 產生 session
- 解開 session
- UserRole驗證

session
內含
- alias
- username
- user role
- user id
- expires times // session到期時間, 假設超過到期時間, 當用戶又呼叫 api時後端就可以知道是因為超時, 所以禁止呼叫. redis key的 expires也是設定這個時間.
