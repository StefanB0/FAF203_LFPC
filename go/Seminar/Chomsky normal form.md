# 0

P = {
    S -> AC
    S -> bA
    S -> B
    S -> aA
    **A -> epsolon**
    A -> aS
    A -> ABAb
    B -> a
    B -> AbSA
    C -> abC
    D -> AB
}

# 1. eliminate epsilon production

P = {
    S -> AC
    S -> bA
    S -> B
    S -> aA
    **A -> epsolon**
    A -> aS
    A -> ABAb
    B -> a
    B -> AbSA
    C -> abC
    D -> AB
    S -> C
    S -> b
    S -> a
    A -> BAb
    A -> ABb
    B -> AbS
    B -> bSA
    D -> B
}

# 2. Eliminate renamings

???

# 3. Eliminate nonproductive

Prod(G) = {A | A := Vn, all A => V, V;=Vt}

# 4. Eliminate non-accesible states
