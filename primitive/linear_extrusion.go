// Copyright 2021-2022 Lukasz Janyst <lukasz@jany.st>
// Licensed under the MIT license, see the LICENSE file for details.

package primitive

import (
	"bufio"
	"fmt"
)

type LinearExtrusion struct {
	ParentImpl
	Height    float64
	Center    bool    "optional"
	Convexity uint16  "optional"
	Twist     uint16  "optional"
	Slices    uint16  "optional"
	Scale     float64 "optional"
	Fn        uint16  "optional"
	Items     *List   "forward:Add"
	prefix    string  "prefix"
	NoOption  bool
}

func NewLinearExtrusion(height float64, items ...Primitive) *LinearExtrusion {
	ret := &LinearExtrusion{
		Height:    height,
		Center:    true,
		Convexity: 10,
		Slices:    20,
		Scale:     1.0,
		Fn:        16,
		Items:     NewList(),
	}
	ret.Items.SetParent(ret)
	ret.Items.Add(items...)
	return ret
}

func NewZeroLinearExtrusion(height float64, items ...Primitive) *LinearExtrusion {
	ret := &LinearExtrusion{
		Height:   height,
		Scale:    1.0,
		Items:    NewList(),
		NoOption: true,
	}
	ret.Items.SetParent(ret)
	ret.Items.Add(items...)
	return ret
}

func (o *LinearExtrusion) Render(w *bufio.Writer) {
	w.WriteString(o.Prefix())
	w.WriteString("linear_extrude(")
	w.WriteString(fmt.Sprintf("height=%f", o.Height))
	if o.NoOption {
		w.WriteString(fmt.Sprintf(", center=%t", o.Center))
	}
	if o.NoOption {
		w.WriteString(fmt.Sprintf(", convexity=%d", o.Convexity))
	}
	if o.NoOption {
		w.WriteString(fmt.Sprintf(", twist=%d", o.Twist))
	}
	if o.NoOption {
		w.WriteString(fmt.Sprintf(", slices=%d", o.Slices))
	}
	if o.NoOption {
		w.WriteString(fmt.Sprintf(", scale=%f", o.Scale))
	}
	if o.NoOption {
		w.WriteString(fmt.Sprintf(", $fn=%d", o.Fn))
	}
	w.WriteString(")")
	o.Items.Render(w)
}
