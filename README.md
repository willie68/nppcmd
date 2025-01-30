# nppcmd
Maybe you know this: you make lists of command lines in your favorite editor that you need again and again. For example, I have several lists, commands for Kubernetes management, port forwarding, Mongoshell commands, backup calls, etc.

But each time you first have to open a command shell, copy out the command and execute it.

This mini tool simplifies exactly this process by calling it via the command line with the file name and the line to be executed. The tool then loads the current file, gets the corresponding line and starts a Windows shell with it. Nothing more and nothing less.

Calling `nppcmd.exe C:\Users\w.klaas\Desktop\Docker.txt" 56`

would execute line 56 in a command shell.

In notepad++ the command is called like this: `<path to exe>\nppcmd.exe $(FULL_CURRENT_PATH) $(CURRENT_LINE)`
