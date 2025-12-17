import numpy as np
import scipy as sp

A = np.array([
    [1,1,0,1,0,0],
    [0,0,1,1,1,0],
    [0,1,0,0,0,1],
    [0,0,0,0,1,1],
])

# w = np.array([7,4,5,3])

# def f(b: np.ndarray) -> float:
#     c = A @ b - w
#     return np.linalg.norm(c, 1) + np.linalg.norm(b, 1)

# whateverthisis = sp.optimize.minimize(f, x0=np.zeros(6), bounds=[(0, 7),(0, 5),(0, 4),(0, 4),(0, 3),(0, 3)], method="L-BFGS-B")

# s = [round(x,0) for x in whateverthisis.x]
# result = sum(s)
# print(whateverthisis)
# print(s)
# print(result)

c = np.array([1,1,1,1,1,1])
b_u = np.array([7,4,5,3])
b_l = np.array([7,4,5,3])

from scipy.optimize import LinearConstraint
constraints = LinearConstraint(A, b_l, b_u)

from scipy.optimize import milp
from scipy.optimize import Bounds
bounds = Bounds([0,0,0,0,0,0], [7,5,4,4,3,3])
res = milp(c=c, constraints=constraints, bounds=bounds)
print(res)
print(res.x)
print(sum(res.x))
