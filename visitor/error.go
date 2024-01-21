package visitor

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"strconv"
)

func Invalid(pctx antlr.Token, msg string) error {
	return fmt.Errorf(
		"%s (%s in line %d pos %d)",
		msg,
		strconv.Quote(pctx.GetText()),
		pctx.GetLine(),
		pctx.GetColumn())
}
