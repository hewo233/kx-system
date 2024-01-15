// RAM[2] = RAM[0] * RAM[1]    
    
    
    @R0
    D=M
    @n
    M=D // n = RAM[0]


    @i
    M=0 //init i = 0

    @sum
    M=0 //init sum = 0

(LOOP)

    @i
    D=M
    @n
    D=D-M
    @ANS
    D;JEQ //if i == n, go to ANS

    @R1
    D=M
    @sum
    M=D+M //sum = sum + RAM[1]

    @i
    M=M+1 //i++

    @LOOP
    0;JMP

(ANS)

    @sum
    D=M
    @R2
    M=D //RAM[2] = sum

(END)
    
    @END
    0;JMP