package main

import (
	"advent/pkg/execute"
)

var tests = execute.TestCases{
	{
		testpuzzle,
		`26397`,
		`288957`,
	},
}

var testpuzzle = `[({(<(())[]>[[{[]{<()<>>
[(()[<>])]({[<{<<[]>>(
{([(<{}[<>[]}>{[]{[(<()>
(((({<>}<{<{<>}{[]{[]{}
[[<[([]))<([[{}[[()]]]
[{[{({}]{}}([{[{{{}}([]
{<[[]]>}<{[{[{[]{()[[[]
[<(<(<(<{}))><([]([]()
<{([([[(<>()){}]>(<<{{
<{([{{}}[<[[[<>{}]]]>[]]`
var puzzle = `([<{(<{(({<{{{[]<>}<<>()>}<<()()>[()()])}{({()[]}[<>{}])<(<><>)<[]{}>>}>{{((()<>)<[]()>)[{[]<>}(()[])]
<<[[<[(<[<[[[{()<>}(<>{})]]][{(<{}{}>(<>))[[<>](<><>)]}[({{}<>}<[]()>){<{}{}>[{}[]]}]]><{(<([][]
{<([[[({{(<[([{}<>]<{}{}>)[[<>[]]]]{{(()<>)[()]}}>){((([{}<>]{[]{}})<<[][]><<>{}>>){([[]{}]<<>()>){
{{[{({{(({<[<{[]{}](<>)>][<(()[])(()[])>[[<>{}](<>())]]><<([{}]{{}{}})[{[]}{{}[]}]>[{<()<>><<>>}]>
(<<<[(({[[[{<[{}<>]{<><>}><({})>}{{{()}{(){}}}}][{[({}{})[<><>]]<({}[])[()()]>}(<([]<>)<[](
<[{[((<<[[<[({[]()}{(){}})[{<>[]}<()()>]]>]]<[{{[[[][]]<[]()>][<<>{}><<><>>]}}({[<[]()><{}[]>]{<[]
(({((({{[[(([<[]<>>(())]{{()<>}<<>[]>}))]]([<([{{}()}(<>{})][({}<>)])>](({{<()<>>[{}())}{((){}){[]()}}})
({{({(<<{<<([<{}>])(<[()<>][{}[])>{[()()][{}{}]})>[[[{{}[]}[(){}]][{<>{}}{()<>}]]({<<>[]>{{
<[<<{[(<{[{({({}[])(()[])}[{{}<>}{{}<>}]){<{<>[]}[<>()]>{<{}()>}}}{{{{()()}}[{{}<>}({}<>)]}}}}>{{[[{<[
[{([<{{<<<[([<[]()>])]([{{<>{}}}{(()[])[<>[]]}])><[(<{()<>}{<>}>[([]{})<<>{}>]}([{()<>}({}[])]<<{}[]>({}{}
<([<([{{<<{<(([]<>)<{}{}>)>({[(){}]}((<>[])<{}[]>))}<{<(()[])<[]>>({[][]}(()[]))}(<{()[]}{[]()}})>>[([({()[
{{{[({(<[{[[{{()<>}[<>()]}{({}<>)<{}()>}]{{[()()][(){}]}[{<>()}]}]}]>)(({<<(({{}<>}([]<>))({<>[]}<[]{
({{({<([{[{{[[[]()><{}[]>]({{}[]}[{}()])}}<{[{[]()}[<>]]<[{}<>]<()<>>>}>]{<<(((){}){<>()})[[[][]]]>[([<><>]{
(<[[{{[(([{{<((){}){<>{}}>{<{}[]><<>[]>}}(({<>[]}<<><>>)({()[]}<()<>>))}<{(<()()>[{}{}]){{
{{{[{<<(<<[{({(){}}[{}{}])}[{[<>[]]{()()}}{[()]{(){}}}]]>>)><{((<<{<[]{}>([]{})}{{{}{}}}><([{}[]](()()))>
[<(<<<[(([((<<<><>>>[<[]<>>(()[])])[{((){}){<><>}}<<{}{}>(()[])>]){(([[]<>]<[]>)<{[]}[[]<>]>)<<{{}{}}[[
<(<<({(([[({[[[])(()[])][<[]()>{{}[]}]}{({{}}{(){}})(((){})([]<>))})([[([][])<{}<>>]<<()><<>()>>](
[{[[<{<<{<{<{<()()><{}[]>}([(){}](<>()))>(<{<>[]}([])><[[][]](<>{})>)}(({<[]<>>{()()}}<<[][]>>){({()[]
<([<{((({({[<(()[])[{}]>({[]()})]<<[[]]>>}<[[{[]{}}[{}{}]]([{}[]]{(){}})]({(<>())})>)}({<[{{{}
({{([(<<{({(<(<><>)<{}<>>><<{}()>[()()]>){<{{}()}(<>[]>>}})}>>[<<<{(({{}<>}(()[])){<[][]><{}>})
[<((<<{{<[[[[{[]()}][{[]{}}(<>())]]]]{<<{(()<>)((){})}<[<>[]]{[][]}>>(<(()())(()[])><{{}<>}({}())>)>[<<[(
[(<<(<<[(([{((()())<<><>>}[{()<>}({}<>)]}]{{<(<>[]){(){}}>([<>()]({}{}))}})<<(<[[][]][[][]]>[([][]){(){
<[([[<<[{(<([([]{}){{}<>}][{()}([][])])>)({<({(){}}){[[]{}][()()]}>{{{[]()}({}{})}(<[]{}><<>()>)}}[{[<[]()><
{<{<[[{({<[<<({})<()[]>>((<>{})<[]{}>)>]>})}((<[[(<<{}{}>{[][]}>)<<[()[]]{<>{}}>{{()[]}[()]}>]<({[
[({({(<<<<<[[<[]{}>]{<<>{}><[]{}>}]><<<([]{})<{}[]>><[{}[]][()<>]>>[[<{}[]>{[]}](<{}()>[{}<>])]>>>>{<{([((
<({[({<{[<({{<()[]]}[{()[]}[{}<>]]}{[{()<>}{<><>}]([[][]]<[]<>>)})>[[[(<()<>>({}[]))<[(){}]<<><>>>]][[<({}{}
{<(<<(([[{(<[[()[]]<()[]>]{<()<>><()[]>}><{<[]{}>{()()}}>)([([(){}]{<>{}})[{<>{}}[<>{}]]]<<[[]{}]{()
{[([[{[<[<{(<<[]{}]>({<><>}({})))([{{}<>}]<{()<>}[{}[]]>)}<([(()[])<()()>])<{{[]{}}[{}()]}(<<><>>[()()]
<<[{<[({((<<({{}()}{()[]})>[<<{}{}>{<><>}>]>){{<[{[]()}[[]()]]{[{}[]]}>[{<{}{}>{[]<>}}([{}[]]((){}))]}(
<<{((([[({<{[(()<>)<()[]>]({()[]}({}<>))}{{{[]<>}[{}[]]}[{<><>}]}>}<[[([{}()]{(){}})<({}<>)[[]()
{(([({{(([<{{{<>()}<<>()>}[{<>()}(()())]}>]{(<([()<>]{()<>}){{()<>}}><<{{}}{[][]}>(<(){}><{
{<(({{[<([(<<([][])(<>()]>{{{}{}}{[]<>}}>)[{(<()[]>{<>{}})[{()[]}<[]<>>]}{<{{}[]}><<{}>[{}<>
(<{{([<[[{{[{[[]()][()()]}{<<>{}>}][{{<>()}<()()>}(((){}){<>{}})]}{<(({}()))<[(){}]<(){}>>>(<({}[]>[[][]]>)}}
(({({[({<{{{[[()()]({}())]}[([{}]<{}()>)([{}()]{{}{}})]}[((<{}<>><()()>))]}[{{{(<>{}){<><>)}[<()()
(<({{<([{[([[({}<>)][([]<>){[]{}]]][[<()[]>]])][<{[[{}[]](())]<<()[]>(())>}{<{{}[]}[<>{}]>{<[]<>>
(([[([([<[{[(({}[])<[]{}>)]}][[{({()<>}[{}[]])(([]())(()<>))}<[<[]{}>({}<>)]{{{}{}}{[][]}}>]{{{
<[{[{<<{[[<{(({})(<>{}]){([]())[[]()]}}{(<()()>([]())){([]{})<[][]>}}>]{({[[<><>]{[]()}](<{}[]><[]<>>)}
([{<({{([([[{{()[]}}({(){}}[<>{}])>]({{(())([]<>)}<[()<>][[]()]>}([{()[]}{()()}])))](([[[{<>{}}[<>(
<[[[(<[{{[[{{{<>[]><<>>}[{()[]}({}())]}[<[{}<>]<[][]>>]]]}{{({{{<>()}{()<>}}<<{}[]>({}[])>}[[(<><
[[<[<({<[({{[[<>{}][[]{}]]{<()()>{<>{}}}}[<<<>{}><[]<>>><<{}<>>(()[])>]}){<[(([]())[<>()))[({}[])<
<<(<<<<((([[<<{}[]>(<>())>]]<{{[[]<>]<[]>}{[{}()]{<>[]}}}<[[<>{}]({}())]>>)){{([(({}[])({}[]
([{<([<[[[{[[{<>[]}{[]()}>([[]])]<({{}{}}({}[])){(<>{})<()()>}>}<{(([]()){()<>})[{<>{}}{<>()}]}<(([]
<{[<(({([[((((<>[])[[]{}])({[][]}))(<[()<>]{<><>}>)){<<{<>[]}<<>()}>>{({{}{}}(()()))}}]][{
({<{<<<{{{(((([]())[<>]){<(){}><{}[]>})<{([]<>)[<><>]}[{()<>}({}<>)]>){((<[]()>[<>{}]){<()<>>(
[{<(([(<[((([[(){}]{[]()}])))]>)[<{(<<[<<>[]><()()>]>>){[<{{(){}}({}[])}([{}()][<>])><((()<>)[[]])[[{}()](<>
{<<[<{[(<{<([{()[]}({}<>)]{[{}<>][{}()]})([{<>{}}<<>>])>{(<{{}())<{}()>>(({}[])<<>()>))[({<>()}[()[]])<<{}<>
[<([<<<[({([<<(){}><{}[]>>[[<>{}][{}<>]]]<[<()[]><(){}>]([<>{}][<>[]])>)<<{<[]<>>[[]{}]}<({}(
<[[(<[{[[({{{<<>[]][{}[]]}[[[]{}]]}}){{[[(<>{})(()<>)]({<>()})]({({}{})<[]{}>}[<()<>>])}{<{[()[]]}[<
<(<(<((<[((([[<>()]]){<[<>[]]><([]{})>})[<<{<>()}{{}<>}>(((){})[<>[]])>{[[<>()][{}[]]](([][])([]))}])((
[(<{{<(((<<{[[<>[]]]{([]<>)<{}[]>}}[<{{}{}}[{}()]><{[]()}[{}]>]>{[<{{}[]}([]<>)>[[(){}][()()]]]
<(<(([[{[<<<(({}))[{[][]}]>[[([]{})({}[])]]>[<<(<>[]){<>[]}><[()<>]>)<[{{}<>}(<>{})]{[<>()]{{}{}}}>]>{[{<{
({<({{{[[(([<<()[]><<><>>>])(<((<>()))(<<><>><[]<>>)>)){{({[[]()]{[]<>}}[[{}[]]<[]()>])}({({()()}<[]()
{{{<[<{<{[<(<{[]()}><[()[]]{[][]}>){<<[]>[<><>]>([<>()][<>()])}>)<<<(([]())[<>[]]){[[][]](<>{})}><([(){}]{{}{
<[((<<<({<[((<{}<>>[[]{}})((<>[])))({<()()>{<><>}}{(<><>){(){}}})][[<<[]>>[{<>[]}{[]()}]]([{
(((<<{<({<[[[{()()}(()<>))(<()[]>((){}))][{({}[])[{}()]}<<{}{}>{()()}>]]>})<<{{<[[<>{}]<[][]>]((<><
[<[{[[{{<{[(<<<>[]><{}()>>)({[{}[]]}((<>())<()[]]))]{{({()[]}([]<>)){{<>[]}[[][]]}}}}{({(<<>[]>{[]})[[{}
<<[({[[[[([<(<<>[]>{()[]})>{<[()<>]>[[[]<>]({})]}]<(<{{}}([])>[<{}()>])[{{<>()}<(){}>}<<<>{}}[{}{}]>]
([(((([[<[{({<{}>})<([{}<>][(){}])<<[]()><()[]>>>]{[([<>[]](()()))[{{}[]}[{}<>]]]({({}<>)[<>()]}<{()<>}{[]{
[[<{[{<{{(<(<[<>{}]><([]{})[{}<>]>)>{{[{<>[]}[{}()]]<<()>{[]}>}[[((){})([])]{<<>()>([][])}]})
{<[{{{<{([<[<{{}{}}[[]{}]><{{}[]}>]{(<{}<>><{}[]>)(<(){}><()()>)}>]<<<({{}}(<>()))[<{}[]>(<>{})]>({<[]>{()<
{<[{[<(({[{{({{}[]}[()[]]}}<(([]())[()[]]){(()){<>()}}>}[<<{<><>}>[([]{}){{}<>}]>[{{<>{}}{()[]}}
<{[{([<{<((<[<()[]>{[][]}]<<()[]><<>{}>>><([<><>]<()()>)({{}()}<(){}>)>){<(([][]){[][]})(<[]><{}{}>)>
(([<((((({{{[<()[]]{{}<>}]{<()()>({}<>)}}([[{}()]<<>>][((){})<<>[]>])}}(<{{[(){}]{()()}}}{([{}()
<<<((<[[{[<([<{}<>>(<><>)]<{[]<>}<[]<>>>)>]{([{[[]<>][<>()]}({[]<>}<<>[]>)])[[{{<>[]}(()())}
{<(<<<(({{(((([]<>)(<>[]))))}[(([[[][]][{}()]])<<(<>{}><{}{}>>{(()[])([]())}>)]}([({((<>){[]<>}){<[][]
({(([[[[<{[{(<{}[]>[{}<>])(<[]<>>({}[]))}{([<><>]{{}{}}){{{}[]}[[][]]}}]<[((()[])([]<>))[{[][]
[<([[([[<({({<[]()>}<[[]()]([]<>)>)([[()[]]]{[<>{}][()()]})))>]<({{[{<{}{}>}{{()()}[<>[]]}]{([{
({<<{{({[{[{{[[]<>]<(){}>}<<<>{}>({}{})>}<{[{}<>]}(<[]>[()[]]))]}]<<<<<{(){}}<<>[]>>{<[]{}>{(){}}
(({<([<((<((({{}[]}[[]()])){{[[]()]<[]{}>}(<[][]>{<>})})({{[()()]{{}<>}}(({}{})[<><>])}[[([]{})[
({{[(<{{<<<<[{{}{}}{{}{}}][<[]()>[[]()]]>([{{}}<[]<>>][<[]{}>{()<>}])>[<{({}{})({}<>>}[<{}<>>[<>()]]>]><<
{[((({{(<(<{[{[]()}}<<[][]>(<>[])>}<(<[]()><()<>>){[<><>]<<>()>}>><[{<()[]>}{[{}[]](<>{})}]>)<({<{<><>}>}{
[(([{([<<<([[(()())][([]{})[<>{}]]]({((){})((){})}(([][])[()<>])))[<((<>())[<>{}])>{{[[]<>](
<(({<([(<<<{([{}[]]<[]>)}({{<>{}}[[]{}]}(<{}[]>(()<>)))>>({<{<<>()>((){})}((()())[<>{}])><[<[]()>]{<
[<[{({<[({[(<[{}[]]{[]<>}>(({}{}){{}[]}))(<{()[]}>)]{<([[][]][{}<>])<<<>{}>>>{{((){}){()()}}<{{}}
[<([{<(<(<<{([[]<>]{(){}})((()<>)[[]])}<[<{}()>]<{<><>}<{}{}>>>>([{({}<>){{}()}}]<((()()))<<[][]>
({[{((<{[{([<{()()}{<>()}>[[()[]]]])}[{([[{}]{<>{}}]({[]()}[{}[]]))<{<{}<>>}>}]][[[(<({}{}){{}
[[<{{<{[([(<([[]<>][()<>]))[[[<>[]]]({()()}([][]))])(([[[][]]([]())])<{{{}<>}([]{})}<{[][]}{
<<({<[{[<{<<([<>()]({}<>))][([{}()][{}[]])]>}>]}{[<[[[[[[]{}]<{}>]<<{}()>{()[]}>]][<{({}())
(<<[({[([([[(<<>{}>){[(){}][()]}][[[(){}]{[]<>}]]]<[([<>[]]([]))[{{}<>}{()()}]]<[<{}[]>(()())]<<<><>>{{}[]
{([[[<[(<({<{{{}{}}[[]<>]}<[()[]]{<>[]}>>}<<<(()())<{}[]>>[([][])<<><>>]>[(((){})(()[])){<{}()>(()[])
{[<{[[[((<{[{[[]<>]{{}()}}[{{}<>}{<><>}]][<<()<>>{<>{}}>[{[]{}}((){})>]}(<[<[][]>({}())]<[<>[]](()[])
{({(<<{({({([<(){}>[()()]](<[]><[]<>>))([{<>()}<{}<>>]{<[]<>>[[][]]})})({[(([]<>)[[]])]{([[]{}]{{}()}
{{[<{<<{<{({(({}()))[<<>{}]{()()}]}({{[]{}}([][])}[<<>()>{[]<>}]))}>(<<[{<[]()><<>{}>}]{((<>[]))<(
{{[<[{<(([({([{}{}]({}[]))<<()>(<>[])>})[<<{[][]}<[]()>><{{}[]}<<>{}>>>{{{{}{}}{<>{}}}([{}
[[([(<{{[{{{{{()()}[[][]]}([{}()](<>))}<{<<>[]>[{}[]]}<<{}()>>>}}]<<[{[<{}>({})](<()[]>)}](<<<[][]>(
{<[([{{{((<{((<>())){<[]<>>}}<<{<><>}[<>{}]>{{()[]}(()[])}>>{{{<()()>[[]{}]}[[<><>]<[]()>]}([{{}}{()()}][(<>[
[{<{[([((<(<(((){})[<>]){({}[])<{}<>>}>(((()())([]{}))<[()()]{[]{}}>))([((()<>)<{}()})<{{}[]}<<>
{(({{{<[<({([<{}()>]{<<><>>{<><>}})}){<(((<>[])<[][]>)(<[]<>>{{}<>}))(<{<>{}}<()[]>>[[{}[]]({}())])>
{<([{(<{<<{(<{()}{<>[]}>)([[{}()]{<>{}}]<({}{})>)}[<<<()[]>({}<>)>[<[]>[<>{}]]><(<{}()>[<>{}])>]
<{{{{(({{{[{<[[][]]<()<>>>}[({{}[]}([]<>))({{}{}}{<>()>)]]}}}([{<{{<{}<>><<><>>}}{([{}<>]<<>{}>)
{<[<<{<([[[[[[<>[]](()[])]({{}[]}<{}>)][((()())({}()))<{{}[]}(<>())>]]([<[{}[]][<>[]]>]<{([]{})}{(
{<{[[([<([[[(<()()>({}[]))[{()[]}]]((({}[]><()()>)[({}){[]<>}])]]{<<{<()[]>}{<<>[]>{<>()}}>><[{(()<>)({}<
[[<[[<{{{{<<{{[][]}<(){}>}><{([]())[[]<>]}<<{}[]>(<>())>>>}<{{[[(){}]<<>()>]<{{}()}[[][]]>}}<(<<(
<[[(({({(({{{([]{}){<>{}}}{[<><>]{[]<>}}}}))[({[[({}<>){<><>}]{<<>{}><<><>>}]}[<{<{}<>>(<>{})}(([
({<<([{[<[<([((){})([]())][(()<>)({}[])]){{[{}()]({}[])}[{{}()}{{}()}]}>{{({<>[]}{<>})[<(){}>
<(<{<[<{({{[<{{}[]}{[]<>)>][{[{}{}][()<>]}(<[]<>>[[]{}])]}(([[[]<>]{()()}]<[{}[]]{<>{}}>)((([][]))
<{<(<<((<<[[{<()()><<><>>}<(()<>)({}<>)>>[<<[][]>(<><>)>{(<>[])}]](<{<()[]>([]())}[{<><>}<{}[]>]>)>>{(
{(({(<(<<[[[{<<>()><{}()>}[[<><>]{()<>>]]{[{(){}}]}]]>>)(<<(<({([]())<<><>>}({[]()}({}<>)))>
<{<<[<<{[<([[(()())<<><>>][<<>[]>(<>())]]({<{}{}><<><>>}{[{}]<<>[]>})){[{{[]}<{}[]>}[[(){}]{[]<>}]]
({[[{<[([[[{[[(){}]{[][]}]<({}[])({}())>}{{<()[]>{<><>}}<{<>}>}]{<<<()<>]{<>{}}>[<[]<>>{{}()}]><[<<
<{{{([({<{{([{()}[()<>]][(<>())(())])([({}[])[{}{}]]({{}{})([]())))}{(<[<>()]{()<>}><[()<>][<>[]]
((<[(<[(({{(<<(){}>{(){}}>({()[]}[<>[]]))<(((){})({}()))<<[]{}><()[]>>>}})<[{([<<>{}>]){<{<>{}}{()<>}>}}{{(`