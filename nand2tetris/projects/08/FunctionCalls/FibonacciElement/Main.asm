(Main.fibonacci)
@0
D=A
@R13
M=D
(Main_Main.fibonacci.INIT_LOCAL_LOOP_START)
@R13
D=M
@Main_Main.fibonacci.INIT_LOCAL_LOOP_END
D;JEQ
@SP
A=M
M=0
@SP
M=M+1
@R13
MD=M-1
@Main_Main.fibonacci.INIT_LOCAL_LOOP_START
0;JMP
(Main_Main.fibonacci.INIT_LOCAL_LOOP_END)
@ARG
D=M
@0
A=D+A
D=M
@SP
A=M
M=D
@SP
M=M+1
@2
D=A
@SP
A=M
M=D
@SP
M=M+1
@SP
AM=M-1
D=M
A=A-1
D=M-D
M=-1
@LT_TRUE_Main_0
D;JLT
@SP
A=M-1
M=0
(LT_TRUE_Main_0)
@SP
AM=M-1
D=M
@N_LT_2
D;JNE
@N_GE_2
0;JMP
(N_LT_2)
@ARG
D=M
@0
A=D+A
D=M
@SP
A=M
M=D
@SP
M=M+1
@SP
A=M-1
D=M
@R13
M=D
@LCL
D=M
@5
A=D-A
D=M
@R14
M=D
@ARG
D=M+1
@SP
M=D
@LCL
D=M
@1
A=D-A
D=M
@THAT
M=D
@LCL
D=M
@2
A=D-A
D=M
@THIS
M=D
@LCL
D=M
@3
A=D-A
D=M
@ARG
M=D
@LCL
D=M
@4
A=D-A
D=M
@LCL
M=D
@R13
D=M
@SP
A=M-1
M=D
@R14
A=M
0;JMP
(N_GE_2)
@ARG
D=M
@0
A=D+A
D=M
@SP
A=M
M=D
@SP
M=M+1
@2
D=A
@SP
A=M
M=D
@SP
M=M+1
@SP
AM=M-1
D=M
A=A-1
M=M-D
@Main_RETURN_LABEL_0
D=A
@SP
A=M
M=D
@SP
M=M+1
@LCL
D=M
@SP
A=M
M=D
@SP
M=M+1
@ARG
D=M
@SP
A=M
M=D
@SP
M=M+1
@THIS
D=M
@SP
A=M
M=D
@SP
M=M+1
@THAT
D=M
@SP
A=M
M=D
@SP
M=M+1
@5
D=A
@1
D=D+A
@SP
D=M-D
@ARG
M=D
@SP
D=M
@LCL
M=D
@Main.fibonacci
0;JMP
(Main_RETURN_LABEL_0)
@ARG
D=M
@0
A=D+A
D=M
@SP
A=M
M=D
@SP
M=M+1
@1
D=A
@SP
A=M
M=D
@SP
M=M+1
@SP
AM=M-1
D=M
A=A-1
M=M-D
@Main_RETURN_LABEL_1
D=A
@SP
A=M
M=D
@SP
M=M+1
@LCL
D=M
@SP
A=M
M=D
@SP
M=M+1
@ARG
D=M
@SP
A=M
M=D
@SP
M=M+1
@THIS
D=M
@SP
A=M
M=D
@SP
M=M+1
@THAT
D=M
@SP
A=M
M=D
@SP
M=M+1
@5
D=A
@1
D=D+A
@SP
D=M-D
@ARG
M=D
@SP
D=M
@LCL
M=D
@Main.fibonacci
0;JMP
(Main_RETURN_LABEL_1)
@SP
AM=M-1
D=M
A=A-1
M=M+D
@SP
A=M-1
D=M
@R13
M=D
@LCL
D=M
@5
A=D-A
D=M
@R14
M=D
@ARG
D=M+1
@SP
M=D
@LCL
D=M
@1
A=D-A
D=M
@THAT
M=D
@LCL
D=M
@2
A=D-A
D=M
@THIS
M=D
@LCL
D=M
@3
A=D-A
D=M
@ARG
M=D
@LCL
D=M
@4
A=D-A
D=M
@LCL
M=D
@R13
D=M
@SP
A=M-1
M=D
@R14
A=M
0;JMP