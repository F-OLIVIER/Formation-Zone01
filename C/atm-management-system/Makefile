# Execution en console :
# $: make
# $: ./atm

# objects contiens tous les fichier a build
objects = src/main.o src/system.o src/auth.o src/register.o src/account.o src/transaction.o

atm : $(objects)
	cc -o atm $(objects)

main.o : src/header.h
kbd.o : src/header.h
command.o : src/header.h
display.o : src/header.h
insert.o : src/header.h
search.o : src/header.h
files.o : src/header.h
utils.o : src/header.h

clean :
	rm -f $(objects)
