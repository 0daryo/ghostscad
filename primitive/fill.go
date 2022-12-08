package primitive

import (
	"bufio"
	"fmt"

	. "github.com/go-gl/mathgl/mgl64"
)

type Fill struct {
	ParentImpl
	Resize Vec3
	Auto   bool
	Import Import
}

var _ Primitive = &Fill{}

func NewFill(resize Vec3, auto bool, import_ Import) *Fill {
	return &Fill{
		Resize: resize,
		Auto:   auto,
		Import: import_,
	}
}

func (o *Fill) Render(w *bufio.Writer) {
	w.WriteString(fmt.Sprintf("fill(){resize([%f, %f,%f],auth=%t){\n", o.Resize[0], o.Resize[1], o.Resize[2], o.Auto))
	o.Import.Render(w)
	w.WriteString("}}\n")
}

func (o *Fill) SetParent(Primitive)
func (o *Fill) Parent() Primitive
func (o *Fill) Disable() Primitive
func (o *Fill) Highlight() Primitive
func (o *Fill) ShowOnly() Primitive
func (o *Fill) Transparent() Primitive
func (o *Fill) Prefix() string
