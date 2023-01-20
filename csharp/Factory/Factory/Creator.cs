namespace Factory {
    public class Creator {
        public const int TINT_WINE = 1;
        public const int BEER = 2;

        public static AlcoholicDrink? DrinkCreator(int type) {
            AlcoholicDrink? value = type switch {
                TINT_WINE => new TintWine(),
                BEER => new Beer(),
                _ => null,
            };

            return value;
        }
    }
}