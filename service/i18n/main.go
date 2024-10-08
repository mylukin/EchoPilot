package i18n

import (
	"fmt"
	"io"

	"github.com/labstack/echo/v4"
	"github.com/mylukin/easy-i18n/i18n"
	tele "gopkg.in/telebot.v3"
)

// Domain is domain
type (
	Domain     = i18n.Domain
	Printer    = i18n.Printer
	PluralRule = i18n.PluralRule
	Message    = i18n.Message
)

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
func Make(lang any) *i18n.Printer {
	return i18n.NewPrinter(lang)
}

// Printf is like fmt.Printf, but using language-specific formatting.
func Printf[T any](ctx T, format string, args ...any) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf(format, args...)
		}
	}()
	getPrinter(ctx).Printf(format, args...)
}

// Sprintf is like fmt.Sprintf, but using language-specific formatting.
func Sprintf[T any](ctx T, format string, args ...any) (result string) {
	defer func() {
		if err := recover(); err != nil {
			result = fmt.Sprintf(format, args...)
		}
	}()
	return getPrinter(ctx).Sprintf(format, args...)
}

// Fprintf is like fmt.Fprintf, but using language-specific formatting.
func Fprintf[T any](w io.Writer, ctx T, key string, args ...any) (n int, resErr error) {
	defer func() {
		if err := recover(); err != nil {
			n, resErr = fmt.Fprintf(w, key, args...)
		}
	}()
	return getPrinter(ctx).Fprintf(w, key, args...)
}

// Plural is plural
func Plural(cases ...any) []i18n.PluralRule {
	return i18n.Plural(cases...)
}

// getPrinter 是一个泛型函数，用于获取 i18n.Printer
func getPrinter[T any](ctx T) *i18n.Printer {
	switch c := any(ctx).(type) {
	case echo.Context:
		return c.Get("Language").(*i18n.Printer)
	case tele.Context:
		return c.Get("Language").(*i18n.Printer)
	case *i18n.Printer:
		return c
	default:
		panic("i18n ctx error")
	}
}
