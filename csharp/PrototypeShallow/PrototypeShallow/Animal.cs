namespace PrototypeShallow {
    public class Animal : ICloneable {
        public string? Name { get; set; }
        public int Feets { get; set; }

        public object Clone() {
            return this.MemberwiseClone();
        }
    }
}