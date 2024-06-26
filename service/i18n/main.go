package i18n

import (
	"context"
	"fmt"
	"io"

	"github.com/labstack/echo/v4"

	"github.com/mylukin/easy-i18n/i18n"

	tele "gopkg.in/telebot.v3"
)

// i18nKey is string
type i18nKey string

// Domain is domain
type Domain = i18n.Domain

// I18nCtxKey is context key
const I18nCtxKey i18nKey = "i18n-ctx"

// NewPrinter is new printer
func NewPrinter(lang any) *i18n.Printer {
	return i18n.NewPrinter(lang)
}

// SetLang set language
func SetLang(lang any) *i18n.Printer {
	i18n.SetLang(lang)
	return NewPrinter(lang)
}

// Make is make language printer
func Make(lang any) context.Context {
	ctx, _ := context.WithCancel(context.Background())
	return context.WithValue(ctx, I18nCtxKey, i18n.NewPrinter(lang))
}

// Printf is like fmt.Printf, but using language-specific formatting.
func Printf(ctx any, format string, args ...any) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf(format, args...)
		}
	}()
	var p *i18n.Printer
	switch _ctx := ctx.(type) {
	case echo.Context:
		p = _ctx.Get("Language").(*i18n.Printer)
	case tele.Context:
		p = _ctx.Get("Language").(*i18n.Printer)
	case context.Context:
		p = _ctx.Value(I18nCtxKey).(*i18n.Printer)
	case *i18n.Printer:
		p = _ctx
	default:
		panic("i18n ctx error")
	}
	p.Printf(format, args...)
}

// Sprintf is like fmt.Sprintf, but using language-specific formatting.
func Sprintf(ctx any, format string, args ...any) (result string) {
	defer func() {
		if err := recover(); err != nil {
			result = fmt.Sprintf(format, args...)
		}
	}()
	var p *i18n.Printer
	switch _ctx := ctx.(type) {
	case echo.Context:
		p = _ctx.Get("Language").(*i18n.Printer)
	case tele.Context:
		p = _ctx.Get("Language").(*i18n.Printer)
	case context.Context:
		p = _ctx.Value(I18nCtxKey).(*i18n.Printer)
	case *i18n.Printer:
		p = _ctx
	default:
		panic("i18n ctx error")
	}
	return p.Sprintf(format, args...)
}

// Fprintf is like fmt.Fprintf, but using language-specific formatting.
func Fprintf(w io.Writer, ctx any, key string, args ...any) (n int, resErr error) {
	defer func() {
		if err := recover(); err != nil {
			n, resErr = fmt.Fprintf(w, key, args...)
		}
	}()
	var p *i18n.Printer
	switch _ctx := ctx.(type) {
	case echo.Context:
		p = _ctx.Get("Language").(*i18n.Printer)
	case tele.Context:
		p = _ctx.Get("Language").(*i18n.Printer)
	case context.Context:
		p = _ctx.Value(I18nCtxKey).(*i18n.Printer)
	case *i18n.Printer:
		p = _ctx
	default:
		panic("i18n ctx error")
	}
	return p.Fprintf(w, key, args...)
}

// Plural is plural
func Plural(cases ...any) []i18n.PluralRule {
	return i18n.Plural(cases...)
}
