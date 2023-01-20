namespace Composite {
    public class CakeComposite : Component {
        private List<Component> ingredients = new List<Component>();
        
        public decimal Cost {
            get {
                decimal value = 0;
                ingredients.ForEach(i => {
                    value += i.Price;
                });
                return value;
            }
        }

        public CakeComposite(string name, decimal price) : base(name, price) { }

        public void Add(Component element) {
            ingredients.Add(element);
        }

        public void Remove(Component element) {
            ingredients.Remove(element);
        }
    }
}