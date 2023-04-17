// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/0xJacky/Nginx-UI/server/model"
)

func newSite(db *gorm.DB, opts ...gen.DOOption) site {
	_site := site{}

	_site.siteDo.UseDB(db, opts...)
	_site.siteDo.UseModel(&model.Site{})

	tableName := _site.siteDo.TableName()
	_site.ALL = field.NewAsterisk(tableName)
	_site.ID = field.NewInt(tableName, "id")
	_site.CreatedAt = field.NewTime(tableName, "created_at")
	_site.UpdatedAt = field.NewTime(tableName, "updated_at")
	_site.DeletedAt = field.NewTime(tableName, "deleted_at")
	_site.Path = field.NewString(tableName, "path")
	_site.Advanced = field.NewBool(tableName, "advanced")

	_site.fillFieldMap()

	return _site
}

type site struct {
	siteDo

	ALL       field.Asterisk
	ID        field.Int
	CreatedAt field.Time
	UpdatedAt field.Time
	DeletedAt field.Time
	Path      field.String
	Advanced  field.Bool

	fieldMap map[string]field.Expr
}

func (s site) Table(newTableName string) *site {
	s.siteDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s site) As(alias string) *site {
	s.siteDo.DO = *(s.siteDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *site) updateTableName(table string) *site {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewInt(table, "id")
	s.CreatedAt = field.NewTime(table, "created_at")
	s.UpdatedAt = field.NewTime(table, "updated_at")
	s.DeletedAt = field.NewTime(table, "deleted_at")
	s.Path = field.NewString(table, "path")
	s.Advanced = field.NewBool(table, "advanced")

	s.fillFieldMap()

	return s
}

func (s *site) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *site) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 6)
	s.fieldMap["id"] = s.ID
	s.fieldMap["created_at"] = s.CreatedAt
	s.fieldMap["updated_at"] = s.UpdatedAt
	s.fieldMap["deleted_at"] = s.DeletedAt
	s.fieldMap["path"] = s.Path
	s.fieldMap["advanced"] = s.Advanced
}

func (s site) clone(db *gorm.DB) site {
	s.siteDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s site) replaceDB(db *gorm.DB) site {
	s.siteDo.ReplaceDB(db)
	return s
}

type siteDo struct{ gen.DO }

// FirstByID Where("id=@id")
func (s siteDo) FirstByID(id int) (result *model.Site, err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, id)
	generateSQL.WriteString("id=? ")

	var executeSQL *gorm.DB
	executeSQL = s.UnderlyingDB().Where(generateSQL.String(), params...).Take(&result) // ignore_security_alert
	err = executeSQL.Error

	return
}

// DeleteByID update @@table set deleted_at=strftime('%Y-%m-%d %H:%M:%S','now') where id=@id
func (s siteDo) DeleteByID(id int) (err error) {
	var params []interface{}

	var generateSQL strings.Builder
	params = append(params, id)
	generateSQL.WriteString("update sites set deleted_at=strftime('%Y-%m-%d %H:%M:%S','now') where id=? ")

	var executeSQL *gorm.DB
	executeSQL = s.UnderlyingDB().Exec(generateSQL.String(), params...) // ignore_security_alert
	err = executeSQL.Error

	return
}

func (s siteDo) Debug() *siteDo {
	return s.withDO(s.DO.Debug())
}

func (s siteDo) WithContext(ctx context.Context) *siteDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s siteDo) ReadDB() *siteDo {
	return s.Clauses(dbresolver.Read)
}

func (s siteDo) WriteDB() *siteDo {
	return s.Clauses(dbresolver.Write)
}

func (s siteDo) Session(config *gorm.Session) *siteDo {
	return s.withDO(s.DO.Session(config))
}

func (s siteDo) Clauses(conds ...clause.Expression) *siteDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s siteDo) Returning(value interface{}, columns ...string) *siteDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s siteDo) Not(conds ...gen.Condition) *siteDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s siteDo) Or(conds ...gen.Condition) *siteDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s siteDo) Select(conds ...field.Expr) *siteDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s siteDo) Where(conds ...gen.Condition) *siteDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s siteDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *siteDo {
	return s.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (s siteDo) Order(conds ...field.Expr) *siteDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s siteDo) Distinct(cols ...field.Expr) *siteDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s siteDo) Omit(cols ...field.Expr) *siteDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s siteDo) Join(table schema.Tabler, on ...field.Expr) *siteDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s siteDo) LeftJoin(table schema.Tabler, on ...field.Expr) *siteDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s siteDo) RightJoin(table schema.Tabler, on ...field.Expr) *siteDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s siteDo) Group(cols ...field.Expr) *siteDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s siteDo) Having(conds ...gen.Condition) *siteDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s siteDo) Limit(limit int) *siteDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s siteDo) Offset(offset int) *siteDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s siteDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *siteDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s siteDo) Unscoped() *siteDo {
	return s.withDO(s.DO.Unscoped())
}

func (s siteDo) Create(values ...*model.Site) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s siteDo) CreateInBatches(values []*model.Site, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s siteDo) Save(values ...*model.Site) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s siteDo) First() (*model.Site, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Site), nil
	}
}

func (s siteDo) Take() (*model.Site, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Site), nil
	}
}

func (s siteDo) Last() (*model.Site, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Site), nil
	}
}

func (s siteDo) Find() ([]*model.Site, error) {
	result, err := s.DO.Find()
	return result.([]*model.Site), err
}

func (s siteDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Site, err error) {
	buf := make([]*model.Site, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s siteDo) FindInBatches(result *[]*model.Site, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s siteDo) Attrs(attrs ...field.AssignExpr) *siteDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s siteDo) Assign(attrs ...field.AssignExpr) *siteDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s siteDo) Joins(fields ...field.RelationField) *siteDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s siteDo) Preload(fields ...field.RelationField) *siteDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s siteDo) FirstOrInit() (*model.Site, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Site), nil
	}
}

func (s siteDo) FirstOrCreate() (*model.Site, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Site), nil
	}
}

func (s siteDo) FindByPage(offset int, limit int) (result []*model.Site, count int64, err error) {
	result, err = s.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = s.Offset(-1).Limit(-1).Count()
	return
}

func (s siteDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s siteDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s siteDo) Delete(models ...*model.Site) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *siteDo) withDO(do gen.Dao) *siteDo {
	s.DO = *do.(*gen.DO)
	return s
}
