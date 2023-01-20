/*
Inyeccion de dependencias.- La inyecccion de dependencias es un patron de diseño
usado cuando tenemos una clase la cual tiene un objeto el cual puede tener diferentes
implementaciones y que con cada implementacion tenga un comportamiento diferente
entonces lo que hacemos es que creamos una interfaz a la cual le agregamos los
metodos que queremos que que las diferentes clases tengan y que sean los que 
definan su comportamiento, despues en la clase donde hacemos uso de estos en lugar
de crear el objeto alli mismo lo que hacemos es que inyectamos la interfaz y alli
la usamos, para despues en el punto donde hacemos uso de esta clase, simplemente
le pasamos una instancia que implemente la interfaz y que sea la que queremos que
lleve a cabo el comportamiento, asi, si algun dia queremos usar una nueva clase diferente
ya solo cambiamos en este punto la clase que inyectamos por la nueva que
queremos usar y esta se inyectara en el servicio que haga uso de ella
*/

namespace DependencyInyection {
    class Program {
        static void Main(string[] args) {
            var barMan1 = new BarMan(
                new PiñaColada("Picante")
            );
            barMan1.Prepare();

            var barMan2 = new BarMan(
                new Margarita()
            );
            barMan2.Prepare();
        }
    }
}