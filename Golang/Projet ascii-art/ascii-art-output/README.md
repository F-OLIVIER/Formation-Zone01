
# <p align="center">Ascii Art Output</p>
  
An implementation of an output file for Ascii-Art.

## Color Option
The supported colors are : 

`red`, `green`, `blue`, `yellow`, `cyan`, `purple`, `gray`, `orange`

as well as their rgb codes ex :
red = "rgb(255, 0, 0)"

## Banner option
Our program can also take differents banners, the banners provided are : `standard`, `shadow`, `thinkertoy`

<table align= "center">
    <thead>
        <th align= "center">standard Ascii art</th>
        <th align= "center">shadow Ascii art</th>
        <th align= "center">thinkertoy Ascii art</th>
    </thead>
    <tbody>
        <tr>
            <td><img src="https://i43.servimg.com/u/f43/15/76/70/95/image_11.png"></td>
            <td><img src="https://i43.servimg.com/u/f43/15/76/70/95/captur22.png"></td>
            <td><img src="https://i43.servimg.com/u/f43/15/76/70/95/image_12.png"></td>
        </tr>    
    </tbody>
</table>

##  Usage
The file must be named by using the flag `--output=<fileName.txt>`, in which `--output` is the flag and `<fileName.txt>` is the file name which will contain the output.

The flag must have exactly the same format as above, any other formats must return the following usage message: `go run . [OPTION] [STRING] [BANNER]`

```go
go run . "exemple here"
go run . --color=red t toto --output=sample.txt shadow
```
Our program can also take the color as an option

##  Authors
- Fabien OLIVIER
- Fabien FANISE
- Valerian BOHERS
        
