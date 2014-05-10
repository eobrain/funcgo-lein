# fgoc

A Leiningen plugin to compile [Functional Go][fgo] into Clojure

## Usage

Put `[org.eamonn.funcgo/funcgo-lein-plugin "0.2.0"]` into the
`:dependencies` and `:plugins` vectors of your
`project.clj`, or if you are on Leiningen 1.x do `lein plugin install
org.eamonn.funcgo/funcgo-lein-plugin 0.2.0`.


    :dependencies [...
                   [org.eamonn.funcgo/funcgo-lein-plugin "0.2.0"]
                   ...]
    :plugins [...
              [org.eamonn.funcgo/funcgo-lein-plugin "0.2.0"]
              ...]

## License

Copyright © 2014 Eamonn O'Brien-Strain

Distributed under the Eclipse Public License either version 1.0 or (at
your option) any later version.

[fgo]: http://funcgo.org
