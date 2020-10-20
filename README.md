# paraexp
A tool for parallely exploiting multiple targets concurrently. Mostly used for Attack and Defence CTF exploit automation

## Usage 

1) Create an exploit for a single target, taking the ip as a command line argument.
2) Now make changes in the config file according to your needs.
```{

   "teams":[
      
      "1.1.1.1",

      "2.2.2.2",

      "3.3.3.3"

   ],

   "flag_sub_ip":"10.10.10.10",
   
   "flag_sub_port":5555,

   "regex":"flag(.{32})"

}
``` 
3) just run the build.sh if you want to build the binary. (Remember to install `go` in your system before doing so)
4) `./paraexp exploit.py congif.json``
