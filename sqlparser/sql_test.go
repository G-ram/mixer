package sqlparser

import (
	"testing"
)

func testParse(t *testing.T, sql string) Statement {
	stmt, err := Parse(sql)
	if err != nil {
		t.Fatalf("sql: %s err: %s", sql, err)
	}
	return stmt
}

func TestSet(t *testing.T) {
	sql := "set names gbk"
	testParse(t, sql)
}

func TestSimpleSelect(t *testing.T) {
	sql := "select last_insert_id() as a"
	testParse(t, sql)
}

func TestAliasedWhere(t *testing.T) {
	sql := "select * from foo.tables where a"
	testParse(t, sql)

	sql = "select * from foo.tables where a or b"
	testParse(t, sql)

	sql = "select * from foo.tables where a > b"
	testParse(t, sql)

	sql = "SELECT sum_a_ok AS `sum_a_ok`, sum_b_ok AS `sum_b_ok` FROM ( SELECT SUM(`foo`.`a`) AS `sum_a_ok`, SUM(`foo`.`b`) AS `sum_b_ok`, (COUNT(1) > 0) AS `havclause`, 1 AS `_Tableau_const_expr` FROM `foo` GROUP BY 4 ) `t0` WHERE havclause"
	testParse(t, sql)

	sql = "select * from foo.tables"
	testParse(t, sql)
}

func TestBooleanLiteral(t *testing.T) {
	sql := "select * from foo.tables where a=true"
	testParse(t, sql)

	sql = "select false from foo.tables where false"
	testParse(t, sql)

	sql = "select false from foo.tables order by false"
	testParse(t, sql)
}

func TestTimeConstructors(t *testing.T) {

	sql := "select * from foo.tables where a > (DATE '2014-06-01 00:00:00.000')"
	testParse(t, sql)

	sql = "select * from foo.tables where a > (DATETIME '2014-06-01 00:00:00.000')"
	testParse(t, sql)

	sql = "select * from foo.tables where a > (YEAR '2014-06-01 00:00:00.000')"
	testParse(t, sql)

	sql = "select * from foo.tables where a > (TIME '2014-06-01 00:00:00.000')"
	testParse(t, sql)

	sql = "select * from foo.tables where a > (TIMESTAMP '2014-06-01 00:00:00.000')"
	testParse(t, sql)
}

func TestCastExpr(t *testing.T) {

	sql := "SELECT CAST(3/3 AS int precision) from foo.tables"
	testParse(t, sql)

	sql = "SELECT CAST(3/3 AS datetime) from foo.tables"
	testParse(t, sql)

	sql = "SELECT (100 * (CASE WHEN 1 = 0 THEN NULL ELSE CAST((CASE WHEN (flights201406.cancelled = '1') THEN 1 ELSE 0 END) AS DOUBLE PRECISION) / 1 END)) from foo.tables"
	testParse(t, sql)
}

func TestCaseExpr(t *testing.T) {

	// simple case expression
	sql := "SELECT case x when 1 then 2 when 2 then 3 else 4 end from foo.tables"
	testParse(t, sql)

	// searched case expression
	sql = "SELECT case when a=1 then 2 when a=2 then 3 else 4 end from foo.tables"
	testParse(t, sql)

}

func TestSubqueryComparisons(t *testing.T) {
	sql := "SELECT * FROM foo WHERE a > any (SELECT a FROM foo)"
	testParse(t, sql)

	sql = "SELECT * FROM foo WHERE a = some (SELECT a FROM foo)"
	testParse(t, sql)

	sql = "SELECT * FROM foo WHERE a <> ALL (SELECT a FROM foo)"
	testParse(t, sql)
}

func TestFunnyNames(t *testing.T) {
	sql := "select * from columns"
	testParse(t, sql)

	sql = "select * from foo.columns"
	testParse(t, sql)

	sql = "select * from tables"
	testParse(t, sql)

	sql = "select * from foo.tables"
	testParse(t, sql)
}

func TestMixer(t *testing.T) {
	sql := `admin upnode("node1", "master", "127.0.0.1")`
	testParse(t, sql)

	sql = "show databases"
	testParse(t, sql)

	sql = "show tables from abc"
	testParse(t, sql)

	sql = "show tables from abc like a"
	testParse(t, sql)

	sql = "show tables from abc where a = 1"
	testParse(t, sql)

	sql = "show proxy abc"
	testParse(t, sql)

	sql = "show databases"
	testParse(t, sql)

	sql = "show databases like 'foo'"
	testParse(t, sql)

	sql = "SHOW VARIABLES"
	testParse(t, sql)

	sql = "show variables"
	testParse(t, sql)

	sql = "show variables LIKE 'foo'"
	testParse(t, sql)

	sql = "show columns from 'foo'"
	testParse(t, sql)

	sql = "show columns in 'foo'"
	testParse(t, sql)
	if testParse(t, sql).(*Show).Modifier != "" {
		t.Fatal("modifier wrong")
	}

	sql = "show full columns from 'foo'"
	if testParse(t, sql).(*Show).Modifier != "full" {
		t.Fatal("modifier wrong")
	}

	sql = "show columns from 'foo' from 'bar'"
	testParse(t, sql)
	if String(testParse(t, sql).(*Show).From) != "'foo'" {
		t.Fatalf("table wrong: %s", String(testParse(t, sql).(*Show).From))
	}
	if String(testParse(t, sql).(*Show).DBFilter) != "'bar'" {
		t.Fatalf("db wrong: %s", String(testParse(t, sql).(*Show).DBFilter))
	}

	sql = "show columns in 'foo' in 'bar'"
	testParse(t, sql)
	if String(testParse(t, sql).(*Show).From) != "'foo'" {
		t.Fatalf("table wrong: %s", String(testParse(t, sql).(*Show).From))
	}
	if String(testParse(t, sql).(*Show).DBFilter) != "'bar'" {
		t.Fatalf("db wrong: %s", String(testParse(t, sql).(*Show).DBFilter))
	}
}
