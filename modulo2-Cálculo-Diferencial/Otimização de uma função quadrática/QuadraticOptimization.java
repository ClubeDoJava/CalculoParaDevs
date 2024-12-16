public class QuadraticOptimization {

    // Função quadrática: f(x) = ax^2 + bx + c
    public static double f(double a, double b, double c, double x) {
        return a * x * x + b * x + c;
    }

    // Derivada da função quadrática: f'(x) = 2ax + b
    public static double derivative(double a, double b, double x) {
        return 2 * a * x + b;
    }

    // Função para encontrar o ponto de mínimo
    public static double findMin(double a, double b, double c) {
        // O ponto de mínimo para uma função quadrática ocorre em x = -b / (2a)
        return -b / (2 * a);
    }

    public static void main(String[] args) {
        double a = 1.0;
        double b = -4.0;
        double c = 4.0;

        double minPoint = findMin(a, b, c);
        double minValue = f(a, b, c, minPoint);

        System.out.println("O ponto de mínimo ocorre em x = " + minPoint);
        System.out.println("O valor mínimo da função é f(x) = " + minValue);
    }
}
