# Intro
EMM-MoneyBox is A small tool for recording daily financial affairs.

# Configuration
$project_root/util/config_util.go
## Database
supported mongodb & mysql, read URI form os environment.

# Command
## Cash
### Query
emm-moneybox cash query -h
### Delete
emm-moneybox cash delete -h
### Outcome
refactoring now.
### Income
non-supported yet.
## Category
### add
non-supported yet.
### delete
non-supported yet.
### query
non-supported yet.
### update
non-supported yet.
## Manage
### export
refactoring now.
### import
refactoring now.

# API
non-supported yet.

# Thanks
- [cobra](https://github.com/spf13/cobra): a command-line program framework.
- [zap](https://github.com/uber-go/zap): for logging, specially json format.
- [excelize](https://github.com/qax-os/excelize): Excel handler for export&import.
- [decimal](github.com/shopspring/decimal): calculate the amount more accurately.
- [go-sql-driver](https://github.com/go-sql-driver/mysql): for MySQL DB support.
- [mongo-go-driver](https://github.com/mongodb/mongo-go-driver): for MongoDB support.

# Relate
May put some article here, now empty :lol
