package leiningen/fgoc
import(
        sh "clojure/java/shell"
        "clojure/java/io"
	"leiningen/core/eval"
)

func fgoc(project, args...) {
	eval.evalInProject(project,
		syntax main.Compile(unquotes args),
		quote(
			require(quote([\`funcgo.main`, AS, main]))
		)
	)
}
