package cue

import (
	"strings"

	"cuelang.org/go/cue/ast"
	"cuelang.org/go/cue/token"
	"google.golang.org/protobuf/compiler/protogen"
)

func toLeadingComments(c protogen.Comments) *ast.CommentGroup {
	var cs []*ast.Comment
	for i, l := range strings.Split(strings.TrimSpace(c.String()), "\n") {
		if len(l) == 0 {
			continue
		}
		if i == 0 {
			cs = append(cs, &ast.Comment{Text: l, Slash: token.NewSection.Pos()})
			continue
		}
		cs = append(cs, &ast.Comment{Text: l})
	}
	if len(cs) == 0 {
		return nil
	}
	return &ast.CommentGroup{
		List: cs,
	}
}

func toTrailingComments(c protogen.Comments) *ast.CommentGroup {
	var cs []*ast.Comment
	lines := strings.Split(strings.TrimSpace(c.String()), "\n")
	s := token.NoPos
	if len(lines) > 1 {
		s = token.Newline.Pos()
	}
	for i, l := range lines {
		if len(l) == 0 {
			continue
		}
		if i == 0 {
			cs = append(cs, &ast.Comment{Text: l, Slash: s})
			continue
		}
		cs = append(cs, &ast.Comment{Text: l})
	}
	if len(cs) == 0 {
		return nil
	}
	return &ast.CommentGroup{
		List: cs,
	}
}
