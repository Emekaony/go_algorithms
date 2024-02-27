
# ASCII Picture Renderer

# How it works
1. Every image is made up fundamentally of a bunch of pixels.
2. The goal of this project is to render an image using just a set of ASCII characters on a console output.
3. This is accomplished using the following steps:

# Steps

1. We get a source image and convert it into a 3-Dimentional array filled with the corresponding (R, G, B) values for every pixel in the image.

2. The next step invlved reducing the 3D matriix into an intensity matrix by mapping each set of r, g, b values into a brightness number by taking the average of all three.

3. To wrap things up, we map each brightness number into its corresponding ASCII character using a linrar normalization algorithm. Since pixel values range from 0 -> 255, and we have a small set of ascii characters, we had to use normalization to evenly distrubute brightness numbers from low to high so every number had a corresponding ascii character to it.

4. Last step involved printing the image to the console. For better output, we had to print each ascii character twice and also resize the terminal till the result looked good.


    

## Demo

Source Image: 
![Ascii_Pineapple](https://github.com/Emekaony/Ascii_Art/blob/main/assets/images/ascii-pineapple.jpg)

Rendered Image:

![Rendered_image](https://github.com/Emekaony/Ascii_Art/blob/main/assets/images/correct_ascii.png)



## Installation

```bash
  Fork the project
  cd ascii_art
  go run main.go
```
    
## Lessons Learned

1. I learned a lot about multi-dimensional image processing.
2. I got very comfortable with parsing through documentation for useful informatino
3. I learned to use low level bit manipulation when necessary.

