package primitive

import (
	"bufio"
	"fmt"

	. "github.com/go-gl/mathgl/mgl64"
)

type Fill struct {
	ParentImpl
	Items *List "forward:Add"
}

var _ Primitive = &Fill{}

func NewFill(items ...Primitive) *Fill {
	ret := &Fill{
		Items: NewList(),
	}
	ret.Items.SetParent(ret)
	ret.Items.Add(items...)
	return ret
}

func (o *Fill) Render(w *bufio.Writer) {
	w.WriteString(fmt.Sprintf("fill("))
	w.WriteString(fmt.Sprintf(")"))
	o.Items.Render(w)
}

func (o *Fill) SetParent(Primitive) {
	return
}
func (o *Fill) Parent() Primitive {
	return o
}
func (o *Fill) Disable() Primitive {
	return o
}
func (o *Fill) Highlight() Primitive {
	return o
}
func (o *Fill) ShowOnly() Primitive {
	return o
}
func (o *Fill) Transparent() Primitive {
	return o
}
func (o *Fill) Prefix() string {
	return ""
}

type Resize struct {
	ParentImpl
	Auto  bool
	Dims  Vec3
	Items *List "forward:Add"
}

var _ Primitive = &Resize{}

func NewResize(Auto bool, vec Vec3, items ...Primitive) *Resize {
	ret := &Resize{
		Auto:  Auto,
		Dims:  vec,
		Items: NewList(),
	}
	ret.Items.SetParent(ret)
	ret.Items.Add(items...)
	return ret
}

func (o *Resize) Render(w *bufio.Writer) {
	w.WriteString(fmt.Sprintf("resize(\n"))
	w.WriteString(
		fmt.Sprintf("\t[%f, %f, %f], auto=%t\n", o.Dims[0], o.Dims[1], o.Dims[2], o.Auto),
	)
	w.WriteString(fmt.Sprintf(")"))
	o.Items.Render(w)
}

func (o *Resize) SetParent(Primitive) {
	return
}
func (o *Resize) Parent() Primitive {
	return o
}
func (o *Resize) Disable() Primitive {
	return o
}
func (o *Resize) Highlight() Primitive {
	return o
}
func (o *Resize) ShowOnly() Primitive {
	return o
}
func (o *Resize) Transparent() Primitive {
	return o
}
func (o *Resize) Prefix() string {
	return ""
}
