namespace DependencyInyection {
    public class PiñaColada : IDrink {
        public string Type { get; set; }

        public PiñaColada(string type) {
            Type = type;
        }

        public void Prepare() {
            Console.WriteLine("Se hace una piña colada " + Type);
        }
    }
}