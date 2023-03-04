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

func newWeightUnit(db *gorm.DB, opts ...gen.DOOption) weightUnit {
	_weightUnit := weightUnit{}

	_weightUnit.weightUnitDo.UseDB(db, opts...)
	_weightUnit.weightUnitDo.UseModel(&model.WeightUnit{})

	tableName := _weightUnit.weightUnitDo.TableName()
	_weightUnit.ALL = field.NewAsterisk(tableName)
	_weightUnit.WeightUnitID = field.NewInt32(tableName, "weight_unit_id")
	_weightUnit.Type = field.NewString(tableName, "type")

	_weightUnit.fillFieldMap()

	return _weightUnit
}

type weightUnit struct {
	weightUnitDo

	ALL          field.Asterisk
	WeightUnitID field.Int32
	Type         field.String

	fieldMap map[string]field.Expr
}

func (w weightUnit) Table(newTableName string) *weightUnit {
	w.weightUnitDo.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w weightUnit) As(alias string) *weightUnit {
	w.weightUnitDo.DO = *(w.weightUnitDo.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *weightUnit) updateTableName(table string) *weightUnit {
	w.ALL = field.NewAsterisk(table)
	w.WeightUnitID = field.NewInt32(table, "weight_unit_id")
	w.Type = field.NewString(table, "type")

	w.fillFieldMap()

	return w
}

func (w *weightUnit) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *weightUnit) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 2)
	w.fieldMap["weight_unit_id"] = w.WeightUnitID
	w.fieldMap["type"] = w.Type
}

func (w weightUnit) clone(db *gorm.DB) weightUnit {
	w.weightUnitDo.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w weightUnit) replaceDB(db *gorm.DB) weightUnit {
	w.weightUnitDo.ReplaceDB(db)
	return w
}

type weightUnitDo struct{ gen.DO }

func (w weightUnitDo) Debug() *weightUnitDo {
	return w.withDO(w.DO.Debug())
}

func (w weightUnitDo) WithContext(ctx context.Context) *weightUnitDo {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w weightUnitDo) ReadDB() *weightUnitDo {
	return w.Clauses(dbresolver.Read)
}

func (w weightUnitDo) WriteDB() *weightUnitDo {
	return w.Clauses(dbresolver.Write)
}

func (w weightUnitDo) Session(config *gorm.Session) *weightUnitDo {
	return w.withDO(w.DO.Session(config))
}

func (w weightUnitDo) Clauses(conds ...clause.Expression) *weightUnitDo {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w weightUnitDo) Returning(value interface{}, columns ...string) *weightUnitDo {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w weightUnitDo) Not(conds ...gen.Condition) *weightUnitDo {
	return w.withDO(w.DO.Not(conds...))
}

func (w weightUnitDo) Or(conds ...gen.Condition) *weightUnitDo {
	return w.withDO(w.DO.Or(conds...))
}

func (w weightUnitDo) Select(conds ...field.Expr) *weightUnitDo {
	return w.withDO(w.DO.Select(conds...))
}

func (w weightUnitDo) Where(conds ...gen.Condition) *weightUnitDo {
	return w.withDO(w.DO.Where(conds...))
}

func (w weightUnitDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *weightUnitDo {
	return w.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (w weightUnitDo) Order(conds ...field.Expr) *weightUnitDo {
	return w.withDO(w.DO.Order(conds...))
}

func (w weightUnitDo) Distinct(cols ...field.Expr) *weightUnitDo {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w weightUnitDo) Omit(cols ...field.Expr) *weightUnitDo {
	return w.withDO(w.DO.Omit(cols...))
}

func (w weightUnitDo) Join(table schema.Tabler, on ...field.Expr) *weightUnitDo {
	return w.withDO(w.DO.Join(table, on...))
}

func (w weightUnitDo) LeftJoin(table schema.Tabler, on ...field.Expr) *weightUnitDo {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w weightUnitDo) RightJoin(table schema.Tabler, on ...field.Expr) *weightUnitDo {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w weightUnitDo) Group(cols ...field.Expr) *weightUnitDo {
	return w.withDO(w.DO.Group(cols...))
}

func (w weightUnitDo) Having(conds ...gen.Condition) *weightUnitDo {
	return w.withDO(w.DO.Having(conds...))
}

func (w weightUnitDo) Limit(limit int) *weightUnitDo {
	return w.withDO(w.DO.Limit(limit))
}

func (w weightUnitDo) Offset(offset int) *weightUnitDo {
	return w.withDO(w.DO.Offset(offset))
}

func (w weightUnitDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *weightUnitDo {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w weightUnitDo) Unscoped() *weightUnitDo {
	return w.withDO(w.DO.Unscoped())
}

func (w weightUnitDo) Create(values ...*model.WeightUnit) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w weightUnitDo) CreateInBatches(values []*model.WeightUnit, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w weightUnitDo) Save(values ...*model.WeightUnit) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w weightUnitDo) First() (*model.WeightUnit, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.WeightUnit), nil
	}
}

func (w weightUnitDo) Take() (*model.WeightUnit, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.WeightUnit), nil
	}
}

func (w weightUnitDo) Last() (*model.WeightUnit, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.WeightUnit), nil
	}
}

func (w weightUnitDo) Find() ([]*model.WeightUnit, error) {
	result, err := w.DO.Find()
	return result.([]*model.WeightUnit), err
}

func (w weightUnitDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.WeightUnit, err error) {
	buf := make([]*model.WeightUnit, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w weightUnitDo) FindInBatches(result *[]*model.WeightUnit, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w weightUnitDo) Attrs(attrs ...field.AssignExpr) *weightUnitDo {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w weightUnitDo) Assign(attrs ...field.AssignExpr) *weightUnitDo {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w weightUnitDo) Joins(fields ...field.RelationField) *weightUnitDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w weightUnitDo) Preload(fields ...field.RelationField) *weightUnitDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w weightUnitDo) FirstOrInit() (*model.WeightUnit, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.WeightUnit), nil
	}
}

func (w weightUnitDo) FirstOrCreate() (*model.WeightUnit, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.WeightUnit), nil
	}
}

func (w weightUnitDo) FindByPage(offset int, limit int) (result []*model.WeightUnit, count int64, err error) {
	result, err = w.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = w.Offset(-1).Limit(-1).Count()
	return
}

func (w weightUnitDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w weightUnitDo) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w weightUnitDo) Delete(models ...*model.WeightUnit) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *weightUnitDo) withDO(do gen.Dao) *weightUnitDo {
	w.DO = *do.(*gen.DO)
	return w
}
