package discord

import "fmt"

func GetKnownLanguages() map[string]string {
	return map[string]string{
		"nodemon.json":           "nodemon",
		"package.json":           "npm",
		"turbo.json":             "turbo",
		"/(vercel|now)\\.json/i": "vercel",
		"/\\.prettier((rc)|(\\.(toml|yml|yaml|json|js))?$){2}/i": "prettier",
		"/\\.eslint((rc|ignore)|(\\.(json|js))?$){2}/i":          "eslint",
		"/\\.(now|vercel)ignore/i":                               "vercel",
		"/\\prettier.config.js/i":                                "prettier",
		"/vue.config\\.(js|ts)/i":                                "vueconfig",
		"/vite.config\\.(js|ts)/i":                               "viteconfig",
		"/vitest.config\\.(js|ts|mjs)/i":                         "vitestconfig",
		"/jest.config\\.(js|ts)/i":                               "jest",
		"/tailwind\\.config\\.(js|cjs|mjs|ts|cts|mts)/i":         "tailwind",
		"/gatsby-(browser|node|ssr|config)\\.js/i":               "gatsbyjs",
		"/webpack(\\.dev|\\.development|\\.prod|\\.production)?\\.config(\\.babel)?\\.(js|jsx|coffee|ts|json|json5|yaml|yml)/i": "webpack",
		"babel.config.js":         "babel",
		".ahk":                    "ahk",
		".ahkl":                   "ahk",
		".astro":                  "astro",
		".astro.config.mjs":       "astroconfig",
		"androidmanifest.xml":     "android",
		"/^angular[^.]*\\.js$/i":  "angular",
		".ng":                     "angular",
		".applescript":            "applescript",
		"/(\\.)?appveyor\\.yml/i": "appveyor",
		".ino":                    "arduino",
		".swf":                    "as",
		".as":                     "as",
		".jsfl":                   "as",
		".swc":                    "as",
		".asp":                    "asp",
		".asax":                   "asp",
		".ascx":                   "asp",
		".ashx":                   "asp",
		".asmx":                   "asp",
		".aspx":                   "asp",
		".axd":                    "asp",
		"/\\.(l?a|[ls]?o|out|s|a51|asm|axf|elf|prx|puff|z80)$/i": "assembly",
		".agc": "assembly",
		".ko":  "assembly",
		".kv":  "kivy",
		".ks":  "kag-script",
		".tjs": "kirikiri-tpv-javascript",
		".lst": "assembly",
		"/\\.((c([+px]{2}?)?-?)?objdump|bsdiff|bin|dat|pak|pdb)$/i": "assembly",
		".d-objdump":           "assembly",
		"/\\.gcode|\\.gco/i":   "assembly",
		"/\\.rpy[bc]$/i":       "assembly",
		"/\\.py[co]$/i":        "assembly",
		".swp":                 "assembly",
		".DS_Store":            "assembly",
		".au3":                 "autoit",
		"/\\.babelrc/i":        "babel",
		".bat":                 "bat",
		".batch":               "bat",
		".cmd":                 "bat",
		"/\\.(exe|com|msi)$/i": "bat",
		".reg":                 "bat",
		"/^(BUILD)|(WORKSPACE)|(\\.bzl)|(\\.bazel(rc)?)|(\\.s(tar|ky))$/": "bazel",
		"/^(\\.bowerrc|bower\\.json|Bowerfile)$/i":                        "bower",
		"/\\.bf?$/i":                     "brainfuck",
		"/\\.c$/i":                       "c",
		"/(cargo.toml|cargo.lock)/i":     "cargo",
		".casc":                          "casc",
		".cas":                           "casc",
		".cfc":                           "coldfusion",
		".cfm":                           "coldfusion",
		"circle.yml":                     "circleci",
		".clj":                           "clojure",
		".cl2":                           "clojure",
		".cljc":                          "clojure",
		".cljx":                          "clojure",
		".hic":                           "clojure",
		"/\\.cljs(cm)?$/i":               "clojure",
		".cmake":                         "cmake",
		"/^CMakeLists\\.txt$/":           "cmake",
		"/\\.codeclimate\\.(yml|json)/i": "codeclimate",
		".coffee":                        "coffee",
		".cjsx":                          "coffee",
		".coffee.ecr":                    "coffee",
		".coffee.erb":                    "coffee",
		".litcoffee":                     "coffee",
		".iced":                          "coffee",
		".cos":                           "cosmo",
		".⭐":                             "cosmo",
		"/\\.c[+px]{2}$|\\.cc$/i":        "cpp",
		"/\\.h[+px]{2}$/i":               "cpp",
		"/\\.[it]pp$/i":                  "cpp",
		"/\\.(tcc|inl)$/i":               "cpp",
		".cats":                          "cpp",
		".idc":                           "cpp",
		".w":                             "cpp",
		".nc":                            "cpp",
		".upc":                           "cpp",
		".xpm":                           "cpp",
		"/\\.e?cr$/i":                    "crystal",
		".cs":                            "csharp",
		".csx":                           "csharp",
		".cshtml":                        "razor",
		".csproj":                        "csproj",
		".css":                           "css",
		".css.map":                       "cssmap",
		".cu":                            "cuda",
		".pyx":                           "cython",
		"/\\.di?$/i":                     "d",
		".dart":                          "dart",
		".dfm":                           "delphi",
		".dpr":                           "delphi",
		".dsc":                           "denizen",
		".dm":                            "dm",
		".dme":                           "dm",
		".dmm":                           "dm",
		"/^(Dockerfile|docker-compose)|\\.docker(file|ignore)$/i": "docker",
		"/^docker-sync\\.yml$/i":                                  "docker",
		".editorconfig":                                           "editorconfig",
		".ejs":                                                    "ejs",
		".ex":                                                     "elixir",
		"/\\.(exs|l?eex)$/i":                                      "elixir",
		"/^mix\\.(exs?|lock)$/i":                                  "elixir",
		".elm":                                                    "elm",
		".env":                                                    "env",
		".erl":                                                    "erlang",
		"/\\.([fF])(03|08|18|90|95)$/i":                           "fortran",
		".beam":                                                   "erlang",
		".hrl":                                                    "erlang",
		".xrl":                                                    "erlang",
		".yrl":                                                    "erlang",
		".app.src":                                                "erlang",
		"/^Emakefile$/":                                           "erlang",
		"/^rebar(\\.config)?\\.lock$/i":                           "erlang",
		"/(\\.firebaserc|firebase\\.json)/i":                      "firebase",
		".flowconfig":                                             "flowconfig",
		".fs":                                                     "fsharp",
		".fsi":                                                    "fsharp",
		".fsscript":                                               "fsharp",
		".fsx":                                                    "fsharp",
		"/^Gemfile(\\.lock)?$/i":                                  "gemfile",
		"/^\\.git|^\\.keep$|\\.mailmap$/i":                        "git",
		".go":                                                     "go",
		".gd":                                                     "godot",
		".gr":                                                     "grain",
		".gradle":                                                 "gradle",
		"gradlew":                                                 "gradle",
		"/\\.(g|c)sc$/i":                                          "gamescript",
		".gql":                                                    "graphql",
		".graphql":                                                "graphql",
		".groovy":                                                 "groovy",
		"/\\.gv?y$/i":                                             "groovy",
		".gsh":                                                    "groovy",
		"/[Gg]runtfile\\.(js|coffee)/i":                           "gruntfile",
		"gulpfile.js":                                             "gulp",
		"/\\.(hbs|handlebars|(mu)?stache)$/i":                     "handlebars",
		".prg":                                                    "harbour",
		".ha":                                                     "hare",
		".hbp":                                                    "harbour",
		".hbc":                                                    "harbour",
		".rc":                                                     "harbour",
		".fmg":                                                    "harbour",
		".hs":                                                     "haskell",
		".hsc":                                                    "haskell",
		".c2hs":                                                   "haskell",
		".c3":                                                     "c3",
		".lhs":                                                    "haskell",
		"/\\.(hlsl|cginc|cg|shader|fx)$/i":                        "hlsl",
		"/\\.(glsl|vert|frag|geom|tesc|tese|comp)$/i": "glsl",
		"/\\.hx(ml)?$/i":           "haxe",
		"/^procfile/i":             "heroku",
		".heex":                    "heex",
		"heroku.yml":               "heroku",
		".hjson":                   "hjson",
		"/\\.x?html?$/i":           "html",
		".http":                    "http",
		".rest":                    "http",
		".jar":                     "jar",
		".java":                    "java",
		".j2":                      "jinja",
		".jinja":                   "jinja",
		".js":                      "js",
		".es6":                     "js",
		".es":                      "js",
		".mjs":                     "js",
		".js.map":                  "jsmap",
		".json":                    "json",
		".jsonc":                   "json",
		".jsx":                     "jsx",
		".jule":                    "jule",
		"/\\.(jil|jl)/i":           "julia",
		".ipynb":                   "jupyter",
		"/\\.kt[ms]?$/i":           "kotlin",
		".less":                    "less",
		"/\\.l(i?sp)?$/i":          "lisp",
		"/\\.n[ly]$/i":             "lisp",
		".podsl":                   "lisp",
		"/\\.s([s]|(cm)|(exp))$/i": "lisp",
		".ls":                      "livescript",
		".log":                     "log",
		".lua":   "lua",
		"/\\.luau$/i":              "luau",
		"/\\.rbx(?:lx|l|m|s)?$/i":  "luau",
		"/^Makefile/":              "makefile",
		"/^mk\\.config$/":          "makefile",
		"/\\.(mk|mak|make)$/i":     "makefile",
		"/^BSDmakefile$/i":         "makefile",
		"/^GNUmakefile$/i":         "makefile",
		"/^makefile\\.sco$/i":      "makefile",
		"/^Kbuild$/":               "makefile",
		"/^makefile$/":             "makefile",
		"/^mkfile$/i":              "makefile",
		"/^\\.?qmake$/i":           "makefile",
		"/\\.(geo|topo)$/i":        "manifest",
		".cson":                    "manifest",
		".json5":                   "manifest",
		".ndjson":                  "manifest",
		".fea":                     "manifest",
		".json.eex":                "manifest",
		".pawn":                    "pawn",
		".proto":                   "manifest",
		".pytb":                    "manifest",
		".pydeps":                  "manifest",
		"/\\.pot?$/i":              "manifest",
		".ejson":                   "manifest",
		".edn":                     "manifest",
		".eam.fs":                  "manifest",
		".qml":                     "manifest",
		".qbs":                     "manifest",
		".ston":                    "manifest",
		".ttl":                     "manifest",
		".rviz":                    "manifest",
		".sol":                     "solidity",
		".syntax":                  "manifest",
		".webmanifest":             "manifest",
		"/^pkginfo$/":              "manifest",
		".moon":                    "moonscript",
		"/^mime\\.types$/i":        "manifest",
		"/^METADATA\\.pb$/":        "manifest",
		"/[\\/\\\\](?:magic[\\/\\\\]Magdir|file[\\/\\\\]magic)[\\/\\\\][-.\\w]+$/i":             "manifest",
		"/(\\\\|\\/)dev[-\\w]+\\1(?:[^\\\\\\/]+\\1)*(?!DESC|NOTES)(?:[A-Z][-A-Z]*)(?:\\.in)?$/": "manifest",
		"lib/icons/.icondb.js": "manifest",
		"/\\.git[\\/\\\\](.*[\\/\\\\])?(HEAD|ORIG_HEAD|packed-refs|logs[\\/\\\\](.+[\\/\\\\])?[^\\/\\\\]+)$/": "manifest",
		"/\\.(md|mdown|markdown|mkd|mkdown|mdwn|mkdn|rmd|ron|pmd)$/i":                                         "markdown",
		".mdx":                                  "markdownx",
		".marko":                                "marko",
		".nim":                                  "nim",
		".nims":                                 "nim",
		".nimble":                               "nim",
		".nix":                                  "nix",
		".npmrc":                                "npm",
		".npmignore":                            "npm",
		".nut":                                  "squirrel",
		"/\\.mm?$/i":                            "objective-c",
		".pch":                                  "objective-c",
		".x":                                    "objective-c",
		"/\\.eliom[i]?$/i":                      "ocaml",
		"/\\.ml[4lyi]?$/i":                      "ocaml",
		".mt":                                   "metal",
		".odin":                                 "odin",
		"/\\.pas(cal)?$/i":                      "pascal",
		".lpr":                                  "pascal",
		".inc":                                  "pawn",
		".sma":                                  "pawn",
		"/\\.p(wn)?$/i":                         "pawn",
		".sp":                                   "sourcepawn",
		"/\\.p(er)?l$/i":                        "perl",
		".al":                                   "perl",
		"/\\.p([hm]|(lx))$/i":                   "perl",
		"/\\.(psgi|xs)$/i":                      "perl",
		".pl6":                                  "perl",
		"/\\.[tp]6$|\\.6pl$/i":                  "perl",
		"/\\.(pm6|p6m)$/i":                      "perl",
		".6pm":                                  "perl",
		".nqp":                                  "perl",
		".p6l":                                  "perl",
		".pod6":                                 "perl",
		"/^Rexfile$/":                           "perl",
		"/\\.php([st\\d]|_cs)?$/i":              "php",
		"/^Phakefile/":                          "php",
		".pony":                                 "ponylang",
		".pcss":                                 "postcss",
		"/\\.ps[md]?1$/i":                       "powershell",
		".ps1xml":                               "powershell",
		".prettierignore":                       "prettier",
		"prisma.yml":                            "prisma",
		".pde":                                  "processing",
		".jade":                                 "pug",
		".pug":                                  "pug",
		".purs":                                 "purescript",
		".ipy":                                  "python",
		".isolate":                              "python",
		".pep":                                  "python",
		"/\\.gypi?$/i":                          "python",
		".pyde":                                 "python",
		"/\\.py([wi3tp]|(de))?$/i":              "python",
		".tac":                                  "python",
		".wsgi":                                 "python",
		".xpy":                                  "python",
		".rpy":                                  "python",
		"/\\.?(pypirc|pythonrc|python-venv)$/i": "python",
		"/^(SConstruct|SConscript)$/":           "python",
		"/^(Snakefile|WATCHLISTS)$/":            "python",
		"/^wscript$/":                           "python",
		"/\\.(r|Rprofile|rsx|rd)$/i":            "r",
		".razor":                                "razor",
		"/\\.res?i?$/i":                         "reasonml",
		".rst":                                  "restructuredtext",
		"/\\.(rb|ru|ruby|erb|gemspec|god|mspec|pluginspec|podspec|rabl|rake|opal)$/i":                            "ruby",
		"/^\\.?(irbrc|gemrc|pryrc|ruby-(gemset|version))$/i":                                                     "ruby",
		"/^(Appraisals|(Rake|[bB]uild|Cap|Danger|Deliver|Fast|Guard|Jar|Maven|Pod|Puppet|Snap)file(\\.lock)?)$/": "ruby",
		"/\\.(jbuilder|rbuild|rb[wx]|builder)$/i":                                                                "ruby",
		"/^rails$/":         "ruby",
		".watchr":           "ruby",
		".rs":               "rust",
		"/\\.(sc|scala)$/i": "scala",
		"/\\.s[ac]ss$/i":    "scss",
		"/\\.(sh|rc|bats|bash|tool|install|command)$/i":                                                  "shell",
		"/^(\\.?bash(rc|[-_]?(profile|login|logout|history|prompt))|_osc|config|install-sh|PKGBUILD)$/i": "shell",
		"/\\.(ksh|mksh|pdksh)$/i": "shell",
		".sh-session":             "shell",
		"/\\.zsh(-theme|_history)?$|^\\.?(antigen|zpreztorc|zlogin|zlogout|zprofile|zshenv|zshrc)$/i": "shell",
		"/\\.fish$|^\\.fishrc$/i": "shell",
		"/^\\.?(login|profile)$/": "shell",
		".inputrc":                "shell",
		".tmux":                   "shell",
		"/^(configure|config\\.(guess|rpath|status|sub)|depcomp|libtool|compile)$/": "shell",
		"/^\\/(private\\/)?etc\\/([^\\/]+\\/)*(profile$|nanorc$|rc\\.|csh\\.)/i":    "shell",
		"/^\\.?cshrc$/i":       "shell",
		".profile":             "shell",
		".tcsh":                "shell",
		".csh":                 "shell",
		".sk":                  "skript",
		".sqf":                 "sqf",
		"/\\.(my)?sql$/i":      "sql",
		".ddl":                 "sql",
		".udf":                 "sql",
		".hql":                 "sql",
		".rkt":                 "racket",
		".viw":                 "sql",
		".prc":                 "sql",
		".cql":                 "sql",
		".db2":                 "sql",
		"/\\.(styl|stylus)$/i": "stylus",
		".sln":                 "csproj",
		"/\\.sv(h)?$/i":        "systemverilog",
		".svelte":              "svelte",
		".svg":                 "svg",
		".swift":               "swift",
		"/\\.tex(i)?$/i":       "tex",
		".ltx":                 "tex",
		".aux":                 "tex",
		".sty":                 "tex",
		".dtx":                 "tex",
		".cls":                 "tex",
		".ins":                 "tex",
		".lbx":                 "tex",
		".mkiv":                "tex",
		"/\\.mk[vi]i$/i":       "tex",
		"/^hyphen(ex)?\\.(cs|den|det|fr|sv|us)$/": "tex",
		"/\\.te?xt$/i":        "text",
		".rtf":                "text",
		"/\\.i?nfo$/i":        "text",
		".msg":                "text",
		"/\\.(utxt|utf8)$/i":  "text",
		".toml":               "toml",
		".travis.yml":         "travis",
		".ts.map":             "tsmap",
		"/.*\\.d\\.ts/i":      "typescript-def",
		".ts":                 "ts",
		".tsx":                "tsx",
		".twig":               "twig",
		".v":                  "v",
		".vh":                 "v",
		".vala":               "vala",
		".vapi":               "vala",
		".vb":                 "vb",
		".vbs":                "vb",
		".vbhtml":             "vb",
		".vbproj":             "vb",
		".vba":                "vba",
		".vcxproj":            "vcxproj",
		".verse":              "verse",
		".vscodeignore":       "vscodeignore",
		".vue":                "vue",
		".zu":                 "zura",
		".zura":               "zura",
		".wat":                "wasm",
		".wast":               "wasm",
		".wasm":               "wasm",
		".xml":                "xml",
		".xaml":               "xaml",
		"/\\.ya?ml$/i":        "yaml",
		"/^yarn(\\.lock)?$/i": "yarn",
		".yarnrc":             "yarn",
		".zig":                "zig",
		".maeel":              "maeel",
		"/\\.(tfvars|tf)$/i":  "terraform",
		"/\\.mojo$/i":         "mojo",
		".🔥":                  "mojo",
	}
}

// TODO: Implement this properly
func GetLanguageUrl (extension string, isRedacted bool) string {
  languages := GetKnownLanguages()

  fmt.Print(extension)
  var language = languages[extension]

  if language == "" || isRedacted {
    language = "text"
  }

  if extension == "." {
    language = "idle"
  }

  return fmt.Sprintf("https://raw.githubusercontent.com/circles-00/nvim-discord-status/main/assets/icons/%s.png", language)
}
