import numpy as np

# Definindo a função multivariada f(x, y)
def f(x, y):
    return x**2 + y**2

# Função para calcular o gradiente
def gradient(f, x, y, h=1e-5):
    df_dx = (f(x + h, y) - f(x - h, y)) / (2 * h)
    df_dy = (f(x, y + h) - f(x, y - h)) / (2 * h)
    return np.array([df_dx, df_dy])

# Ponto onde queremos calcular o gradiente
x = 1.0
y = 2.0

grad = gradient(f, x, y)
print(f"O gradiente de f(x, y) em ({x}, {y}) é {grad}")
