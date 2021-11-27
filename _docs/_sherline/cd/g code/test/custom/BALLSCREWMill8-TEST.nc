%



#1 = 1 (assign parameter #1 the value of 20)
o101 while [#1 LT 2]
G01 X90 y70 Z60 F100
G01 X-90 y-70 Z-60 F100
G01 X0 y0 Z0 F100
 #1 = [#1+1] (increment the test counter)
o101 endwhile


%




