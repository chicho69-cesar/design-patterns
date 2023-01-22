import java.util.Scanner;

public class Juego {
    public static void main(String... args) {
        Scanner sc;
        sc = new Scanner(System.in);

        Marcianos marcianos;
        Computadora computadora;
        Jugador jugador;

        marcianos = Marcianos.getMarcianos();
        computadora = new Computadora();
        jugador = new Jugador();

        int disparos;

        do {
            System.out.print("Digite los disparos: ");
            disparos = sc.nextInt();

            for (int i=0; i<disparos; i++) {
                jugador.destruirMarciano();
            }

            computadora.creaMarcianos();
            marcianos.cuantosQuedan();
        } while (disparos != 0);
    }
}