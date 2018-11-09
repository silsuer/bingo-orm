package db

type IBlueprint interface {
	Increments(columnName string, length ...int) IBlueprint
	Nullable() IBlueprint
	Comment(comment string) IBlueprint
	Default(def interface{}) IBlueprint
	String(columnName string, length ...int) IBlueprint
	Integer(columnName string, length ...int) IBlueprint
	BigIncreaments(columnName string, length ...int) IBlueprint      // 递增主键，相当于 unsigned big integer
	BigInteger(columnName string, length ...int) IBlueprint          // BIGINT
	Binary(columnName string) IBlueprint                             // BLOB
	Boolean(columnName string) IBlueprint                            // BOOLEAN
	Char(columnName string, length int) IBlueprint                   // char 带有长度
	Date(columnName string) IBlueprint                               // date
	DateTime(columnName string) IBlueprint                           // Datetime
	Decimal(columnName string, length int, precision int) IBlueprint // 带有精度与基数的 Decimal
	Double(columnName string, lengthPrecision ...int) IBlueprint     // 带有精度与基数的 Double
	Geometry(columnName string) IBlueprint                           // geometry
	GeometryCollection(columnName string) IBlueprint                 // geometryCollection
	IpAddress(columnName string) IBlueprint                          // IP地址
	Json(columnName string) IBlueprint                               // json
	LineString(columnName string) IBlueprint                         // linestring
	LongText(columnName string) IBlueprint                           // longtext
	MediumIncrements(columnName string, length ...int) IBlueprint    // 相当于 unsigned medium integer
	MediumInteger(columnName string, length ...int) IBlueprint
	Enum(columnName string, enum ...string) IBlueprint          // ENUM
	Float(columnName string, lengthPrecision ...int) IBlueprint // 带有精度的float
	MediumText(columnName string) IBlueprint                    // MEDIUMTEXT
	//Morphs(columnName string) IBlueprint                       // 相当于加入递增的 taggable_id 与 字符串 taggable_type
	MultiLineString(columnName string) IBlueprint // 相当于 MULTILINESTRING
	MultiPoint(columnName string) IBlueprint      // multipoint
	MultiPolygon(columnName string) IBlueprint    // multipolygon
	//NullableMorphs(columnName string) IBlueprint               // 相当于可空版本的 morphs字段
	NullableTimestamp(columnName string) IBlueprint                          // 相当于可空版本的 timestamps() 字段
	Point(columnName string) IBlueprint                                      // POINT
	Polygon(columnName string) IBlueprint                                    // Polygon
	RememberToken() IBlueprint                                               // 相当于可空版本的Varchar(100) 的remember_token 字段
	SmallIncrements(columnName string, length ...int) IBlueprint             // 递增id主键，相当于 unsigned small integer
	SmallInteger(columnName string, length ...int) IBlueprint                // smallint
	SoftDeletes() IBlueprint                                                 // 相当于为软删除添加一个可空的 deleted_at 字段
	Text(columnName string) IBlueprint                                       // 相当于 text
	Time(columnName string) IBlueprint                                       // time
	Timestamp(columnName string) IBlueprint                                  // 相当于 timestamp
	Timestamps() IBlueprint                                                  // 自动生成 created_at 和updated_at
	TinyIncrements(columnName string, length ...int) IBlueprint              // 相当于自动递增的 unsigned tinyint
	TinyInteger(columnName string, length ...int) IBlueprint                 // tinyint
	UnsignedBigInteger(columnName string, length ...int) IBlueprint          // 相当于 unsigned bigint
	UnsignedDecimal(columnName string, length int, precision int) IBlueprint // 相当于带有精度和基数的 unsigned decimal
	UnsignedInteger(columnName string, length ...int) IBlueprint             // unsigned int
	UnsignedMediumInteger(columnName string, length ...int) IBlueprint       // unsigned mediumint
	UnsignedSmallInteger(columnName string, length ...int) IBlueprint        // unsigned smallint
	UnsignedTinyInteger(columnName string, length ...int) IBlueprint         // unsigned tinyint
	//Uuid(columnName string) IBlueprint                         // uuid
	Year(columnName string) IBlueprint // YEAR
	Index(columns ...string)           // 传入多个列，给这些列添加索引
	UniqueIndex(column string)         // 唯一索引
	FullTextIndex(column string)       // 全文索引
	PrimaryKey(column string)          // 主键索引
}

// 表格的插入语句
type Blueprint struct {
}

func (bp *Blueprint) Year(columnName string) IBlueprint {
	return nil
}

//func (bp *Blueprint) Uuid(columnName string) IBlueprint  {
//	return nil
//}

func (bp *Blueprint) UnsignedTinyInteger(columnName string, length ...int) IBlueprint {
	return nil
}

func (bp *Blueprint) UnsignedSmallInteger(columnName string, length ...int) IBlueprint {
	return nil
}

func (bp *Blueprint) UnsignedMediumInteger(columnName string, length ...int) IBlueprint {
	return nil
}

func (bp *Blueprint) UnsignedInteger(columnName string, length ...int) IBlueprint {
	return nil
}

func (bp *Blueprint) UnsignedDecimal(columnName string, length int, precision int) IBlueprint {
	return nil
}

func (bp *Blueprint) Text(columnName string) IBlueprint {
	return nil
}

func (bp *Blueprint) Time(columnName string) IBlueprint {
	return nil
}

func (bp *Blueprint) UnsignedBigInteger(columnName string, length ...int) IBlueprint {
	return nil
}

func (bp *Blueprint) SoftDeletes() IBlueprint {
	return nil
}

func (bp *Blueprint) SmallIncrements(columnName string, length ...int) IBlueprint {
	return nil
}

func (bp *Blueprint) SmallInteger(columnName string, length ...int) IBlueprint {
	return nil
}

func (bp *Blueprint) RememberToken() IBlueprint {
	return nil
}

func (bp *Blueprint) Polygon(columnName string) IBlueprint {
	return nil
}

func (bp *Blueprint) Point(columnName string) IBlueprint {
	return nil
}

func (bp *Blueprint) TinyInteger(columnName string, length ...int) IBlueprint {
	return nil
}

func (bp *Blueprint) TinyIncrements(columnName string, length ...int) IBlueprint {
	return nil
}

func (bp *Blueprint) Timestamp(columnName string) IBlueprint {
	return nil
}

func (bp *Blueprint) Timestamps() IBlueprint {
	return nil
}
func (bp *Blueprint) NullableTimestamp(columnName string) IBlueprint {
	return nil
}

func (bp *Blueprint) MultiPolygon(columnName string) IBlueprint {
	return nil
}

func (bp *Blueprint) MultiPoint(columnName string) IBlueprint {
	return nil
}

func (bp *Blueprint) MultiLineString(columnName string) IBlueprint {
	return nil
}

func (bp *Blueprint) MediumText(columnName string) IBlueprint {
	return nil
}

func (bp *Blueprint) Float(columnName string, lengthPrecision ...int) IBlueprint {
	return nil
}

// 枚举类型
func (bp *Blueprint) Enum(columnName string, enum ...string) IBlueprint {
	return nil
}

func (bp *Blueprint) MediumIncrements(columnName string, length ...int) IBlueprint {
	return nil
}

func (bp *Blueprint) MediumInteger(columnName string, length ...int) IBlueprint {
	return nil
}

func (bp *Blueprint) LongText(columnName string) IBlueprint {
	return nil
}

func (bp *Blueprint) LineString(columnName string) IBlueprint {
	return nil
}

func (bp *Blueprint) Json(columnName string) IBlueprint {
	return nil
}

func (bp *Blueprint) IpAddress(columnName string) IBlueprint {
	return nil
}

func (bp *Blueprint) GeometryCollection(columnName string) IBlueprint {
	return nil
}

func (bp *Blueprint) Geometry(columnName string) IBlueprint {
	return nil
}

func (bp *Blueprint) Double(columnName string, lengthPrecision ...int) IBlueprint {
	return nil
}

func (bp *Blueprint) Decimal(columnName string, length int, precision int) IBlueprint {
	return nil
}

func (bp *Blueprint) Date(columnName string) IBlueprint {
	return nil
}

func (bp *Blueprint) DateTime(columnName string) IBlueprint {
	return nil
}

func (bp *Blueprint) Char(columnName string, length int) IBlueprint {
	return nil
}

func (bp *Blueprint) Boolean(columnName string) IBlueprint {
	return nil
}

func (bp *Blueprint) Binary(columnName string) IBlueprint {
	return nil
}

func (bp *Blueprint) BigIncreaments(columnName string, length ...int) IBlueprint {
	return nil
}

func (bp *Blueprint) BigInteger(columnName string, length ...int) IBlueprint {
	return nil
}

func (bp *Blueprint) PrimaryKey(column string) {

}

// 添加普通索引，传入一个值代表普通索引，传入多个值代表聚合索引
func (bp *Blueprint) Index(columns ...string) {

}

// 添加聚合索引
func (bp *Blueprint) UniqueIndex(column string) {

}

// 添加全文索引  只对MyISAM表有效
func (bp *Blueprint) FullTextIndex(column string) {

}

func (bp *Blueprint) Increments(columnName string, length ...int) IBlueprint {
	return bp
}

func (bp *Blueprint) Nullable() IBlueprint {
	return bp
}

func (bp *Blueprint) Comment(comment string) IBlueprint {
	return bp
}

func (bp *Blueprint) Default(def interface{}) IBlueprint {
	return bp
}

func (bp *Blueprint) String(columnName string, length ...int) IBlueprint {
	return bp
}

func (bp *Blueprint) StringWithLength(columnName string, length int) IBlueprint {
	return bp
}

func (bp *Blueprint) Integer(columnName string, length ...int) IBlueprint {
	return bp
}

func (bp *Blueprint) IntegerWithLength(columnName string, length int) IBlueprint {
	return bp
}
