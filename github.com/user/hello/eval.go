package main

import(
    "fmt"
    "math"
)

type Expr interface{
    Eval(env Env) float64
}

type Var string
type literal float64
type unary struct{
    op rune
    x Expr
} 
type binay struct{
    op rune
    x,y Expr
}
type call struct{
    fn string
    args []Expr
}
type Env map[Var]float64

func (v Var) Eval(env Env) float64{
    return env[v]
}

func (l literal) Eval(_ Env) float64{
    return float64(l)
}

func (u unary) Eval(env Env) float64 {
    switch u.op{
    case '+':
        return +u.x.Eval(env)
    case '-':
        return -u.x.Eval(env)
    }
    panic(fmt.Sprintf("unsupported unary operator: %q", u.op))
}


func (b binay) Eval(env Env) float64 {
    switch b.op{
    case '+':
        return b.x.Eval(env) + b.y.Eval(env)
    case '-':
        return b.x.Eval(env) - b.y.Eval(env)
    case '*':
        return b.x.Eval(env) * b.y.Eval(env)
    case '/':
        return b.x.Eval(env) / b.y.Eval(env)
    }
    panic(fmt.Sprintf("unsupported unary operator: %q", b.op))
}

func (c call) Eval(env Env) float64 {
    switch c.fn {
    case "pow":
        return math.Pow(c.args[0].Eval(env), c.args[1].Eval(env))
    case "sin":
        return math.Sin(c.args[0].Eval(env))
    case "sqrt":
        return math.Sqrt(c.args[0].Eval(env))
    }
    panic(fmt.Sprintf("unsupported function call: %s", c.fn))
}

func main() {
    b := binay{'+',literal(101),literal(202)}
    var env Env
    fmt.Println(b.Eval(env))

    env2 := Env{"x":106, "y":203}
    b2 := binay{'+',Var("x"),Var("y")}
    fmt.Println(b2.Eval(env2))
}









