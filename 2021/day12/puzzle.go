package main

import (
	"advent/pkg/execute"
)

var tests = execute.TestCases{
	{
		testpuzzle,
		`10`,
		``,
	},
	{
		testpuzzle2,
		`19`,
		``,
	},
	{
		testpuzzle3,
		`226`,
		``,
	},
}

var testpuzzle = `start-A
start-b
A-c
A-b
b-d
A-end
b-end`
var testpuzzle2 = `dc-end
HN-start
start-kj
dc-start
dc-HN
LN-dc
HN-end
kj-sa
kj-HN
kj-dc`
var testpuzzle3 = `fs-end
he-DX
fs-he
start-DX
pj-DX
end-zg
zg-sl
zg-pj
pj-he
RW-he
fs-DX
pj-RW
zg-RW
start-pj
he-WI
zg-he
pj-fs
start-RW`
var puzzle = `xq-XZ
zo-yr
CT-zo
yr-xq
yr-LD
xq-ra
np-zo
end-LD
np-LD
xq-kq
start-ra
np-kq
LO-end
start-xq
zo-ra
LO-np
XZ-start
zo-kq
LO-yr
kq-XZ
zo-LD
kq-ra
XZ-yr
LD-ws
np-end
kq-yr`
