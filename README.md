A    a    B    b    C
2 -> 3 <- 2 -> 3 <- 2

a requests 3 from A
b requests 3 from B
A has request for 3 from a
B has request for 3 from b
A gives a 2
B gives b 2

a 2 out of 3
b 2 out of 3
A 2 out of 2
B 2 out of 2
C 0 out of 2

a requests 3 from A
b requests 3 from B
a requests 1 from B
b requests 1 from C
A has request for 3 from a
B has request for 3 from b
B has request for 1 from a
C has request for 1 from b
A gives a 2
B gives a 2/3
B gives b 1 1/3
C gives b 1

a 2 2/3 out of 3
b 2 1/3 out of 3
A 2 out of 2
B 2 out of 2
C 1 out of 2

a requests 3 from A
b requests 3 from B
a requests 1 from B
b requests 1 2/3 from C
a requests 1/3 from C
A has request for 3 from a
B has request for 3 from b
B has request for 1 from a
C has request for 1 2/3 from b
C has request for 1/3 from a
A gives a 2
B gives a 2/3
C gives a 1/3
B gives b 1 1/3
C gives b 1 2/3

a 3 out of 3
b 3 out of 3
A 2 out of 2
B 2 out of 2
C 2 out of 2


A -> a for 0 with priority 1
B -> b for 0 with priority 1
B -> a for 0 with priority 2
C -> b for 0 with priority 2
C -> a for 0 with priority 3
A -> b for 0 with priority 3

A -> a for 3 with priority 1 give 2
B -> b for 3 with priority 1 give 2
B -> a for 0 with priority 2
C -> b for 0 with priority 2
C -> a for 0 with priority 3
A -> b for 0 with priority 3

A -> a for 3 with priority 1 give 2
B -> b for 3 with priority 1 give 1 1/3
B -> a for 1 with priority 2 give 2/3
C -> b for 1 with priority 2 give 1
C -> a for 0 with priority 3
A -> b for 0 with priority 3

A -> a for 3 with priority 1 give 2
B -> b for 3 with priority 1 give 1 1/3
B -> a for 1 with priority 2 give 2/3
C -> b for 1 2/3 with priority 2 give 1
C -> a for 1/3 with priority 3 give 2
A -> b for 0 with priority 3

A -> a for 3 with priority 1 give 2
B -> b for 0 with priority 1
B -> a for 0 with priority 2
C -> b for 0 with priority 2
C -> a for 0 with priority 3
A -> b for 0 with priority 3

A -> a for 3 with priority 1 give 2
B -> b for 0 with priority 1
B -> a for 1 with priority 2 give 1
C -> b for 0 with priority 2
C -> a for 0 with priority 3
A -> b for 0 with priority 3

A -> a for 3 with priority 1 give 2
B -> b for 3 with priority 1 give 1 1/3
B -> a for 1 with priority 2 give 2/3
C -> b for 0 with priority 2
C -> a for 0 with priority 3
A -> b for 0 with priority 3

A -> a for 3 with priority 1 give 2
B -> b for 3 with priority 1 give 1 1/3
B -> a for 1 with priority 2 give 2/3
C -> b for 0 with priority 2
C -> a for 1/3 with priority 3 give 1/3
A -> b for 0 with priority 3

A -> a for 3 with priority 1 give 2
B -> b for 3 with priority 1 give 1 1/3
B -> a for 1 with priority 2 give 2/3
C -> b for 1 2/3 with priority 2 give 1 2/3
C -> a for 1/3 with priority 3 give 1/3
A -> b for 0 with priority 3
