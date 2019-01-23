import math
a= int(input('Enter a: '))
b= int(input('Enter b: '))
c= int(input('Enter c: '))
print('ax*x+bx+c=0')
d = b**2 - 4*a*c
print('gia tri cua d:', d)
if (d<0):
    print('Phuong trinh vo nghiem')
elif (d==0):
    x= -b/(2*a)
    print('gia tri cua ','x=%d' %(x) )
else:
    delta = math.sqrt(d)
    x1 = (-b+ delta)/(2*a)
    x2 =(-b- delta)/(2*a)
    print('Gia tri cua x1:', x1)

    print('Gia tri cua x2: ', x2)
