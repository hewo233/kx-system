// This file is part of www.nand2tetris.org
// and the book "The Elements of Computing Systems"
// by Nisan and Schocken, MIT Press.
// File name: projects/04/Fill.asm

// Runs an infinite loop that listens to the keyboard input.
// When a key is pressed (any key), the program blackens the screen
// by writing 'black' in every pixel;
// the screen should remain fully black as long as the key is pressed. 
// When no key is pressed, the program clears the screen by writing
// 'white' in every pixel;
// the screen should remain fully clear as long as no key is pressed.

@8192
D=A
@nums
M=D // Set the nums to 8192

(LOOP)

    @KBD
    D=M // Get the keyboard input

    @WHITE
    D;JEQ

    @BLACK
    0;JMP

(BLACK)
    @i
    M=0 // Set the counter to 0 (use for print)

    @PRINTBLACK
    0;JMP

(WHITE)
    @i
    M=0 // Set the counter to 0 (use for print)

    @PRINTWHITE
    0;JMP

(PRINTBLACK)

    @nums
    D=M // Get the nums
    @i
    D=D-M // D = nums - i
    @LOOP
    D;JEQ // If D = 0, jump to LOOP

    @SCREEN
    D=A // Set the screen address to D

    @i
    D=D+M // D = screen + i

    @nowscreen
    A=D // Set the now address
    M=-1 // Set the now address to black
    
    @i
    M=M+1 // i++

    @PRINTBLACK
    0;JMP // Jump to print

(PRINTWHITE)

    @nums
    D=M // Get the nums
    @i
    D=D-M // D = nums - i
    @LOOP
    D;JEQ // If D = 0, jump to LOOP

    @SCREEN
    D=A // Set the screen address to D

    @i
    D=D+M // D = screen + i

    @nowscreen
    A=D // Set the now address
    M=0 // Set the now address to black
    
    @i
    M=M+1 // i++

    @PRINTWHITE
    0;JMP // Jump to print


    

