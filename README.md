### 自动生成model文件
gentool -dsn "user:pwd@tcp(127.0.0.1:3306)/database?charset=utf8mb4&parseTime=True&loc=Local" -tables "student"  -onlyModel
gormgen --dsn "mors:Htht@2023pg@tcp(1.92.72.96:5432)/mors?currentSchema=mservice&charset=utf8mb4&parseTime=True" --dir models --pkg models --table alert
