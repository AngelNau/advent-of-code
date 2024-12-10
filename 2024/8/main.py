import fractions
dishmap = [[c for c in row.strip('\n')] for row in open('input.txt')]

dishes = []
for i,row in enumerate(dishmap):
    for j,c in enumerate(row):
        if c.isalnum():
            dishes.append((i,j,c))
            
for dish in dishes:
    pair_dishes = [x for x in dishes if x != dish and x[2] == dish[2]]
    for pair_dish in pair_dishes:
        dish_x, dish_y = dish[0], dish[1]
        pair_dish_x, pair_dish_y = pair_dish[0], pair_dish[1]
        #I have to handle the negative sign myself because fractions.Fraction() screws it up and can turn the vector around
        x_vec = pair_dish_x - dish_x
        y_vec = pair_dish_y - dish_y
        
        if x_vec < 0:
            x_sign = -1
        else:
            x_sign = 1
            
        if y_vec < 0:
            y_sign = -1
        else:
            y_sign = 1
        
        vector = fractions.Fraction(abs(pair_dish_x-dish_x),abs(pair_dish_y-dish_y))
        
        i, j, k = vector.numerator * x_sign, vector.denominator * y_sign, 1
        
        while 0 <= dish_x + k*i < len(dishmap) and 0 <= dish_y + k*j < len(dishmap[0]):
            dishmap[dish_x + k*i][dish_y + k*j] = '#'
            k += 1
print(sum([row.count('#') for row in dishmap]))
