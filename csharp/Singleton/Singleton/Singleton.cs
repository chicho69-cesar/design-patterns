namespace Singleton {
    public class Singleton {
        private static Singleton? instance = null;
        public string message = "";

        protected Singleton() {
            message = "Hello World!!!";
        }

        public static Singleton? Instance { 
            get {
                instance ??= new Singleton();
                return instance;
            }
        }
    }
}