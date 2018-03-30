package parser

import (
	"errors"
	"fmt"
)

func parseError(tk *tokenizer, i interface{}) ParseError {
	var pe ParseError
	pe.err = ErrInvalidStat;
	switch i.(type) {
	case *UpdateStmt:
		pe.info = "update error: " + tk.token().val.(string)
	case *WhereStmt:
		pe.info = "where error: " + tk.token().val.(string)
	case *SingleExpStmt:
		pe.info = "singleExpStmt error: " + tk.token().val.(string)
	case *SelectStmt:
		pe.info = "SelectStmt error: " + tk.token().val.(string)
	}
	return pe
}

func (e ParseError) Error() string {
	return e.info
}

var (
	ErrInvalidStat = errors.New("InvalidStat")
)

func Parse(stat string) (interface{}, error) {
	tkn := newTokenizer(stat)
	tkn.getToken()
	token := tkn.token()
	if (token.kind != KEYWORD) {
		// error
		return nil, ErrInvalidStat
	}
	switch token.val {
	case "update":
		return ParseUpdate(tkn)
	case "select":
		return ParseSelect(tkn)
	default:
		return nil, ErrInvalidStat
	}
}

func ParseUpdate(tkn *tokenizer) (*UpdateStmt, error) {
	updateStmt := new(UpdateStmt)
	tkn.popToken()
	updateStmt.TableName = tkn.token().val.(string)
	tkn.popToken()
	set := tkn.token()
	if (set.val.(string) != "set" || set.kind != KEYWORD) {
		return nil, ErrInvalidStat
	}
	tkn.popToken()
	fieldName := tkn.token()
	if (fieldName.kind != IDENTIFIER) {
		return nil, ErrInvalidStat
	}
	updateStmt.FieldName = fieldName.val.(string)
	tkn.popToken()
	equal := tkn.token()
	if (equal.kind != SYMBOL || []byte(equal.val.(string))[0] != '=') {
		return nil, ErrInvalidStat
	}
	tkn.popToken()
	value := tkn.token()
	if (value.kind != LITERAL) {
		return nil, ErrInvalidStat
	}
	updateStmt.Value = value.val.(string)
	tkn.popToken()
	where := tkn.token()
	if (where.kind == EOF) {
		updateStmt.Where = nil
		return updateStmt, nil
	} else {
		updateStmt.Where, _ = parseWhere(tkn)
	}
	return updateStmt, nil
}

func parseWhere(tkn *tokenizer) (*WhereStmt, error) {
	whereStmt := new(WhereStmt)
	where := tkn.token()
	tkn.popToken()
	if where.kind != KEYWORD || where.val.(string) != "where" {
		return nil, ErrInvalidStat
	}
	for {
		singleExp, err := parseExp(tkn)
		if (err != nil) {
			return whereStmt, err
		}
		whereStmt.SingleExp = append(whereStmt.SingleExp, singleExp)
		if (singleExp.LogicOp == "") {
			break
		}
	}
	return whereStmt, nil
}

func parseExp(tkn *tokenizer) (*SingleExpStmt, error) {
	var err error
	field := tkn.token()
	var singleExp *SingleExpStmt
	for {
		if (field.kind == IDENTIFIER) {
			singleExp, err = parseSingleExp(tkn)
		}
		tkn.popToken()
		logicOp := tkn.token()
		if (logicOp.kind != EOF) {
			singleExp.LogicOp = logicOp.val.(string)
		} else {
			break
		}
	}
	return singleExp, err

}

func parseSingleExp(tkn *tokenizer) (*SingleExpStmt, error) {
	singleStmt := new(SingleExpStmt)
	singleStmt.Field = tkn.token().val.(string)
	tkn.popToken()
	comOP := tkn.token()
	if (comOP.kind != SYMBOL || !isComOp(comOP)) {
		return nil, ErrInvalidStat
	}
	singleStmt.CmpOp = comOP.val.(string)
	tkn.popToken()
	value := tkn.token()
	if (value.kind != LITERAL) {
		return nil, ErrInvalidStat
	}
	singleStmt.Value = value.val
	return singleStmt, nil
}

func parseInsert() {

}

func ParseSelect(tkn *tokenizer) (*SelectStmt, error) {
	selectStmt := new(SelectStmt)
	for {
		field, ok := getField(tkn)
		fmt.Println(field,ok)
		if (!ok) {
			if (field.kind==KEYWORD&&field.val=="from") {
				break
			}
			return nil, parseError(tkn, selectStmt)
		}
		selectStmt.Fields = append(selectStmt.Fields, field.val)
	}
	_, ok := getTableName(tkn)
	if (!ok) {
		return nil, nil
	}
	_, ok = getKeyword(tkn, "where")
	where, err := parseWhere(tkn)
	if (err != nil) {
		return nil, err
	}
	selectStmt.Where = where
	return selectStmt, nil

}

func parseDelete() {

}

func isComOp(t token) bool {
	for _, e := range []byte(t.val.(string)) {
		if (e == '>' || e == '=' || e == '<') {
			continue
		} else {
			return false
		}
	}
	return true
}

type ParseError struct {
	err  error
	info string
}

func getField(tkn *tokenizer) (token, bool) {
	tkn.popToken()
	field := tkn.token()
	return field, field.kind == IDENTIFIER
}

func getValue(tkn *tokenizer) (token, bool) {
	tkn.popToken()
	value := tkn.token()
	return value, value.kind == LITERAL
}

func getKeyword(tkn *tokenizer, kw string) (token, bool) {
	tkn.popToken()
	field := tkn.token()
	return field, field.kind == KEYWORD || field.val == kw
}

func getTableName(tkn *tokenizer) (token, bool) {
	tkn.popToken()
	field := tkn.token()
	return field, field.kind == IDENTIFIER
}
