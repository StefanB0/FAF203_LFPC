The language used in this example lexer is called lazymonkey
Lazy monkey is a simple scripting language of managing directories
The main features are the keywords "move", "read", "weight" and "rename", which aid in file and folder manipulation.
"move" moves a file or folder to the specified directory.
"read" returns an array with the names of all the files in the specified folder.
"weight" returns the size of a file or folder.
"rename" changes the name of a file or folder.

Here is the full grammar of the language:

S = {<source code>}

Vn = {<source code>, <function declaration>, <set of affirmations>, <affirmation>, <variable declaration>, <variable>, <identifier>, <value>, <integer>,<non zero digit>, <digit>,<string>, <letter>, <character>, <boolean>, <calculation>, <function call>,<conditional statement>, <condition>, <assignment>, <operation>, <manipulation>, <path>}

Vt = {0, ..., 9, a, ..., z, A, ..., Z, Program, let, for, while, break, if, else, function, return, move, read, weight, rename, true, false, ;, “_”, “&&”, “||”, “:”, “=”, “==”, “;”, “{”, “}”, “<”, “>”, “<=”, “>=”, “!=”, “(”, “)”, “+”, “-”, “*”, “/”, “””, ““”, “,”, “.”}}

P = {
    <source code> -> <function declaration>* Program {<set of affirmations>}
    <function declaration> -> function <indentifier> (<identifier>(, <identifier>*)) {<set of affirmations> return <value>}
    <set of affirmations> -> <affirmation>|<affirmation> <set of affirmations>
    <affirmation> -><variable declaration>;|
                    <array declaration>;|
                    <function call>;|
                    <array assignment>;
                    <calculation>;|
                    <break>;|
                    <manipulation>;|
                    if <conditional statement> {<set of affirmations>} |
                    if <conditional statement> {<set of affirmations>} else {<set of affirmations>} |
                    while (<conditional statement>) {<set of affirmations>} |
                    for (<conditional statement>) {<set of affirmations>}
    <conditional statement> -> <boolean> | <value> (<condition> <value>)*
    <condition> -> && | || | == | < | > | <= | >= | !=
    <variable declaration> -> let <identifier>| let <identifier> = <value>
    <variable> -> <identifier>
    <array> -> <identifier>
    <array declaration> -> let <identifier>[integer] | let <identifier>[integer] = {(<value>, )*}
    <array assignment> -> <identifier>[integer] = <value>
    <array call> -> <identifier>[integer]
    <value> -> <boolean>|<integer>|<string>|<character>|<path>|<variable>|<function call>|<array>|<array call>
    <function call> -> <identifier>(<value>(, <value> )*)
    <assignment> -> <identifier> = <calculation>
    <calculation> -> <value> (<operation> <value>)*
    <operation> -> + | - | * | /
    <manipulation> -> move(<path>, <path>)|
                    read <path>|
                    rename(<path>, <string>|
                    weight <path>
    <break> -> break
    <integer> -> {- | }{0|<non zero digit><digit>*}
    <identifier> -> <letter>|{<letter>|<digit>}*
    <path> -> <string>
    <string> -> <character>+
    <boolean> -> true | false
    <non zero digit> -> 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9
    <digit> -> 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9
    <letter> -> a | b | c | … | A | B | C | … | Z
    <character> -> 0 | … | 9 | a | … | z | A | … | Z | _ | @| # | $ | % | : | = | ; | { | } | < | > | ! | ( | ) | + | - | * | / | .
}

