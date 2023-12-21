# MakeFile for lab 6

# Compiler
CC = gcc

# TODO: Fill in the flags for:
# 	optimization level 0
# 	debugger
# 	turn all optional warnings
# 	turn all warnings into errors
# 	set standard c to c99
CFLAGS =

# default target; prerequisite: lab6
# note: target does not execute anything
all: lab6
	# Don't change this line!
	@echo "Good Luck in Lab6!"

# creates lab6.s file from lab6.c
lab6.s: lab6.c
	# creates .s file from lab6.c
	# renames the output lab6.s
	$(CC) $(CFLAGS) -S lab6.c -o lab6.s

# creates lab6_helper.s file from lab6_helper.c and lab6_helper.h
lab6_helper.s: lab6_helper.c lab6_helper.h
	# TODO: create lab6_helper.s from lab6_helper.c
	# TODO: make sure to name the output lab6_helper.s
	$(CC) $(CLFLAGS) -S [FILL ME IN]

# creates lab6_helper.o file from lab6_helper.s
lab6_helper.o: lab6_helper.s
	# TODO: Fill in the name of the assembly file
	# TODO: Rename the output file to a .o file!
	$(CC) $(CFLAGS) -c [FileName] -o [NameOfOutput]

# creates lab6.o file from lab6.s
lab6.o: lab6.s
	# TODO: Fill in the missing flags & names!
	$(CC) $(CFLAGS) [FILL ME IN]

# creates lab6 executable from lab6.o and lab6_helper.o
lab6_final: lab6.o lab6_helper.o
	# creates an executable file that links lab6.o and lab6_helper.o
	# renames the output to lab6.o
	# TODO: Tell the compiler which .o files to link.
	$(CC) $(CFLAGS) -o lab6 [FileNames]

# Concise compilation rule
# Creates a lab6 executable file that depends on lab6.c lab6_helper.c and
# lab6_helper.h
lab6: lab6.c lab6_helper.c lab6_helper.h
	# creates an executable file named lab6 from lab6.c and lab6_helper.c files
	# TODO: Tell the compiler which .c files to compile.
	$(CC) $(CFLAGS) -o lab6 [FileNames]

# fill in the .PHONY statement!
.PHONY: clean
clean:
	# TODO: Tell the compiler to remove all .o files and all executables that
	# start with lab6!
	# HINT: Maybe try looking through the PowerPoint
	rm -f [FILL ME IN]
