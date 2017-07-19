package base

import (
	"github.com/astaxie/beego/orm"
	"strings"
	"fmt"
)

const (
	SP     string = " "
	SELECT string = "SELECT" + SP
)

func parseSelectSQLToCount(sql *string, fields string) {
	strings.Replace(*sql, fields, "COUNT(*) count", 1)
}

type Page struct {
	Pi    int
	Ps    int
	Total int
}

func (page *Page) GetBeginNum() int {
	return (page.Pi - 1) * page.Ps
}

func Query(sql string, fields string, page Page, params ...interface{}) interface{} {
	var limit string = SP + "LIMIT" + SP + "(" + string(page.GetBeginNum()) + "," + string(page.Ps) + ")"
	o := orm.NewOrm()
	rs := o.Raw(sql+limit, params)
	var rows []orm.Params
	rs.Values(&rows)
	parseSelectSQLToCount(&sql, fields)
	rs = o.Raw(sql, params)
	count := make(orm.Params)
	rs.RowsToMap(&count, "name", "value")
	fmt.Println(count["count"])
	return rows
}

func QueryByBuilder(sqlBuilder orm.QueryBuilder, fields string, page Page, params ...interface{}) interface{} {
	o := orm.NewOrm()
	sqlBuilder.Limit(page.GetBeginNum()).Offset(page.Ps)
	sql := sqlBuilder.String()
	rs := o.Raw(sql, params)
	var rows []orm.Params
	rs.Values(&rows)
	parseSelectSQLToCount(&sql, fields)
	rs = o.Raw(sql, params)
	count := make(orm.Params)
	rs.RowsToMap(&count, "name", "value")
	fmt.Println(count["count"])
	return rows
}

func GetQueryBuilder() orm.QueryBuilder {
	qb, _ := orm.NewQueryBuilder("mysql")
	return qb
}