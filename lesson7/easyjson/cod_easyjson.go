// TEMPORARY AUTOGENERATED FILE: easyjson stub code to make the package
// compilable during generation.

package main

import (
	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
)

func (Fruit) MarshalJSON() ([]byte, error)       { return nil, nil }
func (*Fruit) UnmarshalJSON([]byte) error        { return nil }
func (Fruit) MarshalEasyJSON(w *jwriter.Writer)  {}
func (*Fruit) UnmarshalEasyJSON(l *jlexer.Lexer) {}

type EasyJSON_exporter_Fruit *Fruit

func (Product) MarshalJSON() ([]byte, error)       { return nil, nil }
func (*Product) UnmarshalJSON([]byte) error        { return nil }
func (Product) MarshalEasyJSON(w *jwriter.Writer)  {}
func (*Product) UnmarshalEasyJSON(l *jlexer.Lexer) {}

type EasyJSON_exporter_Product *Product
