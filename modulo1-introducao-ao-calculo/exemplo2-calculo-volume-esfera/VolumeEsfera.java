import java.util.Scanner;

public class VolumeEsfera {
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);

        System.out.print("Digite o raio da esfera: ");
        double raio = scanner.nextDouble();

        double volume = (4.0 / 3.0) * Math.PI * Math.pow(raio, 3);
        System.out.println("Volume da Esfera: " + volume);

        scanner.close();
    }
}

//Para calcular o volume de uma esfera, utilizamos a fórmula:        4      3 
//                                                           V =  ------- 𝜋𝑟    
//                                                                   3
//
//       onde 𝑟  é o raio da esfera e 𝜋 é uma constante com valor aproximadamente igual a 3,14159.