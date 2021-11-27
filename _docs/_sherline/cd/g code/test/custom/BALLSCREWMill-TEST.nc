%



#1 = 1 (assign parameter #1 the value of 1)
#2 = 60 (assing parameter #2 the limit max)
o101 while [#1 LT #2]
G01 X5 y4.5 Z3.4 F100
G01 X-5 y-4.5 Z-3.4 F100
(debug,Counter:#1 of #2)
G01 X0 y0 Z0 F100
 #1 = [#1+1] (increment the test counter)
o101 endwhile


%






