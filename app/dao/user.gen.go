// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

import (
	"context"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gen/helper"

	"gorm.io/plugin/dbresolver"

	"skeleton/app/model"
)

func newUser(db *gorm.DB, opts ...gen.DOOption) user {
	_user := user{}

	_user.userDo.UseDB(db, opts...)
	_user.userDo.UseModel(&model.User{})

	tableName := _user.userDo.TableName()
	_user.ALL = field.NewAsterisk(tableName)
	_user.ID = field.NewInt32(tableName, "id")
	_user.Nickname = field.NewString(tableName, "nickname")
	_user.Name = field.NewString(tableName, "name")
	_user.Avatar = field.NewString(tableName, "avatar")
	_user.Mobile = field.NewString(tableName, "mobile")
	_user.Role = field.NewString(tableName, "role")
	_user.Age = field.NewInt32(tableName, "age")
	_user.Intro = field.NewString(tableName, "intro")
	_user.Sex = field.NewInt8(tableName, "sex")
	_user.CreatedAt = field.NewField(tableName, "created_at")
	_user.UpdatedAt = field.NewField(tableName, "updated_at")
	_user.Address = field.NewString(tableName, "address")

	_user.fillFieldMap()

	return _user
}

type user struct {
	userDo

	ALL       field.Asterisk
	ID        field.Int32
	Nickname  field.String // 昵称
	Name      field.String
	Avatar    field.String // 头像
	Mobile    field.String // 手机号
	Role      field.String
	Age       field.Int32
	Intro     field.String // 描述
	Sex       field.Int8   // 性别【0：男：1女2：保密】
	CreatedAt field.Field  // 创建时间
	UpdatedAt field.Field  // 修改时间
	Address   field.String

	fieldMap map[string]field.Expr
}

func (u user) Table(newTableName string) *user {
	u.userDo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u user) As(alias string) *user {
	u.userDo.DO = *(u.userDo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *user) updateTableName(table string) *user {
	u.ALL = field.NewAsterisk(table)
	u.ID = field.NewInt32(table, "id")
	u.Nickname = field.NewString(table, "nickname")
	u.Name = field.NewString(table, "name")
	u.Avatar = field.NewString(table, "avatar")
	u.Mobile = field.NewString(table, "mobile")
	u.Role = field.NewString(table, "role")
	u.Age = field.NewInt32(table, "age")
	u.Intro = field.NewString(table, "intro")
	u.Sex = field.NewInt8(table, "sex")
	u.CreatedAt = field.NewField(table, "created_at")
	u.UpdatedAt = field.NewField(table, "updated_at")
	u.Address = field.NewString(table, "address")

	u.fillFieldMap()

	return u
}

func (u *user) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *user) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 12)
	u.fieldMap["id"] = u.ID
	u.fieldMap["nickname"] = u.Nickname
	u.fieldMap["name"] = u.Name
	u.fieldMap["avatar"] = u.Avatar
	u.fieldMap["mobile"] = u.Mobile
	u.fieldMap["role"] = u.Role
	u.fieldMap["age"] = u.Age
	u.fieldMap["intro"] = u.Intro
	u.fieldMap["sex"] = u.Sex
	u.fieldMap["created_at"] = u.CreatedAt
	u.fieldMap["updated_at"] = u.UpdatedAt
	u.fieldMap["address"] = u.Address
}

func (u user) clone(db *gorm.DB) user {
	u.userDo.ReplaceConnPool(db.Statement.ConnPool)
	return u
}

func (u user) replaceDB(db *gorm.DB) user {
	u.userDo.ReplaceDB(db)
	return u
}

type userDo struct{ gen.DO }

type IUserDo interface {
	gen.SubQuery
	Debug() IUserDo
	WithContext(ctx context.Context) IUserDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IUserDo
	WriteDB() IUserDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IUserDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IUserDo
	Not(conds ...gen.Condition) IUserDo
	Or(conds ...gen.Condition) IUserDo
	Select(conds ...field.Expr) IUserDo
	Where(conds ...gen.Condition) IUserDo
	Order(conds ...field.Expr) IUserDo
	Distinct(cols ...field.Expr) IUserDo
	Omit(cols ...field.Expr) IUserDo
	Join(table schema.Tabler, on ...field.Expr) IUserDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IUserDo
	RightJoin(table schema.Tabler, on ...field.Expr) IUserDo
	Group(cols ...field.Expr) IUserDo
	Having(conds ...gen.Condition) IUserDo
	Limit(limit int) IUserDo
	Offset(offset int) IUserDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IUserDo
	Unscoped() IUserDo
	Create(values ...*model.User) error
	CreateInBatches(values []*model.User, batchSize int) error
	Save(values ...*model.User) error
	First() (*model.User, error)
	Take() (*model.User, error)
	Last() (*model.User, error)
	Find() ([]*model.User, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.User, err error)
	FindInBatches(result *[]*model.User, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.User) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IUserDo
	Assign(attrs ...field.AssignExpr) IUserDo
	Joins(fields ...field.RelationField) IUserDo
	Preload(fields ...field.RelationField) IUserDo
	FirstOrInit() (*model.User, error)
	FirstOrCreate() (*model.User, error)
	FindByPage(offset int, limit int) (result []*model.User, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IUserDo
	UnderlyingDB() *gorm.DB
	schema.Tabler

	FindByNameAndAge(name string, age int) (result model.User, err error)
	FindBySimpleName() (result []model.User, err error)
	FindByIDOrName(cond1 bool, id int, col string, name string) (result model.User, err error)
	FindAll() (result []map[string]interface{}, err error)
	FindOne() (result map[string]interface{})
	FindAddress() (result model.User, err error)
	FindByID(id int) (result model.User, err error)
	FindAdult() (result []model.User, err error)
	FindByRole(role string, id int)
	UpdateUserName(name string, id int) (err error)
}

// Where("name=@name and age=@age")
func (u userDo) FindByNameAndAge(name string, age int) (result model.User, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, name)
	params = append(params, age)
	generateSQL.WriteString("name=? and age=? ")

	var executeSQL *gorm.DB
	executeSQL = u.UnderlyingDB().Where(generateSQL.String(), params...).Take(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// sql(select id,name,age from user where age>18)
func (u userDo) FindBySimpleName() (result []model.User, err error) {
	var generateSQL strings.Builder
	generateSQL.WriteString("select id,name,age from user where age>18 ")

	var executeSQL *gorm.DB
	executeSQL = u.UnderlyingDB().Raw(generateSQL.String()).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// sql(select id,name,age from @@table where age>18
// {{if cond1}}and id=@id {{end}}
// {{if name == ""}}and @@col=@name{{end}})
func (u userDo) FindByIDOrName(cond1 bool, id int, col string, name string) (result model.User, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("select id,name,age from user where age>18 ")
	if cond1 {
		params = append(params, id)
		generateSQL.WriteString("and id=? ")
	}
	if name == "" {
		params = append(params, name)
		generateSQL.WriteString("and " + u.Quote(col) + "=? ")
	}

	var executeSQL *gorm.DB
	executeSQL = u.UnderlyingDB().Raw(generateSQL.String(), params...).Take(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// sql(select * from user)
func (u userDo) FindAll() (result []map[string]interface{}, err error) {
	var generateSQL strings.Builder
	generateSQL.WriteString("select * from user ")

	var executeSQL *gorm.DB
	executeSQL = u.UnderlyingDB().Raw(generateSQL.String()).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// sql(select * from user limit 1)
func (u userDo) FindOne() (result map[string]interface{}) {
	var generateSQL strings.Builder
	generateSQL.WriteString("select * from user limit 1 ")

	result = make(map[string]interface{})
	var executeSQL *gorm.DB
	executeSQL = u.UnderlyingDB().Raw(generateSQL.String()).Take(result) // ignore_security_alert
	_ = executeSQL

	return
}

// sql(select address from user limit 1)
func (u userDo) FindAddress() (result model.User, err error) {
	var generateSQL strings.Builder
	generateSQL.WriteString("select address from user limit 1 ")

	var executeSQL *gorm.DB
	executeSQL = u.UnderlyingDB().Raw(generateSQL.String()).Take(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// where(id=@id)
func (u userDo) FindByID(id int) (result model.User, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, id)
	generateSQL.WriteString("id=? ")

	var executeSQL *gorm.DB
	executeSQL = u.UnderlyingDB().Where(generateSQL.String(), params...).Take(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// select * from users where age>18
func (u userDo) FindAdult() (result []model.User, err error) {
	var generateSQL strings.Builder
	generateSQL.WriteString("select * from users where age>18 ")

	var executeSQL *gorm.DB
	executeSQL = u.UnderlyingDB().Raw(generateSQL.String()).Find(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// select * from @@table
//
//	{{where}}
//		{{if role=="user"}}
//			id=@id
//		{{else if role=="admin"}}
//			role="user" or rule="normal-admin"
//		{{else}}
//			role="user" or role="normal-admin" or role="admin"
//		{{end}}
//	{{end}}
func (u userDo) FindByRole(role string, id int) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("select * from user ")
	var whereSQL0 strings.Builder
	if role == "user" {
		params = append(params, id)
		whereSQL0.WriteString("id=? ")
	} else if role == "admin" {
		whereSQL0.WriteString("role=\"user\" or rule=\"normal-admin\" ")
	} else {
		whereSQL0.WriteString("role=\"user\" or role=\"normal-admin\" or role=\"admin\" ")
	}
	helper.JoinWhereBuilder(&generateSQL, whereSQL0)

	var executeSQL *gorm.DB
	executeSQL = u.UnderlyingDB().Exec(generateSQL.String(), params...) // ignore_security_alert
	_ = executeSQL

	return
}

// update users
//
//	{{set}}
//		update_time=now(),
//		{{if name != ""}}
//			name=@name
//		{{end}}
//	{{end}}
//
// where id=@id
func (u userDo) UpdateUserName(name string, id int) (err error) {
	var params []interface{}

	var generateSQL strings.Builder
	generateSQL.WriteString("update users ")
	var setSQL0 strings.Builder
	setSQL0.WriteString("update_time=now(), ")
	if name != "" {
		params = append(params, name)
		setSQL0.WriteString("name=? ")
	}
	helper.JoinSetBuilder(&generateSQL, setSQL0)
	params = append(params, id)
	generateSQL.WriteString("where id=? ")

	var executeSQL *gorm.DB
	executeSQL = u.UnderlyingDB().Exec(generateSQL.String(), params...) // ignore_security_alert
	err = executeSQL.Error

	return
}

func (u userDo) Debug() IUserDo {
	return u.withDO(u.DO.Debug())
}

func (u userDo) WithContext(ctx context.Context) IUserDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u userDo) ReadDB() IUserDo {
	return u.Clauses(dbresolver.Read)
}

func (u userDo) WriteDB() IUserDo {
	return u.Clauses(dbresolver.Write)
}

func (u userDo) Session(config *gorm.Session) IUserDo {
	return u.withDO(u.DO.Session(config))
}

func (u userDo) Clauses(conds ...clause.Expression) IUserDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u userDo) Returning(value interface{}, columns ...string) IUserDo {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u userDo) Not(conds ...gen.Condition) IUserDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u userDo) Or(conds ...gen.Condition) IUserDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u userDo) Select(conds ...field.Expr) IUserDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u userDo) Where(conds ...gen.Condition) IUserDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u userDo) Order(conds ...field.Expr) IUserDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u userDo) Distinct(cols ...field.Expr) IUserDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u userDo) Omit(cols ...field.Expr) IUserDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u userDo) Join(table schema.Tabler, on ...field.Expr) IUserDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u userDo) LeftJoin(table schema.Tabler, on ...field.Expr) IUserDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u userDo) RightJoin(table schema.Tabler, on ...field.Expr) IUserDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u userDo) Group(cols ...field.Expr) IUserDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u userDo) Having(conds ...gen.Condition) IUserDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u userDo) Limit(limit int) IUserDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u userDo) Offset(offset int) IUserDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u userDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IUserDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u userDo) Unscoped() IUserDo {
	return u.withDO(u.DO.Unscoped())
}

func (u userDo) Create(values ...*model.User) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u userDo) CreateInBatches(values []*model.User, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u userDo) Save(values ...*model.User) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u userDo) First() (*model.User, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.User), nil
	}
}

func (u userDo) Take() (*model.User, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.User), nil
	}
}

func (u userDo) Last() (*model.User, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.User), nil
	}
}

func (u userDo) Find() ([]*model.User, error) {
	result, err := u.DO.Find()
	return result.([]*model.User), err
}

func (u userDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.User, err error) {
	buf := make([]*model.User, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u userDo) FindInBatches(result *[]*model.User, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u userDo) Attrs(attrs ...field.AssignExpr) IUserDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u userDo) Assign(attrs ...field.AssignExpr) IUserDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u userDo) Joins(fields ...field.RelationField) IUserDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Joins(_f))
	}
	return &u
}

func (u userDo) Preload(fields ...field.RelationField) IUserDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Preload(_f))
	}
	return &u
}

func (u userDo) FirstOrInit() (*model.User, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.User), nil
	}
}

func (u userDo) FirstOrCreate() (*model.User, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.User), nil
	}
}

func (u userDo) FindByPage(offset int, limit int) (result []*model.User, count int64, err error) {
	result, err = u.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = u.Offset(-1).Limit(-1).Count()
	return
}

func (u userDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u userDo) Scan(result interface{}) (err error) {
	return u.DO.Scan(result)
}

func (u userDo) Delete(models ...*model.User) (result gen.ResultInfo, err error) {
	return u.DO.Delete(models)
}

func (u *userDo) withDO(do gen.Dao) *userDo {
	u.DO = *do.(*gen.DO)
	return u
}