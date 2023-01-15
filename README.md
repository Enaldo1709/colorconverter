# Color Converter

Color converter is a small command line tool used to convert colors between 3 of the most common color models: 
* RGB
* Hexadecimal
* X11 (256 Xterm)
# Getting Started
First step to work with color converter is to get the executable file:
## Building from source:
These steps are used to build the tool from source:
### Pre-requisites
There are some dependencies needed to build the project, install it before start the building process:
- make
- go v1.19
### Building
1. Clone the repository:
```bash
git clone https://github.com/Enaldo1709/colorconverter.git 
```
2. Get into the repository folder and build the project:
```bash
cd colorconverter && make build 
```
3. Executable file will be in the folder /out created in the current directory and will be ready to be executed. Generated files:
```bash
📂out
┗ 📦colorconverter
```

# Usage
## Color conversion use cases:
There are two use cases in the color converter tool usage:
### Convert Color Format
Used to convert a given color format to another:
```bash
colorconverter convert -i [color format] -o [color format] <color value>
```
Example: 
```bash
❯ colorconverter convert -i rgb -o hex 'rgb(95,175,255)'
#5FAFFF
```
### Get All Conversions
Used to get all supported color formats conversions from a specific input color format:
```bash
colorconverter all -i [color format] <color value>
```
Example:
```bash
❯ colorconverter convert -i hex '#5fafff'
Color conversions:
  Hex     #5FAFFF
  RGB     rgb(95,175,255)
  X11     Recommended closest value: 75
```

## Supported color formats:
Reference for **[color format]** | Description                    | Example  
---------------------------------|--------------------------------|----------
hex                              | Hexadecimal color format.      | #5fafff
rgb                              | Red Green Blue color format.   | rgb(95,175,255)
x11                              | X11 (256 Xterm) color format.  | 75

# Contributors:
Enaldo Narvaez Yepes - @Enaldo1709

Thanks!!