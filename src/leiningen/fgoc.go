//////
// This file is part of the Funcgo Leiningen plugin.
//
// Copyright (c) 2014 Eamonn O'Brien-Strain All rights
// reserved. This program and the accompanying materials are made
// available under the terms of the Eclipse Public License v1.0 which
// accompanies this distribution, and is available at
// http://www.eclipse.org/legal/epl-v10.html
//
// Contributors:
// Eamonn O'Brien-Strain e@obrain.com - initial author
//////

package fgoc
import(
        "clojure/java/io"
	"leiningen/core/eval"
)
import extern(
	main
)

// Extract source and test paths from the project description.
// TODO(eob) Add tasks directory if it exists.
func dirs(project) {
	roots      := for [k,v] := lazy project if k==SOURCE_PATHS || k==TEST_PATHS {
		if isVector(v) {
			v[0]   // TODO(eob) handle >1 file in vector
		} else {
			v
		}
	}
	childRoots := for [_, v] := lazy project if isMap(v) { dirs(v) }

	roots concat childRoots
}

// Compile Functional Go files in the project. (entry point into the plugin)
func fgoc(project, args...) {
	opts    := func{first($1) == '-'} filter args
	nonOpts := func{first($1) != '-'} filter args
	roots   := opts concat (if seq(nonOpts) {
		nonOpts
	} else {
		set(func(path){io.file(path)->getAbsolutePath()} map flatten(dirs(project)))
	})

	eval.evalInProject(project,
		syntax main.Compile(unquotes roots),
		quote(
			require(quote([\`funcgo.main`, AS, main]))
		)
	)
}
