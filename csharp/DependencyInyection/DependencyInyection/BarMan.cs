namespace DependencyInyection {
    public class BarMan {
        IDrink _drink;

        public BarMan(IDrink drink) {
            _drink = drink;
        }

        public void Prepare() {
            this._drink.Prepare();
        }
    }
}