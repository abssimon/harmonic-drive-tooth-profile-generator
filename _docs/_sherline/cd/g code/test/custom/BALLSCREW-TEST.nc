%

#1 = 1 (assign parameter #1 the value of 20)
o101 while [#1 LT 100]
G01 X10 Z-6 F100
G01 X0 Z0 F100
 #1 = [#1+1] (increment the test counter)
o101 endwhile

%
