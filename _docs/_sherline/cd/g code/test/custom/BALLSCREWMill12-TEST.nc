%




#1 = 1 (assign parameter #1 the value of 100)
o101 while [#1 LT 100]
G01 X3.5 y1.75 Z3.4 F100
G01 X-3.5 y-1.75 Z-3.4 F100
G01 X0 y0 Z0 F100
 #1 = [#1+1] (increment the test counter)
o101 endwhile

%




