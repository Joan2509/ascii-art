# Ascii-Art-fs
This project takes input text and a specified banner file then generates and displays the ASCII-Art graphical representation of the input text on the terminal using the  banner file format.
## Documentation
This section illustrates how to make use of this program.

### Installation
To run this program, download and install the latest version of Go from [here](https://go.dev/doc/install)

### Usage
1. Clone this repository on to the terminal by using the following command:
```bash
git clone https://learn.zone01kisumu.ke/git/josotieno/ascii-art-fs.git
```
2. Navigate into the ascii-art directory by using the command:
```bash
cd ascii-art-fs
```
3. To run the program, execute the command below:
```bash
go run . [STRING] [BANNER]

EX: go run . something standard
```
For example:
```bash
go run . "hello" standard
```
It should print out the following ASCII art on the terminal:
```bash
 _              _   _          
| |            | | | |         
| |__     ___  | | | |   ___   
|  _ \   / _ \ | | | |  / _ \  
| | | | |  __/ | | | | | (_) | 
|_| |_|  \___| |_| |_|  \___/  
                               
```


## Features
- This program can write the ascii-art graphical representation to a file by specifying an output flag.
- This program can as well display the ascii-art grapgical representation in a specified color using a color flag.
- The program is also able to run with a single string argument.
- This program has our own extra customized banner file. 

## Contributions
Pull requests are welcome where users of this program are allowed to contribute to this project in terms of adding features, or fixing bugs.

For major changes, please open an issue first to discuss what you would like to change.
## Authors
[josotieno](https://learn.zone01kisumu.ke/git/josotieno/)

[kada](https://learn.zone01kisumu.ke/git/kada/)

[jwambugu](https://learn.zone01kisumu.ke/git/jwambugu/)
## Licence
[MIT](https://choosealicense.com/licenses/mit/)
## Credits
[Zone01Kisumu](https://www.zone01kisumu.ke/)
