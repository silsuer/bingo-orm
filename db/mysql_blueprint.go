package db

import (
	"strconv"
	"strings"
)

const (
	TypeIncrements           = "increments"
	TypeInteger              = "int"
	TypeString               = "varchar"
	TypeStringDefaultLength  = 255
	TypeIntegerDefaultLength = 11

	CreateDefaultType = iota
	CreateIfNotExists
)

// mysql实现 blueprint
type MysqlBlueprint struct {
	Blueprint
	name       string // 表名
	engine     string
	columns    []*MysqlColumn
	currentCol *MysqlColumn
	charset    string
}

type MysqlColumn struct {
	name         string // 列名
	columnType   string // 列类型
	defaultValue string // 默认值
	comment      string // 备注
	nullable     bool   // 是否允许为null
	length       int    // 列的长度
}

func NewColumn(name string) *MysqlColumn {
	c := new(MysqlColumn)
	c.nullable = false
	c.name = name
	return c
}

func (mb *MysqlBlueprint) Increments(columnName string) IBlueprint {
	// 创建一个列
	column := NewColumn(columnName)
	column.length = TypeIntegerDefaultLength
	column.columnType = TypeIncrements
	mb.columns = append(mb.columns, column)
	mb.currentCol = column // 设置当前的列名
	return mb
}

func (mb *MysqlBlueprint) Nullable() IBlueprint {
	mb.currentCol.nullable = true
	return mb
}

func (mb *MysqlBlueprint) Comment(comment string) IBlueprint {
	mb.currentCol.comment = comment
	return mb
}

func (mb *MysqlBlueprint) Default(def interface{}) IBlueprint {
	switch def.(type) {
	case int:
		if mb.currentCol.columnType == TypeInteger {
			mb.currentCol.defaultValue = strconv.Itoa(def.(int))
		} else {
			panic("the default value of column " + mb.currentCol.name + " need " + TypeInteger)
		}
	case string:
		mb.currentCol.defaultValue = def.(string)
	default:
		panic("wrong default value type : column " + mb.currentCol.name)
	}
	return mb
}

func (mb *MysqlBlueprint) String(columnName string) IBlueprint {
	column := NewColumn(columnName)
	column.length = TypeStringDefaultLength
	// 字符串类型
	column.columnType = TypeString
	mb.columns = append(mb.columns, column)
	mb.currentCol = column
	return mb
}

func (mb *MysqlBlueprint) StringWithLength(columnName string, length int) IBlueprint {
	column := NewColumn(columnName)
	column.length = length
	column.columnType = TypeString
	mb.columns = append(mb.columns, column)
	mb.currentCol = column
	return mb
}

func (mb *MysqlBlueprint) Integer(columnName string) IBlueprint {
	column := NewColumn(columnName)
	column.length = TypeIntegerDefaultLength
	column.columnType = TypeInteger
	mb.columns = append(mb.columns, column)
	mb.currentCol = column
	return mb
}

func (mb *MysqlBlueprint) Charset(charset string) IBlueprint {
	mb.charset = charset
	return mb
}

func (mb *MysqlBlueprint) IntegerWithLength(columnName string, length int) IBlueprint {
	column := NewColumn(columnName)
	column.length = length
	column.columnType = TypeInteger
	mb.columns = append(mb.columns, column)
	mb.currentCol = column
	return mb
}

// 设置引擎
func (mb *MysqlBlueprint) Engine(engine string) IBlueprint {
	mb.engine = engine
	return mb
}

// 将表结构拼装成sql语句
func Assembly(createType int, schema *MysqlBlueprint) string {
	var sql string
	if createType == CreateDefaultType {
		sql = `CREATE TABLE ` + schema.name + "("
	} else if createType == CreateIfNotExists {
		sql = `CREATE TABLE IF NOT EXISTS ` + schema.name + "("
	}

	var columnSql []string
	// 拼接列
	for k := range schema.columns {
		columnSql = append(columnSql, columnAssembly(schema.columns[k]))
	}
	sql += strings.Join(columnSql, ",") + ")ENGINE=" + schema.engine

	if schema.charset != "" {
		sql += " DEFAULT CHARSET=" + schema.charset
	}

	return sql
}

// 将列拼接成语句
func columnAssembly(column *MysqlColumn) string {
	var sql string

	switch column.columnType {
	case TypeIncrements:
		sql = column.name + ` INT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT `
	case TypeString:
		// name varchar(11) not null default "dd"
		sql = column.name + ` varchar(` + strconv.Itoa(column.length) + `)`
		if !column.nullable {
			sql += " not null "
		}
		if column.defaultValue != "" {
			// 默认值
			sql += " default \"" + column.defaultValue + "\""
		}
	case TypeInteger:
		// age int(11) not null default 0
		sql = column.name + ` int(` + strconv.Itoa(column.length) + `)`
		if !column.nullable {
			sql += " not null "
		}
		if column.defaultValue != "" {
			sql += " default " + column.defaultValue
		}
	}

	// 添加comment
	if column.comment != "" {
		sql += ` comment "` + column.comment + `"`
	}

	return sql
}
