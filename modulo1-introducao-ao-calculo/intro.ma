(* Cálculo da Área de um Retângulo *)

largura = 10;
altura = 5;
areaRetangulo = largura * altura;
Print["Área do Retângulo: ", areaRetangulo];

Graphics[{EdgeForm[Black], FaceForm[LightBlue], Rectangle[{0, 0}, {largura, altura}]}]

(* Derivada de uma Função *)
f[x_] := x^2 + 3*x + 2;

derivada = D[f[x], x];
Print["Derivada de f(x): ", derivada];

Plot[{f[x], derivada}, {x, -10, 10}, PlotLegends -> {"f(x)", "f'(x)"}, 
PlotStyle -> {Blue, Red}, AxesLabel -> {"x", "y"}, 
PlotLabel -> "Função e sua Derivada"]

(* Área de uma Esfera *)

raio = 3;
areaEsfera = 4 * Pi * raio^2;
Print["Área da Esfera: ", areaEsfera];

(* Visualização da Esfera *)
Graphics3D[{EdgeForm[Black], FaceForm[LightGreen], Sphere[{0, 0, 0}, raio]}]
