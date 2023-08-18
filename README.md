# Mandelbrot
## 1. Run the code:
In order to run the code, you need to run the following command: 
``` run main . ```

## 2. Functions:
### 2.1 Function Color():
This function is used in order to calculate the number of iterations needed to see if a point belongs or not to the Mandelbrot set.
This is where the Mandelbrot algorithm is used. The idea is to check if after n iterations the complex point (Cx, Cy) is still close to the point (0,0), which
means that this point belongs to the Mandelbrot, or it is far away which will mean that it's outside the Mandelbrot set. <br>

This function takes as parameters:
<ul>
  <li> n_max : The maximum number of iterations we want to do.</li>
  <li> Cx and Cy : Which are the coordinates of the point in the complex plane.</li>
</ul>

### 2.2 Function handlerGG():
This function enables the parallelization of the work. This function is used to calculate and determine the color of a pixel and making the different CPU's
of the machine to work simultaneously in order to do these calculations.

This function takes as parameters:
<ul>
  <li> hd *sync.WaitGroup : Pointer used in order to announce that the job has been done by one the threads.</li>
  <li> id : An Id for the point (pixel) we are working on.</li>
  <li> max_Iter: The maximum number of iterations that can be done. </li>
  <li> Cx and Cy: The coordinates of the point in the complex plane.</li>
  <li> i and y: which are the coordinates of the pixel in the picture.</li>
  <li> dc *gg.Context : Is the context in which we are working on.</li>
</ul>

### 2.3 Function getColor():
This function is going to give the color of a pixel based on the number of iterations that have been done by the Color function. <br>
If the number of iterations is equal to n_max, then this means that we are still relatively close to the point (0,0) which means that we are inside
the Mandelbrot Set so we will color the pixel as black. <br>
If the number of iterations is >= n_max - 10 this means that we are outside the Mandelbrot set, so we will color the pixel as red. <br>
Finally, we also use the HSV conversion in order to give a color transition inside the Mandelbrot, putting the corners as green/yellow for example. <br>

This function takes as parameters:
<ul>
  <li> iter: The number of iterations done.</li>
  <li> max_Iter: The maximum number of iterations that can be done</li>
</ul>

### 2.4 Function main():
This function is going to set up the HTTP server and display the Mandelbrot set. <br>
In order to do the calculations and then display the image it follows these different steps: <br>
  1. Set the number of workers you want to use simultaneously (= the number of CPU's of your PC)
  2. Loops around the different columns of the image.
  3. For each column is going to loop around the different lines, adding up the number of workers, so that we are not working on a line that another thread is
     working on.
  4. For each line, we are going to look each pixel.
  5. Now, we will do the different calculations in order to get the position in the complex plane and the color of this pixel.
  6. Once we have done that, we need to wait that all the workers have done their job before rendering the image.
  7. The image is rendered and sent as a HTTP response in order to display it on our server.

There is also an HTTP request for an interactive page where there are some calculations in order that when we press the buttons we can zoom in/out or move
around the image. It also calls the interactive.html file in order to display the different buttons with the image.
## 3. API:
There are two routes:
<ul>
  <li> http://localhost:8080/ : Display's the Mandelbrot set after all the calculations have been made. </li>
  <img width="612" alt="image" src="https://github.com/Pedropinto29/mandelbrot/assets/61003408/b9851ced-38f1-461b-9138-3eefeae33f13">

  <li> http://localhost:8080/interactive : Allows to play around with the buttons, like zooming in/out or move left/right/up/down. </li>
  <img width="611" alt="image" src="https://github.com/Pedropinto29/mandelbrot/assets/61003408/0aac172d-1328-4285-a4cd-5b2f409ce9d6">
</ul>

## 4. Used libraries:
<ul>
  <li> fmt : Formatting and printing strings. </li>
  <li> html/template : Implements data-driven templates for generating HTML output.</li>
  <li> image/color : Working with colors in images.</li>
  <li> net/http : Standard HTTP library. </li>
  <li> sync : Provides basic synchronization primitives. </li>
  <li> github.com/fogleman/gg : For rendering graphics. </li>
</ul>

## Done by:
Roquero Da Costa Pinto Pedro 17010























