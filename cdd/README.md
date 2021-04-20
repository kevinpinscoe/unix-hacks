# cdd - change directory deeper

Uses ncurses via the pick command to select the next level 
directory to change in to

##Requisite:

* https://github.com/mptre/pick

##Use:

Put in your Bourne like shell's startup file:

cf. https://en.wikipedia.org/wiki/Bash_(Unix_shell)#Startup_scripts

`alias cdd='cd $(find . -maxdepth 1 -type d | pick)'`

1. Then re-source yur shells startup file (ex:  . ~/.bashrc)
2. run cdd command: (ex: $ cdd)

## Example:

$ ls -F
Projects/  code/  files/  work/
$ cdd

`.` <----

`./Projects`

`./code`

`./work`

`./files`



