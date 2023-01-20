namespace Composite {
    public class Ingredient : Component {
        public int Quantity { get; set; }
        public string Unit { get; set; }

        public Ingredient(string name, decimal price, int quantity, string unit) 
            : base(name, price) {
            Quantity = quantity;
            Unit = unit;
        }
    }
}