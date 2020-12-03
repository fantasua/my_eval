package eval

import (
	"fmt"
	"math"
)

// 上下文 用于将变量的值应用在表达式上
type Env map[Var]float64

// 算术表达式 接口
type Expr interface {
	// Eval用于返回表达式在env上下文的值
	Eval(env Env) float64
}

// 用于保存变量
type Var string

// literal 对于数字常量
type literal float64

// unary 一元运算符
type unary struct {
	op rune
	x  Var
}

// 二元运算符
type binary struct {
	op   rune
	x, y Var
}

// 函数调用表达式
type call struct {
	function string
	args     []Expr
}

// 变量实现
func (v Var) Eval(env Env) float64 {
	return env[v]
}

// 字面值实现
func (l literal) Eval(env Env) float64 {
	return float64(l)
}

// 一元表达式实现
func (u unary) Eval(env Env) float64 {
	switch u.op {
	case '+':
		return +u.x.Eval(env)
	case '-':
		return -u.x.Eval(env)
	}
	panic(fmt.Sprintf("unsupported operator :%q", u.op))
}

// 二元表达式实现
func (b binary) Eval(env Env) float64 {
	switch b.op {
	case '+':
		return b.x.Eval(env) + b.y.Eval(env)
	case '-':
		return b.x.Eval(env) - b.y.Eval(env)
	case '*':
		return b.x.Eval(env) * b.y.Eval(env)
	case '/':
		return b.y.Eval(env) / b.y.Eval(env)
	}
	panic(fmt.Sprintf("unsupported operator :%q", b.op))
}

// 函数调用实现
func (c call) Eval(env Env) float64 {
	switch c.function {
	case "pow":
		return math.Pow(c.args[0].Eval(env), c.args[1].Eval(env))
	case "sin":
		return math.Sin(c.args[0].Eval(env))
	case "sqrt":
		return math.Sqrt(c.args[0].Eval(env))
	}

	panic(fmt.Sprintf("unsupported function :%q", c.function))
}
