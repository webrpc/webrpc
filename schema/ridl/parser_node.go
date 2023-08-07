package ridl

// NodeType represents the type of a parser tree node.
type NodeType uint

const (
	RootNodeType NodeType = iota
	TokenNodeType
	DefinitionNodeType
	ImportNodeType
	AliasNodeType
	EnumNodeType
	StructNodeType
	ErrorNodeType
	ArgumentNodeType
	MethodNodeType
	ServiceNodeType
)

// Node represents a parser tree node
type Node interface {
	Start() int
	End() int
	Type() NodeType
}

type node struct {
	children []Node

	start int
	end   int
}

func (n *node) End() int {
	return n.end
}

func (n *node) Start() int {
	return n.start
}

func (n *node) Push(child Node) {
	n.children = append(n.children, child)
}

func (n *node) Children() []Node {
	return n.children
}

func (n *node) Filter(nt NodeType) []Node {
	matches := make([]Node, 0, len(n.children))
	for i := 0; i < len(n.children); i++ {
		if n.children[i].Type() == nt {
			matches = append(matches, n.children[i])
		}
	}
	return matches
}

type RootNode struct {
	node
}

func (rn RootNode) Definitions() []*DefinitionNode {
	nodes := rn.Filter(DefinitionNodeType)

	definitionNodes := make([]*DefinitionNode, 0, len(nodes))
	for i := range nodes {
		definitionNodes = append(definitionNodes, nodes[i].(*DefinitionNode))
	}

	return definitionNodes
}

func (rn RootNode) Imports() []*ImportNode {
	nodes := rn.Filter(ImportNodeType)

	importNodes := make([]*ImportNode, 0, len(nodes))
	for i := range nodes {
		importNodes = append(importNodes, nodes[i].(*ImportNode))
	}

	return importNodes
}

func (rn RootNode) Structs() []*StructNode {
	nodes := rn.Filter(StructNodeType)

	structNodes := make([]*StructNode, 0, len(nodes))
	for i := range nodes {
		structNodes = append(structNodes, nodes[i].(*StructNode))
	}

	return structNodes
}

func (rn RootNode) Errors() []*ErrorNode {
	nodes := rn.Filter(ErrorNodeType)

	errorNodes := make([]*ErrorNode, 0, len(nodes))
	for i := range nodes {
		errorNodes = append(errorNodes, nodes[i].(*ErrorNode))
	}

	return errorNodes
}

func (rn RootNode) Enums() []*EnumNode {
	nodes := rn.Filter(EnumNodeType)

	enumNodes := make([]*EnumNode, 0, len(nodes))
	for i := range nodes {
		enumNodes = append(enumNodes, nodes[i].(*EnumNode))
	}

	return enumNodes
}

func (rn RootNode) Aliases() []*AliasNode {
	nodes := rn.Filter(AliasNodeType)

	aliasNodes := make([]*AliasNode, 0, len(nodes))
	for i := range nodes {
		aliasNodes = append(aliasNodes, nodes[i].(*AliasNode))
	}

	return aliasNodes
}

func (rn RootNode) Services() []*ServiceNode {
	nodes := rn.Filter(ServiceNodeType)

	serviceNodes := make([]*ServiceNode, 0, len(nodes))
	for i := range nodes {
		serviceNodes = append(serviceNodes, nodes[i].(*ServiceNode))
	}

	return serviceNodes
}

func (rn RootNode) Type() NodeType {
	return RootNodeType
}

type DefinitionNode struct {
	node

	leftNode  *TokenNode
	rightNode *TokenNode

	optional bool

	meta []*DefinitionNode
}

func (dn DefinitionNode) Meta() []*DefinitionNode {
	return dn.meta
}

func (dn DefinitionNode) Type() NodeType {
	return DefinitionNodeType
}

func (dn DefinitionNode) Left() *TokenNode {
	if dn.leftNode == nil {
		return invalidToken
	}
	return dn.leftNode
}

func (dn DefinitionNode) Right() *TokenNode {
	if dn.rightNode == nil {
		return invalidToken
	}
	return dn.rightNode
}

func (dn DefinitionNode) Optional() bool {
	return dn.optional
}

type TokenNode struct {
	node

	tok *token
}

var invalidToken = &TokenNode{
	tok: &token{tt: tokenInvalid},
}

func newTokenNode(tok *token) *TokenNode {
	return &TokenNode{
		node: node{
			start: tok.pos,
			end:   tok.pos + len(tok.val),
		},
		tok: tok,
	}
}

func (tn TokenNode) String() string {
	if tn.tok == nil {
		return ""
	}
	return tn.tok.val
}

func (tn TokenNode) Type() NodeType {
	return TokenNodeType
}

type ImportNode struct {
	node

	path    *TokenNode
	members []*TokenNode
}

func (in ImportNode) Members() []*TokenNode {
	return in.members
}

func (in ImportNode) Path() *TokenNode {
	if in.path == nil {
		return invalidToken
	}
	return in.path
}

func (in ImportNode) Type() NodeType {
	return ImportNodeType
}

type AliasNode struct {
	node

	name      *TokenNode
	aliasType *TokenNode
	extra     *DefinitionNode
}

func (tn AliasNode) Type() NodeType {
	return AliasNodeType
}

func (tn AliasNode) Name() *TokenNode {
	return tn.name
}

func (tn AliasNode) TypeName() *TokenNode {
	return tn.aliasType
}

func (tn AliasNode) Extra() *DefinitionNode {
	return tn.extra
}

type EnumNode struct {
	node

	name     *TokenNode
	enumType *TokenNode
	values   []*DefinitionNode
}

func (en EnumNode) Type() NodeType {
	return EnumNodeType
}

func (en EnumNode) Name() *TokenNode {
	return en.name
}

func (en EnumNode) TypeName() *TokenNode {
	return en.enumType
}

func (en EnumNode) Values() []*DefinitionNode {
	return en.values
}

type StructNode struct {
	node

	name   *TokenNode
	fields []*DefinitionNode
}

func (mn StructNode) Name() *TokenNode {
	return mn.name
}

func (mn *StructNode) Type() NodeType {
	return StructNodeType
}

func (mn *StructNode) Fields() []*DefinitionNode {
	return mn.fields
}

type ErrorNode struct {
	node

	code       *TokenNode
	name       *TokenNode
	message    *TokenNode
	httpStatus *TokenNode
}

func (en ErrorNode) Type() NodeType {
	return ErrorNodeType
}

func (en ErrorNode) Name() *TokenNode {
	return en.name
}

type ArgumentNode struct {
	node

	name         *TokenNode
	argumentType *TokenNode
	optional     bool

	inlineStruct *TokenNode
}

func (an *ArgumentNode) Name() *TokenNode {
	if an.name == nil {
		return invalidToken
	}
	return an.name
}

func (an *ArgumentNode) TypeName() *TokenNode {
	if an.argumentType == nil {
		return invalidToken
	}
	return an.argumentType
}

func (an *ArgumentNode) Optional() bool {
	return an.optional
}

func (an *ArgumentNode) Type() NodeType {
	return ArgumentNodeType
}

type MethodNode struct {
	name *TokenNode

	proxy bool

	inputs  argumentList
	outputs argumentList
}

func (mn *MethodNode) Name() *TokenNode {
	return mn.name
}

func (mn *MethodNode) Proxy() bool {
	return mn.proxy
}

func (mn *MethodNode) StreamInput() bool {
	return mn.inputs.stream
}

func (mn *MethodNode) StreamOutput() bool {
	return mn.outputs.stream
}

func (mn *MethodNode) Inputs() []*ArgumentNode {
	return mn.inputs.arguments
}

func (mn *MethodNode) Outputs() []*ArgumentNode {
	return mn.outputs.arguments
}

type ServiceNode struct {
	node

	name *TokenNode

	methods []*MethodNode
}

func (sn ServiceNode) Type() NodeType {
	return ServiceNodeType
}

func (sn ServiceNode) Name() *TokenNode {
	return sn.name
}

func (sn ServiceNode) Methods() []*MethodNode {
	return sn.methods
}

type argumentList struct {
	stream bool

	arguments []*ArgumentNode
}
