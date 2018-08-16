package v1

import (
	"github.com/mlab-lattice/lattice/pkg/definition/component"
	"github.com/mlab-lattice/lattice/pkg/definition/tree"
)

type ReferenceNode struct {
	parent    tree.Node
	path      tree.Path
	reference *Reference
}

func NewReferenceNode(name string, parent tree.Node, reference *Reference) *ReferenceNode {
	return &ReferenceNode{
		parent:    parent,
		path:      parent.Path().Child(name),
		reference: reference,
	}
}

func (n *ReferenceNode) Path() tree.Path {
	return n.path
}

func (n *ReferenceNode) Value() interface{} {
	return n.reference
}

func (n *ReferenceNode) Parent() tree.Node {
	return n.parent
}

func (n *ReferenceNode) Children() map[string]tree.Node {
	return nil
}

func (n *ReferenceNode) Component() component.Interface {
	return n.reference
}

func (n *ReferenceNode) Reference() *Reference {
	return n.reference
}