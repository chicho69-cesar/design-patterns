namespace PrototypeDeep {
    public class Animal : ICloneable {
        public string? Name { get; set; }
        public int Feets { get; set; }
        public Details? Details { get; set; }

        public object Clone() {
            Animal clone = (Animal)this.MemberwiseClone();
            clone.Details = new Details {
                Color = this.Details!.Color,
                Race = this.Details!.Race,
            };

            return clone;
        }
    }
}