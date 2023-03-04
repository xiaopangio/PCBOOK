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

func newMemory(db *gorm.DB, opts ...gen.DOOption) memory {
	_memory := memory{}

	_memory.memoryDo.UseDB(db, opts...)
	_memory.memoryDo.UseModel(&model.Memory{})

	tableName := _memory.memoryDo.TableName()
	_memory.ALL = field.NewAsterisk(tableName)
	_memory.MemoryID = field.NewInt32(tableName, "memory_id")
	_memory.Value = field.NewInt32(tableName, "value")
	_memory.MemoryUnitID = field.NewInt32(tableName, "memory_unit_id")
	_memory.Unit = memoryHasOneUnit{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Unit", "model.MemoryUnit"),
	}

	_memory.fillFieldMap()

	return _memory
}

type memory struct {
	memoryDo

	ALL          field.Asterisk
	MemoryID     field.Int32
	Value        field.Int32
	MemoryUnitID field.Int32
	Unit         memoryHasOneUnit

	fieldMap map[string]field.Expr
}

func (m memory) Table(newTableName string) *memory {
	m.memoryDo.UseTable(newTableName)
	return m.updateTableName(newTableName)
}

func (m memory) As(alias string) *memory {
	m.memoryDo.DO = *(m.memoryDo.As(alias).(*gen.DO))
	return m.updateTableName(alias)
}

func (m *memory) updateTableName(table string) *memory {
	m.ALL = field.NewAsterisk(table)
	m.MemoryID = field.NewInt32(table, "memory_id")
	m.Value = field.NewInt32(table, "value")
	m.MemoryUnitID = field.NewInt32(table, "memory_unit_id")

	m.fillFieldMap()

	return m
}

func (m *memory) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := m.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (m *memory) fillFieldMap() {
	m.fieldMap = make(map[string]field.Expr, 4)
	m.fieldMap["memory_id"] = m.MemoryID
	m.fieldMap["value"] = m.Value
	m.fieldMap["memory_unit_id"] = m.MemoryUnitID

}

func (m memory) clone(db *gorm.DB) memory {
	m.memoryDo.ReplaceConnPool(db.Statement.ConnPool)
	return m
}

func (m memory) replaceDB(db *gorm.DB) memory {
	m.memoryDo.ReplaceDB(db)
	return m
}

type memoryHasOneUnit struct {
	db *gorm.DB

	field.RelationField
}

func (a memoryHasOneUnit) Where(conds ...field.Expr) *memoryHasOneUnit {
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

func (a memoryHasOneUnit) WithContext(ctx context.Context) *memoryHasOneUnit {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a memoryHasOneUnit) Model(m *model.Memory) *memoryHasOneUnitTx {
	return &memoryHasOneUnitTx{a.db.Model(m).Association(a.Name())}
}

type memoryHasOneUnitTx struct{ tx *gorm.Association }

func (a memoryHasOneUnitTx) Find() (result *model.MemoryUnit, err error) {
	return result, a.tx.Find(&result)
}

func (a memoryHasOneUnitTx) Append(values ...*model.MemoryUnit) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a memoryHasOneUnitTx) Replace(values ...*model.MemoryUnit) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a memoryHasOneUnitTx) Delete(values ...*model.MemoryUnit) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a memoryHasOneUnitTx) Clear() error {
	return a.tx.Clear()
}

func (a memoryHasOneUnitTx) Count() int64 {
	return a.tx.Count()
}

type memoryDo struct{ gen.DO }

func (m memoryDo) Debug() *memoryDo {
	return m.withDO(m.DO.Debug())
}

func (m memoryDo) WithContext(ctx context.Context) *memoryDo {
	return m.withDO(m.DO.WithContext(ctx))
}

func (m memoryDo) ReadDB() *memoryDo {
	return m.Clauses(dbresolver.Read)
}

func (m memoryDo) WriteDB() *memoryDo {
	return m.Clauses(dbresolver.Write)
}

func (m memoryDo) Session(config *gorm.Session) *memoryDo {
	return m.withDO(m.DO.Session(config))
}

func (m memoryDo) Clauses(conds ...clause.Expression) *memoryDo {
	return m.withDO(m.DO.Clauses(conds...))
}

func (m memoryDo) Returning(value interface{}, columns ...string) *memoryDo {
	return m.withDO(m.DO.Returning(value, columns...))
}

func (m memoryDo) Not(conds ...gen.Condition) *memoryDo {
	return m.withDO(m.DO.Not(conds...))
}

func (m memoryDo) Or(conds ...gen.Condition) *memoryDo {
	return m.withDO(m.DO.Or(conds...))
}

func (m memoryDo) Select(conds ...field.Expr) *memoryDo {
	return m.withDO(m.DO.Select(conds...))
}

func (m memoryDo) Where(conds ...gen.Condition) *memoryDo {
	return m.withDO(m.DO.Where(conds...))
}

func (m memoryDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *memoryDo {
	return m.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (m memoryDo) Order(conds ...field.Expr) *memoryDo {
	return m.withDO(m.DO.Order(conds...))
}

func (m memoryDo) Distinct(cols ...field.Expr) *memoryDo {
	return m.withDO(m.DO.Distinct(cols...))
}

func (m memoryDo) Omit(cols ...field.Expr) *memoryDo {
	return m.withDO(m.DO.Omit(cols...))
}

func (m memoryDo) Join(table schema.Tabler, on ...field.Expr) *memoryDo {
	return m.withDO(m.DO.Join(table, on...))
}

func (m memoryDo) LeftJoin(table schema.Tabler, on ...field.Expr) *memoryDo {
	return m.withDO(m.DO.LeftJoin(table, on...))
}

func (m memoryDo) RightJoin(table schema.Tabler, on ...field.Expr) *memoryDo {
	return m.withDO(m.DO.RightJoin(table, on...))
}

func (m memoryDo) Group(cols ...field.Expr) *memoryDo {
	return m.withDO(m.DO.Group(cols...))
}

func (m memoryDo) Having(conds ...gen.Condition) *memoryDo {
	return m.withDO(m.DO.Having(conds...))
}

func (m memoryDo) Limit(limit int) *memoryDo {
	return m.withDO(m.DO.Limit(limit))
}

func (m memoryDo) Offset(offset int) *memoryDo {
	return m.withDO(m.DO.Offset(offset))
}

func (m memoryDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *memoryDo {
	return m.withDO(m.DO.Scopes(funcs...))
}

func (m memoryDo) Unscoped() *memoryDo {
	return m.withDO(m.DO.Unscoped())
}

func (m memoryDo) Create(values ...*model.Memory) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Create(values)
}

func (m memoryDo) CreateInBatches(values []*model.Memory, batchSize int) error {
	return m.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (m memoryDo) Save(values ...*model.Memory) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Save(values)
}

func (m memoryDo) First() (*model.Memory, error) {
	if result, err := m.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Memory), nil
	}
}

func (m memoryDo) Take() (*model.Memory, error) {
	if result, err := m.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Memory), nil
	}
}

func (m memoryDo) Last() (*model.Memory, error) {
	if result, err := m.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Memory), nil
	}
}

func (m memoryDo) Find() ([]*model.Memory, error) {
	result, err := m.DO.Find()
	return result.([]*model.Memory), err
}

func (m memoryDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Memory, err error) {
	buf := make([]*model.Memory, 0, batchSize)
	err = m.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (m memoryDo) FindInBatches(result *[]*model.Memory, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return m.DO.FindInBatches(result, batchSize, fc)
}

func (m memoryDo) Attrs(attrs ...field.AssignExpr) *memoryDo {
	return m.withDO(m.DO.Attrs(attrs...))
}

func (m memoryDo) Assign(attrs ...field.AssignExpr) *memoryDo {
	return m.withDO(m.DO.Assign(attrs...))
}

func (m memoryDo) Joins(fields ...field.RelationField) *memoryDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Joins(_f))
	}
	return &m
}

func (m memoryDo) Preload(fields ...field.RelationField) *memoryDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Preload(_f))
	}
	return &m
}

func (m memoryDo) FirstOrInit() (*model.Memory, error) {
	if result, err := m.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Memory), nil
	}
}

func (m memoryDo) FirstOrCreate() (*model.Memory, error) {
	if result, err := m.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Memory), nil
	}
}

func (m memoryDo) FindByPage(offset int, limit int) (result []*model.Memory, count int64, err error) {
	result, err = m.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = m.Offset(-1).Limit(-1).Count()
	return
}

func (m memoryDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = m.Count()
	if err != nil {
		return
	}

	err = m.Offset(offset).Limit(limit).Scan(result)
	return
}

func (m memoryDo) Scan(result interface{}) (err error) {
	return m.DO.Scan(result)
}

func (m memoryDo) Delete(models ...*model.Memory) (result gen.ResultInfo, err error) {
	return m.DO.Delete(models)
}

func (m *memoryDo) withDO(do gen.Dao) *memoryDo {
	m.DO = *do.(*gen.DO)
	return m
}
