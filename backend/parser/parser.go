package parser

import (
	"errors"
)

var(
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
		return nil,ErrInvalidStat
	}
	tkn.popToken()
	fieldName := tkn.token()
	if (fieldName.kind != IDENTIFIER) {
		return nil,ErrInvalidStat
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
		updateStmt.Where,_=parseWhere(tkn)
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
		whereStmt.SingleExp= append(whereStmt.SingleExp,singleExp)
		if (singleExp.LogicOp == "") {
			break
		}
	}
	return whereStmt, nil
}

func parseExp(tkn *tokenizer) (*SingleExpStmt, error) {
	var err error
	tkn.popToken()
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
		}else {
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

func isComOp(t token) bool {
	for _,e:=range []byte(t.val.(string)){
		if(e=='>'||e=='='||e=='<'){
			continue
		}else{
			return false
		}
	}
	return true
}
