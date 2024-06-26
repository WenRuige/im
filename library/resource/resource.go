package resource

import (
	"database/sql"
	"go.uber.org/zap"
)

// IMDB 初始化数据库
var IMDB *sql.DB

// Logger 初始化zapper
var Logger *zap.Logger
