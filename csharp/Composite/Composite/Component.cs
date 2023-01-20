namespace Composite {
    public abstract class Component {
        public string Name { get; set; }
        public decimal Price { get; set; }

        public Component(string name, decimal price) {
            Name = name;
            Price = price;
        }
    }
}