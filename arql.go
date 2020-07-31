package main

type ArqlEquals struct {
	op    string
	expr1 string
	expr2 string
}

type ArqlAndOr struct {
	op    string
	expr1 *ArqlAndOr
	expr2 *ArqlAndOr
}

func or(exprs ...ArqlAndOr) ArqlAndOr {
	if len(exprs) == 1 {
		return exprs[0]
	}
	var expr2 ArqlAndOr = exprs[0]
	if len(exprs) > 1 {
		expr2 = or(exprs...)
	}
	var expr1 ArqlAndOr
	expr1, exprs = exprs[0], exprs[1:]
	var op = ArqlAndOr{
		"or",
		&expr1,
		&expr2,
	}
	return op
}

func and(exprs ...ArqlAndOr) ArqlAndOr {
	if len(exprs) == 1 {
		return exprs[0]
	}
	var expr2 ArqlAndOr = exprs[0]
	if len(exprs) > 1 {
		expr2 = or(exprs...)
	}
	var expr1 ArqlAndOr
	expr1, exprs = exprs[0], exprs[1:]
	var op = ArqlAndOr{
		"and",
		&expr1,
		&expr2,
	}
	return op
}

func equals(expr1, expr2 string) ArqlEquals {
	return ArqlEquals{"equals", expr1, expr2}
}
