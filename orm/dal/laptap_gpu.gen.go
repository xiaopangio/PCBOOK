// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dal

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/xiaopangio/pcbook/orm/model"
)

func newLaptapGpu(db *gorm.DB, opts ...gen.DOOption) laptapGpu {
	_laptapGpu := laptapGpu{}

	_laptapGpu.laptapGpuDo.UseDB(db, opts...)
	_laptapGpu.laptapGpuDo.UseModel(&model.LaptapGpu{})

	tableName := _laptapGpu.laptapGpuDo.TableName()
	_laptapGpu.ALL = field.NewAsterisk(tableName)
	_laptapGpu.LaptapID = field.NewString(tableName, "laptap_id")
	_laptapGpu.GpuID = field.NewInt32(tableName, "gpu_id")
	_laptapGpu.Gpu = laptapGpuHasOneGpu{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Gpu", "model.Gpu"),
	}

	_laptapGpu.fillFieldMap()

	return _laptapGpu
}

type laptapGpu struct {
	laptapGpuDo

	ALL      field.Asterisk
	LaptapID field.String
	GpuID    field.Int32
	Gpu      laptapGpuHasOneGpu

	fieldMap map[string]field.Expr
}

func (l laptapGpu) Table(newTableName string) *laptapGpu {
	l.laptapGpuDo.UseTable(newTableName)
	return l.updateTableName(newTableName)
}

func (l laptapGpu) As(alias string) *laptapGpu {
	l.laptapGpuDo.DO = *(l.laptapGpuDo.As(alias).(*gen.DO))
	return l.updateTableName(alias)
}

func (l *laptapGpu) updateTableName(table string) *laptapGpu {
	l.ALL = field.NewAsterisk(table)
	l.LaptapID = field.NewString(table, "laptap_id")
	l.GpuID = field.NewInt32(table, "gpu_id")

	l.fillFieldMap()

	return l
}

func (l *laptapGpu) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := l.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (l *laptapGpu) fillFieldMap() {
	l.fieldMap = make(map[string]field.Expr, 3)
	l.fieldMap["laptap_id"] = l.LaptapID
	l.fieldMap["gpu_id"] = l.GpuID

}

func (l laptapGpu) clone(db *gorm.DB) laptapGpu {
	l.laptapGpuDo.ReplaceConnPool(db.Statement.ConnPool)
	return l
}

func (l laptapGpu) replaceDB(db *gorm.DB) laptapGpu {
	l.laptapGpuDo.ReplaceDB(db)
	return l
}

type laptapGpuHasOneGpu struct {
	db *gorm.DB

	field.RelationField
}

func (a laptapGpuHasOneGpu) Where(conds ...field.Expr) *laptapGpuHasOneGpu {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a laptapGpuHasOneGpu) WithContext(ctx context.Context) *laptapGpuHasOneGpu {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a laptapGpuHasOneGpu) Model(m *model.LaptapGpu) *laptapGpuHasOneGpuTx {
	return &laptapGpuHasOneGpuTx{a.db.Model(m).Association(a.Name())}
}

type laptapGpuHasOneGpuTx struct{ tx *gorm.Association }

func (a laptapGpuHasOneGpuTx) Find() (result *model.Gpu, err error) {
	return result, a.tx.Find(&result)
}

func (a laptapGpuHasOneGpuTx) Append(values ...*model.Gpu) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a laptapGpuHasOneGpuTx) Replace(values ...*model.Gpu) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a laptapGpuHasOneGpuTx) Delete(values ...*model.Gpu) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a laptapGpuHasOneGpuTx) Clear() error {
	return a.tx.Clear()
}

func (a laptapGpuHasOneGpuTx) Count() int64 {
	return a.tx.Count()
}

type laptapGpuDo struct{ gen.DO }

func (l laptapGpuDo) Debug() *laptapGpuDo {
	return l.withDO(l.DO.Debug())
}

func (l laptapGpuDo) WithContext(ctx context.Context) *laptapGpuDo {
	return l.withDO(l.DO.WithContext(ctx))
}

func (l laptapGpuDo) ReadDB() *laptapGpuDo {
	return l.Clauses(dbresolver.Read)
}

func (l laptapGpuDo) WriteDB() *laptapGpuDo {
	return l.Clauses(dbresolver.Write)
}

func (l laptapGpuDo) Session(config *gorm.Session) *laptapGpuDo {
	return l.withDO(l.DO.Session(config))
}

func (l laptapGpuDo) Clauses(conds ...clause.Expression) *laptapGpuDo {
	return l.withDO(l.DO.Clauses(conds...))
}

func (l laptapGpuDo) Returning(value interface{}, columns ...string) *laptapGpuDo {
	return l.withDO(l.DO.Returning(value, columns...))
}

func (l laptapGpuDo) Not(conds ...gen.Condition) *laptapGpuDo {
	return l.withDO(l.DO.Not(conds...))
}

func (l laptapGpuDo) Or(conds ...gen.Condition) *laptapGpuDo {
	return l.withDO(l.DO.Or(conds...))
}

func (l laptapGpuDo) Select(conds ...field.Expr) *laptapGpuDo {
	return l.withDO(l.DO.Select(conds...))
}

func (l laptapGpuDo) Where(conds ...gen.Condition) *laptapGpuDo {
	return l.withDO(l.DO.Where(conds...))
}

func (l laptapGpuDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *laptapGpuDo {
	return l.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (l laptapGpuDo) Order(conds ...field.Expr) *laptapGpuDo {
	return l.withDO(l.DO.Order(conds...))
}

func (l laptapGpuDo) Distinct(cols ...field.Expr) *laptapGpuDo {
	return l.withDO(l.DO.Distinct(cols...))
}

func (l laptapGpuDo) Omit(cols ...field.Expr) *laptapGpuDo {
	return l.withDO(l.DO.Omit(cols...))
}

func (l laptapGpuDo) Join(table schema.Tabler, on ...field.Expr) *laptapGpuDo {
	return l.withDO(l.DO.Join(table, on...))
}

func (l laptapGpuDo) LeftJoin(table schema.Tabler, on ...field.Expr) *laptapGpuDo {
	return l.withDO(l.DO.LeftJoin(table, on...))
}

func (l laptapGpuDo) RightJoin(table schema.Tabler, on ...field.Expr) *laptapGpuDo {
	return l.withDO(l.DO.RightJoin(table, on...))
}

func (l laptapGpuDo) Group(cols ...field.Expr) *laptapGpuDo {
	return l.withDO(l.DO.Group(cols...))
}

func (l laptapGpuDo) Having(conds ...gen.Condition) *laptapGpuDo {
	return l.withDO(l.DO.Having(conds...))
}

func (l laptapGpuDo) Limit(limit int) *laptapGpuDo {
	return l.withDO(l.DO.Limit(limit))
}

func (l laptapGpuDo) Offset(offset int) *laptapGpuDo {
	return l.withDO(l.DO.Offset(offset))
}

func (l laptapGpuDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *laptapGpuDo {
	return l.withDO(l.DO.Scopes(funcs...))
}

func (l laptapGpuDo) Unscoped() *laptapGpuDo {
	return l.withDO(l.DO.Unscoped())
}

func (l laptapGpuDo) Create(values ...*model.LaptapGpu) error {
	if len(values) == 0 {
		return nil
	}
	return l.DO.Create(values)
}

func (l laptapGpuDo) CreateInBatches(values []*model.LaptapGpu, batchSize int) error {
	return l.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (l laptapGpuDo) Save(values ...*model.LaptapGpu) error {
	if len(values) == 0 {
		return nil
	}
	return l.DO.Save(values)
}

func (l laptapGpuDo) First() (*model.LaptapGpu, error) {
	if result, err := l.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.LaptapGpu), nil
	}
}

func (l laptapGpuDo) Take() (*model.LaptapGpu, error) {
	if result, err := l.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.LaptapGpu), nil
	}
}

func (l laptapGpuDo) Last() (*model.LaptapGpu, error) {
	if result, err := l.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.LaptapGpu), nil
	}
}

func (l laptapGpuDo) Find() ([]*model.LaptapGpu, error) {
	result, err := l.DO.Find()
	return result.([]*model.LaptapGpu), err
}

func (l laptapGpuDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.LaptapGpu, err error) {
	buf := make([]*model.LaptapGpu, 0, batchSize)
	err = l.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (l laptapGpuDo) FindInBatches(result *[]*model.LaptapGpu, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return l.DO.FindInBatches(result, batchSize, fc)
}

func (l laptapGpuDo) Attrs(attrs ...field.AssignExpr) *laptapGpuDo {
	return l.withDO(l.DO.Attrs(attrs...))
}

func (l laptapGpuDo) Assign(attrs ...field.AssignExpr) *laptapGpuDo {
	return l.withDO(l.DO.Assign(attrs...))
}

func (l laptapGpuDo) Joins(fields ...field.RelationField) *laptapGpuDo {
	for _, _f := range fields {
		l = *l.withDO(l.DO.Joins(_f))
	}
	return &l
}

func (l laptapGpuDo) Preload(fields ...field.RelationField) *laptapGpuDo {
	for _, _f := range fields {
		l = *l.withDO(l.DO.Preload(_f))
	}
	return &l
}

func (l laptapGpuDo) FirstOrInit() (*model.LaptapGpu, error) {
	if result, err := l.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.LaptapGpu), nil
	}
}

func (l laptapGpuDo) FirstOrCreate() (*model.LaptapGpu, error) {
	if result, err := l.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.LaptapGpu), nil
	}
}

func (l laptapGpuDo) FindByPage(offset int, limit int) (result []*model.LaptapGpu, count int64, err error) {
	result, err = l.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = l.Offset(-1).Limit(-1).Count()
	return
}

func (l laptapGpuDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = l.Count()
	if err != nil {
		return
	}

	err = l.Offset(offset).Limit(limit).Scan(result)
	return
}

func (l laptapGpuDo) Scan(result interface{}) (err error) {
	return l.DO.Scan(result)
}

func (l laptapGpuDo) Delete(models ...*model.LaptapGpu) (result gen.ResultInfo, err error) {
	return l.DO.Delete(models)
}

func (l *laptapGpuDo) withDO(do gen.Dao) *laptapGpuDo {
	l.DO = *do.(*gen.DO)
	return l
}