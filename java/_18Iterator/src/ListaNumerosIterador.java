public class ListaNumerosIterador implements Iterador {
    private int numeros[];
    private int posicion;

    public ListaNumerosIterador(int num[]) {
        this.numeros = num;
        this.posicion = 0;
    }

    @Override
    public Object siguiente() {
        return numeros[posicion++];
    }

    @Override
    public boolean tieneSiguiente() {
        return posicion < numeros.length && numeros[posicion] != 0;
    }
}