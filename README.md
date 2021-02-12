A    a    B    b    C  
2 -> 3 <- 2 -> 3 <- 2  

Loop 1  
a requests 3 from A  
A gives a 2  
A -> a 2  
A 2 out of 2  

b requests 3 from B  
B gives b 2  

A -> a 2  
B -> b 2  
a 2 out of 3  
b 2 out of 3  
A 2 out of 2  
B 2 out of 2  

a requests 1 from B  
B gives a 1/2  
B adjusts b to 1 1/2  

Loop 2  
A -> a 2  
B -> b 1 1/2  
B -> a 1/2  
a 2 1/2 out of 3  
b 1 1/2 out of 3  
A 2 out of 2  
B 2 out of 2  

b requests 1 1/2 from C  
C gives b 1 1/2  

A -> a 2  
B -> b 1 1/2  
B -> a 1/2  
C -> b 1 1/2  
a 2 1/2 out of 3  
b 3 out of 3  
A 2 out of 2  
B 2 out of 2  
C 1 1/2 out of 2  


Loop 3
a requests 1/2 from C  
C gives a 1/2  

A -> a 2  
B -> b 1 1/2  
B -> a 1/2  
C -> b 1 1/2  
C -> a 1/2  
a 3 out of 3  
b 3 out of 3  
A 2 out of 2  
B 2 out of 2  
C 2 out of 2  

b has no more requests to make

Loop 4
a has no more requests to make
b has no more requests to make

cost
A -> a 2 distance 1 weight 2 cost of 2
B -> b 1 1/2 distance 1 weight 1 1/2 cost of 1 1/2
B -> a 1/2 distance 1 weight 1/2 cost of 1/2
C -> b 1 1/2 distance 1 weight 1 1/2 cost of 1 1/2
C -> a 1/2 distance 3 weight 1/2 cost of 1 1/2
total cost: 7
