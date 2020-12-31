package selector

import (
	"fmt"
)

//Fields - Custom type
type Fields []string

//Formatter - Helps create a query for specific databases by implementing a Visitor pattern.
type Formatter interface {
	Format(fld, op, val string) string
	FormatArray(op string, val ...string) string
}

/*DataSelectorBuilder - Builds a 'where' like clause to restricts results returned by query.
Specific drivers should implement this interface.
*/
type DataSelectorBuilder interface {
	Build(expr Expr) string
}

/*Expr - Represents a single expression of a query
 */
type Expr interface {
	//Expand - Expands expression to required format
	Expand(f Formatter) string
}

type (
	//Int  Custom type that implements the Expr interface
	Int int
	//String  Custom type that implements the Expr interface
	String string
	//Bool  Custom type that implements the Expr interface
	Bool bool
	//Float  Custom type that implements the Expr interface
	Float float32
	//Null - Custom type that implements the Expr interface
	Null string
	//Empty - Custom type, empty selector
	Empty string
)

func (e Int) Expand(f Formatter) string    { return fmt.Sprintf(`%d`, e) }
func (e String) Expand(f Formatter) string { return fmt.Sprintf(`"%s"`, e) }
func (e Bool) Expand(f Formatter) string   { return fmt.Sprintf(`%t`, e) }
func (e Float) Expand(f Formatter) string  { return fmt.Sprintf(`%f`, e) }
func (e Null) Expand(f Formatter) string   { return fmt.Sprintf(`null`, e) }
func (e Empty) Expand(f Formatter) string  { return `{}` }

//CmpExpr - Base expression for comparison operators
type CmpExpr struct {
	Field string
	Op    string
	Ex    Expr
}

func (e CmpExpr) Expand(f Formatter) string { return f.Format(e.Field, e.Op, e.Ex.Expand(f)) }

func Eq(fld string, expr Expr) Expr  { return &CmpExpr{Field: fld, Ex: expr, Op: "$eq"} }
func Ne(fld string, expr Expr) Expr  { return &CmpExpr{Field: fld, Ex: expr, Op: "$ne"} }
func Lt(fld string, expr Expr) Expr  { return &CmpExpr{Field: fld, Ex: expr, Op: "$lt"} }
func Lte(fld string, expr Expr) Expr { return &CmpExpr{Field: fld, Ex: expr, Op: "$lte"} }
func Gt(fld string, expr Expr) Expr  { return &CmpExpr{Field: fld, Ex: expr, Op: "$gt"} }
func Gte(fld string, expr Expr) Expr { return &CmpExpr{Field: fld, Ex: expr, Op: "$gte"} }

//LogExpr - Base expression for logical operators
type LogExpr struct {
	Ex []Expr
	Op string
}

func And(exprA Expr, expr ...Expr) Expr {

	e := make([]Expr, 0)
	e = append(e, exprA)
	e = append(e, expr...)

	return &LogExpr{Ex: e, Op: "$and"}
}
func Or(exprA Expr, expr ...Expr) Expr {

	e := make([]Expr, 0)
	e = append(e, exprA)
	e = append(e, expr...)

	return &LogExpr{Ex: e, Op: "$or"}
}
func Nor(exprA Expr, expr ...Expr) Expr {

	e := make([]Expr, 0)
	e = append(e, exprA)
	e = append(e, expr...)

	return &LogExpr{Ex: e, Op: "$nor"}
}
func Not(exprA Expr, expr ...Expr) Expr {

	e := make([]Expr, 0)
	e = append(e, exprA)
	e = append(e, expr...)

	return &LogExpr{Ex: e, Op: "$not"}
}

func (e LogExpr) Expand(f Formatter) string {

	expr := make([]string, 0)
	for _, e := range e.Ex {
		expr = append(expr, e.Expand(f))
	}
	return f.FormatArray(e.Op, expr...)
}
