/*
Copyright Arata Furukawa

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

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
